package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/responses"
	"gitlab.com/eAuction/buying-power-svc/internal/services/deposits"
	"gitlab.com/eAuction/events/go/kafka"

	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/bouncer/allow"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
)

type depositPatchRequest struct {
	DepositID string                 `json:"-"`
	Data      resources.DepositPatch `json:"data"`
}

func newPatchDepositRequest(r *http.Request) (*depositPatchRequest, error) {
	result := depositPatchRequest{
		DepositID: chi.URLParam(r, "deposit"),
		Data:      resources.DepositPatch{},
	}
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal body")
	}

	return &result, validation.Errors{
		"/data/attributes/state": validation.Validate(result.Data.Attributes.State, validation.In(kafka.DepositV2_STATE_PAID.String(), kafka.DepositV2_STATE_RETURNING.String())),
	}.Filter()
}

func PatchDeposits(w http.ResponseWriter, r *http.Request) {
	request, err := newPatchDepositRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var deposit *data.Deposit
	err = Storage(r).Transaction(func() error {
		deposit, err = Storage(r).Deposits().GetForUpdate(request.DepositID)
		if err != nil {
			return errors.Wrap(err, "failed to get deposit for update")
		}

		if deposit == nil {
			return problems.NotFound()
		}

		depositor, err := Storage(r).Identities().Get(deposit.Depositor)
		if err != nil {
			return errors.Wrap(err, "failed to get depositor")
		}

		if depositor == nil {
			return errors.New("expected depositor to exist")
		}

		_, err = Check(r, allow.Admin{
			Platform: depositor.Platform,
		})
		if err != nil {
			return problems.NotAllowed(err)
		}

		if request.Data.Attributes.State != nil {
			switch *request.Data.Attributes.State {
			case kafka.DepositV2_STATE_RETURNING.String():
				err = handleReturn(r, *deposit)
				if err != nil {
					return errors.Wrap(err, "failed to handle deposit return")
				}
			case kafka.DepositV2_STATE_PAID.String():
				err = handleManualApprove(r, *deposit)
				if err != nil {
					return errors.Wrap(err, "failed to manually approve deposit")
				}
			default:
				return errors.From(errors.New("unexpected state in patch"), logan.F{
					"state": *request.Data.Attributes.State,
				})
			}
		}

		// reread deposit to get it's latest state
		deposit, err = Storage(r).Deposits().GetForUpdate(request.DepositID)
		if err != nil {
			return errors.Wrap(err, "failed to reread deposit")
		}

		if deposit == nil {
			return problems.NotFound()
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

	response := responses.NewDeposit(*deposit, nil)
	json.NewEncoder(w).Encode(&resources.DepositResponse{
		Data:     response,
		Included: resources.Included{},
	})
}

func handleManualApprove(r *http.Request, deposit data.Deposit) error {
	allowedStates := map[kafka.DepositV2_State]struct{}{
		kafka.DepositV2_STATE_PENDING: {},
		kafka.DepositV2_STATE_FAILED:  {},
	}

	if _, ok := allowedStates[kafka.DepositV2_State(deposit.State)]; !ok {
		return problems.BadRequest(validation.Errors{
			"deposit": errors.New("must be in pending or failed state"),
		})[0]
	}

	err := deposits.ManuallyApprove(Storage(r), deposit)
	if err != nil {
		return errors.Wrap(err, "failed to manually approve deposit")
	}

	return nil
}

func handleReturn(r *http.Request, deposit data.Deposit) error {
	allowedStates := map[kafka.DepositV2_State]struct{}{
		kafka.DepositV2_STATE_PAID:                 {},
		kafka.DepositV2_STATE_REQUESTED_WITHDRAWAL: {},
	}
	_, ok := allowedStates[kafka.DepositV2_State(deposit.State)]
	if !ok || deposit.LotID != nil {
		return problems.BadRequest(validation.Errors{
			"deposit": errors.New("must be in paid or requested withdrawal state and must not be attached to any lots"),
		})[0]
	}

	withdrawals, err := Storage(r).Withdrawals().Select(data.WithdrawalsSelector{
		States:    []int32{int32(kafka.Withdrawal_STATE_PENDING)},
		ForUpdate: true,
		Deposits:  []string{deposit.ID},
	})

	if err != nil {
		return errors.Wrap(err, "failed to select withdrawals")
	}

	for _, withdrawal := range withdrawals {
		err = Storage(r).Withdrawals().SetState(withdrawal.ID, int32(kafka.Withdrawal_STATE_APPROVED))
		if err != nil {
			return errors.Wrap(err, "failed to set state for withdrawal")
		}
	}

	err = Storage(r).Deposits().SetState(deposit.ID, int32(kafka.DepositV2_STATE_RETURNING))
	if err != nil {
		return errors.Wrap(err, "failed to set deposit into state returning")
	}

	return nil
}
