package kafka

import (
	"github.com/segmentio/kafka-go"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (m CallbackCreated) Event() CallbackEvent {
	return CallbackEvent{
		Type:    CallbackEvent_CREATED,
		Created: &m,
	}
}

func (m CallbackEvent) Message() *CallbackEventMessage {
	var key string
	switch m.Type {
	case CallbackEvent_CREATED:
		key = m.Created.User + ":" + m.Created.User
	default:
		panic(errors.Errorf("unknown CallbackEvent Type %s", m.Type.String()))
	}

	return &CallbackEventMessage{
		Message: Message{msg: kafka.Message{Key: []byte(key)}},
		Value:   &m,
	}
}
