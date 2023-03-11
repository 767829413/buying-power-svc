package api

import (
	"context"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/handlers"
	"gitlab.com/eAuction/buying-power-svc/internal/services/power"
)

func Run(ctx context.Context, cfg config.Config) {
	startCopartBidderKeepAlive(ctx, cfg)
	r := chi.NewRouter()

	ape.DefaultMiddlewares(r, cfg.Log())

	r.Use(
		ape.CtxMiddleware(
			handlers.CtxConfig(cfg),
			handlers.WithStorage(cfg.Storage()),
			handlers.WithBuyingPowerCalc(power.NewCalculator(cfg)),
			handlers.WithRunTimeInfo(handlers.NewRunTimeInfo()),
		),
	)

	r.Get("/info", handlers.Info)
	r.Get("/ws/", handlers.WS)
	r.Post("/deposits", handlers.CreateDeposit)
	r.Get("/deposits", handlers.ListDeposits)
	r.Patch("/deposits/{deposit}", handlers.PatchDeposits)
	r.Get("/deposit_options", handlers.ListDepositOptions)
	r.Get("/deposit_options/{lot}", handlers.ListDepositOptionsLot)
	r.Post("/withdrawals", handlers.CreateWithdrawal)

	r.Route("/lots/{lot}", func(r chi.Router) {
		r.Put("/relationships/deposits", handlers.AttachDeposits)
		r.Put("/winner", handlers.PutWinnerDetails)
		r.Delete("/participants", handlers.DeleteParticipants)
		r.Delete("/participants/{account}", handlers.DeleteParticipant)

	})

	r.Route("/participants", func(r chi.Router) {
		r.Put("/", handlers.CreateParticipant)
		r.Put("/auto_bid_limit", handlers.CreateAutoBidLimit)
		r.Put("/buy_now", handlers.CreateBuyNow)
		r.Put("/bid_limit", handlers.CreateBidLimit)
		r.Put("/bid", handlers.CreateBid)
	})

	cfg.Log().WithField("service", "api").Info("service started")
	ape.Serve(ctx, r, cfg, ape.ServeOpts{})

}
