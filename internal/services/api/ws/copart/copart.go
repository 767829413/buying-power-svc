package copart

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/robfig/cron"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	router "gitlab.com/eAuction/buying-power-svc/internal/services/api/ws/router"
	"gitlab.com/eAuction/copart-parser-v2/pkg/copart"
)

var rng = rand.New(rand.NewSource(42))

type Sale struct {
	CopartSaleID string
	ShelfSaleID  string
}

//LiveBidder - wrapper over copart.LiveBidder
type LiveBidder struct {
	bidder    copart.LiveBidder
	log       *logan.Entry
	ctx       context.Context
	cancel    context.CancelFunc
	platforms []string
	sales     Sales

	router router.Router
}

// log *logan.Entry, storage data.Storage, lots data.LotsCache, platforms []string,
//	isSimulation bool, copartPlatformID, buyerNumber string, sales Sales

func NewLiveBidder(ctx context.Context, log *logan.Entry, bidder copart.LiveBidder, platforms []string,
	storage data.Storage, lots data.LotsCache, isSimulation bool, copartPlatformID string) (*LiveBidder, error) {
	ctx, cancel := context.WithCancel(ctx)
	log = log.WithField("service", "buying-power-copart-live-bidder")
	r := router.New(ctx, log)
	s := newSales(log, r)
	h := newHandler(log, storage, lots, platforms, isSimulation, copartPlatformID, bidder.BuyerNumber(), s)
	listener := newEventsListener(ctx, log, bidder.Subscribe(ctx, "events_listener"), h, r, s)
	result := &LiveBidder{
		bidder:    bidder,
		log:       log,
		ctx:       ctx,
		cancel:    cancel,
		router:    r,
		platforms: platforms,
		sales:     s,
	}

	// Reload the client once a day at ~ 10 AM UTC
	// This helps to mitigate possible memory leaks on Copart's side and to spread
	// load among proxies more evenly.
	// Reload time is randomized to avoid simultaneous reload of all clients
	c := cron.New()
	hour := 10
	minute := rng.Intn(60)
	if time.Now().Hour() == hour && time.Now().Minute() > minute {
		// go back to avoid repeated reloads at 10 AM
		hour--
	}
	c.AddFunc(fmt.Sprintf("0 %d %d * * *", minute, hour), func() {
		log.Info("reloading the client on schedule")
		cancel()
	})
	c.Start()

	go func() {
		defer result.Close()
		select {
		case <-listener.Done():
		case <-bidder.Done():
		case <-ctx.Done():
			c.Stop()
		}
	}()

	return result, nil
}

//IsPlatformSupported - returns true if platform is supported
func (r *LiveBidder) IsPlatformSupported(id string) bool {
	for _, p := range r.platforms {
		if p == id {
			return true
		}
	}

	return false
}

func (r *LiveBidder) Close() {
	r.cancel()
	r.router.Close()
}

//Subscribe - subscribes for copart events converted to json marshaled resources.EnvelopeResponse
func (r *LiveBidder) Subscribe(ctx context.Context, connectionID string) <-chan []byte {
	return r.router.AddNewListener(ctx, connectionID)
}

//ctx - returns context of the router
func (r *LiveBidder) Ctx() context.Context {
	return r.ctx
}

//UnsubscribeFromSale - unsubscribe user from sale
func (r *LiveBidder) UnsubscribeFromSale(connectionID, saleID string) {
	r.router.Unsubscribe(connectionID, saleID)
}

func (r *LiveBidder) MakeBid(ctx context.Context, requestID string, saleID, lotID string, amount int64) error {
	ok := r.sales.IsSubscribed(saleID)
	if !ok {
		return errors.New("need to subscribe to sale first")
	}

	err := r.bidder.Send(ctx, copart.ParserCommandBid{
		SaleID: saleID,
		LotID:  lotID,
		MaxBid: copart.Amount(amount),
		ID:     requestID,
	})
	return errors.Wrap(err, "failed to send bid command")
}

//SubscribeForSale - subscribe connection for sale. Returns true, nil if subscription has been performed right await.
// on false, nil need to wait for response from copart
func (r *LiveBidder) SubscribeForSale(ctx context.Context, connectionID, requestID string, sale Sale) {
	log := r.log.WithFields(logan.F{
		"sale_id":       sale.CopartSaleID,
		"shelf_sale_id": sale.ShelfSaleID,
		"connection_id": connectionID,
		"request_id":    requestID,
	})

	shouldSend := r.sales.AcquireSubscribeLock(connectionID, sale)
	if !shouldSend {
		log.Debug("already subscribed to the sale")
		return
	}

	log.Debug("create sale subscription")

	err := r.bidder.Send(ctx, copart.ParserCommandJoin{SaleID: sale.CopartSaleID})
	if err != nil {
		log.WithError(err).Error("failed to subscribe for sale - sending sale end msg for all listeners")
		r.sales.SaleEnded(sale.CopartSaleID, err.Error())
		return
	}

	return
}
