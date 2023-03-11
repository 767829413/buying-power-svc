package handlers

import (
	"gitlab.com/eAuction/bouncer/allow"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/events/go/kafka"
)

type deleteParticipants struct {
	LotID   string
}

func newDeleteParticipants(r *http.Request) (*deleteParticipants, error) {
	return &deleteParticipants{
		LotID:   chi.URLParam(r, "lot"),
	}, nil
}

func DeleteParticipants(w http.ResponseWriter, r *http.Request) {
	request, err := newDeleteParticipants(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	claims, err := Check(r, allow.Admin{})
	if err != nil || claims == nil {
		ape.RenderErr(w, problems.NotAllowed())
		return
	}

	err = Storage(r).Transaction(func() error {
		var participants []data.Participant
		participants, err = Storage(r).Participants().Select(data.ParticipantsSelector{
			State:         []kafka.Participant_State{kafka.Participant_STATE_PENDING, kafka.Participant_STATE_DEPOSIT_PENDING},
			Platforms:     []string{claims.PlatformID},
			LotID:         &request.LotID,
			ForUpdate:     true,
		})
		if err != nil {
			return errors.Wrap(err, "failed to select participants")
		}

		for _, participant := range participants {
			err = Storage(r).Participants().SetState(participant.ID, kafka.Participant_STATE_LOST)
			if err != nil {
				return errors.Wrap(err, "failed to mark participant lost")
			}
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
