package services

import (
	"context"
	"strconv"
	"time"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/eAuction/events/go/kafka"

	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
)

// RunLotCloser periodically checks if there are old pending lots and publish it to auctions topic.
func RunLotCloser(ctx context.Context, cfg config.Config) {
	l := lotCloser{
		lots:   cfg.Storage().Lots(),
		writer: cfg.AuctionsWriter(),
	}
	// such period should allow the system to not post duplicate events.
	// in any case, the state will be ingested eventually and the events
	// for the lot will stop to be produced.
	running.WithBackOff(ctx, cfg.Log(), "lot_closer", l.runOnce, 10*time.Minute, 10*time.Second, time.Minute)
}

type lotCloser struct {
	lots   data.Lots
	writer kafka.AuctionEventWriter
}

func (l lotCloser) runOnce(ctx context.Context) error {
	lots, err := l.lots.SelectEndedPending()
	if err != nil {
		return errors.Wrap(err, "failed to select old pending lots")
	}
	for _, lot := range lots {
		lotID, err := strconv.ParseInt(lot.ID, 10, 64)
		if err != nil {
			return errors.Wrap(err, "failed to parse lot id")
		}
		event := kafka.AuctionEvent{
			Type: kafka.AuctionEvent_TYPE_LOT_STATE_UPDATED,
			LotStateUpdated: &kafka.LotStateUpdated{
				LotId: lotID,
				State: kafka.Lot_STATE_FINISHED,
			},
		}
		err = l.writer.WriteMessages(ctx, event.Message())
		if err != nil {
			return errors.Wrap(err, "failed to write closing event")
		}
	}
	return nil
}
