package power

import (
	"sort"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/types"
	"gitlab.com/eAuction/events/go/kafka"
	"gitlab.com/eAuction/exrates-svc/pkg/exrates"
	"gitlab.com/eAuction/platformer-svc/pkg/platformer"
	"gitlab.com/eAuction/platformer-svc/resources"
)

type fixedDepositsAndFee struct {
	deposits           config.Deposits
	converter          *exrates.Converter
	platforms          platformer.Platformer
	storage            data.Storage
	depositIDGenerator data.DepositIDGenerator
	cfg                config.Config
}

//BuyingPower - returns max limit for specified deposit in lot currency
func (c *fixedDepositsAndFee) BuyingPower(depositorID string, depositAmount int64, depositCurrency, lotCurrency string) (int64, error) {
	depositor, err := c.getDepositor(depositorID)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get depositor")
	}

	limitDepositCurrency := c.deposits.GetLimits(depositor.Platform).FindMatchingLimit(depositAmount)

	limitLotCurrency, _, err := c.converter.ConvertCents(depositCurrency, lotCurrency, limitDepositCurrency)
	if err != nil {
		return 0, errors.Wrap(err, "failed to convert limit into lot currency", logan.F{
			"lot_currency":           lotCurrency,
			"limit_deposit_currency": limitDepositCurrency,
		})
	}

	const roundTo = 10000
	limitLotCurrency = (limitLotCurrency / roundTo) * roundTo
	return limitLotCurrency, nil
}

func (c *fixedDepositsAndFee) getDepositor(depositorID string) (data.Identity, error) {
	depositor, err := c.storage.Identities().Get(depositorID)
	if err != nil {
		return data.Identity{}, errors.Wrap(err, "failed to get depositor")
	}

	if depositor == nil {
		return data.Identity{}, errors.From(errors.New("expected depositor to exist"), logan.F{
			"depositor_id": depositorID,
		})
	}

	return *depositor, nil
}

//GetFeeChangedFromDeposit -
func (c *fixedDepositsAndFee) GetFeeChangedFromDeposit(depositorID string, lot data.Lot, _ int64, _ string) (*resources.Fee, error) {
	depositor, err := c.getDepositor(depositorID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get depositor")
	}

	platform := c.platforms.MustPlatform(depositor.Platform)
	key := platform.Data.Relationships.IndirectFees
	if key == nil {
		return nil, nil
	}

	indirectFees := platform.Included.MustIndirectFees(*key.Data)
	for _, fee := range indirectFees.Attributes.Data {
		if fee.Platform == lot.Details.PlatformId {
			return fee.ShelfFee, nil
		}
	}

	return nil, nil
}

func (c *fixedDepositsAndFee) getDepositLimits(depositorID string) ([]config.Limit, error) {
	depositor, err := c.getDepositor(depositorID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get depositor")
	}

	limits := make([]config.Limit, 0)
	for _, limit := range c.deposits.GetLimits(depositor.Platform).Limits {
		if limit.IsRechargerOnly && !c.cfg.IsAutoRechargeable(depositor.Type) {
			continue
		}
		limits = append(limits, limit)
	}

	return limits, nil
}

// GetDepositLimits returns limits for given lot or all limits if the lot is nil.
func (c *fixedDepositsAndFee) GetDepositLimits(depositorID string, _ *data.Lot, _ int64, _ string) ([]config.Limit, error) {
	return c.getDepositLimits(depositorID)
}

//GetDepositsOptions - returns options for deposit state, deposits to attach, deposit to create
func (c *fixedDepositsAndFee) GetDepositsOptions(depositor string, balance int64, attachedDeposits, availableDeposits []data.Deposit, limit config.Limit) (types.DepositOptionState, []data.Deposit, *data.Deposit) {
	var depositsToAttach []data.Deposit
	depositsToAttach = append(depositsToAttach, attachedDeposits...)
	targetDeposit := limit.Deposit - MustTotalDepositAmount(depositsToAttach, c.deposits.Currency)
	if targetDeposit == 0 {
		return types.DepositOptionStateTaken, depositsToAttach, nil
	}

	// it's now allowed to lower attached deposit
	if targetDeposit < 0 {
		return types.DepositOptionStateNotAvailable, nil, nil
	}

	state, isExactMatch, depositsToAttach, newDeposit := c.getDepositsToAttach(depositor, balance, targetDeposit, availableDeposits)
	if state != types.DepositOptionStateAvailable {
		return types.DepositOptionStateNotAvailable, nil, nil
	}

	if isExactMatch {
		return types.DepositOptionStateAvailable, append(attachedDeposits, depositsToAttach...), newDeposit
	}

	// match is not exact try to improve result, by creating new deposit
	state, isExactMatch, _, newDeposit = c.getDepositsToAttach(depositor, balance, targetDeposit, nil)
	// we've tried but failed, fall back to prev result
	if state != types.DepositOptionStateAvailable || !isExactMatch {
		return types.DepositOptionStateAvailable, append(attachedDeposits, depositsToAttach...), newDeposit
	}

	return types.DepositOptionStateAvailable, attachedDeposits, newDeposit
}

func (c *fixedDepositsAndFee) getDepositsToAttach(depositor string, balance, targetDeposit int64, availableDeposits []data.Deposit) (types.DepositOptionState, bool, []data.Deposit, *data.Deposit) {
	depositsToAttach := getDepositsForTarget(targetDeposit, availableDeposits)
	targetDeposit = targetDeposit - MustTotalDepositAmount(depositsToAttach, c.deposits.Currency)
	if targetDeposit <= 0 {
		return types.DepositOptionStateAvailable, targetDeposit == 0, depositsToAttach, nil
	}

	newDeposit, isExactMatch := c.findNewDepositForTarget(depositor, targetDeposit, balance)
	if newDeposit == nil {
		return types.DepositOptionStateNotAvailable, false, nil, nil
	}

	return types.DepositOptionStateAvailable, isExactMatch, depositsToAttach, newDeposit
}

func (c *fixedDepositsAndFee) findNewDepositForTarget(depositorID string, target, currentBalance int64) (*data.Deposit, bool) {
	limits, err := c.getDepositLimits(depositorID)
	if err != nil {
		panic(errors.Wrap(err, "failed to get deposit limits", logan.F{
			"depositorID": depositorID,
		}))
	}

	for _, limit := range limits {
		if limit.Deposit < target {
			continue
		}

		newBalance := currentBalance + limit.Deposit
		if newBalance < currentBalance {
			panic(errors.New("overflow"))
		}

		maxAmount := c.deposits.MaxAmount
		if newBalance > maxAmount {
			continue
		}

		return &data.Deposit{
			ID:        c.depositIDGenerator.Generate(),
			State:     int32(kafka.DepositV2_STATE_PENDING),
			Amount:    limit.Deposit,
			Currency:  c.deposits.Currency,
			Depositor: depositorID,
		}, limit.Deposit == target
	}

	return nil, false
}

func getDepositsForTarget(target int64, deposits []data.Deposit) []data.Deposit {
	sort.Sort(data.DepositsByCreatedAtAsc(deposits))
	exactMatch := tryFindExactAmount(target, deposits)
	if exactMatch != nil {
		return []data.Deposit{*exactMatch}
	}

	var result []data.Deposit
	for _, deposit := range deposits {
		result = append(result, deposit)
		target -= deposit.Amount
		if target <= 0 {
			return result
		}
	}

	return result
}

func tryFindExactAmount(target int64, deposits []data.Deposit) *data.Deposit {
	for i := range deposits {
		if deposits[i].Amount == target {
			return &deposits[i]
		}
	}

	return nil
}
