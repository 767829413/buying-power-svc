package copart

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

// Message is a struct returned by copart from ws.
type Message struct {
	RequestID uint64

	Type           MessageType
	Event          *Event          `json:"event,omitempty"`
	BatchSubscribe *BatchSubscribe `json:"batch_subscribe,omitempty"`
	SessionStart   *SessionStart   `json:"session_start,omitempty"`
	Bid            *Bid            `json:"bid,omitempty"`
}

// MessageType - defines message types received/sent by copart.
type MessageType int

const (
	MessageTypeStartSession        MessageType = 1
	MessageTypeResourceSubscribe   MessageType = 2
	MessageTypeKeepAlive           MessageType = 3
	MessageTypeBatchSubscribe      MessageType = 5
	MessageTypePacketReceive       MessageType = 7
	MessageTypeEventReceive        MessageType = 8
	MessageTypeSetResourceAlias    MessageType = 10
	MessageTypeResourcePublish     MessageType = 14
	MessageTypeResourceUnsubscribe MessageType = 15
	MessageTypeQueueCommit         MessageType = 16
	MessageTypeQueueRollback       MessageType = 17
	MessageTypeTxPublish           MessageType = 20
	MessageTypeTxIsCommitted       MessageType = 24
	MessageTypeClientClose         MessageType = 30
)

// SessionStart - defines copart response for session start msg.
type SessionStart struct {
	SessionID string
	IsSuccess bool
}

//Bid - defines copart response for bid
type Bid struct {
	IsSuccess bool
	Details   []string
}

func (b Bid) ToError() error {
	if b.IsSuccess {
		return nil
	}

	return errors.From(errors.New("unexpected error"), logan.F{
		"details": b.Details,
	})
}

// BatchSubscribe - response for copart batch subscribe.
type BatchSubscribe struct {
	IsSuccess     bool
	ResourceName  string
	ResourceAlias string
}

//SaleID - returns sale ID
func (b BatchSubscribe) SaleID() string {
	saleID := b.ResourceName
	saleID = strings.Trim(saleID, "outbound")
	saleID = strings.Trim(saleID, "/")
	return saleID
}

func (b BatchSubscribe) ToError() error {
	if b.IsSuccess {
		return nil
	}

	if strings.Contains(strings.ToLower(b.ResourceAlias), "channel not found") {
		return ErrSaleNotFound
	}

	return errors.From(errors.New("received unknown error when trying to subscribe to sale"), logan.F{
		"alias": b.ResourceAlias,
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (m *Message) UnmarshalJSON(b []byte) error {
	type rawResponse struct {
		RequestID uint64            `json:"r"`
		Data      []json.RawMessage `json:"d"`
	}

	var raw rawResponse
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	if len(raw.Data) == 0 {
		return errors.From(errors.New("no data in response"), logan.F{"rawData": string(b)})
	}

	messageType, err := strconv.Atoi(string(raw.Data[0]))
	if err != nil {
		return errors.New("failed to parse message type")
	}

	m.RequestID = raw.RequestID
	m.Type = MessageType(messageType)

	return errors.Wrap(m.unmarshalData(raw.Data), "failed to unmarshal data", logan.F{
		"raw_data": string(b),
	})
}

func (m *Message) unmarshalData(data []json.RawMessage) error {
	switch m.Type {
	case MessageTypeStartSession:
		if len(data) < 3 {
			return errors.New("unexpected number of data elements of start session")
		}

		var isSuccess bool
		err := json.Unmarshal(data[1], &isSuccess)
		if err != nil {
			return errors.Wrap(err,
				"failed to unmarshal isSuccessful flag of start session",
				logan.F{"isSuccess": data[1]})
		}

		m.SessionStart = &SessionStart{
			IsSuccess: isSuccess,
			SessionID: string(data[2]),
		}
	case MessageTypeEventReceive:
		return m.unmarshalEventReceived(data)
	case MessageTypeBatchSubscribe:
		if len(data) != 2 {
			return errors.New("unexpected number of elements for batch subscribe")
		}

		rawData := []interface{}{"", false, ""}
		err := json.Unmarshal(data[1], &rawData)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal subscribe")
		}

		success, ok := rawData[1].(bool)
		if !ok {
			return errors.New("unexpected type of second data element of batch subscribe: expected bool")
		}

		m.BatchSubscribe = &BatchSubscribe{
			IsSuccess:     success,
			ResourceName:  rawData[0].(string),
			ResourceAlias: rawData[2].(string),
		}
	case MessageTypeResourcePublish:
		if len(data) < 2 {
			return errors.New("unexpected number of elements for bid")
		}

		var isSuccess bool
		err := json.Unmarshal(data[1], &isSuccess)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal bid is success")
		}

		var dataAsString []string
		for _, d := range data {
			dataAsString = append(dataAsString, string(d))
		}

		m.Bid = &Bid{
			IsSuccess: isSuccess,
			Details:   dataAsString,
		}
	}

	return nil
}

func (m *Message) unmarshalEventReceived(data []json.RawMessage) error {
	if len(data) != 2 {
		return errors.New("unexpected number of elements of data for event receive type")
	}

	var eventReceive struct {
		Event      json.RawMessage `json:"Data"`
		EventType  EventType       `json:"Tag"`
		Dictionary struct {
			Sale []json.RawMessage `json:"sale"`
		} `json:"Dictionary"`
	}

	if err := json.Unmarshal(data[1], &eventReceive); err != nil {
		return errors.Wrap(err, "failed to unmarshal event receive data")
	}

	if len(eventReceive.Dictionary.Sale) != 2 {
		return errors.New("unexpected number of elements of Dictionary->sale")
	}

	m.Event = &Event{
		SaleID: strings.Trim(string(bytes.Trim(eventReceive.Dictionary.Sale[1], `"`)), "/"),
		Type:   eventReceive.EventType,
	}

	if m.Event.Type == EventTypeUnknow {
		m.Event.RawPayload = string(eventReceive.Event)
		return nil
	}

	err := json.Unmarshal(eventReceive.Event, m.Event)
	return errors.Wrap(err, "failed to unmarshal event")
}
