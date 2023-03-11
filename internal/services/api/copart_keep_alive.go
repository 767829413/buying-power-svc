package api

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"time"
)

//startCopartBidderKeepAlive - starts goroutines that keep alive live bidder
func startCopartBidderKeepAlive(ctx context.Context, cfg config.Config) {
	cfg.Log().Info("initializing live bidder")
	if len(cfg.GetSupportedCountries()) == 0 {
		cfg.Log().Warn("live bidding is disabled - no countries specified")
		return
	}
	for _, country := range cfg.GetSupportedCountries() {
		go runCopartBidderKeepAlive(ctx, cfg, country)
	}

	cfg.Log().Info("initialized all live bidders keep alive")
}

func runCopartBidderKeepAlive(ctx context.Context, cfg config.Config, country string) {
	log := cfg.Log().WithField("service", "copart-bidder-keep-alive").WithField("country", country)
	running.WithBackOff(ctx, log, "keep-alive", func(ctx context.Context) error {
		copart, err := cfg.Platformer().GetPlatformByCode("COPART")
		if err != nil {
			return errors.Wrap(err, "failed to get copart platform")
		}

		if copart == nil {
			return errors.New("expected copart platform to exist")
		}

		to := time.Now().UTC().Add(40*time.Minute)
		from := time.Now().UTC().Add(-4*time.Hour)
		sale, err := cfg.Storage().Sales().GetSaleWithinPeriod(copart.Data.ID, from, to)
		if err != nil {
			return errors.Wrap(err, "failed to get sale")
		}

		if sale == nil {
			log.WithFields(logan.F{
				"from": from.String(),
				"to": to.String(),
			}).Info("no active sales found - not starting live bidder")
			return nil
		}
		initCtx, cancel := context.WithTimeout(ctx, 2*time.Minute)
		defer cancel()
		bidder, err := cfg.Real(initCtx, country)
		if err != nil {
			return errors.Wrap(err, "failed to initialize live bidder")
		}
		cfg.Log().Info("initialized live bidder")
		cancel()
		select {
		case <-bidder.Ctx().Done():
			return errors.Wrap(bidder.Ctx().Err(), "bidder stopped")
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "received signal to stop")
		}

	}, 10*time.Second, time.Second, 5*time.Minute)
}
