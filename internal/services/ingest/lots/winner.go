package lots

import (
	"time"

	"gitlab.com/eAuction/buying-power-svc/internal/data"

	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/events/go/kafka"
)

type winner struct {
	log *logan.Entry
}

func newWinner(log *logan.Entry) *winner {
	return &winner{log: log}
}

// Handle - handles creation of new lot winner or updating the existing one.
// We don't differ created/updated events and just upsert data to the storage,
// using lot ID as the key.
func (h *winner) Handle(storage data.Storage, msg *kafka.AuctionEventMessage) error {
	kafkaWinner := msg.Value.GetWinner()

	dbWinner := data.Winner{
		LotID:                       cast.ToString(kafkaWinner.LotId),
		Fee:                         kafkaWinner.FeeAmount,
		FeeCurrency:                 kafkaWinner.FeeCurrency,
		TransportationPrice:         kafkaWinner.TransportationPrice,
		TransportationPriceCurrency: kafkaWinner.TransportationPriceCurrency,
		State:                       int16(kafkaWinner.State),
		IsFromEvent:                 true,
		Published:                   true,
		Version:                     0,
	}

	if kafkaWinner.InvoiceId != "" {
		dbWinner.InvoiceID = &kafkaWinner.InvoiceId
	}

	if kafkaWinner.InvoiceCreatedAt != 0 {
		invoiceCreatedAt := time.Unix(kafkaWinner.InvoiceCreatedAt, 0).UTC()
		dbWinner.InvoiceCreatedAt = &invoiceCreatedAt
	}

	if kafkaWinner.ContainerId != "" {
		dbWinner.ContainerID = &kafkaWinner.ContainerId
	}

	if kafkaWinner.ContainerLink != "" {
		dbWinner.ContainerLink = &kafkaWinner.ContainerLink
	}

	lotExists, err := storage.Lots().Exists(dbWinner.LotID)
	if err != nil {
		return errors.Wrap(err, "failed to check if lot exists")
	}

	if !lotExists {
		return ErrLotNotFound
	}

	currentWinner, err := storage.Winners().GetByLotID(dbWinner.LotID)
	if err != nil {
		return err
	}
	
	if currentWinner != nil && !currentWinner.IsFromEvent {
		return nil
	}

	if err = storage.Winners().Store(dbWinner); err != nil {
		return errors.Wrap(err, "failed to store winner", logan.F{
			"lot_id": kafkaWinner.LotId,
		})
	}

	return nil
}
