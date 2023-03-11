package kafka

func (m DialogEvent) Message() DialogEventMessage {
	var msg DialogEventMessage
	msg.msg.Key = []byte(m.Id)
	msg.Value = &m

	return msg
}
