package kafka

import (
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

func (m IdentityEvent) Message() *IdentityEventMessage {
	// build compaction key based on atomic sub-entry
	var key string
	pk := strings.Join([]string{m.Address, m.Platform}, ":")
	switch m.Type {
	case Identity_EVENT_TYPE_CREATED:
		key = strings.Join([]string{"created", pk}, ":")
	case Identity_EVENT_TYPE_UPDATED:
		key = strings.Join([]string{"updated", pk, m.Updated.Type.String()}, ":")
	case Identity_EVENT_TYPE_DELETED:
		key = strings.Join([]string{"deleted", pk}, ":")
	default:
		panic(fmt.Errorf("unexpected identity message type: %v", m.Type))
	}

	return &IdentityEventMessage{
		Message: Message{msg: kafka.Message{Key: []byte(key)}},
		Value:   &m,
	}
}

func (e *IdentityEvent) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		panic(fmt.Errorf("unsupported type: %T", value))
	}

	if err := proto.Unmarshal(bytes, e); err != nil {
		return errors.Wrap(err, "failed to unmarshal event")
	}

	return nil
}
