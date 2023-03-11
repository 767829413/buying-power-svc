package identities

import (
	"context"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/events/go/kafka"
)

func Run(ctx context.Context, cfg config.Config) {
	identities := kafka.NewIngest(&identitiesHandler{
		cfg,
		cfg.Log().WithField("ingest", "identities"),
	}, cfg.IdentitiesIngester())

	identities.Run(ctx)
}

type identitiesHandler struct {
	cfg config.Config
	log *logan.Entry
}

func identityCreated(storage data.Storage, event *kafka.IdentityEvent) error {
	if err := storage.Identities().Store(data.Identity{
		ID:       event.Address,
		Platform: event.Platform,
		Type:     event.Created.Type,
	}); err != nil {
		return errors.Wrap(err, "failed to store identity")
	}
	return nil
}

func (h *identitiesHandler) Handle(ctx context.Context, messages []kafka.Message) error {
	storage := h.cfg.Storage()
	err := storage.Transaction(func() error {
		for _, raw := range messages {

			if ctx.Err() != nil {
				return ctx.Err()
			}

			msg := raw.MustIdentityEventMessage()
			event := msg.Value
			switch event.Type {
			case kafka.Identity_EVENT_TYPE_CREATED:
				if err := identityCreated(storage, event); err != nil {
					return errors.Wrap(err, "failed to process created event")
				}
			case kafka.Identity_EVENT_TYPE_DELETED:
				if err := storage.Identities().MarkDeleted(event.Address); err != nil {
					return errors.Wrap(err, "failed to mark identity as deleted", logan.F{
						"identity": event.Address,
					})
				}
			}
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to commit tx")
	}

	return nil
}
