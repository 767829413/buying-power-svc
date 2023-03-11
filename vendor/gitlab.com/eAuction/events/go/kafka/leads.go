package kafka

func (m LeadEvent) Message() LeadEventMessage {
	var msg LeadEventMessage
	msg.msg.Key = []byte(m.Id)
	msg.Value = &m

	return msg
}
