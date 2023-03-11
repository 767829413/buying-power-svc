package handlers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/eAuction/bouncer/allow"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/events/go/kafka"

	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
)

// PutWinnerDetails is an endpoint of upserting winner details. Requires lot-id path segment.
func PutWinnerDetails(w http.ResponseWriter, r *http.Request) {
	lotID := chi.URLParam(r, "lot")
	if lotID == "" {
		ape.RenderErr(w, problems.BadRequest(validation.Errors{"path": errors.New("should contain lot id segment")})...)
		return
	}

	body := resources.PutWinnerResponse{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(validation.Errors{"body": err})...)
		return
	}

	newState, err := validatePutWinnerDetails(r, body, lotID)
	if err != nil {
		switch errors.Cause(err).(type) {
		case validation.Errors:
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		default:
			ape.Log(r).WithError(err).Error("something went wrong during request validation")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	}

	identity, err := Storage(r).Identities().Get(body.Data.Relationships.Account.Data.ID)
	if err != nil {
		panic(errors.Wrap(err, "failed to get identity"))
	}

	if identity == nil {
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"data/relationships/account/data/id": errors.New("not found"),
		})...)
	}

	_, err = Check(r, allow.Admin{
		Platform: identity.Platform,
	})
	if err != nil {
		ape.RenderErr(w, problems.NotAllowed(err))
		return
	}

	attributes := body.Data.Attributes
	newWinner := data.Winner{
		LotID:                       lotID,
		InvoiceID:                   attributes.InvoiceId,
		InvoiceCreatedAt:            attributes.InvoiceCreatedAt,
		ContainerID:                 attributes.ContainerId,
		ContainerLink:               attributes.ContainerLink,
		Fee:                         int64(attributes.Fee),
		FeeCurrency:                 attributes.FeeCurrency,
		TransportationPrice:         int64(attributes.TransportationPrice),
		TransportationPriceCurrency: attributes.TransportationPriceCurrency,
		State:                       int16(newState),
		DeliveredAt:                 attributes.DeliveredAt,
		Published:                   false,
		IsFromEvent:                 false,
	}

	err = Storage(r).Transaction(func() error {
		_, err = Storage(r).Lots().ByIDForUpdate(lotID)
		if err != nil {
			panic(errors.Wrap(err, "failed to select lot"))
		}

		err = markWinner(r, lotID, identity.ID)
		if err != nil {
			return errors.Wrap(err, "failed to mark winner")
		}

		err = Storage(r).Winners().Store(newWinner)
		if err != nil {
			return errors.Wrap(err, "failed to store winner")
		}

		if newState == kafka.Winner_STATE_DEAL_CANCELED {
			err = handleDealCanceled(r, lotID)
			if err != nil {
				return errors.Wrap(err, "failed to handle deal canceled")
			}
		}

		return nil
	})
	if err != nil {
		panic(errors.Wrap(err, "transaction failed"))
	}
	w.WriteHeader(http.StatusNoContent)
}

func markWinner(r *http.Request, lotID string, accountID string) error {
	winners, err := Storage(r).Participants().Select(data.ParticipantsSelector{
		State: []kafka.Participant_State{kafka.Participant_STATE_WINNER, kafka.Participant_STATE_BUY_NOW_WINNER},
		LotID: &lotID,
	})
	if err != nil {
		return errors.Wrap(err, "failed to select winning participants")
	}

	for _, w := range winners {
		if w.AccountAddress != accountID {
			return problems.Conflict()
		}
	}

	winner, err := Storage(r).Participants().GetForUpdate(lotID, accountID)
	if err != nil {
		return errors.Wrap(err, "failed to select participant")
	}

	if winner == nil {
		return problems.BadRequest(validation.Errors{
			"data/relationships/account/data/id": errors.New("not found"),
		})[0]
	}

	state := kafka.Participant_STATE_WINNER
	if winner.RequestedBuyNow {
		state = kafka.Participant_STATE_BUY_NOW_WINNER
	}

	err = Storage(r).Participants().SetState(winner.ID, state)
	if err != nil {
		return errors.Wrap(err, "failed to set state for winner")
	}

	participants, err := Storage(r).Participants().Select(data.ParticipantsSelector{
		State:     []kafka.Participant_State{kafka.Participant_STATE_PENDING, kafka.Participant_STATE_DEPOSIT_PENDING},
		LotID:     &lotID,
		ForUpdate: true,
	})
	if err != nil {
		return errors.Wrap(err, "failed to select pending participants")
	}

	for _, p := range participants {
		err = Storage(r).Participants().SetState(p.ID, kafka.Participant_STATE_LOST)
		if err != nil {
			return errors.Wrap(err,"failed to mark participant as lost")
		}
	}

	return nil
}

func handleDealCanceled(r *http.Request, lotID string) error {
	winningParticipants, err := Storage(r).Participants().Select(data.ParticipantsSelector{
		State:     []kafka.Participant_State{kafka.Participant_STATE_WINNER, kafka.Participant_STATE_BUY_NOW_WINNER},
		LotID:     &lotID,
		ForUpdate: true,
	})
	if err != nil {
		return errors.Wrap(err, "failed to select winning participants")
	}

	for _, wp := range winningParticipants {
		err = Storage(r).Participants().SetState(wp.ID, kafka.Participant_STATE_DEAL_CANCELED)
		if err != nil {
			return errors.Wrap(err, "failed to set deal cancelled state for participant")
		}
	}

	return nil
}

func validatePutWinnerDetails(r *http.Request, winner resources.PutWinnerResponse, lotID string) (kafka.Winner_State, error) {
	errs := validation.Errors{}

	newState := kafka.Winner_State(kafka.Winner_State_value[winner.Data.Attributes.State])
	if newState == kafka.Winner_STATE_UNKNOWN {
		errs["data/attributes/state"] = errors.New("unknown state " + winner.Data.Attributes.State)
	}

	lotExist, err := Storage(r).Lots().Exists(lotID)
	if err != nil {
		return 0, err
	}
	if !lotExist {
		errs["path/lot_id"] = errors.New("unknown lot")
	}

	if winner.Data.Relationships.Account.Data == nil {
		errs["/data/relationships/account/data"] = errors.New("must be present")
	}

	return newState, errs.Filter()
}
