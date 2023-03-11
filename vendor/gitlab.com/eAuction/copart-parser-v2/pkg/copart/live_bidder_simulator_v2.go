package copart

/*

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"strconv"
)

type liveBidderSimulationV2 struct {
	ctx    context.Context
	cancel context.CancelFunc
	lots   Lots
	log    *logan.Entry
}

//NewLiveBidderSimulationV2 - returns new instance of live bidder simulation
func NewLiveBidderSimulationV2(ctx context.Context, lots Lots, log *logan.Entry) LiveBidderV2 {
	log = log.WithField("service", "live-bidder-simulator")
	ctx, cancel := context.WithCancel(ctx)
	return &liveBidderSimulationV2{
		lots:   lots,
		ctx:    ctx,
		cancel: cancel,
		log:    log,
	}
}

// SubscribeToSale - subscribes to a copart sale. Use context cancel to unsubscribe
func (l *liveBidderSimulationV2) SubscribeToSale(ctx context.Context, saleID string) (SaleHandler, error) {
	yard, lane, err := parseSaleID(saleID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse sale ID")
	}

	lots, err := l.lots.SelectExternalIDForSale(strconv.FormatInt(yard, 10), lane)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select random lots")
	}

	if len(lots) == 0 {
		return nil, ErrSaleNotFound
	}

	return newSaleHandlerSimulatorV2(l.ctx, l.log, saleID, lots), nil
}

func (l *liveBidderSimulationV2) BuyerNumber() string {
	return strconv.FormatInt(simulationBuyerNumber, 10)
}

func (l *liveBidderSimulationV2) Close() {
	l.cancel()
}

func (l *liveBidderSimulationV2) Done() <-chan struct{} {
	return l.ctx.Done()
}
*/