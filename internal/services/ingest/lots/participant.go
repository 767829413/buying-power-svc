package lots

import (
	"time"

	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/events/go/kafka"
	"gitlab.com/eAuction/platformer-svc/pkg/platformer"
)

type participant struct {
	log *logan.Entry

	platformer platformer.Platformer

	platfromIDs []string
	updatedPlatformIDsAt time.Time
}

func newParticipant(log *logan.Entry, platformer platformer.Platformer) *participant {
	return &participant{
		platformer: platformer,
		log:        log,
	}
}

// HandleRequest handles the participant created/updated events. HandleRequest will intentionally
// skip participations from external platforms (Copart/Iaai) and only use them
// for updating highest bid value.
func (h *participant) Handle(storage data.Storage, msg *kafka.AuctionEventMessage) error {
	var event interface {
		GetLotId() int64
		GetParticipant() *kafka.Participant
	}

	switch msg.Value.Type {
	case kafka.AuctionEvent_TYPE_PARTICIPANT_CREATED:
		event = msg.Value.ParticipantCreated
	case kafka.AuctionEvent_TYPE_PARTICIPANT_UPDATED:
		event = msg.Value.ParticipantUpdated
	default:
		return errors.Errorf("unexpected event type: %d", msg.Value.Type)
	}

	lotID := cast.ToString(event.GetLotId())
	p := event.GetParticipant()
	if p.State == kafka.Participant_STATE_PENDING {
		err := h.updateLotHighestBid(storage, lotID, p.CurrentBid)
		if err != nil {
			return errors.Wrap(err, "failed to update lot highest bid")
		}
	}

	if h.isPlatform(p) {
		return nil
	}

	err := h.handle(storage, lotID, p)
	if err != nil {
		return errors.Wrap(err, "failed to handle create/update of participant")
	}

	return nil
}

func (h *participant) getLot(storage data.Storage, id string) (*data.Lot, error) {
	lot, err := storage.Lots().ByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get lot by ID", logan.F{
			"lot_id": id,
		})
	}
	if lot == nil {
		return nil, errors.From(ErrLotNotFound, logan.F{
			"lot_id": id,
		})
	}

	return lot, nil
}

func (h *participant) updateLotHighestBid(storage data.Storage, lotID string, bid int64) error {
	if bid == 0 {
		return nil
	}

	if err := storage.Lots().SetHighestBid(lotID, bid); err != nil {
		return errors.Wrap(err, "failed to update lot prices", logan.F{
			"lot_id": lotID,
		})
	}

	return nil
}

func (h *participant) handle(storage data.Storage, lotID string, p *kafka.Participant) error {
	if err := h.updateParticipant(storage, lotID, p); err != nil {
		return errors.Wrap(err, "failed to update lot participant", logan.F{
			"lot_id":              lotID,
			"participant_address": p.AccountId,
		})
	}

	return nil
}

func (h *participant) updateParticipant(storage data.Storage, lotID string, p *kafka.Participant) error {
	existingParticipant, err := storage.Participants().ByLotIDAccountAddress(lotID, p.AccountId)
	if err != nil {
		return errors.Wrap(err, "failed to get existing participant")
	}

	if existingParticipant != nil && existingParticipant.IsLocal {
		return nil
	}

	if existingParticipant == nil {
		var lotExists bool
		lotExists, err = storage.Lots().Exists(lotID)
		if err != nil {
			return errors.Wrap(err, "failed to check if lot exists")
		}

		if !lotExists {
			return ErrLotNotFound
		}
	}

	return storage.Participants().Store(data.Participant{
		AccountAddress:  p.AccountId,
		LotID:           lotID,
		State:           int32(p.State),
		CurrentBid:      p.CurrentBid,
		RequestedBuyNow: p.RequestedBuyNow,
		BidLimit:        p.BidLimit,
		AutoBidLimit:    p.AutoBidLimit,
		UpdatedAt:       time.Now(),
		CreatedAt:       time.Now().UTC(),
		IsLocal:         false,
		IsSynced:        true,
	})
}


func (h *participant) isPlatform(p *kafka.Participant) bool {
	if h.updatedPlatformIDsAt.Add(30*time.Minute).Before(time.Now()) {
		h.platfromIDs = h.platformer.MustPlatformIDs()
		h.updatedPlatformIDsAt = time.Now()
	}
	for _, platformId := range h.platfromIDs  {
		if p.AccountId == platformId {
			return true
		}
	}

	return false
}
