package handlers

import (
	"encoding/json"
	"time"

	"github.com/google/jsonapi"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/responses"
	"gitlab.com/eAuction/events/go/kafka"

	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/bouncer/allow"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
)

func newCreateWithdrawal(r *http.Request) (string, error) {
	result := struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		return "", errors.Wrap(err, "failed to unmarshal body")
	}

	return result.Data.ID, nil
}

func CreateWithdrawal(w http.ResponseWriter, r *http.Request) {
	depositID, err := newCreateWithdrawal(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var withdrawal data.Withdrawal
	storage := Storage(r)
	err = storage.Transaction(func() error {
		deposit, err := storage.Deposits().GetForUpdate(depositID)
		if err != nil {
			return errors.Wrap(err, "failed to select deposit")
		}

		if deposit == nil {
			return problems.BadRequest(validation.Errors{
				"/data/id": errors.New("deposit not found"),
			})[0]
		}

		_, err = Check(r, allow.Account{
			Address: deposit.Depositor,
		})
		if err != nil {
			return problems.NotAllowed(err)
		}

		if deposit.State != int32(kafka.DepositV2_STATE_PAID) && deposit.LotID != nil {
			return problems.BadRequest(validation.Errors{
				"/data/id": errors.New("deposit must be in state paid and must not be attached to any lots"),
			})[0]
		}

		// we have lock on deposits so we are safe to do unique checks
		withdrawals, err := storage.Withdrawals().Select(data.WithdrawalsSelector{
			States:   []int32{int32(kafka.Withdrawal_STATE_PENDING)},
			Deposits: []string{deposit.ID},
		})
		if err != nil {
			return errors.Wrap(err, "failed to select withdrawals")
		}

		if len(withdrawals) > 0 {
			return problems.Conflict()
		}

		withdrawal, err = storage.Withdrawals().Create(deposit.ID)
		if err != nil {
			return errors.Wrap(err, "failed to create withdrawal")
		}

		err = storage.Movements().Insert(data.Movement{
			Identity: deposit.Depositor,
			Amount:   deposit.Amount,
			Currency: deposit.Currency,
			Action:   kafka.Movement_ACTION_LOCKED,
			CreatedAt: time.Now().UTC(),
		})
		if err != nil {
			return errors.Wrap(err, "failed to insert movement")
		}

		depositState := kafka.DepositV2_STATE_REQUESTED_WITHDRAWAL
		if Config(r).Deposits().AutoApproveWithdrawal {
			err = storage.Withdrawals().SetState(withdrawal.ID, int32(kafka.Withdrawal_STATE_APPROVED))
			if err != nil {
				return errors.Wrap(err, "failed to mark withdrawal as approved")
			}
			depositState = kafka.DepositV2_STATE_RETURNING
		}
		err = storage.Deposits().SetState(deposit.ID, int32(depositState))
		if err != nil {
			return errors.Wrap(err, "failed to set deposit to state request withdrawal")
		}

		return nil
	})
	if err != nil {
		switch cause := errors.Cause(err).(type) {
		case *jsonapi.ErrorObject:
			ape.RenderErr(w, cause)
			return
		default:
			panic(errors.Wrap(err, "unexpected error"))
		}
	}

	response := responses.NewWithdrawal(withdrawal)
	json.NewEncoder(w).Encode(&resources.WithdrawalResponse{
		Data:     response,
		Included: resources.Included{},
	})
}
