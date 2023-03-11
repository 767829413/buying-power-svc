package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/bouncer"
	"gitlab.com/eAuction/bouncer/allow"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
)

type attachDeposits struct {
	lotID    string              `json:"-"`
	Data     []resources.Key     `json:"data"`
	deposits map[string]struct{} `json:"-"`
}

func newAttachDepositRequest(r *http.Request) (*attachDeposits, error) {
	result := attachDeposits{
		lotID:    chi.URLParam(r, "lot"),
		deposits: map[string]struct{}{},
	}

	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal body")
	}

	for _, depositKey := range result.Data {
		if depositKey.Type != resources.DEPOSITS {
			return nil, validation.Errors{
				"data": errors.New("must be deposits only"),
			}
		}

		result.deposits[depositKey.ID] = struct{}{}
	}

	return &result, nil
}

func AttachDeposits(w http.ResponseWriter, r *http.Request) {
	request, err := newAttachDepositRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	claims, err := Check(r, allow.Account{})
	if err != nil || claims == nil {
		if err == nil {
			err = bouncer.ErrNotSigned
		}
		ape.RenderErr(w, problems.NotAllowed(err))
		return
	}

	depositor := claims.AccountID

	err = Storage(r).Transaction(func() error {
		return AttachDepositsV2(r, attachDepositsV2{
			lotID:     request.lotID,
			depositor: depositor,
			deposits:  request.deposits,
		})
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

	w.WriteHeader(http.StatusNoContent)
}
