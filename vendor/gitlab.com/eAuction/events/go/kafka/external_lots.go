package kafka

import (
	"github.com/pkg/errors"
)

func (m ExternalLotClosed) Event() ExternalLotEvent {
	return ExternalLotEvent{
		Type:   ExternalLotEvent_TYPE_CLOSED,
		Closed: &m,
	}
}

func (m ExternalLotCreated) Event() ExternalLotEvent {
	return ExternalLotEvent{
		Type:    ExternalLotEvent_TYPE_CREATED,
		Created: &m,
	}
}

func (m ExternalLotCreatedV2) Event() ExternalLotEvent {
	return ExternalLotEvent{
		Type:      ExternalLotEvent_TYPE_CREATED_V2,
		CreatedV2: &m,
	}
}

func (m ExternalLotBidUpdated) Event() ExternalLotEvent {
	return ExternalLotEvent{
		Type:       ExternalLotEvent_TYPE_BID_UPDATED,
		BidUpdated: &m,
	}
}

func (m ExternalLotLiveBiddingDetails) Event() ExternalLotEvent {
	return ExternalLotEvent{
		Type:               ExternalLotEvent_TYPE_LIVE_BIDDING_DETAILS,
		LiveBiddingDetails: &m,
	}
}

func (m ExternalLotEvent) Message() (msg ExternalLotEventMessage) {
	var lotID string
	switch m.Type {
	case ExternalLotEvent_TYPE_CREATED:
		lotID = m.Created.Lot.Id
	case ExternalLotEvent_TYPE_CREATED_V2:
		lotID = m.CreatedV2.Lot.Id
	case ExternalLotEvent_TYPE_BID_UPDATED:
		lotID = m.BidUpdated.Id
	case ExternalLotEvent_TYPE_CLOSED:
		lotID = m.Closed.Id
	case ExternalLotEvent_TYPE_LIVE_BIDDING_DETAILS:
		lotID = m.LiveBiddingDetails.Id
	default:
		panic(errors.Errorf("unknown ExternalLotEvent type: %d", m.Type))
	}

	msg.msg.Key = []byte(lotID)
	msg.Value = &m

	return msg
}
