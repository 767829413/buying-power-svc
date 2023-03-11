package lots

import (
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/events/go/kafka"
)

type lotStateUpdated struct {
	log *logan.Entry
}

func newLotStateUpdated(log *logan.Entry) *lotStateUpdated {
	return &lotStateUpdated{log: log}
}

// HandleRequest - handles updates of the lot state
func (h *lotStateUpdated) Handle(storage data.Storage, msg *kafka.AuctionEventMessage) error {
	event := msg.Value
	state := event.GetLotStateUpdated().State
	lotID := cast.ToString(event.GetLotStateUpdated().LotId)

	switch state {
	case kafka.Lot_STATE_FINISHED, kafka.Lot_STATE_SOLD:
		return h.updateLotState(storage, lotID, int32(state))
	case kafka.Lot_STATE_PENDING, kafka.Lot_STATE_NEGOTIATIONS:
		return nil
	default:
		return errors.From(errors.New("recovered unknown lot state"), logan.F{
			"state": state,
		})
	}
}

func (h *lotStateUpdated) updateLotState(storage data.Storage, lotID string, state int32) error {
	err := storage.Lots().SetState(lotID, state)
	if err != nil {
		if errors.Cause(err) == data.ErrorLotNotFound {
			return errors.From(ErrLotNotFound, logan.F{
				"lot_id": lotID,
				"state":  state,
			})
		}
		return errors.Wrap(err, "failed to update lot state", logan.F{
			"state_to_set": state,
		})
	}

	if state == int32(kafka.Lot_STATE_FINISHED) || state == int32(kafka.Lot_STATE_SOLD) {
		lostParticipantsState := []kafka.Participant_State{kafka.Participant_STATE_DEPOSIT_PENDING}
		if state == int32(kafka.Lot_STATE_FINISHED) {
			lostParticipantsState = append(lostParticipantsState, kafka.Participant_STATE_PENDING)
		}
		participants, err := storage.Participants().Select(data.ParticipantsSelector{
			State:     lostParticipantsState,
			LotID:     &lotID,
			ForUpdate: true,
		})
		if err != nil {
			return errors.Wrap(err, "failed to select participants")
		}

		for _, p := range participants {
			err = storage.Participants().SetState(p.ID, kafka.Participant_STATE_LOST)
			if err != nil {
				return errors.Wrap(err, "failed to set state for participant")
			}
		}
	}

	return nil
}
