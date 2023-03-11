package kafka

import (
	"fmt"

	"github.com/go-errors/errors"
)

func (m LeasingState) Event() LeasingEvent {
	return LeasingEvent{
		Type:  LeasingEvent_TYPE_STATE,
		State: &m,
	}
}

func (m LeasingEvent) Message() (msg LeasingEventMessage) {
	msg.Value = &m
	switch m.Type {
	case LeasingEvent_TYPE_STATE:
		leasing := m.State.Leasing
		msg.msg.Key = []byte(fmt.Sprintf("%d:%s", leasing.State, leasing.Id))
	default:
		panic(errors.Errorf("unexpected event type: %d", m.Type))
	}
	return
}
