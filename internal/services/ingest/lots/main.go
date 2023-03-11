package lots

import (
	"context"
	"time"

	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/events/go/kafka"
)

func Run(ctx context.Context, cfg config.Config) {
	ingest := kafka.NewIngest(&messageHandler{
		log:     cfg.Log(),
		storage: cfg.Storage(),
		specificMessageHandlers: map[kafka.AuctionEvent_Type]specificMessageHandler{
			kafka.AuctionEvent_TYPE_LOT_CREATED:          newLotCreated(cfg.Log()),
			kafka.AuctionEvent_TYPE_LOT_STATE_UPDATED:    newLotStateUpdated(cfg.Log()),
			kafka.AuctionEvent_TYPE_PARTICIPANT_CREATED:  newParticipant(cfg.Log(), cfg.Platformer()),
			kafka.AuctionEvent_TYPE_PARTICIPANT_UPDATED:  newParticipant(cfg.Log(), cfg.Platformer()),
			kafka.AuctionEvent_TYPE_DEPOSIT:              &noOpHandler{},
			kafka.AuctionEvent_TYPE_WINNER:               newWinner(cfg.Log()),
			kafka.AuctionEvent_TYPE_LIVE_BIDDING_DETAILS: newLiveBiddingDetails(cfg.Log(), cfg),
		},
	}, cfg.LotsIngester())

	ingest.Run(ctx)
}

type specificMessageHandler interface {
	Handle(storage data.Storage, message *kafka.AuctionEventMessage) error
}

type noOpHandler struct{}

func (n *noOpHandler) Handle(storage data.Storage, message *kafka.AuctionEventMessage) error {
	return nil
}

type messageHandler struct {
	log     *logan.Entry
	storage data.Storage

	specificMessageHandlers map[kafka.AuctionEvent_Type]specificMessageHandler
}

// HandleRequest handles kafka messages by passing them to correct event handlers.
func (h *messageHandler) Handle(ctx context.Context, messages []kafka.Message) (err error) {
	storage := h.storage.Clone()
	err = storage.Transaction(func() error {
		for _, raw := range messages {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			msg := raw.MustAuctionEventMessage()
			if err := h.handleMessage(storage, msg); err != nil {
				return errors.Wrap(err, "failed to handle message")
			}
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "batch handle failed")
	}

	return nil
}

func (h *messageHandler) handleMessage(storage data.Storage, msg kafka.AuctionEventMessage) error {
	specificHandler, ok := h.specificMessageHandlers[msg.Value.Type]
	if !ok {
		h.log.WithField("offset", msg.Offset()).WithField("event_type", msg.Value.Type).Info("no handler for event type")
		return nil
	}

	start := time.Now()
	err := specificHandler.Handle(storage, &msg)
	h.log.WithField("duration", time.Since(start)).WithField("msg_type", msg.Value.Type.String()).Info("handled msg")
	switch errors.Cause(err) {
	case nil:
		return nil
	case ErrLotNotFound, ErrParticipantNotFound:
		errMsg := "Unable to process message because entity that this event depends from is not found " +
			"Probably it's because of a bug in the writer side. Message will be skipped. For being able " +
			"to re-ingest the data from this message - you will have to move writer cursor to this message " +
			"and write all events happened after it again. (Buying-power handling them in idempotently)"
		h.log.
			WithFields(logan.F{
				"msg_offset": msg.Offset(),
				"msg_time":   msg.Time(),
			}).
			WithError(err).
			Warn(errMsg)
		return nil
	default:
		return errors.Wrap(err, "failed to handle specific message")
	}
}
