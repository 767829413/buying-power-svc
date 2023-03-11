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

// ExternalLotEventMessage wraps original kafka-go message and
// parsed protobuf generated ExternalLotEvent.
type ExternalLotEventMessage struct {
	// raw is set directly from kafka event and left private intentionally.
	// it's advisable to implement adapter methods for both producer and consumer
	// to have key serialization semantics available in one place.
	// helpers might look something like:
	//
	// func (m ExternalLotEventMessage) Key() uint64 {
	//     key, err := proto.DecodeVarint(m.rawkey)
	//     if err != nil {
	//         panic(errors.Wrap(err, "failed to decode event key"))
	//     }
	//     return key
	// }
	//
	// func (m ExternalLotEvent) Message(key uint64) ExternalLotEventMessage {
	// return ExternalLotEventMessage{
	//        Key: proto.EncodeVarint(key),
	//        Value: &m,
	//     }
	// }
	//
	// it also prevents kafka-go (implementation detail) to leak outside of package
	Message
	Value *ExternalLotEvent
}

//ExternalLotEventMessage - unmarshal msg value into ExternalLotEventMessage.
func (m Message) ExternalLotEventMessage() (ExternalLotEventMessage, error) {
	var event ExternalLotEvent
	err := proto.Unmarshal(m.msg.Value, &event)
	if err != nil {
		return ExternalLotEventMessage{}, errors.Wrap(err, "failed to unmarshal msg value into ExternalLotEvent")
	}

	return ExternalLotEventMessage{
		Message: m,
		Value:   &event,
	}, nil
}

//MustExternalLotEventMessage - unmarshal msg value into ExternalLotEventMessage. Panics if unmarshal fails
func (m Message) MustExternalLotEventMessage() ExternalLotEventMessage {
	result, err := m.ExternalLotEventMessage()
	if err != nil {
		panic(err)
	}

	return result
}

// DEPRECATED
type ExternalLotEventTopic struct {
	cfg Config
}

func (t ExternalLotEventTopic) Reader() ExternalLotEventReader {
	return NewExternalLotEventReader(t.cfg.reader(Topic_external_lots.String()))
}

func (t ExternalLotEventTopic) Writer() ExternalLotEventWriter {
	return NewExternalLotEventWriter(t.cfg.writer(Topic_external_lots.String()))
}

// DEPRECATED
func NewExternalLotEventTopic(cfg Config) ExternalLotEventTopic {
	return ExternalLotEventTopic{
		cfg,
	}
}

// ExternalLotEventReader reads LotUpdatedEvent from kafka.
// DEPERCATED: use ingester
type ExternalLotEventReader ProtoReader

// NewExternalLotEventReader inits an instance of ExternalLotEventReader.
// DEPRECATED
func NewExternalLotEventReader(cfg ReaderConfig) ExternalLotEventReader {
	return ExternalLotEventReader(NewProtoReader(cfg))
}

// WithAutoReset returns copy of ExternalLotEventReader that in case of ReadMessage error will reset its offset to last commited.
// DEPRECATED: NOT THREAD SAFE & LEAKS RESOURCES
func (r ExternalLotEventReader) WithAutoReset() ExternalLotEventReader {
	r.autoReset = true
	return r
}

func (r ExternalLotEventReader) readMessage(ctx context.Context, reader protoReaderFunc) (ExternalLotEventMessage, error) {
	e := new(ExternalLotEvent)
	msg, err := reader(ctx, e)
	if err != nil {
		if r.autoReset {
			r.Reset()
		}
		return ExternalLotEventMessage{}, err
	}
	return ExternalLotEventMessage{Message: Message{msg: msg.raw}, Value: e}, nil
}

// DEPRECATED: use FetchMessage/CommitMessage instead
func (r ExternalLotEventReader) ReadMessage(ctx context.Context) (ExternalLotEventMessage, error) {
	return r.readMessage(ctx, ProtoReader(r).ReadMessage)
}

// DEPRECATED
func (r ExternalLotEventReader) FetchMessage(ctx context.Context) (ExternalLotEventMessage, error) {
	return r.readMessage(ctx, ProtoReader(r).FetchMessage)
}

// DEPRECATED
func (r ExternalLotEventReader) CommitMessages(ctx context.Context, msgs ...ExternalLotEventMessage) error {
	protoMsgs := make([]ProtoMessage, 0, len(msgs))
	for _, msg := range msgs {
		protoMsgs = append(protoMsgs, ProtoMessage{raw: msg.msg, Value: msg.Value})
	}
	return ProtoReader(r).CommitMessages(ctx, protoMsgs...)
}

// Reset defaults underlying kafka.Reader
// DEPRECATED: leaks resources
func (r ExternalLotEventReader) Reset() {
	*r.Reader = *kafka.NewReader(r.Reader.Config())
}

// ExternalLotEventWriter writes ExternalLotEvent to kafka.
type ExternalLotEventWriter ProtoWriter

// NewExternalLotEventWriter inits an instance of ExternalLotEventWriter.
func NewExternalLotEventWriter(cfg WriterConfig) ExternalLotEventWriter {
	return ExternalLotEventWriter(NewProtoWriter(cfg))
}

// WriteMessages to kafka.
func (w ExternalLotEventWriter) WriteMessages(ctx context.Context, msgs ...ExternalLotEventMessage) error {
	protoMsgs := make([]ProtoMessage, len(msgs))
	for idx, msg := range msgs {
		protoMsgs[idx] = ProtoMessage{raw: msg.msg, Value: msg.Value}
	}
	return ProtoWriter(w).WriteMessages(ctx, protoMsgs...)
}

// DEPRECATED
func (cfg Config) ExternalLotEventTopic() ExternalLotEventTopic {
	return NewExternalLotEventTopic(cfg)
}
