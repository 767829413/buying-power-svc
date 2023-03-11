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

// AuctionEventMessage wraps original kafka-go message and
// parsed protobuf generated AuctionEvent.
type AuctionEventMessage struct {
	// raw is set directly from kafka event and left private intentionally.
	// it's advisable to implement adapter methods for both producer and consumer
	// to have key serialization semantics available in one place.
	// helpers might look something like:
	//
	// func (m AuctionEventMessage) Key() uint64 {
	//     key, err := proto.DecodeVarint(m.rawkey)
	//     if err != nil {
	//         panic(errors.Wrap(err, "failed to decode event key"))
	//     }
	//     return key
	// }
	//
	// func (m AuctionEvent) Message(key uint64) AuctionEventMessage {
	// return AuctionEventMessage{
	//        Key: proto.EncodeVarint(key),
	//        Value: &m,
	//     }
	// }
	//
	// it also prevents kafka-go (implementation detail) to leak outside of package
	Message
	Value *AuctionEvent
}

//AuctionEventMessage - unmarshal msg value into AuctionEventMessage.
func (m Message) AuctionEventMessage() (AuctionEventMessage, error) {
	var event AuctionEvent
	err := proto.Unmarshal(m.msg.Value, &event)
	if err != nil {
		return AuctionEventMessage{}, errors.Wrap(err, "failed to unmarshal msg value into AuctionEvent")
	}

	return AuctionEventMessage{
		Message: m,
		Value:   &event,
	}, nil
}

//MustAuctionEventMessage - unmarshal msg value into AuctionEventMessage. Panics if unmarshal fails
func (m Message) MustAuctionEventMessage() AuctionEventMessage {
	result, err := m.AuctionEventMessage()
	if err != nil {
		panic(err)
	}

	return result
}

// DEPRECATED
type AuctionEventTopic struct {
	cfg Config
}

func (t AuctionEventTopic) Reader() AuctionEventReader {
	return NewAuctionEventReader(t.cfg.reader(Topic_auctions_v2.String()))
}

func (t AuctionEventTopic) Writer() AuctionEventWriter {
	return NewAuctionEventWriter(t.cfg.writer(Topic_auctions_v2.String()))
}

// DEPRECATED
func NewAuctionEventTopic(cfg Config) AuctionEventTopic {
	return AuctionEventTopic{
		cfg,
	}
}

// AuctionEventReader reads LotUpdatedEvent from kafka.
// DEPERCATED: use ingester
type AuctionEventReader ProtoReader

// NewAuctionEventReader inits an instance of AuctionEventReader.
// DEPRECATED
func NewAuctionEventReader(cfg ReaderConfig) AuctionEventReader {
	return AuctionEventReader(NewProtoReader(cfg))
}

// WithAutoReset returns copy of AuctionEventReader that in case of ReadMessage error will reset its offset to last commited.
// DEPRECATED: NOT THREAD SAFE & LEAKS RESOURCES
func (r AuctionEventReader) WithAutoReset() AuctionEventReader {
	r.autoReset = true
	return r
}

func (r AuctionEventReader) readMessage(ctx context.Context, reader protoReaderFunc) (AuctionEventMessage, error) {
	e := new(AuctionEvent)
	msg, err := reader(ctx, e)
	if err != nil {
		if r.autoReset {
			r.Reset()
		}
		return AuctionEventMessage{}, err
	}
	return AuctionEventMessage{Message: Message{msg: msg.raw}, Value: e}, nil
}

// DEPRECATED: use FetchMessage/CommitMessage instead
func (r AuctionEventReader) ReadMessage(ctx context.Context) (AuctionEventMessage, error) {
	return r.readMessage(ctx, ProtoReader(r).ReadMessage)
}

// DEPRECATED
func (r AuctionEventReader) FetchMessage(ctx context.Context) (AuctionEventMessage, error) {
	return r.readMessage(ctx, ProtoReader(r).FetchMessage)
}

// DEPRECATED
func (r AuctionEventReader) CommitMessages(ctx context.Context, msgs ...AuctionEventMessage) error {
	protoMsgs := make([]ProtoMessage, 0, len(msgs))
	for _, msg := range msgs {
		protoMsgs = append(protoMsgs, ProtoMessage{raw: msg.msg, Value: msg.Value})
	}
	return ProtoReader(r).CommitMessages(ctx, protoMsgs...)
}

// Reset defaults underlying kafka.Reader
// DEPRECATED: leaks resources
func (r AuctionEventReader) Reset() {
	*r.Reader = *kafka.NewReader(r.Reader.Config())
}

// AuctionEventWriter writes AuctionEvent to kafka.
type AuctionEventWriter ProtoWriter

// NewAuctionEventWriter inits an instance of AuctionEventWriter.
func NewAuctionEventWriter(cfg WriterConfig) AuctionEventWriter {
	return AuctionEventWriter(NewProtoWriter(cfg))
}

// WriteMessages to kafka.
func (w AuctionEventWriter) WriteMessages(ctx context.Context, msgs ...AuctionEventMessage) error {
	protoMsgs := make([]ProtoMessage, len(msgs))
	for idx, msg := range msgs {
		protoMsgs[idx] = ProtoMessage{raw: msg.msg, Value: msg.Value}
	}
	return ProtoWriter(w).WriteMessages(ctx, protoMsgs...)
}

// DEPRECATED
func (cfg Config) AuctionEventTopic() AuctionEventTopic {
	return NewAuctionEventTopic(cfg)
}
