package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type BuyingPowerRecharger struct {
	IsEnabled            bool  `fig:"is_enabled"`
	TradersDepositAmount int64 `fig:"target_deposit_amount"`
}

func (c *config) BuyingPowerRecharger() BuyingPowerRecharger {
	return c.buyingPowerRecharger.Do(func() interface{} {
		raw := kv.MustGetStringMap(c.getter, "buying_power_recharger")

		var cfg BuyingPowerRecharger

		if err := figure.Out(&cfg).With(figure.BaseHooks).From(raw).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out buying power recharger config"))
		}

		if cfg.IsEnabled && cfg.TradersDepositAmount <= 0 {
			panic(errors.New("expected target deposit amount to be greater 0"))
		}

		return cfg
	}).(BuyingPowerRecharger)
}
