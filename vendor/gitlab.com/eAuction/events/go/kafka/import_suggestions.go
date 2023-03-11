package kafka

import (
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

func (m ExternalPriceUpdated) Event() ImportSuggestionEvent {
	return ImportSuggestionEvent{
		Type:    ImportSuggestionEvent_TYPE_EXTERNAL_PRICE_UPDATED,
		Updated: &m,
	}
}

func (m ExternalPriceUpdatedV2) Event() ImportSuggestionEvent {
	return ImportSuggestionEvent{
		Type:      ImportSuggestionEvent_TYPE_EXTERNAL_PRICE_UPDATED_V2,
		UpdatedV2: &m,
	}
}

// ExternalPriceDeltaPercent - percent delta of total import price and external_price in terms of external price
func (m ExternalPriceUpdated) ExternalPriceDeltaPercent() float64 {
	return float64((m.ExternalPriceEur-m.TotalImportPriceEur)*100) / float64(m.ExternalPriceEur)
}

// ExternalPriceDeltaPercent - percent delta of total import price and external_price in terms of external price
func (m ExternalPriceUpdatedV2) ExternalPriceDeltaPercent() float64 {
	return float64((m.ExternalPriceEur-m.TotalImportPriceEur)*100) / float64(m.ExternalPriceEur)
}

func (m ImportSuggestionEvent) Message() (msg ImportSuggestionEventMessage) {
	var lotID int64
	switch m.Type {
	case ImportSuggestionEvent_TYPE_EXTERNAL_PRICE_UPDATED:
		lotID = m.Updated.LotId
	case ImportSuggestionEvent_TYPE_EXTERNAL_PRICE_UPDATED_V2:
		lotID = m.UpdatedV2.LotId
	default:
		panic(errors.Errorf("unknown ExternalLotEvent type: %d", m.Type))
	}

	msg.msg.Key = proto.EncodeVarint(uint64(lotID))
	msg.Value = &m

	return msg
}

func (m ImportSuggestionEventMessage) Key() int64 {
	result, _ := proto.DecodeVarint(m.msg.Key)
	return int64(result)
}
