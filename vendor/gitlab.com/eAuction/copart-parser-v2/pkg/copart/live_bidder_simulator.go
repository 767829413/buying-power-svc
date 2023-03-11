package copart

/*

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Lots interface {
	SelectExternalIDForSale(branchID, lane string) ([]string, error)
}

type sale struct {
	bids chan simulationBid
}

type liveBidderSimulator struct {
	lock      sync.Mutex
	requestID uint64
	sales     map[string]sale
	notifier  *notifier
	ctx       context.Context
	cancel    context.CancelFunc
	lots      Lots
	log       *logan.Entry
}

//NewLiveBidderSimulation - returns new instance of live bidder simulation
func NewLiveBidderSimulation(ctx context.Context, lots Lots, log *logan.Entry) LiveBidder {
	log = log.WithField("service", "live-bidder-simulator")
	ctx, cancel := context.WithCancel(ctx)
	return &liveBidderSimulator{
		sales:    map[string]sale{},
		notifier: newNotifier(ctx, log),
		ctx:      ctx,
		cancel:   cancel,
		lots:     lots,
		log:      log,
	}
}

func (r *liveBidderSimulator) Close() {
	r.cancel()
}

func (r *liveBidderSimulator) PingRequest() []byte {
	return nil
}

func (r *liveBidderSimulator) generateRequestID() uint64 {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.requestID++
	return r.requestID
}

func (r *liveBidderSimulator) Done() <-chan struct{} {
	return r.ctx.Done()
}

func (r *liveBidderSimulator) Send(ctx context.Context, rawRequest Request) error {
	switch request := rawRequest.(type) {
	case *subscribe:
		return r.handleSubscribe(ctx, rawRequest, request)
	case *bid:
		return r.handleBid(ctx, rawRequest, request)
	default:
		return nil
	}
}

func (r *liveBidderSimulator) handleBid(ctx context.Context, request Request, bid *bid) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	s, ok := r.sales[bid.SaleID]
	if !ok {
		r.log.WithField("bid", bid).Warn("received bid for unknown sale")
		return nil
	}

	select {
	case s.bids <- simulationBid{
		BuyerNumber: r.GetBuyerNumber(),
		LotID:       LotNumber(strconv.FormatInt(bid.LotID, 10)),
		AmountCents: bid.AmountCents,
		SaleID:      bid.SaleID,
	}:
	default:
	}

	return nil
}

func (r *liveBidderSimulator) handleSubscribe(ctx context.Context, request Request, subscribe *subscribe) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if _, ok := r.sales[subscribe.SaleID]; ok {
		r.notifySubscription(ctx, request.GetID(), subscribe.SaleID)
		return nil
	}

	yard, lane, err := parseSaleID(subscribe.SaleID)
	if err != nil {
		return errors.Wrap(err, "failed to parse sale ID")
	}

	lots, err := r.lots.SelectExternalIDForSale(strconv.FormatInt(yard, 10), lane)
	if err != nil {
		return errors.Wrap(err, "failed to select random lots")
	}

	if len(lots) == 0 {
		go r.notifyAll(ctx, Message{
			RequestID: request.GetID(),
			Type:      MessageTypeBatchSubscribe,
			BatchSubscribe: &BatchSubscribe{
				IsSuccess:     false,
				ResourceName:  subscribe.SaleID,
				ResourceAlias: "channel not found",
			},
		})

		return nil
	}

	newSale := sale{
		bids: make(chan simulationBid, 1),
	}
	r.sales[subscribe.SaleID] = newSale
	go r.runSimulation(request.GetID(), subscribe.SaleID, lots, newSale.bids)
	return nil
}

func (r *liveBidderSimulator) notifySubscription(ctx context.Context, requestID uint64, saleID string) {
	r.notifyAll(ctx, Message{
		RequestID: requestID,
		Type:      MessageTypeBatchSubscribe,
		BatchSubscribe: &BatchSubscribe{
			IsSuccess:    true,
			ResourceName: saleID,
		},
	})
}

func (r *liveBidderSimulator) runSimulation(requestID uint64, saleID string, lots []string, bids <-chan simulationBid) {
	log := r.log.WithFields(logan.F{
		"request_id": requestID,
		"sale_id":    saleID,
	})
	log.Debug("starting new simulation")

	defer func() {
		r.lock.Lock()
		defer r.lock.Unlock()
		// TODO: FIX ME (causes panic as bidder tries to write into closed channel)
		//close(r.sales[saleID].bids)
		delete(r.sales, saleID)
		log.Debug("finished simulation")
	}()
	err := running.RunSafely(r.ctx, "simulation:"+saleID, func(ctx context.Context) error {
		startBidder(r.ctx, log, saleID, r.sales[saleID].bids, r.notifier)
		startBidder(r.ctx, log, saleID, r.sales[saleID].bids, r.notifier)
		r.notifySubscription(ctx, requestID, saleID)
		const unknownLotI = 3
		if len(lots) > unknownLotI {
			lots = append(lots, "copart-nonexistinglot-startime")
			lots[unknownLotI], lots[len(lots)-1] = lots[len(lots)-1], lots[unknownLotI]
		}
		for _, externalID := range lots {
			log.Debug("starting lot simulation")
			err := r.runLotAuction(log, saleID, externalID, bids)
			if err != nil {
				return errors.Wrap(err, "failed to run simulation for lot", logan.F{
					"lot_external_id": externalID,
				})
			}
		}

		r.notifyAll(r.ctx, Message{
			Type: MessageTypeEventReceive,
			Event: &Event{
				SaleID: saleID,
				Type:   EventTypeEndAuction,
			},
		})

		return nil
	})
	if err != nil {
		r.log.WithError(err).Error("simulation failed")
	}
}

func getLotNumber(externalID string) (LotNumber, error) {
	parts := strings.Split(externalID, "-")
	if len(parts) != 3 {
		return "", errors.From(errors.New("unexpected format of external ID"), logan.F{
			"ex_id": externalID,
		})
	}

	return LotNumber(parts[1]), nil
}

func (r *liveBidderSimulator) getSale(saleID string) sale {
	r.lock.Lock()
	defer r.lock.Unlock()
	s, ok := r.sales[saleID]
	if !ok {
		panic(errors.From(errors.New("tried to get non existing sale"), logan.F{
			"sale_id": saleID,
		}))
	}

	return s
}

func (r *liveBidderSimulator) notifyAll(ctx context.Context, msg Message) {
	_ = r.notifier.NotifyAll(ctx, msg)
}

//Subscribe - returns channel with all messages received from  Panics, if same id used twice
func (r *liveBidderSimulator) Subscribe(ctx context.Context, id string) <-chan Message {
	return r.notifier.Subscribe(ctx, id)

}

func (r *liveBidderSimulator) SubscribeRequest(saleID string) (Request, error) {
	return &subscribe{
		request: request{
			ID: r.generateRequestID(),
		},
		SaleID: saleID,
	}, nil
}
func (r *liveBidderSimulator) AuthorizeRequest(saleID, auctionToken string) (Request, error) {
	return &authorize{
		request: request{
			ID: r.generateRequestID(),
		},
		SaleID: saleID,
	}, nil
}
func (r *liveBidderSimulator) BidRequest(saleID string, lotID, amount int64) (Request, error) {
	return &bid{
		request: request{
			ID: r.generateRequestID(),
		},
		SaleID:      saleID,
		LotID:       lotID,
		AmountCents: amount,
	}, nil
}
func (r *liveBidderSimulator) StartSessionRequest() (Request, error) {
	return &startSession{
		request: request{
			ID: r.generateRequestID(),
		},
	}, nil
}

func (r *liveBidderSimulator) GetCookiesJar() http.CookieJar {
	panic(errors.New("not implemented"))
}

const simulationBuyerNumber = 45213414

func (r *liveBidderSimulator) GetBuyerNumber() int64 {
	return simulationBuyerNumber
}

func (r *liveBidderSimulator) GetFullBuyerNumber() BuyerNumber {
	return BuyerNumber{
		Primary:   simulationBuyerNumber,
		Secondary: "secondary buyer number",
	}
}

func (r *liveBidderSimulator) GetAuctionToken(saleID string) (string, error) {
	return "auction_token_for_" + saleID, nil
}

func (r *liveBidderSimulator) GetSessionBuyerToken() string {
	return "session_buyer_token"
}

func (r *liveBidderSimulator) GetSale(saleID string) (*SaleData, error) {
	return &SaleData{
		AuctionKey:         saleID,
		AuctionID:          0,
		Dbdate:             "",
		IsAdsSale:          false,
		YardNumber:         0,
		LaneCode:           "",
		SaleStartTimestamp: "",
	}, nil
}

func (r *liveBidderSimulator) SaveAuctionJoinDetails(saleData SaleData, auctionToken string) error {
	return nil
}

func (r *liveBidderSimulator) AuthenticateSale(saleID string) error {
	return nil
}
*/