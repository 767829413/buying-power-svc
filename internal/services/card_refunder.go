package services

import (
	"context"
	"time"

	"gitlab.com/eAuction/buying-power-svc/internal/services/deposits"
	"gitlab.com/eAuction/payment-service/pkg/payments"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/events/go/kafka"
)

// RunCardRefund - returns expired deposits back to the card
func RunCardRefund(ctx context.Context, cfg config.Config) {
	r := &cardRefund{
		log:      cfg.Log().WithField("service", "card_refund"),
		payments: cfg.PaymentService(),
		deposits: cfg.Deposits(),
		cfg:      cfg,
	}
	r.log.Info("starting card refunder")

	running.WithBackOff(ctx, r.log, "card_refund", r.runOnce, time.Minute, time.Second*10, time.Minute)
}

type cardRefund struct {
	log      *logan.Entry
	payments payments.PaymentService
	deposits config.Deposits
	cfg      config.Config
}

func (c *cardRefund) getDepositsToReturn(storage data.Storage) ([]data.Deposit, error) {
	// // it's not possible to attach deposit to any lot if deposit has been created before now() - expires in + lot processing duration
	// // but Cardeal does not attach deposits to lots for now, so we have to remove lot processing duration for them
	maxCreatedAt := time.Now().Add(-c.deposits.ExpiresIn) //.Add(c.deposits.LotProcessingDuration)
	no := false
	depositStates := []int32{int32(kafka.DepositV2_STATE_PAID), int32(kafka.DepositV2_STATE_REQUESTED_WITHDRAWAL)}
	expired, err := storage.Deposits().Select(data.DepositsSelector{
		States:       depositStates,
		MaxCreatedAt: &maxCreatedAt,
		IsArtificial: &no,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to select expired deposits")
	}

	yes := true
	maxCreatedAt = maxCreatedAt.Add(c.deposits.LotProcessingDuration)
	artificial, err := storage.Deposits().Select(data.DepositsSelector{
		States:       depositStates,
		MaxCreatedAt: &maxCreatedAt,
		IsArtificial: &yes,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to select expired artificial deposits")
	}

	returningDeposits, err := storage.Deposits().Select(data.DepositsSelector{
		States: []int32{int32(kafka.DepositV2_STATE_RETURNING)},
	})
	c.log.Infof("returning %d deposits", len(returningDeposits))

	if err != nil {
		return nil, errors.Wrap(err, "failed to select returning deposits")
	}

	return append(append(expired, returningDeposits...), artificial...), nil
}

func (c *cardRefund) runOnce(ctx context.Context) error {
	storage := c.cfg.Storage()
	returning, err := c.getDepositsToReturn(storage)
	if err != nil {
		return errors.Wrap(err, "failed to select deposits that need to returned")
	}

	for _, deposit := range returning {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		txStorage := storage.Clone()
		err = running.RunSafely(ctx, "deposit-returner", func(ctx context.Context) error {
			return txStorage.Transaction(func() error {
				return c.returnDeposit(txStorage, deposit.ID)
			})
		})
		if err != nil {
			c.log.WithError(err).WithField("deposit.id", deposit.ID).Error("failed to return deposit")
			continue
		}

	}

	return nil
}

func (c *cardRefund) isAllowedToReturn(deposit data.Deposit) bool {
	allowedStates := []kafka.DepositV2_State{kafka.DepositV2_STATE_PAID, kafka.DepositV2_STATE_RETURNING, kafka.DepositV2_STATE_REQUESTED_WITHDRAWAL}
	for _, state := range allowedStates {
		if int32(state) == deposit.State {
			return true
		}
	}

	return false
}

func (c *cardRefund) returnDeposit(storage data.Storage, depositID string) error {
	deposit, err := storage.Deposits().GetForUpdate(depositID)
	if err != nil {
		return errors.Wrap(err, "failed to select deposit by id")
	}

	if deposit == nil || !c.isAllowedToReturn(*deposit) {
		return errors.New("expected deposit to exist in allowed state")
	}

	err = c.cleanupWithdrawals(storage, depositID)
	if err != nil {
		return errors.Wrap(err, "failed to cleanup withdrawals")
	}

	err = c.revertTransaction(*deposit)
	if err != nil {
		return errors.Wrap(err, "failed to revert transaction")
	}

	return deposits.Return(storage, *deposit)
}

func (c *cardRefund) cleanupWithdrawals(storage data.Storage, depositID string) error {
	withdrawals, err := storage.Withdrawals().Select(data.WithdrawalsSelector{
		States:    []int32{int32(kafka.Withdrawal_STATE_PENDING)},
		ForUpdate: true,
		Deposits:  []string{depositID},
	})
	if err != nil {
		return errors.Wrap(err, "failed to select withdrawals")
	}

	for _, withdrawal := range withdrawals {
		err = storage.Withdrawals().SetState(withdrawal.ID, int32(kafka.Withdrawal_STATE_APPROVED))
		if err != nil {
			return errors.Wrap(err, "failed to set withdrawal state approved")
		}
	}

	return nil
}

func (c *cardRefund) revertTransaction(deposit data.Deposit) error {
	if deposit.IsArtificial {
		return nil
	}

	err := c.payments.RefundTransactionByInvoiceID(payments.TxPurposeDeposit, deposit.ID)
	if err == payments.ErrNotFound {
		c.log.WithField("deposit_id", deposit.ID).Error("expected transaction for non artificial deposit to exist")
		return nil
	}
	if err != nil {
		return errors.Wrap(err, "failed to reverse tx for expired deposit")
	}

	return nil
}
