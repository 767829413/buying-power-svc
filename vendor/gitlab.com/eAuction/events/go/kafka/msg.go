package kafka

import (
	"time"

	"github.com/segmentio/kafka-go"
)

//Message - represent general kafka message
type Message struct {
	msg kafka.Message
}

//Offset - returns position of the message in topic
func (m Message) Offset() int64 {
	return m.msg.Offset
}

func (m Message)Partition() int {
	return m.msg.Partition
}

//Topic - returns topic of the message
func (m Message) Topic() Topic {
	return Topic(Topic_value[m.msg.Topic])
}

//Time - returns time specified by producer of msg or time, when msg was added into topic
func (m Message) Time() time.Time {
	return m.msg.Time
}

// SetTime - sets specified time as time of the message.
func (m *Message) SetTime(t time.Time) {
	m.msg.Time = t
}

//Key - returns key of the message
func (m Message) Key() string {
	return string(m.msg.Key)
}
