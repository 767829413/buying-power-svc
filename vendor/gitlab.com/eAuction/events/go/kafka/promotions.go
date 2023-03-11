package kafka

import (
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

func (m PromotionEvent) Message() PromotionEventMessage {
	return PromotionEventMessage{
		Message:   Message{msg:kafka.Message{Key: proto.EncodeVarint(uint64(m.Promotion.Id))}},
		Value: &m,
	}
}

func (m PromotionEventMessage) Key() uint64 {
	key, n := proto.DecodeVarint(m.msg.Key)
	if n == 0 {
		panic(errors.New("failed to decode PromotionEventMessage key"))
	}

	return key
}

func (m Promotion) event(t PromotionEvent_Type) PromotionEvent {
	return PromotionEvent{
		Type:      t,
		Promotion: &m,
	}
}

func (m Promotion) Started() PromotionEvent {
	return m.event(PromotionEvent_STARTED)
}

func (m Promotion) Finished() PromotionEvent {
	return m.event(PromotionEvent_FINISHED)
}

func (m Promotion) Declined() PromotionEvent {
	return m.event(PromotionEvent_DECLINED)
}
