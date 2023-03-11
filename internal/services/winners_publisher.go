package services

import (
	"context"
	"strconv"
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/eAuction/events/go/kafka"

	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
)

// RunWinnersPublisher starts winner event producing service.
func RunWinnersPublisher(ctx context.Context, cfg config.Config) {
	const serviceName = "winners_publisher"
	log := cfg.Log().WithField("service", serviceName)

	storage := cfg.Storage()
	s := winnersPublisher{
		writer:     cfg.AuctionsWriter(),
		log:        log,
		storage:    storage.Winners(),
		transactor: storage,
	}

	log.Info("service started")
	running.WithBackOff(ctx, log, serviceName, s.runOnce, 100*time.Millisecond, time.Second, time.Minute)
}

type winnersPublisher struct {
	writer     kafka.AuctionEventWriter
	log        *logan.Entry
	storage    data.Winners
	transactor pgdb.Transactor
}

func (s winnersPublisher) runOnce(ctx context.Context) error {
	const winnerBatchLimit = 20
	for {
		winners, err := s.storage.SelectNotPublished(winnerBatchLimit)
		if err != nil {
			return errors.Wrap(err, "failed to select unpublished winners")
		}

		for _, w := range winners {
			if ctx.Err() != nil {
				return ctx.Err()
			}

			err = s.processWinner(ctx, w)
			if err != nil {
				return errors.Wrap(err, "failed to process winner")
			}
		}
		if len(winners) < winnerBatchLimit {
			return nil
		}
	}
}

func (s winnersPublisher) processWinner(ctx context.Context, w data.Winner) error {
	lotID, err := strconv.Atoi(w.LotID)
	if err != nil {
		return errors.Wrap(err, "failed to parse lot_id as number")
	}
	err = s.transactor.Transaction(func() error {
		err := s.storage.MarkAsPublished(w.LotID, w.Version)
		if err != nil {
			return errors.Wrap(err, "failed to mark winner as published")
		}

		invoiceID := ""
		if w.InvoiceID != nil {
			invoiceID = *w.InvoiceID
		}
		containerID := ""
		if w.ContainerID != nil {
			containerID = *w.ContainerID
		}
		containerLink := ""
		if w.ContainerLink != nil {
			containerLink = *w.ContainerLink
		}
		invoiceCreatedAt := int64(0)
		if w.InvoiceCreatedAt != nil {
			invoiceCreatedAt = w.InvoiceCreatedAt.UTC().Unix()
		}
		deliveredAt := int64(0)
		if w.DeliveredAt != nil {
			deliveredAt = w.DeliveredAt.UTC().Unix()
		}
		msgs := []kafka.AuctionEventMessage{
			kafka.AuctionEvent{
				Type: kafka.AuctionEvent_TYPE_WINNER,
				Winner: &kafka.Winner{
					LotId:                       int64(lotID),
					InvoiceId:                   invoiceID,
					InvoicePaid:                 kafka.Winner_State(w.State) == kafka.Winner_STATE_DEAL_COMPLETED,
					InvoiceCreatedAt:            invoiceCreatedAt,
					ContainerId:                 containerID,
					ContainerLink:               containerLink,
					FeeAmount:                   w.Fee,
					FeeCurrency:                 w.FeeCurrency,
					TransportationPrice:         w.TransportationPrice,
					TransportationPriceCurrency: w.TransportationPriceCurrency,
					DeliveredAt:                 deliveredAt,
					State:                       kafka.Winner_State(w.State),
				},
			}.Message(),
		}

		if !w.LotSoldPublished {
			msgs = append(msgs, kafka.AuctionEvent{
				Type: kafka.AuctionEvent_TYPE_LOT_STATE_UPDATED,
				LotStateUpdated: &kafka.LotStateUpdated{
					LotId: int64(lotID),
					State: kafka.Lot_STATE_SOLD,
				},
			}.Message())
		}

		err = s.writer.WriteMessages(ctx, msgs...)
		if err != nil {
			return errors.Wrap(err, "failed to write winner event")
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "transaction failed")
	}
	return nil
}
