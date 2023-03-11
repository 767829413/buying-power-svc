package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"sync"
)

type fetchFn func(ctx context.Context) (kafka.Message, error)

// streamer - helper runner that sends messages into the chanel.
// if error occurs - chanel is closed and new instance of streamer must be created
type streamer struct {
	log       *logan.Entry
	msgStream chan Message
	fetch     fetchFn

	cancel    context.CancelFunc
	wg sync.WaitGroup
}

func startNewStreamer(log *logan.Entry, fetch fetchFn, batchSize int) *streamer {
	ctx, cancel := context.WithCancel(context.Background())
	s := &streamer{
		log:       log,
		msgStream: make(chan Message, batchSize),
		fetch:     fetch,
		cancel:    cancel,
	}

	s.log.Info("starting new streamer")
	go s.run(ctx)
	return s
}

//Close - stops streamer
func (s *streamer) Close() {
	s.cancel()
	s.wg.Wait()
}

//Stream - returns chanel with kafka msgs
func (s *streamer) Stream() <-chan Message {
	return s.msgStream
}

func (s *streamer) run(ctx context.Context) {
	s.wg.Add(1)
	defer s.wg.Done()
	defer func() {
		recovered := recover()
		if recovered != nil {
			s.log.WithRecover(recovered).Error("batch reader panicked - stopping")
		}
	}()

	defer func() {
		close(s.msgStream)
	}()

	for {
		err := s.runOnce(ctx)
		switch errors.Cause(err) {
		case context.Canceled, context.DeadlineExceeded:
			return
		case nil:
			continue
		default:
			s.log.WithError(err).Error("streamer failed - stopping")
			return
		}
	}
}

func (s *streamer) runOnce(ctx context.Context) (err error) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			err = errors.Wrap(errors.FromPanic(recovered), "paniced")
		}
	}()
	msg, err := s.fetch(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to fetch msg")
	}

	select {
	case s.msgStream <- Message{msg: msg}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
