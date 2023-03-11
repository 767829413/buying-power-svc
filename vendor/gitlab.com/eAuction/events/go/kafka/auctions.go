package kafka

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

func (m LotCreated) Event() AuctionEvent {
	return AuctionEvent{
		Type:       AuctionEvent_TYPE_LOT_CREATED,
		LotCreated: &m,
	}
}

func (m LotStateUpdated) Event() AuctionEvent {
	return AuctionEvent{
		Type:            AuctionEvent_TYPE_LOT_STATE_UPDATED,
		LotStateUpdated: &m,
	}
}

func (m ParticipantCreated) Event() AuctionEvent {
	return AuctionEvent{
		Type:               AuctionEvent_TYPE_PARTICIPANT_CREATED,
		ParticipantCreated: &m,
	}
}

func (m ParticipantUpdated) Event() AuctionEvent {
	return AuctionEvent{
		Type:               AuctionEvent_TYPE_PARTICIPANT_UPDATED,
		ParticipantUpdated: &m,
	}
}

func (m Deposit) Event() AuctionEvent {
	return AuctionEvent{
		Type:    AuctionEvent_TYPE_DEPOSIT,
		Deposit: &m,
	}
}

func (m Winner) Event() AuctionEvent {
	return AuctionEvent{
		Type:   AuctionEvent_TYPE_WINNER,
		Winner: &m,
	}
}

func (m LiveBiddingDetails) Event() AuctionEvent {
	return AuctionEvent{
		Type:               AuctionEvent_TYPE_LIVE_BIDDING_DETAILS,
		LiveBiddingDetails: &m,
	}
}

func (m FinalBid) Event() AuctionEvent {
	return AuctionEvent{
		Type:     AuctionEvent_TYPE_FINAL_BID,
		FinalBid: &m,
	}
}

func (m AuctionEvent) Message() (msg AuctionEventMessage) {
	switch m.Type {
	case AuctionEvent_TYPE_LOT_CREATED:
		msg.msg.Key = proto.EncodeVarint(uint64(m.LotCreated.Lot.Id))
	case AuctionEvent_TYPE_LOT_STATE_UPDATED:
		msg.msg.Key = proto.EncodeVarint(uint64(m.LotStateUpdated.LotId))
	case AuctionEvent_TYPE_PARTICIPANT_CREATED:
		msg.msg.Key = proto.EncodeVarint(uint64(m.ParticipantCreated.LotId))
	case AuctionEvent_TYPE_PARTICIPANT_UPDATED:
		msg.msg.Key = proto.EncodeVarint(uint64(m.ParticipantUpdated.LotId))
	case AuctionEvent_TYPE_DEPOSIT:
		msg.msg.Key = []byte(fmt.Sprintf("deposit:%d:%s", m.Deposit.LotId, m.Deposit.AccountId))
	case AuctionEvent_TYPE_WINNER:
		msg.msg.Key = []byte(fmt.Sprintf("winner:%d", m.Winner.LotId))
	case AuctionEvent_TYPE_FINAL_BID:
		msg.msg.Key = []byte(fmt.Sprintf("finalbid:%d", m.FinalBid.LotId))
	case AuctionEvent_TYPE_LIVE_BIDDING_DETAILS:
		msg.msg.Key = []byte(fmt.Sprintf("live_bidding_details:%d", m.LiveBiddingDetails.LotId))
	default:
		panic(errors.New("unexpected auction event type"))
	}

	msg.Value = &m
	return msg
}

func (m AuctionEventMessage) Key() string {
	switch m.Value.Type {
	case AuctionEvent_TYPE_LOT_CREATED, AuctionEvent_TYPE_LOT_STATE_UPDATED, AuctionEvent_TYPE_PARTICIPANT_CREATED, AuctionEvent_TYPE_PARTICIPANT_UPDATED:
		intID, _ := proto.DecodeVarint(m.msg.Key)
		return strconv.FormatUint(intID, 10)
	default:
		return string(m.msg.Key)
	}
}

//EncodeSaleID - encodes sale ID
func EncodeSaleID(platform, branch, lane string, startTime time.Time) string {
	return fmt.Sprintf("%s-%s-%s-%d", platform, branch, lane, startTime.Unix())
}

//DecodeSaleID - decodes sale ID
func DecodeSaleID(id string) (string, string, string, time.Time, error) {
	raw := strings.Split(id, "-")
	if len(raw) != 4 {
		return "", "", "", time.Time{}, errors.New("invalid format of the id: expected contain 4 parts")
	}

	startTime, err := strconv.ParseInt(raw[3], 10, 64)
	if err != nil {
		return "", "", "", time.Time{}, errors.Wrap(err, "failed to parse start time")
	}

	return raw[0], raw[1], raw[2], time.Unix(startTime, 0).UTC(), nil
}