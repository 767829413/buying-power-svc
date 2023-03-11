package main

import (
	"context"
	"gitlab.com/eAuction/buying-power-svc/internal/services/ingest/leads"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"gitlab.com/eAuction/buying-power-svc/internal/services"
	"gitlab.com/eAuction/buying-power-svc/internal/services/ingest/identities"
	"gitlab.com/eAuction/buying-power-svc/internal/services/ingest/lots"
	"gitlab.com/eAuction/buying-power-svc/internal/services/ingest/tbc"

	migrate "github.com/rubenv/sql-migrate"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gopkg.in/alecthomas/kingpin.v2"

	"gitlab.com/eAuction/buying-power-svc/internal/assets/assets"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api"
)

var migrations = &migrate.PackrMigrationSource{
	Box: assets.Migrations,
}

func main() {
	defer func() {
		if rvr := recover(); rvr != nil {
			logan.New().WithRecover(rvr).Fatal("app panicked")
		}
		os.Stdout.Sync()
		os.Stderr.Sync()
	}()

	app := kingpin.New("buying-power", "")

	runCmd := app.Command("run", "")

	runApiOnly := app.Command("api", "")
	readTBC := app.Command("read_tbc", "")

	migrateCmd := app.Command("migrate", "")
	migrateDownCmd := migrateCmd.Command("down", "")
	migrateUpCmd := migrateCmd.Command("up", "")
	migrateRedoCmd := migrateCmd.Command("redo", "")

	cmd, err := app.Parse(os.Args[1:])
	if err != nil {
		logan.New().WithError(err).Fatal("failed to parse arguments")
	}

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	cfg := config.New(ctx, kv.MustFromEnv())
	cfg.InjectRemoteStateGetter()

	run := func(f func(context.Context, config.Config)) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f(ctx, cfg)
		}()
	}

	switch cmd {
	case readTBC.FullCommand():
		cfg.Log().Info("reading tbc transactions")
		tbc.PrintTBC(ctx, cfg.Log())
	// run all services
	case runApiOnly.FullCommand():
		cfg.Log().Info("starting api only")
		run(api.Run)
	case runCmd.FullCommand():
		cfg.Log().Info("starting all services")
		run(api.Run)
		run(services.RunParticipantsSyncer)
		run(tbc.Run)
		run(lots.Run)
		run(identities.Run)
		run(services.RunLimitsSetter)
		run(services.RunCardRefund)
		run(services.RunBalanceRefund)
		run(services.RunFeeCharger)
		run(services.RunDepositsPendingCleaner)
		run(services.RunLotsCleaner)
		run(services.RunWinnersPublisher)
		run(services.RunBuyingPowerRecharger)
		run(services.RunParticipantsPendingCleaner)
		run(services.RunDepositsSyncer)
		run(services.RunLotCloser)
		run(leads.Run)
	// migrate
	case migrateDownCmd.FullCommand():
		applied, err := migrate.Exec(cfg.RawDB(), "postgres", migrations, migrate.Down)
		if err != nil {
			cfg.Log().WithError(err).Fatal("failed to apply migrations")
		}
		cfg.Log().WithFields(logan.F{
			"direction": "down",
			"applied":   applied,
		}).Info("migrations applied")
		return
	case migrateUpCmd.FullCommand():
		applied, err := migrate.Exec(cfg.RawDB(), "postgres", migrations, migrate.Up)
		if err != nil {
			cfg.Log().WithError(err).Fatal("failed to apply migrations")
		}
		cfg.Log().WithFields(logan.F{
			"direction": "up",
			"applied":   applied,
		}).Info("migrations applied")
		return
	case migrateRedoCmd.FullCommand():
		applied, err := migrate.Exec(cfg.RawDB(), "postgres", migrations, migrate.Down)
		if err != nil {
			cfg.Log().WithError(err).Fatal("failed to apply migrations")
		}
		cfg.Log().WithFields(logan.F{
			"direction": "down",
			"applied":   applied,
		}).Info("migrations applied")
		applied, err = migrate.Exec(cfg.RawDB(), "postgres", migrations, migrate.Up)
		if err != nil {
			cfg.Log().WithError(err).Fatal("failed to apply migrations")
		}
		cfg.Log().WithFields(logan.F{
			"direction": "up",
			"applied":   applied,
		}).Info("migrations applied")
		return
	}

	var gracefulStop = make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	// making WaitGroup usable in select
	wgch := make(chan struct{})
	go func() {
		wg.Wait()
		close(wgch)
	}()

	select {
	// listening for runners stop
	case <-wgch:
		cfg.Log().Warn("all services stopped")
	// listening for OS signals
	case <-gracefulStop:
		cfg.Log().Info("received signal to stop")
		cancel()
		<-wgch
	}
}
