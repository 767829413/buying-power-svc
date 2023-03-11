package lots

import (
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"strings"
	"time"

	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/events/go/kafka"
)

type liveBiddingDetails struct {
	log *logan.Entry
	cfg config.Config
}

func newLiveBiddingDetails(log *logan.Entry, cfg config.Config) *liveBiddingDetails {
	return &liveBiddingDetails{log: log, cfg: cfg}
}

// Handle - handles setting of live bidding details
func (h *liveBiddingDetails) Handle(storage data.Storage, msg *kafka.AuctionEventMessage) error {
	event := msg.Value
	details := event.LiveBiddingDetails
	if details == nil {
		h.log.
			WithFields(logan.F{"offset": msg.Offset(), "time": msg.Time()}).
			Error("unexpected state: live bidding details is nil for live bidding details event - skipping message")
		return nil
	}

	lotID := cast.ToString(details.LotId)

	saleID := ""
	if details.Lane != "" {
		saleID = kafka.EncodeSaleID(details.Platform, details.Branch, details.Lane, time.Unix(details.StartTime, 0))
	}

	if err := storage.Lots().SetLiveBiddingDetails(lotID, saleID, details.Lane, details.ItemNumber); err != nil {
		return errors.Wrap(err, "failed to store live bidding details")
	}

	err := h.storeSale(storage, saleID, msg)
	if err != nil {
		return errors.Wrap(err, "failed to store sale")
	}

	return nil
}

func (h *liveBiddingDetails) storeSale(storage data.Storage, saleID string, msg *kafka.AuctionEventMessage) error {
	if saleID == "" {
		return nil
	}

	details := msg.Value.LiveBiddingDetails
	platform, err := h.cfg.Platformer().GetPlatform(details.Platform)
	if err != nil {
		return errors.Wrap(err, "failed to get platform")
	}

	if platform == nil || strings.ToUpper(platform.Data.Attributes.Code) != "COPART" {
		h.log.WithFields(logan.F{
			"platform":  details.Platform,
			"offset":    msg.Offset(),
			"partition": msg.Partition(),
		}).Error("received live bidding details for not supported platform")
		return nil
	}

	copartSaleID, err := data.ExtractCopartSaleID(saleID)
	if err != nil {
		return errors.Wrap(err, "failed to extract copart sale ID", logan.F{
			"sale_id": saleID,
		})
	}

	err = storage.Sales().Upsert(data.Sale{
		ID:         saleID,
		StartsAt:   time.Unix(details.StartTime, 0),
		ExternalID: copartSaleID,
		Platform:   details.Platform,
	})
	if err != nil {
		return errors.Wrap(err, "failed to upsert sale")
	}

	return nil
}
