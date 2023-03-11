package kafka

import (
	"github.com/pkg/errors"
)

func (m MessengerCommand_BrokersUpdated) Event() MessengerCommand {
	return MessengerCommand{
		Type:           MessengerCommand_TYPE_BROKERS_UPDATED,
		BrokersUpdated: &m,
	}
}

func (m MessengerCommand) Message() MessengerCommandMessage {
	var dialogID string
	switch m.Type {
	case MessengerCommand_TYPE_BROKERS_UPDATED:
		dialogID = m.BrokersUpdated.DialogId
	default:
		panic(errors.Errorf("Unknown MessengerCommand type %s", m.Type.String()))
	}

	var msg MessengerCommandMessage
	msg.msg.Key = []byte(dialogID)
	msg.Value = &m

	return msg
}
