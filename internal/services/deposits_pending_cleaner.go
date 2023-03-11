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

//RunDepositsPendingCleaner - marks deposits that are too long in pending state as failed
func RunDepositsPendingCleaner(ctx context.Context, cfg config.Config) {
	r := &depositsPendingCleaner{
		log:           cfg.Log().WithField("service", "pending_deposits_cleaner"),
		cfg:           cfg,
	}

	running.WithBackOff(ctx, r.log, "pending_deposits_cleaner", r.runOnce, time.Minute, time.Second*10, time.Minute)
}

type depositsPendingCleaner struct {
	log           *logan.Entry
	cfg           config.Config
}

func (c *depositsPendingCleaner) runOnce(ctx context.Context) error {
	storage := c.cfg.Storage().Clone()
	err := storage.Transaction(func() error {
		expired := time.Now().UTC().Add(-30 * time.Minute)
		deposits, err := storage.Deposits().Select(data.DepositsSelector{
			States:       []int32{int32(kafka.DepositV2_STATE_PENDING)},
			MaxCreatedAt: &expired,
		})

		if err != nil {
			return errors.Wrap(err, "failed to select deposits to mark failed")
		}

		for _, d := range deposits {
			err = storage.Deposits().SetState(d.ID, int32(kafka.DepositV2_STATE_FAILED))
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
