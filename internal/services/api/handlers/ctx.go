package handlers

import (
	"context"
	"net/http"

	"gitlab.com/eAuction/bouncer"
	"gitlab.com/eAuction/bouncer/jwt"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/power"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	configCtxKey ctxKey = iota
	storageCtxKey
	buyingPowerCalcCtxKey
	runTimeInfoCtxKey
)

func WithStorage(storage data.Storage) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, storageCtxKey, storage.Clone())
	}
}

func WithRunTimeInfo(i *RunTimeInfo) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, runTimeInfoCtxKey, i)
	}
}

func GetRunTimeInfo(r *http.Request) *RunTimeInfo {
	return r.Context().Value(runTimeInfoCtxKey).(*RunTimeInfo)
}

func WithBuyingPowerCalc(calculator *power.Calculator) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, buyingPowerCalcCtxKey, calculator)
	}
}

func CtxConfig(cfg config.Config) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, configCtxKey, cfg)
	}
}

func Storage(r *http.Request) data.Storage {
	return r.Context().Value(storageCtxKey).(data.Storage)
}

func Log(r *http.Request) *logan.Entry {
	return ape.Log(r)
}

func BuyingPowerCalculator(r *http.Request) *power.Calculator {
	return r.Context().Value(buyingPowerCalcCtxKey).(*power.Calculator)
}

func Config(r *http.Request) config.Config {
	return r.Context().Value(configCtxKey).(config.Config)
}

func Check(r *http.Request, rules ...bouncer.Rule) (*jwt.Claims, error) {
	return Config(r).Bouncer().Check(r, rules...)
}
