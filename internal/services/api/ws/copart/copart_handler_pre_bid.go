package copart

import (
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
	"gitlab.com/eAuction/copart-parser-v2/pkg/copart"
	"gitlab.com/eAuction/events/go/kafka"
	"time"
)

func (h *handler) handlePreBid(event copart.Event) (string, []resources.EnvelopeResponse, error) {
	preBid := event.PreBid
	log := h.log.WithFields(logan.F{
		"sale_id":    event.SaleID,
		"lot_number": preBid.LotNumber,
	})
	log.Debug("received preBid")
	lot, err := h.getLot(event.SaleID, string(preBid.LotNumber))
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to load lot prebid sold")
	}

	if lot == nil {
		log.Debug("failed to find lot for prebid - skipping")
		return event.SaleID, nil, nil
	}

	log.Debug("creating participant for prebid")
	err = h.storage.Participants().Store(data.Participant{
		AccountAddress:  lot.Details.PlatformId,
		LotID:           lot.ID,
		State:           int32(kafka.Participant_STATE_PENDING),
		RequestedBuyNow: false,
		CurrentBid:      int64(preBid.PreBid),
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
	})
	if err != nil {
		return event.SaleID, nil, errors.Wrap(err, "failed to insert platform participant for pre bid")
	}

	log.Debug("participant for prebid created")

	return event.SaleID, nil, nil
}
