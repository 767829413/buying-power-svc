package deposits

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/events/go/kafka"
	"time"
)

//Return - marks deposit as returned to card
func Return(storage data.Storage, deposit data.Deposit) error {
	if deposit.State == int32(kafka.DepositV2_STATE_RETURNED) || deposit.State == int32(kafka.DepositV2_STATE_REVERSED) {
		return nil
	}

	err := storage.Deposits().SetState(deposit.ID, int32(kafka.DepositV2_STATE_RETURNED))
	if err != nil {
		return errors.Wrap(err, "failed to set state returned")
	}

	err = storage.Movements().Insert(data.Movement{
		Identity: deposit.Depositor,
		Amount:   deposit.Amount,
		Currency: deposit.Currency,
		Action:   kafka.Movement_ACTION_WITHDRAWN,
		CreatedAt: time.Now().UTC(),
	})
	return errors.Wrap(err, "failed to create withdrawn movement")
}
