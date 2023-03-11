package copart

/*

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/lorem"
	"math/rand"
	"strconv"
	"time"
)

// bidder - sends maxNumberOfBids bids for each lot in the sale
type bidder struct {
	saleID          string
	maxNumberOfBids int64
	bidsCount       map[LotNumber]int64
	bids            chan simulationBid
	buyerNumber     int64
	log             *logan.Entry
	ctx             context.Context
	cancel          context.CancelFunc
}

func startBidder(ctx context.Context, log *logan.Entry, saleID string, bids chan simulationBid, n *notifier) {
	buyerNumber := rand.Int63()
	ctx, cancel := context.WithCancel(ctx)
	b := bidder{
		saleID:          saleID,
		maxNumberOfBids: 3,
		buyerNumber:     buyerNumber,
		bidsCount:       map[LotNumber]int64{},
		log:             log.WithField("service", "bidder_simulation").WithField("buyer_number", buyerNumber),
		ctx:             ctx,
		cancel:          cancel,
		bids:            bids,
	}

	msgs := n.Subscribe(ctx, b.saleID+":"+lorem.ULID())
	go b.run(ctx, msgs)
}

func (b *bidder) run(ctx context.Context, msgs <-chan Message) {
	defer func() {
		_ = recover()
	}()
	b.log.Debug("starting new bidder")
	const maxInactivity = time.Minute
	timer := time.NewTimer(maxInactivity)
	for {
		select {
		case <-timer.C:
			return
		case <-ctx.Done():
			return
		case msg, ok := <-msgs:
			if !ok {
				return
			}

			if !b.handelNewMsg(msg) {
				continue
			}

			if !timer.Stop() {
				<-timer.C
			}

			timer.Reset(maxInactivity)
		}
	}
}

func (b *bidder) handelNewMsg(msg Message) bool {
	event := msg.Event
	if event == nil || event.SaleID != b.saleID {
		return false
	}

	if event.CurrentItem != nil {
		go b.sendBid(event.SaleID, event.CurrentItem.LotNumber, int64(event.CurrentItem.StartBid))
		return true
	}

	newBid := event.NewBid
	if newBid == nil {
		return true
	}

	if newBid.BuyerNumber == strconv.FormatInt(b.buyerNumber, 10) {
		b.bidsCount[newBid.LotNumber] = b.bidsCount[newBid.LotNumber] + 1
	}

	numberOfBids := b.bidsCount[newBid.LotNumber]
	if numberOfBids >= b.maxNumberOfBids {
		return true
	}

	go b.sendBid(event.SaleID, newBid.LotNumber, int64(newBid.NextBidCents))
	return true
}

func (b *bidder) sendBid(saleID string, rawLotID LotNumber, nextBid int64) {
	time.Sleep(time.Duration(rand.Int63n(5))*time.Second + 3*time.Second)
	select {
	case b.bids <- simulationBid{
		BuyerNumber: b.buyerNumber,
		LotID:       rawLotID,
		AmountCents: nextBid,
		SaleID:      saleID,
	}:
	default:
	}
}*/
