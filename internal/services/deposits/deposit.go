package deposits

import (
	"fmt"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/events/go/kafka"
	"time"
)

//ManuallyApprove - approves deposit and marks it artificial
func ManuallyApprove(storage data.Storage, deposit data.Deposit) error {
	err := storage.Deposits().ManuallyApprove(deposit.ID)
	if err != nil {
		return errors.Wrap(err, "failed to manually approve deposit")
	}

	return newDeposit(storage, deposit)
}

//Approve - marks deposit as paid
func Approve(storage data.Storage, deposit data.Deposit) error {
	err := storage.Deposits().SetState(deposit.ID, int32(kafka.DepositV2_STATE_PAID))
	if err != nil {
		return errors.Wrap(err, "failed to set paid state for deposit")
	}

	return newDeposit(storage, deposit)
}

func ReturnDepositForParticipant(storage data.Storage, participant data.Participant, deposits config.Deposits) error {
	dbDeposits, err := storage.Deposits().Select(data.DepositsSelector{
		Lots:       []string{participant.LotID},
		Depositors: []string{participant.AccountAddress},
		ForUpdate:  true,
		States:     []int32{int32(kafka.DepositV2_STATE_LOCKED), int32(kafka.DepositV2_STATE_PAID)},
	})

	if err != nil {
		return errors.Wrap(err, "failed to select deposits to return to the balance")
	}

	target := "balance"
	if deposits.ReturnDepositsToCard {
		target = "card"
	}
	for _, d := range dbDeposits {
		if deposits.ReturnDepositsToCard {
			err = storage.Deposits().ScheduleReturnToCard(d.ID)
		} else {
			err = Unlock(storage, d)
		}
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to return deposit to %s", target), logan.F{
				"deposit_id": d.ID,
				"target":     target,
			})
		}

	}

	return nil
}

func newDeposit(storage data.Storage, deposit data.Deposit) error {
	if deposit.State == int32(kafka.DepositV2_STATE_PAID) {
		return nil
	}
	err := storage.Movements().Insert(data.Movement{
		Identity: deposit.Depositor,
		Amount:   deposit.Amount,
		Currency: deposit.Currency,
		Action:   kafka.Movement_ACTION_DEPOSIT,
		CreatedAt: time.Now().UTC(),
	})
	if err != nil {
		return errors.Wrap(err, "failed to insert movement")
	}

	return nil
}
