package copart

import (
	"context"
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/lorem"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/responses"
	"gitlab.com/eAuction/copart-parser-v2/pkg/copart"
	"time"
)

//Handler - handles copart messages
type Handler interface {
	Handle(ctx context.Context, event copart.Event) (string, [][]byte, error)
}

type handler struct {
	lots             data.LotsCache
	buyerNumber      string
	storage          data.Storage
	log              *logan.Entry
	platforms        []string
	copartPlatformID string
	isSimulation     bool
	sales            Sales
}

//newHandler - creates new instance of copart handler
func newHandler(log *logan.Entry, storage data.Storage, lots data.LotsCache, platforms []string,
	isSimulation bool, copartPlatformID, buyerNumber string, sales Sales) Handler {
	return &handler{
		lots:             lots,
		buyerNumber:      buyerNumber,
		storage:          storage,
		log:              log,
		platforms:        platforms,
		isSimulation:     isSimulation,
		copartPlatformID: copartPlatformID,
		sales:            sales,
	}
}

//Handle - converts msg into response
func (h *handler) Handle(_ context.Context, event copart.Event) (string, [][]byte, error) {
	saleID, rawResults, err := h.handleEvent(event)
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to handle event", logan.F{
			"event":   event,
		})
	}

	results := make([][]byte, len(rawResults))
	for i, rawResult := range rawResults {
		results[i], err = json.Marshal(rawResult)
		if err != nil {
			panic(errors.Wrap(err, "failed to marshal resposne"))
		}
	}

	return saleID, results, nil
}

func (h *handler) handleEvent(event copart.Event) (string, []resources.EnvelopeResponse, error) {
	switch event.Type {
	case copart.EventTypeNewBid:
		return h.handleNewBid(event)
	case copart.EventTypeAuctioneer:
		return h.handleBonusTime(event)
	case copart.EventTypeCurrentItem:
		return h.handleCurrentItem(event)
	case copart.EventTypeSoldPending, copart.EventTypeSold:
		return h.handleSold(event)
	case copart.EventTypeStartItem:
		return h.handleStartItem(event)
	case copart.EventTypePreBid:
		return h.handlePreBid(event)
	default:
		return "", nil, nil
	}
}

func (h *handler) handleStartItem(event copart.Event) (string, []resources.EnvelopeResponse, error) {
	currentItem := event.StartItem
	lot, err := h.getLot(event.SaleID, string(currentItem.LotNumber))
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to load lot for bonus time")
	}

	if lot == nil {
		return "", []resources.EnvelopeResponse{responses.PopulateEnvelopWithUnknownLotEvent(string(currentItem.LotNumber),
			h.sales.GetShelfSaleID(event.SaleID))}, nil
	}

	result := []resources.EnvelopeResponse{
		responses.PopulateEnvelopeWithLot(*lot),
		h.newBid(currentItem.StartBid, currentItem.NextBid, *lot, lot.Details.PlatformId, h.sales.GetShelfSaleID(event.SaleID)),
	}
	return event.SaleID, result, nil
}

func (h *handler) handleCurrentItem(event copart.Event) (string, []resources.EnvelopeResponse, error) {
	currentItem := event.CurrentItem
	lot, err := h.getLot(event.SaleID, string(currentItem.LotNumber))
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to load lot for bonus time")
	}

	if lot == nil {
		return "", []resources.EnvelopeResponse{responses.PopulateEnvelopWithUnknownLotEvent(string(currentItem.LotNumber),
			h.sales.GetShelfSaleID(event.SaleID))}, nil
	}

	result := []resources.EnvelopeResponse{
		responses.PopulateEnvelopeWithLot(*lot),
		h.newBid(0, currentItem.StartBid, *lot, lot.Details.PlatformId, h.sales.GetShelfSaleID(event.SaleID)),
	}
	return event.SaleID, result, nil
}

func (h *handler) handleBonusTime(event copart.Event) (string, []resources.EnvelopeResponse, error) {
	bonusTime := event.Auctioneer
	lot, err := h.getLot(event.SaleID, string(bonusTime.LotNumber))
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to load lot for bonus time")
	}

	if lot == nil {
		return event.SaleID, []resources.EnvelopeResponse{responses.PopulateEnvelopWithUnknownLotEvent(string(bonusTime.LotNumber),
			h.sales.GetShelfSaleID(event.SaleID))}, nil
	}

	saleKey := resources.Key{
		ID:   h.sales.GetShelfSaleID(event.SaleID),
		Type: resources.SALES,
	}.AsRelation()
	bonusTimeResource := resources.BonusTime{
		Key: resources.Key{
			ID:   lorem.ULID(),
			Type: resources.BONUS_TIMES,
		},
		Attributes: resources.BonusTimeAttributes{
			DurationS: int64(time.Duration(bonusTime.BonusTime) / time.Second),
		},
		Relationships: resources.BonusTimeRelationships{
			Lot: resources.Key{
				ID:   lot.ID,
				Type: resources.LOTS,
			}.AsRelation(),
			Sale: *saleKey,
		},
	}

	env := resources.EnvelopeResponse{
		Data: resources.Envelope{
			Key: resources.Key{
				Type: resources.ENVELOPES,
			},
			Relationships: resources.EnvelopeRelationships{
				BonusTime: bonusTimeResource.Key.AsRelation(),
			},
		},
		Included: resources.Included{},
	}

	env.Included.Add(&bonusTimeResource)
	return event.SaleID, []resources.EnvelopeResponse{env}, nil

}

func (h *handler) getLot(copartSaleID, copartLotID string) (*data.Lot, error) {
	lot, err := h.lots.Get("copart-"+copartLotID, h.isSimulation)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get copart lot")
	}

	if lot == nil {
		return nil, nil
	}

	if !lot.SaleID.Valid {
		h.log.WithFields(logan.F{
			"lot_id":  lot.ID,
			"sale_id": copartSaleID,
		}).Warn("live bidding requested lot without valid sale id")
		return lot, nil
	}

	copartSaleIDFromLot, _ := data.ExtractCopartSaleID(lot.SaleID.String)
	if copartSaleIDFromLot != copartSaleID {
		h.log.WithFields(logan.F{
			"copart_sale_id":          copartSaleID,
			"lot_sale_id":             lot.SaleID.String,
			"copart_sale_id_from_lot": copartSaleIDFromLot,
		}).Debug("sale id from lot does not match sale id from event - returning nil lot")
		return nil, nil
	}

	return lot, nil
}
