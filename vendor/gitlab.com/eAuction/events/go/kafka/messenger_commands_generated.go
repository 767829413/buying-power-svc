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

// MessengerCommandMessage wraps original kafka-go message and
// parsed protobuf generated MessengerCommand.
type MessengerCommandMessage struct {
	// raw is set directly from kafka event and left private intentionally.
	// it's advisable to implement adapter methods for both producer and consumer
	// to have key serialization semantics available in one place.
	// helpers might look something like:
	//
	// func (m MessengerCommandMessage) Key() uint64 {
	//     key, err := proto.DecodeVarint(m.rawkey)
	//     if err != nil {
	//         panic(errors.Wrap(err, "failed to decode event key"))
	//     }
	//     return key
	// }
	//
	// func (m MessengerCommand) Message(key uint64) MessengerCommandMessage {
	// return MessengerCommandMessage{
	//        Key: proto.EncodeVarint(key),
	//        Value: &m,
	//     }
	// }
	//
	// it also prevents kafka-go (implementation detail) to leak outside of package
	Message
	Value *MessengerCommand
}

//MessengerCommandMessage - unmarshal msg value into MessengerCommandMessage.
func (m Message) MessengerCommandMessage() (MessengerCommandMessage, error) {
	var event MessengerCommand
	err := proto.Unmarshal(m.msg.Value, &event)
	if err != nil {
		return MessengerCommandMessage{}, errors.Wrap(err, "failed to unmarshal msg value into MessengerCommand")
	}

	return MessengerCommandMessage{
		Message: m,
		Value:   &event,
	}, nil
}

//MustMessengerCommandMessage - unmarshal msg value into MessengerCommandMessage. Panics if unmarshal fails
func (m Message) MustMessengerCommandMessage() MessengerCommandMessage {
	result, err := m.MessengerCommandMessage()
	if err != nil {
		panic(err)
	}

	return result
}

// DEPRECATED
type MessengerCommandTopic struct {
	cfg Config
}

func (t MessengerCommandTopic) Reader() MessengerCommandReader {
	return NewMessengerCommandReader(t.cfg.reader(Topic_messenger_commands.String()))
}

func (t MessengerCommandTopic) Writer() MessengerCommandWriter {
	return NewMessengerCommandWriter(t.cfg.writer(Topic_messenger_commands.String()))
}

// DEPRECATED
func NewMessengerCommandTopic(cfg Config) MessengerCommandTopic {
	return MessengerCommandTopic{
		cfg,
	}
}

// MessengerCommandReader reads LotUpdatedEvent from kafka.
// DEPERCATED: use ingester
type MessengerCommandReader ProtoReader

// NewMessengerCommandReader inits an instance of MessengerCommandReader.
// DEPRECATED
func NewMessengerCommandReader(cfg ReaderConfig) MessengerCommandReader {
	return MessengerCommandReader(NewProtoReader(cfg))
}

// WithAutoReset returns copy of MessengerCommandReader that in case of ReadMessage error will reset its offset to last commited.
// DEPRECATED: NOT THREAD SAFE & LEAKS RESOURCES
func (r MessengerCommandReader) WithAutoReset() MessengerCommandReader {
	r.autoReset = true
	return r
}

func (r MessengerCommandReader) readMessage(ctx context.Context, reader protoReaderFunc) (MessengerCommandMessage, error) {
	e := new(MessengerCommand)
	msg, err := reader(ctx, e)
	if err != nil {
		if r.autoReset {
			r.Reset()
		}
		return MessengerCommandMessage{}, err
	}
	return MessengerCommandMessage{Message: Message{msg: msg.raw}, Value: e}, nil
}

// DEPRECATED: use FetchMessage/CommitMessage instead
func (r MessengerCommandReader) ReadMessage(ctx context.Context) (MessengerCommandMessage, error) {
	return r.readMessage(ctx, ProtoReader(r).ReadMessage)
}

// DEPRECATED
func (r MessengerCommandReader) FetchMessage(ctx context.Context) (MessengerCommandMessage, error) {
	return r.readMessage(ctx, ProtoReader(r).FetchMessage)
}

// DEPRECATED
func (r MessengerCommandReader) CommitMessages(ctx context.Context, msgs ...MessengerCommandMessage) error {
	protoMsgs := make([]ProtoMessage, 0, len(msgs))
	for _, msg := range msgs {
		protoMsgs = append(protoMsgs, ProtoMessage{raw: msg.msg, Value: msg.Value})
	}
	return ProtoReader(r).CommitMessages(ctx, protoMsgs...)
}

// Reset defaults underlying kafka.Reader
// DEPRECATED: leaks resources
func (r MessengerCommandReader) Reset() {
	*r.Reader = *kafka.NewReader(r.Reader.Config())
}

// MessengerCommandWriter writes MessengerCommand to kafka.
type MessengerCommandWriter ProtoWriter

// NewMessengerCommandWriter inits an instance of MessengerCommandWriter.
func NewMessengerCommandWriter(cfg WriterConfig) MessengerCommandWriter {
	return MessengerCommandWriter(NewProtoWriter(cfg))
}

// WriteMessages to kafka.
func (w MessengerCommandWriter) WriteMessages(ctx context.Context, msgs ...MessengerCommandMessage) error {
	protoMsgs := make([]ProtoMessage, len(msgs))
	for idx, msg := range msgs {
		protoMsgs[idx] = ProtoMessage{raw: msg.msg, Value: msg.Value}
	}
	return ProtoWriter(w).WriteMessages(ctx, protoMsgs...)
}

// DEPRECATED
func (cfg Config) MessengerCommandTopic() MessengerCommandTopic {
	return NewMessengerCommandTopic(cfg)
}
