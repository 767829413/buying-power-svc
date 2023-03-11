package power

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/types"
	"gitlab.com/eAuction/exrates-svc/pkg/exrates"
	"gitlab.com/eAuction/platformer-svc/resources"
)

//ErrNotAvailable - signals that requested resources is not available
var ErrNotAvailable = errors.New("requested resource is not available")

type base interface {
	BuyingPower(depositorID string, depositAmount int64, depositCurrency, lotCurrency string) (int64, error)
	GetFeeChangedFromDeposit(depositor string, lot data.Lot, customerPrice int64, customerPriceCurrency string) (*resources.Fee, error)
	GetDepositLimits(depositor string, lot *data.Lot, customerPrice int64, customerPriceCurrency string) ([]config.Limit, error)
	//GetDepositsOptions - returns options for deposit state, deposits to attach, deposit to create
	GetDepositsOptions(depositor string, balance int64, attachedDeposits, availableDeposits []data.Deposit, limit config.Limit) (types.DepositOptionState, []data.Deposit, *data.Deposit)
}

//Calculator - calculates buying power for specified deposit amount
type Calculator struct {
	base
	converter *exrates.Converter
	deposits  config.Deposits
}

//NewCalculator - returns new instance of buying power calculator
func NewCalculator(cfg config.Config) *Calculator {
	deposits := cfg.Deposits()
	converter := cfg.ExRatesConverter()
	var b base
	if deposits.DepositFraction != nil {
		b = &dynamicDepositsAndFee{
			depositFraction:    *deposits.DepositFraction,
			converter:          converter,
			deposits:           deposits,
			depositIDGenerator: cfg.DepositIDGenerator(),
		}
	} else {
		b = &fixedDepositsAndFee{
			deposits:           deposits,
			converter:          converter,
			platforms:          cfg.Platformer(),
			storage:            cfg.Storage(),
			depositIDGenerator: cfg.DepositIDGenerator(),
			cfg:                cfg,
		}
	}
	return &Calculator{
		base:      b,
		converter: converter,
		deposits:  deposits,
	}
}

//IsDepositSufficientToPayFee - returns true if deposit amount is sufficient
func (c *Calculator) IsDepositSufficientToPayFee(depositor string, lot data.Lot,
	depositAmount int64, depositCurrency string) (bool, error) {
	fee, err := c.GetFeeChangedFromDeposit(depositor, lot, 0, depositCurrency)
	if err != nil {
		return false, errors.Wrap(err, "failed to get shelf fee")
	}
	if fee == nil {
		return false, errors.Wrap(ErrNotAvailable, "shelf fee is not configured for depositor's and lot's platforms pair")
	}

	feeDepositCurrency, _, err := c.converter.ConvertCents(fee.Currency, depositCurrency, int64(fee.Fixed*100))
	if err != nil {
		return false, errors.Wrap(err, "failed to convert fee into deposit currency")
	}

	return depositAmount >= feeDepositCurrency, nil
}

func (c *Calculator) getPrice(lot data.Lot, isBuyNow bool, customerPrice int64) (int64, error) {
	if lot.Details.AuctionPrices == nil {
		return 0, errors.New("expected lot to be auction")
	}

	if isBuyNow {
		if lot.Details.AuctionPrices.BuyNow <= 0 {
			return 0, errors.New("expected lot to have buy now")
		}

		return maxInt64(lot.Details.AuctionPrices.BuyNow, customerPrice), nil
	}

	return maxInt64(lot.HighestBid.Int64, customerPrice), nil
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}

//IsLimitSufficientForLot - returns true if limit is sufficient to place a bid or buy now request
func (c *Calculator) IsLimitSufficientForLot(lot data.Lot, limit int64, limitCurrency string, isBuyNow bool, customerPrice int64) (bool, error) {
	if lot.Details.AuctionPrices == nil {
		return false, errors.New("expected lot to be auction")
	}

	price, err := c.getPrice(lot, isBuyNow, customerPrice)
	if err != nil {
		return false, errors.Wrap(err, "failed to get price")
	}

	priceLimitCurrency, _, err := c.converter.ConvertCents(lot.Currency, limitCurrency, price)
	if err != nil {
		return false, errors.Wrap(err, "failed to convert lot price to limit currency")
	}

	return priceLimitCurrency <= limit, nil
}
