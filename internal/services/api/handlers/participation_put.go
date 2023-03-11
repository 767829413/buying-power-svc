package handlers

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
	"gitlab.com/eAuction/events/go/kafka"
)

type putParticipation struct {
	Deposits map[string]struct{} `json:"-"`
	Account  string              `json:"-"`
	LotID    string              `json:"-"`
}

func newCreateParticipation(r *http.Request) (*putParticipation, error) {
	var request resources.ParticipantPutResponse
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal body")
	}

	var result putParticipation
	if request.Data.Relationships.Account.Data != nil {
		result.Account = request.Data.Relationships.Account.Data.ID
	}

	if request.Data.Relationships.Lot.Data != nil {
		result.LotID = request.Data.Relationships.Lot.Data.ID
	}

	result.Deposits = map[string]struct{}{}
	for _, d := range request.Data.Relationships.Deposits.Data {
		result.Deposits[d.ID] = struct{}{}
	}

	return &result, nil
}

func CreateParticipant(w http.ResponseWriter, r *http.Request) {
	request, err := newCreateParticipation(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	identity, isAdmin := GetAuthorizedAccount(r, w, request.Account, false)
	if identity == nil {
		return
	}

	err = Storage(r).Transaction(func() error {
		_, err = getLotValidForParticipation(r, request.LotID, isAdmin)
		if err != nil {
			return errors.Wrap(err, "failed to ensure lot is valid for participation")
		}

		err = ensureParticipantExists(r, *request, isAdmin)
		if err != nil {
			return errors.Wrap(err, "failed to ensure particopant exists")
		}

		return AttachDepositsV2(r, attachDepositsV2{
			lotID:     request.LotID,
			deposits:  request.Deposits,
			depositor: identity.ID,
			isAdmin:   isAdmin,
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

func ensureParticipantExists(r *http.Request, request putParticipation, isAdmin bool) error {
	p, err := Storage(r).Participants().ByLotIDAccountAddress(request.LotID, request.Account)
	if err != nil {
		return errors.Wrap(err, "failed to select participant")
	}

	if p == nil {
		err = Storage(r).Participants().Create(request.Account, request.LotID)
		if err != nil {
			return errors.Wrap(err, "failed to store participant")
		}

		return nil
	}

	if !isAdmin {
		return nil
	}

	states := map[kafka.Participant_State]struct{}{
		kafka.Participant_STATE_REJECTED:      {},
		kafka.Participant_STATE_LOST:          {},
		kafka.Participant_STATE_DEAL_CANCELED: {},
	}
	if _, requiresRecreate := states[kafka.Participant_State(p.State)]; !requiresRecreate {
		return nil
	}

	Log(r).WithFields(logan.F{
		"identity":    p.AccountAddress,
		"lot":         p.LotID,
		"participant": *p,
		"state":       p.State,
	}).Info("resetting participant to deposit pending state")

	err = Storage(r).Participants().SetState(p.ID, kafka.Participant_STATE_DEPOSIT_PENDING)
	return errors.Wrap(err, "failed to set state for participant")
}

func getLotValidForParticipation(r *http.Request, lotID string, isAdmin bool) (data.Lot, error) {
	lot, err := Storage(r).Lots().ByID(lotID)
	if err != nil {
		return data.Lot{}, errors.Wrap(err, "failed to select lot")
	}

	if lot == nil {
		return data.Lot{}, problems.BadRequest(validation.Errors{
			"lot": errors.New("not found"),
		})[0]
	}

	if lot.Details.Type == kafka.Lot_TYPE_NO_AUCTION || lot.Details.AuctionPrices == nil {
		return data.Lot{}, problems.BadRequest(validation.Errors{
			"lot": errors.New("it is not allowed to participate in no auction lot"),
		})[0]
	}

	if !isAdmin && lot.State != int32(kafka.Lot_STATE_PENDING) {
		return data.Lot{}, problems.BadRequest(validation.Errors{
			"lot": errors.New("already sold"),
		})[0]
	}

	// time for live bidding and for admin to set actual bid
	if !isAdmin && lot.EndsAt.Add(12*time.Hour).UTC().Before(time.Now().UTC()) {
		return data.Lot{}, problems.BadRequest(validation.Errors{
			"lot": errors.New("too late to participate"),
		})[0]
	}

	return *lot, nil
}
