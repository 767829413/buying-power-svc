package copart

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"strconv"
	"strings"
	"time"
)

const DefaultWaitTime = time.Second*9 + time.Second/2

//Amount - defines copart amount
type Amount int64

func (a *Amount) UnmarshalJSON(rawData []byte) error {
	data := strings.TrimSpace(strings.Trim(string(rawData), `"`))
	if data == "" {
		return nil
	}

	result, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return errors.Wrap(err, "failed to parse bid amount")
	}

	*a = Amount(result * 1e2)
	return nil
}

type LotNumber string

func (n *LotNumber) UnmarshalJSON(rawData []byte) error {
	data := strings.TrimSpace(strings.Trim(string(rawData), `"`))
	data = strings.TrimLeft(string(data), "0")
	*n = LotNumber(data)
	return nil
}

//NewBid - signal new bid has been performed
type NewBid struct {
	CurrentBidCents Amount    `json:"CURBID"`
	LotNumber       LotNumber `json:"LOTNO"`
	NextBidCents    Amount    `json:"NEXT"`
	BuyerNumber     string    `json:"BUYERNO"`
	Type            string    `json:"TYPE"`
}

type CopartDuration time.Duration

func (d *CopartDuration) UnmarshalJSON(data []byte) error {
	strData := strings.TrimSpace(strings.Trim(string(data), `"`))
	if strData == "" {
		*d = CopartDuration(DefaultWaitTime)
		return nil
	}

	durSec, err := strconv.ParseInt(strData, 10, 64)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal copart duration")
	}

	*d = CopartDuration(time.Duration(durSec) * time.Second)
	return nil
}

//Auctioneer - defines event issued by auctioneer
type Auctioneer struct {
	LotNumber LotNumber      `json:"LOTNUMBER"`
	BonusTime CopartDuration `json:"BONUSTIME"`
}

//Sold - event issued when bidding for lot has been finished
type Sold struct {
	LotNumber   LotNumber `json:"LOTNO"`
	BuyerNumber string    `json:"BUYERNO"`
	Bid         Amount    `json:"BID"`
}

//CurrentItem - event issued when next item in the lane is ready for bidding
type CurrentItem struct {
	LotNumber LotNumber `json:"LOTNO"`
	StartBid  Amount    `json:"STARTBID"`
}

type StartItem struct {
	StartBid    Amount    `json:"STARTAMT"`
	LotNumber   LotNumber `json:"LOTNO"`
	BuyerNumber string    `json:"BUYERNO"`
	NextBid     Amount    `json:"NEXT"`
}

type PreBid struct {
	BuyerNumber string    `json:"BUYERNO"`
	PreBid      Amount    `json:"PREBID"`
	LotNumber   LotNumber `json:"LOTNO"`
}

// Event - defines structure of event receive msg.
type Event struct {
	SaleID string
	Type   EventType

	NewBid      *NewBid      `json:"new_bid,omitempty"`
	Auctioneer  *Auctioneer  `json:"auctioneer,omitempty"`
	Sold        *Sold        `json:"sold,omitempty"`
	CurrentItem *CurrentItem `json:"current_item,omitempty"`
	StartItem   *StartItem   `json:"start_item,omitempty"`
	PreBid      *PreBid      `json:"pre_bid"`
	// set only for Unknown event type to be able to log it
	RawPayload string `json:"raw_payload,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *Event) UnmarshalJSON(data []byte) error {
	b := bytes.Trim(data, `"`)

	decodedData, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		return errors.Wrap(err, "failed to decode base64")
	}

	var payload interface{}
	switch e.Type {
	case EventTypeNewBid:
		e.NewBid = new(NewBid)
		payload = e.NewBid
	case EventTypeAuctioneer:
		e.Auctioneer = new(Auctioneer)
		payload = e.Auctioneer
	case EventTypePreBid:
		e.PreBid = new(PreBid)
		payload= e.PreBid
	case EventTypeSoldPending, EventTypeSold:
		e.Sold = new(Sold)
		payload = e.Sold
	case EventTypeCurrentItem:
		e.CurrentItem = new(CurrentItem)
		payload = e.CurrentItem
	case EventTypeStartItem:
		e.StartItem = new(StartItem)
		payload = e.StartItem
	case EventTypeEndAuction:
		return nil
	default:
		return errors.From(errors.New("unexpected event type"), logan.F{
			"event_type": e.Type,
		})
	}

	err = json.Unmarshal(decodedData, payload)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal decoded data to json")
	}

	return nil
}
