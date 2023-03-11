package services

import (
	"context"
	"gitlab.com/eAuction/buying-power-svc/internal/services/deposits"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
)

//RunBuyingPowerRecharger - ensures that dealer accounts has sufficient amount of available deposits
func RunBuyingPowerRecharger(ctx context.Context, cfg config.Config) {
	const name = "buying_power_recharger"
	r := &buyingPowerRecharger{
		log:           cfg.Log().WithField("service", name),
		cfg:           cfg,
	}

	if !cfg.BuyingPowerRecharger().IsEnabled {
		r.log.Warn("disabled")
		return
	}

	running.WithBackOff(ctx, r.log, name, r.runOnce, time.Minute, time.Second*10, time.Minute)
}

type buyingPowerRecharger struct {
	log           *logan.Entry
	cfg           config.Config
}

func (c *buyingPowerRecharger) runOnce(ctx context.Context) error {
	storage := c.cfg.Storage()
	identities, err := storage.Identities().SelectNeedRecharge(c.cfg.AutoRechargeableIdentityTypes(),
		c.cfg.BuyingPowerRecharger().TradersDepositAmount)
	if err != nil {
		return errors.Wrap(err, "failed to get identities that required buying power recharge")
	}

	c.log.WithField("identities", len(identities)).Debug("found identities that require recharge")
	for _, identity := range identities {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		c.log.WithField("identity", identity.ID).Debug("starting recharge")
		txStorage := storage.Clone()
		err = running.RunSafely(ctx, "", func(ctx context.Context) error {
			return txStorage.Transaction(func() error {
				return c.rechargeBalance(ctx, txStorage, identity)
			})
		})
		if err != nil {
			c.log.WithError(err).WithFields(logan.F{
				"identity":  identity.ID,
			}).Error("failed to recharge buying power for identity")
			continue
		}

	}

	return nil
}

func getNumberOfDepositsToCreate(tradersDepositAmount, depositAmount int64, identity data.IdentityWithDeposiAmount) int64 {
	delta := tradersDepositAmount - identity.DepositAmount
	if delta <= 0 {
		panic(errors.New("expected traders deposit to > deposit amount"))
	}

	n := delta/depositAmount
	if delta%depositAmount != 0 {
		n++
	}

	return n
}

func (c *buyingPowerRecharger) rechargeBalance(ctx context.Context, storage data.Storage, identity data.IdentityWithDeposiAmount) error {
	limits := c.cfg.Deposits().GetLimits(identity.Platform)
	limit := limits.GetMaxLimit()
	n := getNumberOfDepositsToCreate(c.cfg.BuyingPowerRecharger().TradersDepositAmount, limit.Deposit, identity)
	for i := int64(0); i < n; i++ {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		deposit, err := storage.Deposits().Create(limit.Deposit, c.cfg.Deposits().Currency, identity.ID)
		if err != nil {
			return errors.Wrap(err, "failed to create deposit")
		}

		err = deposits.ManuallyApprove(storage, deposit)
		if err != nil {
			return errors.Wrap(err, "failed to manually approve deposit")
		}
	}

	return nil
}
