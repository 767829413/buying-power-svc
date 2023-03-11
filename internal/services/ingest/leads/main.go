package leads

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/events/go/kafka"
)

func Run(ctx context.Context, cfg config.Config) {
	leads := kafka.NewIngest(&leadsHandler{
		cfg,
		cfg.Log().WithField("ingest", "leads"),
	}, cfg.LeadsIngester())

	leads.Run(ctx)
}

type leadsHandler struct {
	cfg config.Config
	log *logan.Entry
}

func leadUpdated(storage data.Storage, event *kafka.LeadEvent) error {
	lead := event.LeadUpdated
	if lead.Identity == "" {
		return nil
	}

	err := storage.IdentitiesToBrokers().Delete(lead.Identity)
	if err != nil {
		return errors.Wrap(err, "failed to delete existing lead to identity relations")
	}

	for _, broker := range lead.Brokers {
		err = storage.IdentitiesToBrokers().Insert(data.IdentityToBroker{
			Identity: lead.Identity,
			Broker:   broker,
		})
		if err != nil {
			return errors.Wrap(err, "failed to insert identity to broker")
		}
	}

	return nil
}

func (h *leadsHandler) Handle(ctx context.Context, messages []kafka.Message) error {
	storage := h.cfg.Storage()
	err := storage.Transaction(func() error {
		for _, raw := range messages {

			if ctx.Err() != nil {
				return ctx.Err()
			}

			msg := raw.MustLeadEventMessage()
			event := msg.Value
			switch event.Type {
			case kafka.LeadEvent_TYPE_LEAD_UPDATED:
				if err := leadUpdated(storage, event); err != nil {
					return errors.Wrap(err, "failed to process lead updated event")
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
