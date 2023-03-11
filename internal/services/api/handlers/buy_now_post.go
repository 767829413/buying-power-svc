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
)

type postBuyNow struct {
	LotID   string
	Account string
}

func newPostBuyNow(r *http.Request) (*postBuyNow, error) {
	var request resources.BuyNowPostResponse
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal body")
	}

	var result postBuyNow
	if request.Data.Relationships.Account.Data != nil {
		result.Account = request.Data.Relationships.Account.Data.ID
	}

	if request.Data.Relationships.Lot.Data != nil {
		result.LotID = request.Data.Relationships.Lot.Data.ID
	}

	return &result, nil
}

func CreateBuyNow(w http.ResponseWriter, r *http.Request) {
	request, err := newPostBuyNow(r)
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
		// no need to allow admins to shortcut checks here
		lot, err = getLotValidForParticipation(r, request.LotID, false)
		if err != nil {
			return errors.Wrap(err, "failed to ensure lot is valid for participation")
		}

		if lot.Details.AuctionPrices == nil || lot.Details.AuctionPrices.BuyNow <= 0 {
			return problems.BadRequest(validation.Errors{
				"lot": errors.New("does not support buy now"),
			})[0]
		}

		var participant data.Participant
		participant, err = getValidParticipant(Storage(r), request.LotID, request.Account, isAdmin)
		if err != nil {
			return errors.Wrap(err, "failed to get participant")
		}

		if participant.RequestedBuyNow {
			return problems.Conflict()
		}

		var isSufficient bool
		isSufficient, err = BuyingPowerCalculator(r).IsLimitSufficientForLot(lot, participant.BidLimit, lot.Currency,
			true, 0)
		if err != nil {
			return errors.Wrap(err, "failed to check is limit sufficient")
		}

		if !isSufficient {
			return problems.BadRequest(validation.Errors{
				"lot": errors.New("bid limit is not sufficient"),
			})[0]
		}

		var taken bool
		taken, err = isBuyNowAlreadyTaken(r, request.LotID, identity.ID)
		if err != nil {
			return errors.Wrap(err, "failed to check if buy now is already taken")
		}

		if taken {
			return problems.Conflict()
		}

		err = Storage(r).Participants().SetBuyNow(participant.ID)
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

func getValidParticipant(storage data.Storage, lotID, identityID string, isAdmin bool) (data.Participant, error) {
	participant, err := storage.Participants().GetForUpdate(lotID, identityID)
	if err != nil {
		return data.Participant{}, errors.Wrap(err, "failed to get participant for update")
	}

	if participant == nil {
		return data.Participant{}, problems.BadRequest(validation.Errors{
			"lot": errors.New("participant must exist"),
		})[0]
	}

	if !isAdmin && participant.State != int32(kafka.Participant_STATE_PENDING) {
		return data.Participant{}, problems.BadRequest(validation.Errors{
			"lot": errors.New("participant must be in state pending"),
		})[0]
	}

	return *participant, nil
}

func getParticipantValidForBidding(storage data.Storage, lotID, identityID string, isAdmin bool) (data.Participant, error) {
	participant, err := getValidParticipant(storage, lotID, identityID, isAdmin)
	if err != nil {
		return participant, err
	}

	if participant.RequestedBuyNow {
		return data.Participant{}, problems.BadRequest(validation.Errors{
			"data/attributes/amount": errors.New("requested buy now"),
		})[0]
	}

	return participant, nil
}

func isBuyNowAlreadyTaken(r *http.Request, lotID, identityID string) (bool, error) {
	// lock lot to ensure we do not endup with two participation requesting buy now
	_, err := Storage(r).Lots().ByIDForUpdate(lotID)
	if err != nil {
		return false, errors.Wrap(err, "failed to select lot for update")
	}

	participants, err := Storage(r).Participants().Select(data.ParticipantsSelector{
		State: []kafka.Participant_State{kafka.Participant_STATE_PENDING, kafka.Participant_STATE_DEPOSIT_PENDING},
		LotID: &lotID,
	})
	if err != nil {
		return false, errors.Wrap(err, "failed to select participants")
	}

	for _, participant := range participants {
		if participant.RequestedBuyNow && participant.AccountAddress != identityID {
			return true, nil
		}
	}

	return false, nil
}
