package copart

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/lorem"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/responses"
	"gitlab.com/eAuction/copart-parser-v2/pkg/copart"
	"gitlab.com/eAuction/events/go/kafka"
	"gitlab.com/eAuction/go/amount"
)

func (h *handler) handleSold(event copart.Event) (string, []resources.EnvelopeResponse, error) {
	sold := event.Sold
	lot, err := h.getLot(event.SaleID, string(sold.LotNumber))
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to load lot for sold")
	}

	if lot == nil {
		return event.SaleID, []resources.EnvelopeResponse{responses.PopulateEnvelopWithUnknownLotEvent(string(sold.LotNumber),
			h.sales.GetShelfSaleID(event.SaleID))}, nil
	}

	isPending := event.Type == copart.EventTypeSoldPending
	winnerID, err := h.getWinnerID(*lot, *event.Sold, isPending)
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to get winner id")
	}

	saleKey := resources.Key{
		ID:   h.sales.GetShelfSaleID(event.SaleID),
		Type: resources.SALES,
	}.AsRelation()

	lotKey := resources.Key{
		ID:   lot.ID,
		Type: resources.LOTS,
	}.AsRelation()

	winnerKey := resources.Key{
		ID:   winnerID,
		Type: resources.IDENTITIES,
	}.AsRelation()
	soldResource := resources.Sold{
		Key: resources.Key{
			ID:   lorem.ULID(),
			Type: resources.SOLDS,
		},
		Attributes: resources.SoldAttributes{
			Amount:    amount.Fiat(sold.Bid),
			IsPending: isPending,
		},
		Relationships: resources.SoldRelationships{
			Lot:   *lotKey,
			Sale:  *saleKey,
			Buyer: *winnerKey,
		},
	}

	env := resources.EnvelopeResponse{
		Data: resources.Envelope{
			Key: resources.Key{
				Type: resources.ENVELOPES,
			},
			Relationships: resources.EnvelopeRelationships{
				Sold: soldResource.Key.AsRelation(),
			},
		},
		Included: resources.Included{},
	}

	env.Included.Add(&soldResource)
	return event.SaleID, []resources.EnvelopeResponse{env}, nil

}

func (h *handler) markLost(event copart.Sold, lot data.Lot) error {
	txStorage := h.storage.Clone()
	err := txStorage.Transaction(func() error {
		participants, err := txStorage.Participants().Select(data.ParticipantsSelector{
			State:     []kafka.Participant_State{kafka.Participant_STATE_PENDING, kafka.Participant_STATE_DEPOSIT_PENDING},
			LotID:     &lot.ID,
			ForUpdate: true,
			Platforms: h.platforms,
		})

		if err != nil {
			return errors.Wrap(err, "failed to select participants")
		}

		for _, p := range participants {
			err = txStorage.Participants().SetOutbidedWith(p.ID, int64(event.Bid))
			if err != nil {
				return errors.Wrap(err, "failed to set outbided with")
			}
			err = txStorage.Participants().SetState(p.ID, kafka.Participant_STATE_LOST)
			if err != nil {
				return errors.Wrap(err, "failed to set state")
			}
		}

		return nil
	})

	return errors.Wrap(err, "failed to commit tx")
}

func (h *handler) getWinnerID(lot data.Lot, sold copart.Sold, isPending bool) (string, error) {
	if !isPending {
		err := h.storage.Lots().SetHighestBid(lot.ID, int64(sold.Bid))
		if err != nil {
			return "", errors.Wrap(err, "failed to set highest bid for the lot")
		}
	}


	if sold.BuyerNumber != h.buyerNumber {
		if isPending {
			return lot.Details.PlatformId, nil
		}

		err := h.markLost(sold, lot)
		if err != nil {
			return "", errors.Wrap(err, "failed to mark lost")
		}

		return lot.Details.PlatformId, nil

	}

	txStorage := h.storage.Clone()
	var winner string
	err := txStorage.Transaction(func() error {
		participant, err := txStorage.Participants().GetClosestForUpdate(lot.ID, int64(sold.Bid), h.platforms)
		if err != nil {
			return errors.Wrap(err, "failed to select winning participant")
		}

		if participant == nil {
			return nil
		}

		err = txStorage.Participants().SetBid(participant.ID, int64(sold.Bid))
		if err != nil {
			return errors.Wrap(err, "failed to set bid for participant")
		}

		winner = participant.AccountAddress
		return nil
	})
	return winner, errors.Wrap(err, "failed to commit tx")
}
