package copart

import (
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/lorem"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/responses"
	"gitlab.com/eAuction/copart-parser-v2/pkg/copart"
	"gitlab.com/eAuction/go/amount"
)

func (h *handler) handleNewBid(event copart.Event) (string, []resources.EnvelopeResponse, error) {
	newBid := event.NewBid
	lot, err := h.getLot(event.SaleID, string(newBid.LotNumber))
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to load lot for new bid")
	}

	if lot == nil {
		return event.SaleID, []resources.EnvelopeResponse{responses.PopulateEnvelopWithUnknownLotEvent(string(newBid.LotNumber),
			h.sales.GetShelfSaleID(event.SaleID))}, nil
	}

	bidder, err := h.tryUpdateBidder(*lot, *newBid)
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to update bidder")
	}

	return event.SaleID, []resources.EnvelopeResponse{h.newBid(newBid.CurrentBidCents, newBid.NextBidCents, *lot,
		bidder, h.sales.GetShelfSaleID(event.SaleID))}, nil
}

func (h *handler) tryUpdateBidder(lot data.Lot, newBid copart.NewBid) (string, error) {
	bidder := lot.Details.PlatformId
	if newBid.BuyerNumber != h.buyerNumber {
		return bidder, nil
	}

	txStorage := h.storage.Clone()
	err := txStorage.Transaction(func() error {
		participant, err := txStorage.Participants().GetClosestForUpdate(lot.ID, int64(newBid.CurrentBidCents), h.platforms)
		if err != nil {
			return errors.Wrap(err, "failed to get participant closest to bid")
		}

		if participant == nil {
			h.log.WithFields(logan.F{
				"lot_id":  lot.ID,
				"new_bid": newBid,
			}).Warn("received bid from controlled copart account, but failed to find matching participant")
			return nil
		}

		bidder = participant.AccountAddress

		err = txStorage.Participants().SetBid(participant.ID, int64(newBid.CurrentBidCents))
		if err != nil {
			return errors.Wrap(err, "failed to set bid")
		}

		return nil
	})

	return bidder, errors.Wrap(err, "failed to commit tx")
}

func (h *handler) newBid(currentBid, nextBid copart.Amount, lot data.Lot, bidder, saleID string) resources.EnvelopeResponse {
	// NOTE: we need to explicitly specify sale id from event, as one in the lot could be nil or different
	lotRelation := resources.Key{
		ID:   lot.ID,
		Type: resources.LOTS,
	}.AsRelation()

	bidRelation := resources.Key{
		ID:   bidder,
		Type: resources.IDENTITIES,
	}.AsRelation()
	bidResource := resources.Bid{
		Key: resources.Key{
			ID:   lorem.ULID(),
			Type: resources.BIDS,
		},
		Attributes: resources.BidAttributes{
			Amount:        amount.Fiat(currentBid),
			NextBidAmount: amount.Fiat(nextBid),
			Currency:      lot.Currency,
		},
		Relationships: resources.BidRelationships{
			Lot:    *lotRelation,
			Bidder: *bidRelation,
			Sale: resources.Key{
				ID:   saleID,
				Type: resources.SALES,
			}.AsRelation(),
		},
	}
	env := resources.EnvelopeResponse{
		Data: resources.Envelope{
			Key: resources.Key{
				Type: resources.ENVELOPES,
			},
			Relationships: resources.EnvelopeRelationships{
				Bid: bidResource.Key.AsRelation(),
			},
		},
		Included: resources.Included{},
	}

	env.Included.Add(&bidResource)
	return env
}
