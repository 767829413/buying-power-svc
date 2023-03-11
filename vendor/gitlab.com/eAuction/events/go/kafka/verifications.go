package kafka

import (
	"fmt"

	"github.com/segmentio/kafka-go"
)

func (m Verification_Phone) Event() VerificationEvent {
	return VerificationEvent{
		Type:  Verification_EVENT_TYPE_PHONE,
		Phone: &m,
	}
}

func (m Verification_Email) Event() VerificationEvent {
	return VerificationEvent{
		Type:  Verification_EVENT_TYPE_EMAIL,
		Email: &m,
	}
}

func (e VerificationEvent) Message() VerificationEventMessage {
	key := ""
	switch e.Type {
	case Verification_EVENT_TYPE_PHONE:
		key = "phone:" + e.Phone.AccountAddress
	case Verification_EVENT_TYPE_EMAIL:
		key = "email:" + e.Email.AccountAddress
	default:
		panic(fmt.Sprintf("unexpected event type %s", e.Type))
	}
	return VerificationEventMessage{
		Message: Message{msg: kafka.Message{Key: []byte(key)}},
		Value:   &e,
	}
}
