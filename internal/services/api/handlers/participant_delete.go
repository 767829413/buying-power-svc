package handlers

import (
	"gitlab.com/eAuction/buying-power-svc/internal/services/deposits"
	"net/http"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/events/go/kafka"
)

type deleteParticipant struct {
	LotID   string
	Account string
}

func newDeleteParticipant(r *http.Request) (*deleteParticipant, error) {
	return &deleteParticipant{
		LotID:   chi.URLParam(r, "lot"),
		Account: chi.URLParam(r, "account"),
	}, nil
}

func DeleteParticipant(w http.ResponseWriter, r *http.Request) {
	request, err := newDeleteParticipant(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// TODO: do not allow user to delete
	identity, _ := GetAuthorizedAccount(r, w, request.Account, false)
	if identity == nil {
		return
	}

	err = Storage(r).Transaction(func() error {
		var participant *data.Participant
		participant, err = Storage(r).Participants().GetForUpdate(request.LotID, request.Account)
		if err != nil {
			return errors.Wrap(err, "failed to get participant")
		}

		if participant == nil {
			return problems.NotFound()
		}

		if participant.State != int32(kafka.Participant_STATE_PENDING) &&
			participant.State != int32(kafka.Participant_STATE_DEPOSIT_PENDING) {
			return problems.BadRequest(validation.Errors{
				"account": errors.New("must be in state pending or deposit pending"),
			})[0]
		}

		err = Storage(r).Participants().SetState(participant.ID, kafka.Participant_STATE_REJECTED)
		if err != nil {
			return errors.Wrap(err, "failed to mark participant rejected")
		}

		err = deposits.ReturnDepositForParticipant(Storage(r), *participant, Config(r).Deposits())
		if err != nil {
			return errors.Wrap(err, "failed to return deposits for participants")
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
