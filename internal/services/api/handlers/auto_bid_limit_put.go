package handlers

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
	"gitlab.com/eAuction/go/amount"
)

type putAutoBidLimit struct {
	LotID   string
	Account string
	Amount  amount.Fiat
}

func newCreateAutoBidLimit(r *http.Request) (*putAutoBidLimit, error) {
	var request resources.AutoBidLimitPutResponse
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal body")
	}

	var result putAutoBidLimit
	if request.Data.Relationships.Lot.Data != nil {
		result.LotID = request.Data.Relationships.Lot.Data.ID
	}

	if request.Data.Relationships.Account.Data != nil {
		result.Account = request.Data.Relationships.Account.Data.ID
	}

	result.Amount = request.Data.Attributes.Amount

	return &result, nil
}

func CreateAutoBidLimit(w http.ResponseWriter, r *http.Request) {
	request, err := newCreateAutoBidLimit(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	identity, isAdmin := GetAuthorizedAccount(r, w, request.Account, false)
	if identity == nil {
		return
	}

	err = Storage(r).Transaction(func() error {
		var lot data.Lot
		lot, err = getLotValidForParticipation(r, request.LotID, isAdmin)
		if err != nil {
			return errors.Wrap(err, "failed to ensure lot is valid for participation")
		}

		var participant data.Participant
		participant, err = getParticipantValidForBidding(Storage(r), request.LotID, request.Account, isAdmin)
		if err != nil {
			return errors.Wrap(err, "failed to get participant for update")
		}

		participant.AutoBidLimit = int64(request.Amount)
		if participant.AutoBidLimit != 0 && lot.GetHighestBid() >= participant.AutoBidLimit {
			return problems.BadRequest(validation.Errors{
				"data/attributes/amount": errors.New("below start price or highest bid"),
			})[0]
		}

		if participant.AutoBidLimit > participant.BidLimit {
			return problems.BadRequest(validation.Errors{
				"data/attributes/amount": errors.New("exceeds bid limit"),
			})[0]
		}

		err = Storage(r).Participants().SetAutoBidLimit(participant.ID, participant.AutoBidLimit)
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
