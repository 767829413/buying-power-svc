package responses

import (
	"gitlab.com/distributed_lab/lorem"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
)

//PopulateEnvelopWithUnknownLotEvent - crafts response
func PopulateEnvelopWithUnknownLotEvent(copartLotNumber, saleID string) resources.EnvelopeResponse {
	key := resources.Key{
		ID:   lorem.ULID(),
		Type: resources.UNKNOWN_LOT_EVENTS,
	}
	body := resources.EnvelopeResponse{
		Data: resources.Envelope{
			Key: resources.Key{
				ID:   lorem.ULID(),
				Type: resources.ENVELOPES,
			},
			Relationships: resources.EnvelopeRelationships{
				UnknownLotEvent: key.AsRelation(),
			},
		},
		Included: resources.Included{},
	}

	sale := resources.Key{
		ID:   saleID,
		Type: resources.SALES,
	}.AsRelation()
	body.Included.Add(&resources.UnknowLotEvent{
		Key: key,
		Attributes: resources.UnknowLotEventAttributes{
			ExternalLotId: copartLotNumber,
		},
		Relationships: resources.UnknowLotEventRelationships{
			Sale: *sale,
		},
	})

	return body
}
