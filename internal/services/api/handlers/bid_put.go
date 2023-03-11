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
	"gitlab.com/eAuction/events/go/kafka"
	"gitlab.com/eAuction/go/amount"
)

type putBid struct {
	LotID   string
	Account string
	Amount  amount.Fiat
}

func newCreateBid(r *http.Request) (*putBid, error) {
	var request resources.BidPutResponse
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal body")
	}

	var result putBid
	if request.Data.Relationships.Lot.Data != nil {
		result.LotID = request.Data.Relationships.Lot.Data.ID
	}

	if request.Data.Relationships.Account.Data != nil {
		result.Account = request.Data.Relationships.Account.Data.ID
	}

	result.Amount = request.Data.Attributes.Amount

	return &result, nil
}

func CreateBid(w http.ResponseWriter, r *http.Request) {
	request, err := newCreateBid(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	identity, isAdminOperationAllow := GetAuthorizedAccount(r, w, request.Account, false)
	if identity == nil {
		return
	}

	err = Storage(r).Transaction(func() error {
		var lot data.Lot
		lot, err = getLotValidForParticipation(r, request.LotID, isAdminOperationAllow)
		if err != nil {
			return errors.Wrap(err, "failed to ensure lot is valid for participation")
		}

		const amountPath = "data/attributes/amount"
		if lot.Details.AuctionPrices.Start > int64(request.Amount) {
			return problems.BadRequest(validation.Errors{
				amountPath: errors.New("must be greater or equal to start price"),
			})[0]
		}

		var participant data.Participant
		participant, err = getParticipantValidForBidding(Storage(r), request.LotID, request.Account, isAdminOperationAllow)
		if err != nil {
			return errors.Wrap(err, "failed to get participant for update")
		}

		if participant.BidLimit < int64(request.Amount) {
			return problems.BadRequest(validation.Errors{
				amountPath: errors.New("exceeds bid limit"),
			})[0]
		}

		var participants []data.Participant
		participants, err = Storage(r).Participants().Select(data.ParticipantsSelector{
			State: []kafka.Participant_State{kafka.Participant_STATE_PENDING},
			LotID: &lot.ID,
		})
		if err != nil {
			return errors.Wrap(err, "failed to select participant")
		}

		if len(participants) == 0 {
			if !isAdminOperationAllow && int64(request.Amount) != lot.GetHighestBid() && lot.GetHighestBid() != 0 {
				return problems.BadRequest(validation.Errors{
					amountPath: errors.New("first bid must be equal to start price"),
				})[0]
			}
		} else {
			if !isAdminOperationAllow && lot.GetHighestBid()+lot.Details.MinBidStep > int64(request.Amount) {
				return problems.BadRequest(validation.Errors{
					amountPath: errors.New("bid must be greater or equal to highest bid + min bid step"),
				})[0]
			}
		}

		err = Storage(r).Participants().SetBid(participant.ID, int64(request.Amount))
		if err != nil {
			return errors.Wrap(err, "failed to set auto bid limit")
		}

		err = Storage(r).Lots().SetHighestBid(lot.ID, int64(request.Amount))
		if err != nil {
			return errors.Wrap(err, "failed to set highest bid")
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
