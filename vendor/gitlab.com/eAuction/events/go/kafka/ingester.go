package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

//Handler - processes kafka messages. MUST be able idempotent.
// If returns error, will be called with the same batch of messages
type Handler interface {
	Handle(ctx context.Context, msgs []Message) error
}

//Ingest - wrapper of kafka reader that reads msgs from specified topic and passes them into Handler.
// Guaranties that all messages are delivered to handler at least once and in the same order as they are present in kafka.
// Automatically commits successfully processed messages.
type Ingest struct {
	log      *logan.Entry
	handler  Handler
	cfg      IngestConfig
	reader   *kafka.Reader
	streamer *streamer
}

var streamerDiedErr = errors.New("streamer died")

//NewIngest - creates new instance of ingest
func NewIngest(handler Handler, cfg IngestConfig) *Ingest {
	return &Ingest{
		log:     cfg.Log,
		handler: handler,
		cfg:     cfg,
	}
}

func (i *Ingest) init() {
	readerCfg := i.cfg.Config.reader(i.cfg.Topic.String())
	if i.cfg.ingestConfig.GroupID != "" {
		readerCfg.GroupID = i.cfg.ingestConfig.GroupID
	}
	i.reader = kafka.NewReader(kafka.ReaderConfig(readerCfg))
	i.streamer = startNewStreamer(i.cfg.Log.WithField("runner", "streamer"),
		i.reader.FetchMessage, i.cfg.ingestConfig.IngestBatchSize)
}

func (i *Ingest) cleanup() {
	i.streamer.Close()
	err := i.reader.Close()
	if err != nil {
		panic(errors.Wrap(err, "unexpected error on closing kafka reader"))
	}
}

//Run - runs ingest. Must be called only once
func (i *Ingest) Run(ctx context.Context) {
	i.init()

	defer func() {
		recovered := recover()
		if recovered != nil {
			i.log.WithRecover(recovered).Error("Ingest panic")
		}
	}()

	defer func() {
		i.cleanup()
	}()

	for {
		if ctx.Err() != nil {
			i.log.Info("context signaled to stop")
			return
		}

		running.UntilSuccess(ctx, i.log, "run_once", func(ctx context.Context) (bool, error) {
			err := i.runOnce(ctx)
			switch errors.Cause(err) {
			case context.DeadlineExceeded, context.Canceled:
				return false, nil
			case nil:
				return true, nil
			default:
				i.cleanup()
				i.init()
				return false, err
			}
		}, i.cfg.MinRetryPeriod, i.cfg.MaxRetryPeriod)
	}
}

func (i *Ingest) readBatch(ctx context.Context) ([]Message, error) {
	var msg Message
	var ok bool
	// block until msg received or ctx canceled
	stream := i.streamer.Stream()
	select {
	case msg, ok = <-stream:
		if !ok {
			return nil, streamerDiedErr
		}
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	result := make([]Message, 0, len(stream)+1)
	result = append(result, msg)

	// read whole chanel up to batch size
	for {
		if len(result) >= i.cfg.IngestBatchSize {
			return result, nil
		}

		select {
		case msg, ok = <-stream:
			if !ok {
				return nil, streamerDiedErr
			}
			result = append(result, msg)
		default:
			return result, nil
		}
	}
}

func (i *Ingest) runOnce(ctx context.Context) (err error) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			err = errors.Wrap(errors.FromPanic(recovered), "panic")
		}
	}()

	batch, err := i.readBatch(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to read batch")
	}

	if len(batch) == 0 {
		return nil
	}

	log := i.log.WithFields(logan.F{
		"from":           batch[0].Offset(),
		"from_partition": batch[0].Partition(),
		"to":             batch[len(batch)-1].Offset(),
		"partition":      batch[len(batch)-1].Partition(),
		"to_time":        batch[len(batch)-1].Time().UTC(),
	})

	log.Info("Received new batch")

	skipMessage := i.getSkipUntilMessage(batch)
	if skipMessage != nil {
		log.WithField("skip_until_offset", skipMessage.Offset).WithField("partition", skipMessage.Partition).
			Warn("received old event with offset that should be skipped - starting skip process. " +
				"note: might occur to number of partitions specified in config")
		err = i.reader.CommitMessages(ctx, *skipMessage)

		if err != nil {
			return errors.Wrap(err, "failed to commit offset shift")
		}

		return errors.From(errors.New("shifted offset - recreating reader"), logan.F{
			"new_offset": skipMessage.Offset,
			"partition":  skipMessage.Partition,
		})
	}
	running.UntilSuccess(ctx, log, "handle", func(ctx context.Context) (bool, error) {
		err = i.handler.Handle(ctx, batch)
		if err != nil {
			return false, errors.Wrap(err, "failed to handle msg batch")
		}

		return true, nil
	}, i.cfg.MinRetryPeriod, i.cfg.MaxRetryPeriod)

	if ctx.Err() != nil {
		return ctx.Err()
	}

	log.Info("Successfully processed batch")

	running.UntilSuccess(ctx, log, "commit", func(ctx context.Context) (bool, error) {
		err = i.reader.CommitMessages(ctx, msgBatchToRaw(batch)...)
		if err != nil {
			return false, errors.Wrap(err, "failed to commit msg")
		}

		return true, nil
	}, i.cfg.MinRetryPeriod, i.cfg.MaxRetryPeriod)
	log.Info("Successfully committed batch")
	return nil
}

func (i *Ingest) getSkipUntilMessage(batch []Message) *kafka.Message {
	for _, msg := range batch {
		skipUntilOffset, ok := i.cfg.SkipUntilOffset[msg.Partition()]
		if !ok {
			continue
		}

		if skipUntilOffset <= msg.Offset() {
			continue
		}

		return &kafka.Message{
			Offset:    skipUntilOffset,
			Topic:     msg.Topic().String(),
			Partition: msg.Partition(),
		}
	}

	return nil
}

func msgBatchToRaw(batch []Message) []kafka.Message {
	result := make([]kafka.Message, len(batch))
	for i := range batch {
		result[i] = batch[i].msg
	}

	return result
}
