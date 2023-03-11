package lots

import (
	"database/sql"
	"time"

	"gitlab.com/eAuction/buying-power-svc/internal/data"

	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/events/go/kafka"
)

type lotCreated struct {
	log *logan.Entry
}

func newLotCreated(log *logan.Entry) *lotCreated {
	return &lotCreated{log: log}
}

// HandleRequest - handles creation of the new lot
func (h *lotCreated) Handle(storage data.Storage, msg *kafka.AuctionEventMessage) error {
	event := msg.Value
	lot := event.GetLotCreated().GetLot()
	if lot == nil {
		h.log.
			WithFields(logan.F{"offset": msg.Offset(), "time": msg.Time()}).
			Error("unexpected state: lot is nil for lot created event - skipping message")
		return nil
	}

	lotID := cast.ToString(lot.Id)

	dbLot := data.Lot{
		ID:         lotID,
		Currency:   lot.Currency,
		EndsAt:     time.Unix(lot.StartTime+int64(lot.DurationSec), 0).UTC(),
		State:      int32(kafka.Lot_STATE_PENDING),
		Details:    data.LotDetails(*lot),
		ExternalID: sql.NullString{
			String: lot.ExternalId,
			Valid:  lot.ExternalId != "",
		},
	}

	if err := storage.Lots().Store(dbLot); err != nil {
		return errors.Wrap(err, "failed to store lot")
	}

	return nil
}
