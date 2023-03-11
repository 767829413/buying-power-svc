package deposits

import (
	"database/sql"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/events/go/kafka"
	"time"
)

//Lock - locks deposit
func Lock(storage data.Storage, deposit data.Deposit) error {
	err := storage.Deposits().SetState(deposit.ID, int32(kafka.DepositV2_STATE_LOCKED))
	if err != nil {
		return errors.Wrap(err, "failed to mark deposit as locked")
	}

	if deposit.State == int32(kafka.DepositV2_STATE_LOCKED) {
		return nil
	}

	err = storage.Movements().Insert(data.Movement{
		Identity: deposit.Depositor,
		Amount:   deposit.Amount,
		Currency: deposit.Currency,
		Action:   kafka.Movement_ACTION_LOCKED,
		CreatedAt: time.Now(),
		Lot: sql.NullString{
			String: *deposit.LotID,
			Valid:  true,
		},
	})
	if err != nil {
		return errors.Wrap(err, "failed to create locked movement")
	}

	return nil
}

//Unlock - returns deposit to balance
func Unlock(storage data.Storage, deposit data.Deposit) error {
	err := storage.Deposits().ReturnDepositToBalance(deposit.ID)
	if err != nil {
		return errors.Wrap(err, "failed to return deposit to balance")
	}

	if deposit.State != int32(kafka.DepositV2_STATE_LOCKED) {
		return nil
	}

	err = storage.Movements().Insert(data.Movement{
		Identity: deposit.Depositor,
		Amount:   deposit.Amount,
		Currency: deposit.Currency,
		Action:   kafka.Movement_ACTION_UNLOCKED,
		CreatedAt: time.Now(),
		Lot: sql.NullString{
			String: *deposit.LotID,
			Valid:  true,
		},
	})
	if err != nil {
		return errors.Wrap(err, "failed to create unlocked movement")
	}

	return nil
}
