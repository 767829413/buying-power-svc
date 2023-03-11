package kafka

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/segmentio/kafka-go"
	"gitlab.com/distributed_lab/logan"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type (
	ProtoReader struct {
		*kafka.Reader
		autoReset bool
	}
	ProtoWriter struct{ *kafka.Writer }

	readerFunc      func(context.Context) (kafka.Message, error)
	protoReaderFunc func(context.Context, proto.Message) (ProtoMessage, error)

	ProtoMessage struct {
		raw   kafka.Message
		Value proto.Message
	}
)

func (m ProtoMessage) Key() []byte {
	return m.raw.Key
}

func NewProtoMessage(key []byte, value proto.Message) ProtoMessage {
	return ProtoMessage{
		raw: kafka.Message{
			Key: key,
		},
		Value: value,
	}
}

func NewProtoReader(config ReaderConfig) ProtoReader {
	return ProtoReader{Reader: kafka.NewReader(kafka.ReaderConfig(config))}
}

func NewProtoWriter(config WriterConfig) ProtoWriter {
	return ProtoWriter{kafka.NewWriter(kafka.WriterConfig(config))}
}

func (r ProtoReader) getMessage(ctx context.Context, value proto.Message, f readerFunc) (ProtoMessage, error) {
	msg, err := f(ctx)
	if err != nil {
		return ProtoMessage{}, errors.Wrap(err, "failed to read message")
	}

	if err := proto.Unmarshal(msg.Value, value); err != nil {
		return ProtoMessage{}, errors.Wrap(err, "failed to proto unmarshal")
	}

	return ProtoMessage{raw: msg, Value: value}, nil
}

func (r ProtoReader) ReadMessage(ctx context.Context, value proto.Message) (ProtoMessage, error) {
	return r.getMessage(ctx, value, r.Reader.ReadMessage)
}

func (r ProtoReader) FetchMessage(ctx context.Context, value proto.Message) (ProtoMessage, error) {
	return r.getMessage(ctx, value, r.Reader.FetchMessage)
}

func (r ProtoReader) CommitMessages(ctx context.Context, msgs ...ProtoMessage) error {
	messages := make([]kafka.Message, 0, len(msgs))
	for _, msg := range msgs {
		messages = append(messages, msg.raw)
	}

	return r.Reader.CommitMessages(ctx, messages...)
}

func (w ProtoWriter) WriteMessages(ctx context.Context, msgs ...ProtoMessage) error {
	for _, msg := range msgs {
		value, err := proto.Marshal(msg.Value)
		if err != nil {
			return errors.Wrap(err, "failed to proto marshal")
		}

		msg.raw.Value = value
		if err = w.Writer.WriteMessages(ctx, msg.raw); err != nil {
			return errors.Wrap(err, "failed to write message", logan.F{
				"key": string(msg.Key()),
			})
		}
	}

	return nil
}
