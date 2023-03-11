package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
	"gitlab.com/eAuction/go/amount"
)

type putBidLimit struct {
	LotID   string
	Account string
	Amount  amount.Fiat
}

func newCreateBidLimit(r *http.Request) (*putBidLimit, error) {
	var request resources.AutoBidLimitPutResponse
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal body")
	}

	var result putBidLimit
	if request.Data.Relationships.Lot.Data != nil {
		result.LotID = request.Data.Relationships.Lot.Data.ID
	}

	if request.Data.Relationships.Account.Data != nil {
		result.Account = request.Data.Relationships.Account.Data.ID
	}

	result.Amount = request.Data.Attributes.Amount

	return &result, nil
}

func CreateBidLimit(w http.ResponseWriter, r *http.Request) {
	request, err := newCreateBidLimit(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	identity, isAdmin := GetAuthorizedAccount(r, w, request.Account, true)
	if identity == nil {
		return
	}

	err = Storage(r).Transaction(func() error {
		var participant data.Participant
		participant, err = getValidParticipant(Storage(r), request.LotID, request.Account, isAdmin)
		if err != nil {
			return errors.Wrap(err, "failed to get participant")
		}

		err = Storage(r).Participants().SetBidLimit(participant.ID, int64(request.Amount))
		if err != nil {
			return errors.Wrap(err, "failed to set auto bid limit")
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

	w.WriteHeader(http.StatusNoContent)
}
