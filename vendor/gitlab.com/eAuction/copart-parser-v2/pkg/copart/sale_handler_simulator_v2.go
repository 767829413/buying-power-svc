package copart
/*

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type saleHandlerSimulatorV2 struct {
	ctx      context.Context
	cancel   context.CancelFunc
	msgs     <-chan Message
	once     sync.Once
	bids     chan simulationBid
	saleID   string
	log      *logan.Entry
	notifier *notifier
	lots     []string
}

func newSaleHandlerSimulatorV2(ctx context.Context, log *logan.Entry, saleID string, lots []string) *saleHandlerSimulatorV2 {
	ctx, cancel := context.WithCancel(ctx)
	log = log.WithField("sale_id", saleID)
	n := newNotifier(ctx, log)
	result := &saleHandlerSimulatorV2{
		ctx:      ctx,
		cancel:   cancel,
		msgs:     n.Subscribe(ctx, "main"),
		once:     sync.Once{},
		bids:     make(chan simulationBid, 1),
		saleID:   saleID,
		log:      log,
		notifier: n,
		lots:     lots,
	}

	go result.run()
	return result
}

func (s *saleHandlerSimulatorV2) Close() {
	s.cancel()
	s.notifier.CloseAll()
	s.once.Do(func() {
		close(s.bids)
	})
}

func (s *saleHandlerSimulatorV2) Messages() <-chan Message {
	return s.msgs
}

func (s *saleHandlerSimulatorV2) MakeBid(ctx context.Context, requestID, lotID string, amount int64) error {
	defer func() {
		// suppress panic
		_ = recover()
	}()
	select {
	case s.bids <- simulationBid{
		BuyerNumber: simulationBuyerNumber,
		LotID:       LotNumber(lotID),
		AmountCents: amount,
		SaleID:      s.saleID,
	}:
		return nil
	case <-ctx.Done():
	case <-s.ctx.Done():
	}

	return ErrSaleNotFound
}

func (s *saleHandlerSimulatorV2) Snipe(ctx context.Context, requestID, lotID string, maxBid int64) error {
	return errors.New("snipe is not implemented")
}

func (s *saleHandlerSimulatorV2) run() {
	s.log.Debug("starting new simulation")

	defer func() {
		s.log.Debug("finished simulation")
		s.Close()
	}()
	err := running.RunSafely(s.ctx, "simulation:"+s.saleID, func(ctx context.Context) error {
		startBidder(s.ctx, s.log, s.saleID, s.bids, s.notifier)
		startBidder(s.ctx, s.log, s.saleID, s.bids, s.notifier)
		s.notifySubscription(ctx, 1, s.saleID)
		const unknownLotI = 3
		if len(s.lots) > unknownLotI {
			s.lots = append(s.lots, "copart-nonexistinglot-startime")
			s.lots[unknownLotI], s.lots[len(s.lots)-1] = s.lots[len(s.lots)-1], s.lots[unknownLotI]
		}
		for _, externalID := range s.lots {
			s.log.Debug("starting lot simulation")
			err := s.runLotAuction(externalID)
			if err != nil {
				return errors.Wrap(err, "failed to run simulation for lot", logan.F{
					"lot_external_id": externalID,
				})
			}
		}

		s.notifyAll(s.ctx, Message{
			Type: MessageTypeEventReceive,
			Event: &Event{
				SaleID: s.saleID,
				Type:   EventTypeEndAuction,
			},
		})

		return nil
	})
	if err != nil {
		s.log.WithError(err).Error("simulation failed")
	}
}

func (s *saleHandlerSimulatorV2) notifyAll(ctx context.Context, msg Message) {
	_ = s.notifier.NotifyAll(ctx, msg)
}

func (s *saleHandlerSimulatorV2) notifySubscription(ctx context.Context, requestID uint64, saleID string) {
	s.notifyAll(ctx, Message{
		RequestID: requestID,
		Type:      MessageTypeBatchSubscribe,
		BatchSubscribe: &BatchSubscribe{
			IsSuccess:    true,
			ResourceName: saleID,
		},
	})
}



func (s *saleHandlerSimulatorV2) runLotAuction(externalID string) error {
	lotNumber, err := getLotNumber(externalID)
	if err != nil {
		return errors.Wrap(err, "failed to get lot number")
	}

	log := s.log.WithField("lot_number", lotNumber)

	log.Debug("sending current item")
	startBid := rand.Int63n(200) * 50 * 100
	s.notifyAll(s.ctx, Message{
		RequestID: 0,
		Type:      MessageTypeEventReceive,
		Event: &Event{
			SaleID: s.saleID,
			Type:   EventTypeCurrentItem,
			CurrentItem: &CurrentItem{
				LotNumber: lotNumber,
				StartBid:  Amount(startBid),
			},
		},
	})

	highestBid := s.notifyNewBids(simulationBid{
		BuyerNumber: rand.Int63(),
		LotID:       lotNumber,
		AmountCents: startBid - bidStep,
		SaleID:      s.saleID,
	})

	log.Debug("sending sold pending")
	s.notifyAll(s.ctx, Message{
		Type: MessageTypeEventReceive,
		Event: &Event{
			SaleID: s.saleID,
			Type:   EventTypeSoldPending,
			Sold: &Sold{
				LotNumber:   lotNumber,
				Bid:         Amount(highestBid.AmountCents),
				BuyerNumber: strconv.FormatInt(highestBid.BuyerNumber, 10),
			},
		},
	})

	time.Sleep(time.Second * 5)
	log.Debug("sending sold")
	s.notifyAll(s.ctx, Message{
		Type: MessageTypeEventReceive,
		Event: &Event{
			SaleID: s.saleID,
			Type:   EventTypeSold,
			Sold: &Sold{
				LotNumber:   lotNumber,
				Bid:         Amount(highestBid.AmountCents),
				BuyerNumber: strconv.FormatInt(highestBid.BuyerNumber, 10),
			},
		},
	})
	time.Sleep(time.Second * 3)

	return nil
}



func (s *saleHandlerSimulatorV2) notifyNewBids(highestBid simulationBid) simulationBid {
	timer := time.NewTimer(DefaultWaitTime)

	bonusTimePassed := false
	for {
		select {
		case <-s.ctx.Done():
			return highestBid
		case <-timer.C:
			if bonusTimePassed {
				return highestBid
			}
			s.log.Debug("sending bonus time")
			s.notifyAll(s.ctx, Message{
				Type: MessageTypeEventReceive,
				Event: &Event{
					SaleID: highestBid.SaleID,
					Type:   EventTypeAuctioneer,
					Auctioneer: &Auctioneer{
						LotNumber: highestBid.LotID,
						BonusTime: CopartDuration(DefaultWaitTime),
					},
				},
			})

			bonusTimePassed = true

		case newBid := <-s.bids:
			if newBid.LotID != highestBid.LotID && newBid.SaleID != highestBid.SaleID {
				s.log.WithField("new_bid", newBid).WithField("highest_bid", highestBid).Warn("received bid not matching for the sale")
				continue
			}

			if highestBid.AmountCents >= newBid.AmountCents {
				s.log.WithField("new_bid", newBid).WithField("highest_bid", highestBid).Warn("higher bid already present")
				continue
			}

			bonusTimePassed = false
			highestBid = newBid

			s.notifyAll(s.ctx, Message{
				RequestID: 0,
				Type:      MessageTypeEventReceive,
				Event: &Event{
					SaleID: highestBid.SaleID,
					Type:   EventTypeNewBid,
					NewBid: &NewBid{
						CurrentBidCents: Amount(highestBid.AmountCents),
						LotNumber:       highestBid.LotID,
						NextBidCents:    Amount(highestBid.AmountCents + bidStep),
						BuyerNumber:     strconv.FormatInt(highestBid.BuyerNumber, 10),
						Type:            "I",
					},
				},
			})
		}

		s.log.Debug("reseting wait time timer")
		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}

		timer.Reset(DefaultWaitTime)
	}
}*/


