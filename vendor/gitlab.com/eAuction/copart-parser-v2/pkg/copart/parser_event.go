package copart

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

//ParserEventType - defines event types from parser
type ParserEventType string

const (
	ParserEventTypeSessionStarted    ParserEventType = "session_started"
	ParserEventTypeSessionFailed     ParserEventType = "session_failed"
	ParserEventTypeJoinedSale        ParserEventType = "joined_sale"
	ParserEventTypeJoinSaleFailed    ParserEventType = "join_sale_failed"
	ParserEventTypeSaleEnded         ParserEventType = "sale_ended"
	ParserEventTypeRawWSFrame        ParserEventType = "raw_ws_frame"
	ParserEventTypeBidAck            ParserEventType = "bid_ack"
	ParserEventTypeBuyerNumberIssued ParserEventType = "buyer_number_issued"
)

//ParserEvent - defines events from copart parser
type ParserEvent struct {
	Type ParserEventType `json:"event"`
	Data json.RawMessage `json:"data"`
}

func (p ParserEvent) mustParse(eventType ParserEventType, payload interface{}) {
	if p.Type != eventType {
		panic(errors.Errorf("trying to get %s on event type %s", eventType, p.Type))
	}

	err := json.Unmarshal(p.Data, payload)
	if err != nil {
		panic(errors.Wrap(err, "failed to parse data", logan.F{
			"type": eventType,
			"data": string(p.Data),
		}))
	}
}

func (p ParserEvent) SessionStarted() ParserEventSessionStarted {
	var result ParserEventSessionStarted
	p.mustParse(ParserEventTypeSessionStarted, &result)
	return result
}

func (p ParserEvent) SessionFailed() ParserEventSessionFailed {
	var result ParserEventSessionFailed
	p.mustParse(ParserEventTypeSessionFailed, &result)
	return result
}

func (p ParserEvent) JoinedSale() ParserEventSale {
	var result ParserEventSale
	p.mustParse(ParserEventTypeJoinedSale, &result)
	return result
}

func (p ParserEvent) JoinSaleFailed() ParserEventSale {
	var result ParserEventSale
	p.mustParse(ParserEventTypeJoinSaleFailed, &result)
	return result
}

func (p ParserEvent) SaleEnded() ParserEventSale {
	var result ParserEventSale
	p.mustParse(ParserEventTypeSaleEnded, &result)
	return result
}

func (p ParserEvent) RawWSMessage() RawWSFrame {
	var result RawWSFrame
	p.mustParse(ParserEventTypeRawWSFrame, &result)
	return result
}

func (p ParserEvent) BidAck() ParserEventBidAck {
	var result ParserEventBidAck
	p.mustParse(ParserEventTypeBidAck, &result)
	return result
}

func (p ParserEvent) BuyerNumberIssued() ParserEventBuyerNumberIssued {
	var result ParserEventBuyerNumberIssued
	p.mustParse(ParserEventTypeBuyerNumberIssued, &result)
	return result
}

type ParserEventSessionStarted struct {
	ID string `json:"id"`
}

type ParserEventSessionFailed struct {
	Error string `json:"error"`
}

type ParserEventSale struct {
	SaleID string `json:"id"`
}

type ParserEventBidAck struct {
	EndToEndID    string    `json:"e2e_id"`
	ActualLotID   LotNumber `json:"actual_lot_id"`
	ActualBid     Amount    `json:"actual_bid"`
	OutBid        bool      `json:"out_bid"`
	LotIDMismatch bool      `json:"lot_id_mismatch"`
	Error         string    `json:"error"`
}

type RawWSFrame struct {
	SaleID string `json:"sale_id"`
	Frame  string `json:"frame"`
}

func (r RawWSFrame) Messages() []Message {
	var result []Message
	err := json.Unmarshal([]byte(r.Frame), &result)
	if err != nil {
		panic(errors.Wrap(err, "failed to unmarshal frame", logan.F{
			"raw": r.Frame,
		}))
	}

	return result
}

type ParserEventBuyerNumberIssued struct {
	BuyerNumber string `json:"buyer_number"`
}
