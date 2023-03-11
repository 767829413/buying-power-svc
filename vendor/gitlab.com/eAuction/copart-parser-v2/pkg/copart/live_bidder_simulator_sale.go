package copart

/*

import (
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"math/rand"
	"strconv"
	"time"
)

type simulationBid struct {
	BuyerNumber int64
	LotID       LotNumber
	AmountCents int64
	SaleID      string
}

func (r *liveBidderSimulator) runLotAuction(log *logan.Entry, saleID, externalID string, bids <-chan simulationBid) error {
	lotNumber, err := getLotNumber(externalID)
	if err != nil {
		return errors.Wrap(err, "failed to get lot number")
	}

	log = log.WithField("lot_number", lotNumber)

	log.Debug("sending current item")
	startBid := rand.Int63n(200) * 50 * 100
	r.notifyAll(r.ctx, Message{
		RequestID: 0,
		Type:      MessageTypeEventReceive,
		Event: &Event{
			SaleID: saleID,
			Type:   EventTypeCurrentItem,
			CurrentItem: &CurrentItem{
				LotNumber: lotNumber,
				StartBid:  Amount(startBid),
			},
		},
	})

	highestBid := r.notifyNewBids(log, simulationBid{
		BuyerNumber: rand.Int63(),
		LotID:       lotNumber,
		AmountCents: startBid - bidStep,
		SaleID:      saleID,
	}, bids)

	log.Debug("sending sold pending")
	r.notifyAll(r.ctx, Message{
		Type: MessageTypeEventReceive,
		Event: &Event{
			SaleID: saleID,
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
	r.notifyAll(r.ctx, Message{
		Type: MessageTypeEventReceive,
		Event: &Event{
			SaleID: saleID,
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

const bidStep = int64(50 * 100)

func (r *liveBidderSimulator) notifyNewBids(log *logan.Entry, highestBid simulationBid, bids <-chan simulationBid) simulationBid {
	timer := time.NewTimer(DefaultWaitTime)

	bonusTimePassed := false
	for {
		select {
		case <-timer.C:
			if bonusTimePassed {
				return highestBid
			}
			log.Debug("sending bonus time")
			r.notifyAll(r.ctx, Message{
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

		case newBid := <-bids:
			if newBid.LotID != highestBid.LotID && newBid.SaleID != highestBid.SaleID {
				log.WithField("new_bid", newBid).WithField("highest_bid", highestBid).Warn("received bid not matching for the sale")
				continue
			}

			if highestBid.AmountCents >= newBid.AmountCents {
				log.WithField("new_bid", newBid).WithField("highest_bid", highestBid).Warn("higher bid already present")
				continue
			}

			bonusTimePassed = false
			highestBid = newBid

			r.notifyAll(r.ctx, Message{
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

		log.Debug("reseting wait time timer")
		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}

		timer.Reset(DefaultWaitTime)
	}
}*/
