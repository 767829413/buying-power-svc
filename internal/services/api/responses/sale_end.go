package responses

import (
	"gitlab.com/distributed_lab/lorem"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
)

func PopulateSaleEnd(saleID, reason string) resources.EnvelopeResponse {
	saleKey := resources.Key{
		ID:   saleID,
		Type: resources.SALES,
	}.AsRelation()
	endSale := resources.EndSale{
		Key: resources.Key{
			ID:   lorem.ULID(),
			Type: resources.END_SALES,
		},
		Relationships: resources.EndSaleRelationships{
			Sale: *saleKey,
		},
		Attributes: resources.EndSaleAttributes{
			Reason: reason,
		},
	}

	env := resources.EnvelopeResponse{
		Data: resources.Envelope{
			Key: resources.Key{
				Type: resources.ENVELOPES,
			},
			Relationships: resources.EnvelopeRelationships{
				SaleEnd: endSale.Key.AsRelation(),
			},
		},
		Included: resources.Included{},
	}

	env.Included.Add(&endSale)
	return env
}
