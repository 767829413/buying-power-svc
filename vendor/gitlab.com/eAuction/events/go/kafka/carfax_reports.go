package kafka

func (m CarfaxReportFulfilled) Event() CarfaxReportEvent {
	return CarfaxReportEvent{
		Type:      CarfaxReportEvent_TYPE_FULFILLED,
		Fulfilled: &m,
	}
}

func (e CarfaxReportEvent) Message(id string) (msg CarfaxReportEventMessage) {
	msg.msg.Key = []byte(id)
	msg.Value = &e

	return msg
}