package power

import (
	"gitlab.com/eAuction/buying-power-svc/internal/types"
	"gitlab.com/eAuction/events/go/kafka"
	"math"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/exrates-svc/pkg/exrates"
	"gitlab.com/eAuction/platformer-svc/resources"

	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
)

type dynamicDepositsAndFee struct {
	depositFraction    float64
	converter          *exrates.Converter
	deposits           config.Deposits
	depositIDGenerator data.DepositIDGenerator
}

// BuyingPower returns buying power for given deposit amount.
func (d *dynamicDepositsAndFee) BuyingPower(_ string, depositAmount int64, depositCurrency, lotCurrency string) (int64, error) {
	limit := float64(depositAmount) / d.depositFraction
	if math.IsNaN(limit) || math.IsInf(limit, 0) {
		panic(errors.New("limit is not a number"))
	}

	limitDepositCurrency := int64(math.Floor(limit))
	limitLotCurrency, _, err := d.converter.ConvertCents(depositCurrency, lotCurrency, limitDepositCurrency)
	if err != nil {
		return 0, errors.Wrap(err, "failed to convert limit into lot currency", logan.F{
			"lot_currency":           lotCurrency,
			"limit_deposit_currency": limitDepositCurrency,
		})
	}

	return limitLotCurrency, nil
}

// GetFeeChangedFromDeposit returns shelf fee (deposit amount) for given lot if lot is auction.
func (d *dynamicDepositsAndFee) GetFeeChangedFromDeposit(depositor string, lot data.Lot, customerPrice int64, customerPriceCurrency string) (*resources.Fee, error) {
	if lot.Details.AuctionPrices == nil || lot.Details.AuctionPrices.BuyNow <= 0 {
		return nil, nil
	}
	depositCurrency := d.deposits.Currency
	lotPriceDepositCurrency, _, err := d.converter.ConvertCents(lot.Currency, depositCurrency, lot.Details.AuctionPrices.BuyNow)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert lot price")
	}

	customerPriceDepositCurrency, _, err := d.converter.ConvertCents(customerPriceCurrency, depositCurrency, customerPrice)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert customer price")
	}

	if lotPriceDepositCurrency < customerPriceDepositCurrency {
		lotPriceDepositCurrency = customerPriceDepositCurrency
	}

	reverseCalculatedDeposit := int64(math.Ceil(float64(lotPriceDepositCurrency) * d.depositFraction))
	for { // adjust deposit so it is guaranteed for its buying power to match or be slightly greater than lot price (concerning floating point errors)
		depositPower, err := d.BuyingPower(depositor, reverseCalculatedDeposit, depositCurrency, depositCurrency)
		if err != nil {
			return nil, errors.Wrap(err, "failed to calculate deposit buying power")
		}
		if depositPower >= lotPriceDepositCurrency {
			break
		}
		reverseCalculatedDeposit += 1
	}
	return &resources.Fee{
		Currency: depositCurrency,
		Fixed:    float64(reverseCalculatedDeposit) / 100,
	}, nil
}

// GetDepositLimits returns a limit for given lot or nothing if the lot is nil or its not auction.
func (d *dynamicDepositsAndFee) GetDepositLimits(depositor string, lot *data.Lot, customerPrice int64, customerPriceCurrency string) ([]config.Limit, error) {
	if lot == nil {
		return nil, errors.Wrap(ErrNotAvailable, "non lot based deposit limits are not available for dynamic deposit options")
	}
	fee, err := d.GetFeeChangedFromDeposit("", *lot, customerPrice, customerPriceCurrency)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get shelf fee")
	}
	if fee == nil {
		return nil, nil
	}
	deposit := int64(fee.Fixed * 100)
	power, err := d.BuyingPower(depositor, deposit, fee.Currency, d.deposits.Currency)
	if err != nil {
		return nil, err
	}
	return []config.Limit{{
		Deposit:         deposit,
		Limit:           power,
		IsBuyNowAllowed: true,
	}}, nil
}

//GetDepositsOptions - returns options for deposit state, deposits to attach, deposit to create
func (d *dynamicDepositsAndFee) GetDepositsOptions(depositor string, balance int64, attachedDeposits, availableDeposits []data.Deposit, limit config.Limit) (types.DepositOptionState, []data.Deposit, *data.Deposit) {
	var depositsToAttach []data.Deposit
	depositsToAttach = append(depositsToAttach, attachedDeposits...)
	targetDeposit := limit.Deposit - MustTotalDepositAmount(depositsToAttach, d.deposits.Currency)
	if targetDeposit == 0 {
		return types.DepositOptionStateTaken, depositsToAttach, nil
	}

	// it's now allowed to lower attached deposit
	if targetDeposit < 0 {
		return types.DepositOptionStateNotAvailable, nil, nil
	}

	newBalance := balance + limit.Deposit
	if newBalance < balance {
		panic(errors.New("overflow"))
	}

	maxAmount := d.deposits.MaxAmount
	if newBalance > maxAmount {
		return types.DepositOptionStateNotAvailable, nil, nil
	}

	return types.DepositOptionStateAvailable, nil, &data.Deposit{
		ID:        d.depositIDGenerator.Generate(),
		State:     int32(kafka.DepositV2_STATE_PENDING),
		Amount:    limit.Deposit,
		Currency:  d.deposits.Currency,
		Depositor: depositor,
	}
}
