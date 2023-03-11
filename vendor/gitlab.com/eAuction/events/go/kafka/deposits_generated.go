// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package kafka

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

// DepositEventMessage wraps original kafka-go message and
// parsed protobuf generated DepositEvent.
type DepositEventMessage struct {
	// raw is set directly from kafka event and left private intentionally.
	// it's advisable to implement adapter methods for both producer and consumer
	// to have key serialization semantics available in one place.
	// helpers might look something like:
	//
	// func (m DepositEventMessage) Key() uint64 {
	//     key, err := proto.DecodeVarint(m.rawkey)
	//     if err != nil {
	//         panic(errors.Wrap(err, "failed to decode event key"))
	//     }
	//     return key
	// }
	//
	// func (m DepositEvent) Message(key uint64) DepositEventMessage {
	// return DepositEventMessage{
	//        Key: proto.EncodeVarint(key),
	//        Value: &m,
	//     }
	// }
	//
	// it also prevents kafka-go (implementation detail) to leak outside of package
	Message
	Value *DepositEvent
}

//DepositEventMessage - unmarshal msg value into DepositEventMessage.
func (m Message) DepositEventMessage() (DepositEventMessage, error) {
	var event DepositEvent
	err := proto.Unmarshal(m.msg.Value, &event)
	if err != nil {
		return DepositEventMessage{}, errors.Wrap(err, "failed to unmarshal msg value into DepositEvent")
	}

	return DepositEventMessage{
		Message: m,
		Value:   &event,
	}, nil
}

//MustDepositEventMessage - unmarshal msg value into DepositEventMessage. Panics if unmarshal fails
func (m Message) MustDepositEventMessage() DepositEventMessage {
	result, err := m.DepositEventMessage()
	if err != nil {
		panic(err)
	}

	return result
}

// DEPRECATED
type DepositEventTopic struct {
	cfg Config
}

func (t DepositEventTopic) Reader() DepositEventReader {
	return NewDepositEventReader(t.cfg.reader(Topic_deposits.String()))
}

func (t DepositEventTopic) Writer() DepositEventWriter {
	return NewDepositEventWriter(t.cfg.writer(Topic_deposits.String()))
}

// DEPRECATED
func NewDepositEventTopic(cfg Config) DepositEventTopic {
	return DepositEventTopic{
		cfg,
	}
}

// DepositEventReader reads LotUpdatedEvent from kafka.
// DEPERCATED: use ingester
type DepositEventReader ProtoReader

// NewDepositEventReader inits an instance of DepositEventReader.
// DEPRECATED
func NewDepositEventReader(cfg ReaderConfig) DepositEventReader {
	return DepositEventReader(NewProtoReader(cfg))
}

// WithAutoReset returns copy of DepositEventReader that in case of ReadMessage error will reset its offset to last commited.
// DEPRECATED: NOT THREAD SAFE & LEAKS RESOURCES
func (r DepositEventReader) WithAutoReset() DepositEventReader {
	r.autoReset = true
	return r
}

func (r DepositEventReader) readMessage(ctx context.Context, reader protoReaderFunc) (DepositEventMessage, error) {
	e := new(DepositEvent)
	msg, err := reader(ctx, e)
	if err != nil {
		if r.autoReset {
			r.Reset()
		}
		return DepositEventMessage{}, err
	}
	return DepositEventMessage{Message: Message{msg: msg.raw}, Value: e}, nil
}

// DEPRECATED: use FetchMessage/CommitMessage instead
func (r DepositEventReader) ReadMessage(ctx context.Context) (DepositEventMessage, error) {
	return r.readMessage(ctx, ProtoReader(r).ReadMessage)
}

// DEPRECATED
func (r DepositEventReader) FetchMessage(ctx context.Context) (DepositEventMessage, error) {
	return r.readMessage(ctx, ProtoReader(r).FetchMessage)
}

// DEPRECATED
func (r DepositEventReader) CommitMessages(ctx context.Context, msgs ...DepositEventMessage) error {
	protoMsgs := make([]ProtoMessage, 0, len(msgs))
	for _, msg := range msgs {
		protoMsgs = append(protoMsgs, ProtoMessage{raw: msg.msg, Value: msg.Value})
	}
	return ProtoReader(r).CommitMessages(ctx, protoMsgs...)
}

// Reset defaults underlying kafka.Reader
// DEPRECATED: leaks resources
func (r DepositEventReader) Reset() {
	*r.Reader = *kafka.NewReader(r.Reader.Config())
}

// DepositEventWriter writes DepositEvent to kafka.
type DepositEventWriter ProtoWriter

// NewDepositEventWriter inits an instance of DepositEventWriter.
func NewDepositEventWriter(cfg WriterConfig) DepositEventWriter {
	return DepositEventWriter(NewProtoWriter(cfg))
}

// WriteMessages to kafka.
func (w DepositEventWriter) WriteMessages(ctx context.Context, msgs ...DepositEventMessage) error {
	protoMsgs := make([]ProtoMessage, len(msgs))
	for idx, msg := range msgs {
		protoMsgs[idx] = ProtoMessage{raw: msg.msg, Value: msg.Value}
	}
	return ProtoWriter(w).WriteMessages(ctx, protoMsgs...)
}

// DEPRECATED
func (cfg Config) DepositEventTopic() DepositEventTopic {
	return NewDepositEventTopic(cfg)
}