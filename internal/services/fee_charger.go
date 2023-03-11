package services

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/power"
	"gitlab.com/eAuction/events/go/kafka"
	"gitlab.com/eAuction/exrates-svc/pkg/exrates"
	"gitlab.com/eAuction/go/amount"
	"gitlab.com/eAuction/payment-service/pkg/payments"
)

//RunFeeCharger - charges fees from winners
func RunFeeCharger(ctx context.Context, cfg config.Config) {
	r := &feeCharger{
		log:           cfg.Log().WithField("service", "fee_charger"),
		cfg:           cfg,
		payments:      cfg.PaymentService(),
		depositStates: []int32{int32(kafka.DepositV2_STATE_LOCKED), int32(kafka.DepositV2_STATE_PAID)},
		powerCalc:     power.NewCalculator(cfg),
		converter:     cfg.ExRatesConverter(),
	}

	running.WithBackOff(ctx, r.log, "fee_charger", r.runOnce, time.Minute, time.Second*10, time.Minute)
}

type feeCharger struct {
	log           *logan.Entry
	cfg           config.Config
	payments      payments.PaymentService
	depositStates []int32
	powerCalc     *power.Calculator
	converter     *exrates.Converter
}

func (c *feeCharger) runOnce(ctx context.Context) error {
	storage := c.cfg.Storage()
	participants, err := storage.Participants().Select(data.ParticipantsSelector{
		State:         []kafka.Participant_State{kafka.Participant_STATE_BUY_NOW_WINNER, kafka.Participant_STATE_WINNER, kafka.Participant_STATE_DEAL_CANCELED},
		DepositStates: c.depositStates,
	})

	for _, participant := range participants {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		txStorage := storage.Clone()
		err = txStorage.Transaction(func() error {
			return c.chargeFee(ctx, txStorage, participant)
		})
		if err != nil {
			c.log.WithError(err).WithFields(logan.F{
				"lot_id":  participant.LotID,
				"address": participant.AccountAddress,
			}).Error("failed to return deposit to balance")
			continue
		}

	}

	return nil
}

func (c *feeCharger) getAmountToBePaid(storage data.Storage, winnerID string, winner data.Winner, deposits []data.Deposit) (int64, string, error) {
	// if user refused to pay invoice for the lot, we should change full deposit
	// if we are using deposit fraction calculator for deposits, we should just charge full amount as ShelfAP includes transportation fees, etc. into
	// invoice and want's to charge some percentage of this amount.
	if winner.State == int16(kafka.Winner_STATE_DEAL_CANCELED) || c.cfg.Deposits().DepositFraction != nil {
		am, err := power.GetTotalDepositAmount(deposits, c.cfg.Deposits().Currency)
		return am, c.cfg.Deposits().Currency, err
	}

	lot, err := storage.Lots().ByID(winner.LotID)
	if err != nil {
		return 0, "", errors.Wrap(err, "failed to get ")
	}

	if lot == nil {
		return 0, "", errors.New("expected lot to exist")
	}

	fee, err := c.powerCalc.GetFeeChangedFromDeposit(winnerID, *lot, 0, lot.Currency)
	if err != nil {
		return 0, "", errors.Wrap(err, "failed to get shelf fee")
	}
	if fee == nil {
		return 0, "", errors.From(errors.New("expected fee to be set"), logan.F{
			"winner": winnerID,
		})
	}

	amountToCharge, _, err := c.converter.ConvertCents(fee.Currency, c.cfg.Deposits().Currency, int64(fee.Fixed*100))
	if err != nil {
		return 0, "", errors.Wrap(err, "failed to convert fee into deposit currency")
	}

	availableAmount, err := power.GetTotalDepositAmount(deposits, c.cfg.Deposits().Currency)
	if err != nil {
		return 0, "", errors.Wrap(err, "failed to get total deposit amount")
	}

	if amountToCharge > availableAmount {
		return 0, "", errors.From(errors.New("amount to charge exceeds available amount"), logan.F{
			"amount_to_charge": amountToCharge,
			"available_amount": availableAmount,
		})
	}

	return amountToCharge, fee.Currency, nil
}

func (c *feeCharger) chargeFee(ctx context.Context, storage data.Storage, participant data.Participant) error {
	winner, err := storage.Winners().GetByLotID(participant.LotID)
	if err != nil {
		return errors.Wrap(err, "failed to get winner by lot id")
	}

	if winner == nil {
		return nil
	}

	if winner.State != int16(kafka.Winner_STATE_DEAL_COMPLETED) &&
		winner.State != int16(kafka.Winner_STATE_IN_TRANSIT) &&
		winner.State != int16(kafka.Winner_STATE_DEAL_CANCELED) {
		return nil
	}

	deposits, err := storage.Deposits().Select(data.DepositsSelector{
		Lots:       []string{participant.LotID},
		Depositors: []string{participant.AccountAddress},
		ForUpdate:  true,
		States:     c.depositStates,
	})

	amountToBePaid, currency, err := c.getAmountToBePaid(storage, participant.AccountAddress, *winner, deposits)
	if err != nil {
		return errors.Wrap(err, "failed to get amount to be paid")
	}

	action := kafka.Movement_ACTION_FEE_CHARGED
	if winner.State == int16(kafka.Winner_STATE_DEAL_CANCELED) {
		action = kafka.Movement_ACTION_FINE_CHARGED
	}

	err = storage.Movements().Insert(data.Movement{
		Identity:  participant.AccountAddress,
		Amount:    amountToBePaid,
		Currency:  currency,
		Action:    action,
		CreatedAt: time.Now().UTC(),
		Lot: sql.NullString{
			String: participant.LotID,
			Valid:  true,
		},
	})
	if err != nil {
		return errors.Wrap(err, "failed to insert movement")
	}

	err = c.createWithdrawMovement(storage, deposits, currency, amountToBePaid, participant)
	if err != nil {
		return errors.Wrap(err, "failed to create withdraw movement")
	}

	isArtificial := false
	for _, d := range deposits {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		isArtificial = d.IsArtificial || isArtificial

		var charged int64
		charged, err = c.chargeDeposit(d, amountToBePaid)
		if err != nil {
			return errors.Wrap(err, "failed to charge deposit", logan.F{
				"deposit_id": d.ID,
			})
		}

		amountToBePaid -= charged
	}

	if amountToBePaid > 0 && !isArtificial {
		c.log.WithField("amount_to_be_paid", amountToBePaid).Error("expected amount to be paid to <= 0 after fee charge")
	}

	for _, d := range deposits {
		err = storage.Deposits().SetState(d.ID, int32(kafka.DepositV2_STATE_FEE_CHARGED))
		if err != nil {
			return errors.Wrap(err, "failed to mark deposit as fee charged")
		}
	}

	return nil
}

func (c *feeCharger) createWithdrawMovement(storage data.Storage, deposits []data.Deposit, currency string,
	amountToBePaid int64, participant data.Participant) error {

	returnedAmount, err := power.GetTotalDepositAmount(deposits, c.cfg.Deposits().Currency)
	if err != nil {
		return errors.Wrap(err, "failed to calculate total deposit amount")
	}

	returnedAmount, _, err = c.cfg.ExRatesConverter().ConvertCents(c.cfg.Deposits().Currency, currency, returnedAmount)
	if err != nil {
		return errors.Wrap(err, "failed to convert total deposit amount")
	}

	returnedAmount -= amountToBePaid
	if returnedAmount <= 0 {
		return nil
	}

	err = storage.Movements().Insert(data.Movement{
		Identity:  participant.AccountAddress,
		Amount:    returnedAmount,
		Currency:  currency,
		Action:    kafka.Movement_ACTION_WITHDRAWN,
		CreatedAt: time.Now().UTC(),
		Lot: sql.NullString{
			String: participant.LotID,
			Valid:  true,
		},
	})
	if err != nil {
		return errors.Wrap(err, "failed to insert withdrawn movement")
	}

	return nil
}

func (c *feeCharger) chargeDeposit(deposit data.Deposit, amountToCharge int64) (int64, error) {
	if deposit.IsArtificial {
		return 0, nil
	}

	invoiceID, purpose := deposit.ID, payments.TxPurposeDeposit
	statues := strings.Join([]string{payments.TxStatusConfirmed, payments.TxStatusOk}, ",")
	transactions, err := c.payments.GetTransactions(payments.GetTransactionsParams{
		InvoiceID: &invoiceID,
		Status:    &statues,
		Purpose:   &purpose,
	})

	if err != nil {
		return 0, errors.Wrap(err, "failed to get tx by invoice id", logan.F{
			"invoice_id": deposit.ID,
		})
	}
	if len(transactions.Data) > 1 {
		return 0, fmt.Errorf("unexpected number of payments for deposit %s", deposit.ID)
	}
	if len(transactions.Data) == 0 {
		return 0, errors.From(errors.New("expected tx to be present"), logan.F{
			"invoice_id": deposit.ID,
		})
		return 0, nil
	}

	tx := transactions.Data[0]
	if amountToCharge > int64(tx.Attributes.Amount) {
		amountToCharge = int64(tx.Attributes.Amount)
	}

	if amountToCharge <= 0 {
		_, err = c.payments.RefundTransaction(tx.ID)
		if err != nil {
			return 0, errors.Wrap(err, "failed to reverse tx", logan.F{
				"tx_id": tx.ID,
			})
		}

		return 0, nil
	}

	amount := amount.Fiat(amountToCharge)
	_, err = c.payments.ConfirmTransaction(tx.ID, &amount)
	if err != nil {
		return 0, errors.Wrap(err, "failed to charge amount", logan.F{
			"tx_id": tx.ID,
		})
	}

	return amountToCharge, nil
}
