package kafka

import (
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

func (m Star_Created) Event() StarEvent {
	return StarEvent{
		Type:    Star_EVENT_TYPE_CREATED,
		Created: &m,
	}
}

func (m Star_Deleted) Event() StarEvent {
	return StarEvent{
		Type:    Star_EVENT_TYPE_DELETED,
		Deleted: &m,
	}
}

func (m Star_Expiring) Event() StarEvent {
	return StarEvent{
		Type:     Star_EVENT_TYPE_EXPIRING,
		Expiring: &m,
	}
}

func (m Star_Expired) Event() StarEvent {
	return StarEvent{
		Type:    Star_EVENT_TYPE_EXPIRED,
		Expired: &m,
	}
}

func (m StarEvent) Message() (msg StarEventMessage) {
	var starID int64
	switch m.Type {
	case Star_EVENT_TYPE_CREATED:
		starID = m.Created.Id
	case Star_EVENT_TYPE_DELETED:
		starID = m.Deleted.Id
	case Star_EVENT_TYPE_EXPIRING:
		starID = m.Expiring.Id
	case Star_EVENT_TYPE_EXPIRED:
		starID = m.Expired.Id
	default:
		panic(errors.Errorf("unknown StarEvent type: %d", m.Type))
	}

	msg.msg.Key = proto.EncodeVarint(uint64(starID))
	msg.Value = &m

	return msg
}

func (m StarEventMessage) Key() string {
	key, n := proto.DecodeVarint(m.msg.Key)
	if n == 0 {
		panic(errors.New("failed to decode StarEventMessage key"))
	}

	return strconv.FormatUint(key, 10)
}
