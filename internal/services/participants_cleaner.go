package services

import (
	"context"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/events/go/kafka"
)

//RunParticipantsPendingCleaner - marks participants that are too long in deposit pending state as rejected
func RunParticipantsPendingCleaner(ctx context.Context, cfg config.Config) {
	r := &participantsPendingCleaner{
		log: cfg.Log().WithField("service", "pending_participants_cleaner"),
		cfg: cfg,
	}

	running.WithBackOff(ctx, r.log, "pending_participants_cleaner", r.runOnce, time.Hour, time.Second*10, time.Minute)
}

type participantsPendingCleaner struct {
	log *logan.Entry
	cfg config.Config
}

func (c *participantsPendingCleaner) runOnce(ctx context.Context) error {
	storage := c.cfg.Storage().Clone()
	err := storage.Transaction(func() error {
		expired := time.Now().UTC().Add(-24 * time.Hour)
		participants, err := storage.Participants().Select(data.ParticipantsSelector{
			State:         []kafka.Participant_State{kafka.Participant_STATE_DEPOSIT_PENDING},
			MinLotsEndsAt: &expired,
			ForUpdate:     true,
		})

		if err != nil {
			return errors.Wrap(err, "failed to select participants to mark rejected")
		}

		for _, d := range participants {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			err = storage.Participants().SetState(d.ID, kafka.Participant_STATE_REJECTED)
			if err != nil {
				return errors.Wrap(err, "failed to mark deposit as failed")
			}
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "tx failed")
	}

	return nil
}
