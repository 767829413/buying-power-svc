package copart

import (
	"encoding/base64"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"strings"
)

//EventType - defines type of the event
type EventType string

const (
	EventTypeUnknow EventType = "UNKNOWN"
	//EventTypeNewBid - send when new bid recevied
	EventTypeNewBid EventType = "BIDREC"
	//EventTypeAuctioneer - signals that count down for live auction's end for specified lot started
	EventTypeAuctioneer EventType = "AUCTIONEER"
	//EventTypePreBid - signals that someone created pre live auction bid for lot of the sale
	EventTypePreBid EventType = "PREBID"
	//EventTypeSoldPending -  signals that lot winner has been selected, but it's final
	EventTypeSoldPending EventType = "SOLDPEND"
	//EventTypeSold - signals that lot winner has been finalized
	EventTypeSold EventType = "SOLD"
	//EventTypeCurrentItem - signals next item on the lane
	EventTypeCurrentItem EventType = "CURITEM"
	//EventTypeStartItem - not sure
	EventTypeStartItem  EventType = "STARTITEM"
	//EventTypeEndAuction - signals that lane has ended
	EventTypeEndAuction EventType = "ENDAUC"
)

var allEventTypes = []EventType{
	EventTypeUnknow,
	EventTypeNewBid,
	EventTypeAuctioneer,
	EventTypePreBid,
	EventTypeSoldPending,
	EventTypeSold,
	EventTypeCurrentItem,
	EventTypeStartItem,
	EventTypeEndAuction,
}

func (e *EventType) UnmarshalJSON(raw []byte) error {
	str, err := base64.StdEncoding.DecodeString(strings.Trim(string(raw), `"`))
	if err != nil {
		return errors.Wrap(err, "failed to decode base64 event type")
	}

	*e = EventType(str)
	for _, eventType := range allEventTypes {
		if eventType == *e {
			return nil
		}
	}

	*e = EventTypeUnknow
	return nil
}
