package kafka

func (m RecoveryRequestCreated) Event() RecoveryRequestEvent {
	return RecoveryRequestEvent{
		Type:    RecoveryRequestEvent_TYPE_CREATED,
		Created: &m,
	}
}

func (e RecoveryRequestEvent) Message(id string) (msg RecoveryRequestEventMessage) {
	msg.msg.Key = []byte(id)
	msg.Value = &e
	return msg
}