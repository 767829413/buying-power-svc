package handlers

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/bouncer/allow"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/responses"
	"gitlab.com/eAuction/events/go/kafka"
)

func newCreateDepositRequest(r *http.Request) (*resources.DepositCreateResponse, error) {
	var result resources.DepositCreateResponse
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal body")
	}

	validations := validation.Errors{
		"/data/attributes/currency": validation.Validate(result.Data.Attributes.Currency, validation.Required, validation.In(Config(r).Deposits().Currency)),
		"/data/relationships/depositor/data": validation.Validate(result.Data.Relationships.Depositor.Data, validation.Required),
	}

	return &result, validations.Filter()
}

func CreateDeposit(w http.ResponseWriter, r *http.Request) {
	request, err := newCreateDepositRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	claims, err := Check(r, allow.Account{
		Address: request.Data.Relationships.Depositor.Data.ID,
	})

	if request.Data.Relationships.Depositor.Data.ID == "" || err != nil || claims == nil || !claims.PhoneVerified {
		ape.RenderErr(w, problems.NotAllowed())
		return
	}

	if Config(r).Deposits().DepositFraction == nil {
		limits := Config(r).Deposits().GetLimits(claims.PlatformID)
		validationsErr := validation.Errors{
			"/data/attributes/amount": validation.Validate(int64(request.Data.Attributes.Amount), validation.Required,
				validation.In(limits.SupportedDepositAmounts()...)),
		}.Filter()
		if validationsErr != nil {
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}
	}

	var deposit data.Deposit
	err = Storage(r).Transaction(func() error {
		// lock whole identity to be able to calculate current user balance
		var identity *data.Identity
		identity, err = Storage(r).Identities().GetForUpdate(claims.AccountID)
		if err != nil {
			return errors.Wrap(err, "failed to select identity")
		}

		if identity == nil {
			return errors.From(errors.New("expected identity to exist"), logan.F{
				"address": claims.AccountID})
		}

		var balance int64
		balance, err = getPotentialBalance(r, claims.AccountID)
		if err != nil {
			return errors.Wrap(err, "failed to calculate potential user balance")
		}

		if balance+int64(request.Data.Attributes.Amount) > Config(r).Deposits().MaxAmount {
			return problems.BadRequest(validation.Errors{
				"/data/attributes/amount": errors.New("total balance exceeds max allowed balance"),
			})[0]
		}

		deposit, err = Storage(r).Deposits().Create(int64(request.Data.Attributes.Amount), request.Data.Attributes.Currency, claims.AccountID)
		if err != nil {
			return errors.Wrap(err, "failed to create deposit")
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

	response := responses.NewDeposit(deposit, nil)
	json.NewEncoder(w).Encode(&resources.DepositResponse{
		Data:     response,
		Included: resources.Included{},
	})
}

func getPotentialBalance(r *http.Request, address string) (int64, error) {
	deposits, err := Storage(r).Deposits().Select(data.DepositsSelector{
		States:     []int32{int32(kafka.DepositV2_STATE_PENDING), int32(kafka.DepositV2_STATE_PAID), int32(kafka.DepositV2_STATE_LOCKED)},
		Depositors: []string{address},
	})

	if err != nil {
		return 0, errors.Wrap(err, "failed to select deposits")
	}

	if len(deposits) == 0 {
		return 0, nil
	}

	var total int64
	for _, deposit := range deposits {
		if deposit.Currency != Config(r).Deposits().Currency {
			return 0, errors.From(errors.New("expected deposits in one currenct"), logan.F{
				"expected": Config(r).Deposits().Currency,
				"actual":   deposit.Currency,
				"id":       deposit.ID,
			})
		}

		if deposit.Amount+total < total {
			return 0, errors.New("overflow")
		}

		total += deposit.Amount
	}

	return total, nil
}
