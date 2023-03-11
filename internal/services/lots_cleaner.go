package services

import (
	"context"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
)

type lotsCleaner struct {
	cfg config.Config
	log *logan.Entry
}

//RunLotsCleaner - removes redundant lots
func RunLotsCleaner(ctx context.Context, cfg config.Config) {
	s := lotsCleaner{
		cfg: cfg,
		log: cfg.Log().WithField("service", "lots_cleaner"),
	}
	running.WithBackOff(ctx, s.log, "lots_cleaner", s.runOnce, time.Hour, 5*time.Second, time.Minute)
}

func (s *lotsCleaner) runOnce(ctx context.Context) error {
	for {
		if ctx.Err() != nil {
			return nil
		}
		s.log.Info("removing batch of outdated lots")
		shouldContinue, err := s.cfg.Storage().Lots().DeleteRedundant()
		if err != nil {
			return errors.Wrap(err, "failed to delete redundant lots")
		}

		if !shouldContinue {
			return nil
		}
	}

}
