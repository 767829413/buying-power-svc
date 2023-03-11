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
	"gitlab.com/eAuction/buying-power-svc/internal/services/power"
	"gitlab.com/eAuction/events/go/kafka"
)

type limitsSetter struct {
	cfg             config.Config
	deposits        config.Deposits
	log             *logan.Entry
	buyingPowerCalc *power.Calculator
}

//RunLimitsSetter - runner that finds attached lots and sets limits for them
func RunLimitsSetter(ctx context.Context, cfg config.Config) {
	s := limitsSetter{
		cfg:             cfg,
		log:             cfg.Log().WithField("service", "limits_setter"),
		deposits:        cfg.Deposits(),
		buyingPowerCalc: power.NewCalculator(cfg),
	}
	running.WithBackOff(ctx, s.log, "limits_setter", s.runOnce, time.Second, 5*time.Second, time.Minute)
}

func (s *limitsSetter) runOnce(ctx context.Context) error {
	storage := s.cfg.Storage()
	deposits, err := s.getParticipation(storage)
	if err != nil {
		return errors.Wrap(err, "failed to get deposits that require limits update")
	}

	for d := range deposits {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		txStorage := storage.Clone()
		err = txStorage.Transaction(func() error {
			return s.processParticipant(txStorage, d.Depositor, d.LotID)
		})
		if err != nil {
			s.log.WithError(err).WithFields(logan.F{
				"depositor": d.Depositor,
				"lot_id":    d.LotID,
			}).Error("failed to set limit for participation")
		}
	}

	return nil
}

type participation struct {
	Depositor string
	LotID     string
}

func (s *limitsSetter) getParticipation(storage data.Storage) (map[participation]struct{}, error) {
	yes := true
	deposits, err := storage.Deposits().Select(data.DepositsSelector{
		States: []int32{int32(kafka.DepositV2_STATE_PAID), int32(kafka.DepositV2_STATE_FAILED)},
		HasLot: &yes,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to load deposits that require limits update")
	}

	result := map[participation]struct{}{}
	for _, deposit := range deposits {
		result[participation{
			Depositor: deposit.Depositor,
			LotID:     *deposit.LotID,
		}] = struct{}{}
	}

	return result, nil
}

func (s *limitsSetter) cleanUpFailedAttempt(storage data.Storage, depositor string, lotID string) (bool, error) {
	dbDeposits, err := storage.Deposits().Select(data.DepositsSelector{
		States:     []int32{int32(kafka.DepositV2_STATE_FAILED)},
		Lots:       []string{lotID},
		Depositors: []string{depositor},
		ForUpdate:  true,
	})

	if err != nil {
		return false, errors.Wrap(err, "failed to load deposits")
	}

	if len(dbDeposits) == 0 {
		return false, nil
	}

	for _, d := range dbDeposits {
		err = storage.Deposits().SetLot(d.ID, nil)
		if err != nil {
			return false, errors.Wrap(err, "failed to remove lot from failed deposit")
		}
	}

	err = s.returnPaidDeposit(storage, lotID, depositor)
	if err != nil {
		return false, errors.Wrap(err, "failed to return paid deposits")
	}

	return true, nil
}

func (s *limitsSetter) processParticipant(storage data.Storage, depositor string, lotID string) error {
	isReadyToProcess, err := s.isReadyToProcess(storage, depositor, lotID)
	if err != nil {
		return errors.Wrap(err, "failed to check if ready to process")
	}

	fields := logan.F{
		"depositor": depositor,
		"lot_id":    lotID,
	}

	// there are still some pending deposits so lets wait
	if !isReadyToProcess {
		s.log.WithFields(fields).Info("is not ready to process")
		return nil
	}

	removedFailed, err := s.cleanUpFailedAttempt(storage, depositor, lotID)
	if err != nil {
		return errors.Wrap(err, "failed to perform cleanup")
	}

	// one of the deposits has failed, we've released them, so lets try from the begging
	if removedFailed {
		s.log.WithFields(fields).Info("removed failed")
		return nil
	}

	s.log.WithFields(fields).Info("setting limit")
	err = s.setLimit(storage, lotID, depositor)
	return errors.Wrap(err, "failed to set limit")
}

func (s *limitsSetter) setLimit(storage data.Storage, lotID, depositor string) error {
	dbDeposits, err := storage.Deposits().Select(data.DepositsSelector{
		States:     []int32{int32(kafka.DepositV2_STATE_PAID), int32(kafka.DepositV2_STATE_LOCKED)},
		Lots:       []string{lotID},
		Depositors: []string{depositor},
		ForUpdate:  true,
	})

	if err != nil {
		return errors.Wrap(err, "failed to select deposits")
	}

	totalDeposit, err := power.GetTotalDepositAmount(dbDeposits, s.deposits.Currency)
	if err != nil {
		return errors.Wrap(err, "failed to get total deposit")
	}

	lot, err := storage.Lots().ByID(lotID)
	if err != nil {
		return errors.Wrap(err, "failed to load lot by id")
	}

	if lot == nil {
		return errors.From(errors.New("expected lot to exist"), logan.F{
			"lot_id": lotID,
		})
	}

	isDepositSufficient, err := s.buyingPowerCalc.IsDepositSufficientToPayFee(depositor, *lot, totalDeposit, s.deposits.Currency)
	if err != nil {
		return errors.Wrap(err, "failed to check if deposit amount is sufficient to pay fee")
	}

	fields := logan.F{
		"depositor": depositor,
		"lot_id":    lotID,
	}

	if !isDepositSufficient {
		s.log.WithFields(fields).Info("deposit is not sufficient to pay fees - returning to balance")
		err = s.returnPaidDeposit(storage, lotID, depositor)
		return errors.Wrap(err, "failed to return paid deposits")
	}

	buyingPower, err := s.buyingPowerCalc.BuyingPower(depositor, totalDeposit, s.deposits.Currency, lot.Currency)
	if err != nil {
		return errors.Wrap(err, "failed to calculate buying power")
	}

	participant, err := storage.Participants().GetForUpdate(lotID, depositor)
	if err != nil {
		return errors.Wrap(err, "failed to get participant for update")
	}

	if participant == nil {
		panic(errors.New("expected participant to exist"))
	}

	if participant.State != int32(kafka.Participant_STATE_PENDING) && participant.State != int32(kafka.Participant_STATE_DEPOSIT_PENDING) {
		return errors.New("expected participant to be in pending or deposit pending state")
	}

	err = storage.Participants().SetBidLimit(participant.ID, buyingPower)
	if err != nil {
		return errors.Wrap(err, "failed to set bid limit")
	}

	for _, deposit := range dbDeposits {
		err = deposits.Lock(storage, deposit)
		if err != nil {
			return errors.Wrap(err, "failed to lock deposit")
		}
	}

	return nil
}

func (s *limitsSetter) isReadyToProcess(storage data.Storage, depositor, lotID string) (bool, error) {
	// ensure that there is no pending deposits for this lot
	dbDeposits, err := storage.Deposits().Select(data.DepositsSelector{
		States:     []int32{int32(kafka.DepositV2_STATE_PENDING)},
		Lots:       []string{lotID},
		Depositors: []string{depositor},
		ForUpdate:  true,
	})

	if err != nil {
		return false, errors.Wrap(err, "failed to check if pending deposits attached")
	}

	return len(dbDeposits) == 0, nil
}

func (s *limitsSetter) returnPaidDeposit(storage data.Storage, lotID, depositor string) error {
	dbDeposits, err := storage.Deposits().Select(data.DepositsSelector{
		States:     []int32{int32(kafka.DepositV2_STATE_PAID)},
		Lots:       []string{lotID},
		Depositors: []string{depositor},
		ForUpdate:  true,
	})

	if err != nil {
		return errors.Wrap(err, "failed to get paid deposits to be returned")
	}

	for _, d := range dbDeposits {
		err = deposits.Unlock(storage, d)
		if err != nil {
			return errors.Wrap(err, "failed to return deposit", logan.F{
				"deposit_id": d.ID,
			})
		}
	}

	return nil
}
