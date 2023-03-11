package kafka

import (
	"github.com/pkg/errors"
)

func (m DepositV2) Event() DepositEvent {
	return DepositEvent{
		Type:    DepositEvent_EVENT_TYPE_DEPOSIT,
		Deposit: &m,
	}
}

func (m Withdrawal) Event() DepositEvent {
	return DepositEvent{
		Type:       DepositEvent_EVENT_TYPE_WITHDRAWAL,
		Withdrawal: &m,
	}
}

func (m Movement) Event() DepositEvent {
	return DepositEvent{
		Type:     DepositEvent_EVENT_TYPE_MOVEMENT,
		Movement: &m,
	}
}

func (m DepositEvent) Message() (msg DepositEventMessage) {
	switch m.Type {
	case DepositEvent_EVENT_TYPE_DEPOSIT:
		msg.msg.Key = []byte(m.Deposit.Id)
	case DepositEvent_EVENT_TYPE_WITHDRAWAL:
		msg.msg.Key = []byte(m.Withdrawal.Id)
	case DepositEvent_EVENT_TYPE_MOVEMENT:
		msg.msg.Key = []byte(m.Movement.Id)
	default:
		panic(errors.New("unexpected deposit event type"))
	}

	msg.Value = &m
	return msg
}
