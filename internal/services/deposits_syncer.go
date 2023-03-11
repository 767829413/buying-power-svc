package services

import (
	"context"
	"strconv"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/events/go/kafka"
)

type depositsSyncer struct {
	cfg    config.Config
	log    *logan.Entry
	writer kafka.DepositEventWriter
}

//RunDepositsSyncer - sends deposits to kafka
func RunDepositsSyncer(ctx context.Context, cfg config.Config) {
	s := depositsSyncer{
		cfg:    cfg,
		log:    cfg.Log().WithField("service", "deposits_syncer"),
		writer: cfg.DepositsWriter(),
	}
	running.WithBackOff(ctx, s.log, "deposits_syncer", s.runOnce, time.Second, 5*time.Second, time.Minute)
}

func (s *depositsSyncer) syncDeposits(ctx context.Context, storage data.Storage, deposits []data.Deposit) error {
	var events []kafka.DepositEventMessage
	var ids []string
	for _, deposit := range deposits {
		var lotID int64
		if deposit.LotID != nil {
			var err error
			lotID, err = strconv.ParseInt(*deposit.LotID, 10, 64)
			if err != nil {
				return errors.Wrap(err, "failed to parse lot ID", logan.F{
					"lot_id": *deposit.LotID,
				})
			}
		}
		events = append(events, kafka.DepositV2{
			Id:           deposit.ID,
			Amount:       deposit.Amount,
			Currency:     deposit.Currency,
			Depositor:    deposit.Depositor,
			LotId:        lotID,
			State:        kafka.DepositV2_State(deposit.State),
			CreatedAt:    deposit.CreatedAt.Unix(),
			UpdatedAt:    deposit.UpdatedAt.Unix(),
			IsArtificial: deposit.IsArtificial,
		}.Event().Message())

		ids = append(ids, deposit.ID)
	}

	err := s.writer.WriteMessages(ctx, events...)
	if err != nil {
		return errors.Wrap(err, "failed to write messages to kafka")
	}

	err = storage.Deposits().MarkSynced(ids)
	if err != nil {
		return errors.Wrap(err, "failed to mark deposits synced")
	}

	return nil
}

func (s *depositsSyncer) syncMovements(ctx context.Context, storage data.Storage, movements []data.Movement) error {
	var events []kafka.DepositEventMessage
	var ids []int64
	for _, movement := range movements {
		events = append(events, kafka.Movement{
			Id:                   strconv.FormatInt(movement.ID, 10),
			Amount:               movement.Amount,
			Identity:             movement.Identity,
			Currency:             movement.Currency,
			Action:               movement.Action,
			CreatedAt:            movement.CreatedAt.Unix(),
			LotId:                movement.Lot.String,
		}.Event().Message())

		ids = append(ids, movement.ID)
	}

	err := s.writer.WriteMessages(ctx, events...)
	if err != nil {
		return errors.Wrap(err, "failed to write messages to kafka")
	}

	err = storage.Movements().Delete(ids)
	if err != nil {
		return errors.Wrap(err, "failed to mark movements synced")
	}

	return nil
}

func (s *depositsSyncer) syncWithdrawals(ctx context.Context, storage data.Storage, withdrawals []data.Withdrawal) error {
	var events []kafka.DepositEventMessage
	var ids []string
	for _, withdrawal := range withdrawals {
		deposit, err := storage.Deposits().Get(withdrawal.DepositID, false)
		if err != nil {
			return errors.Wrap(err, "failed to get deposit for withdrawal")
		}

		events = append(events, kafka.Withdrawal{
			Id:        withdrawal.ID,
			DepositId: withdrawal.DepositID,
			State:     kafka.Withdrawal_State(withdrawal.State),
			CreatedAt: withdrawal.CreatedAt.Unix(),
			UpdatedAt: withdrawal.UpdatedAt.Unix(),
			Depositor: deposit.Depositor,
			Amount:    deposit.Amount,
			Currency:  deposit.Currency,
		}.Event().Message())

		ids = append(ids, withdrawal.ID)
	}

	err := s.writer.WriteMessages(ctx, events...)
	if err != nil {
		return errors.Wrap(err, "failed to write withdrawal messages to kafka")
	}

	err = storage.Withdrawals().MarkSynced(ids)
	if err != nil {
		return errors.Wrap(err, "failed to mark withdrawals synced")
	}

	return nil
}

func (s *depositsSyncer) runOnce(ctx context.Context) error {
	storage := s.cfg.Storage()
	err := storage.Transaction(func() error {
		// to ensure correct order of the deposits & withdrawals, we need to first read withdrawals
		no := false
		withdrawals, err := storage.Withdrawals().Select(data.WithdrawalsSelector{
			ForUpdate: true,
			IsSynced:  &no,
		})
		if err != nil {
			return errors.Wrap(err, "failed to select not synced withdrawals")
		}

		movements, err := storage.Movements().SelectNotSynced(100)
		if err != nil {
			return errors.Wrap(err, "failed to select movements")
		}

		deposits, err := storage.Deposits().Select(data.DepositsSelector{
			IsSynced:  &no,
			ForUpdate: true,
		})
		if err != nil {
			return errors.Wrap(err, "failed to select not synced deposits")
		}

		err = s.syncDeposits(ctx, storage, deposits)
		if err != nil {
			return errors.Wrap(err, "failed to sync deposits")
		}

		err = s.syncWithdrawals(ctx, storage, withdrawals)
		if err != nil {
			return errors.Wrap(err, "failed to sync withdrawals")
		}

		err = s.syncMovements(ctx, storage, movements)
		if err != nil {
			return errors.Wrap(err, "failed to sync movements")
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "failed to process deposits batch")
	}

	return nil
}
