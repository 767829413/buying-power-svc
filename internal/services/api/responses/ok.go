package responses

import (
	"gitlab.com/distributed_lab/lorem"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
)

//PopulateOk - populates ok response
func PopulateOk(requestID string) *resources.EnvelopeResponse {
	return &resources.EnvelopeResponse{
		Data: resources.Envelope{
			Key: resources.Key{
				ID:   requestID,
				Type: resources.ENVELOPES,
			},
			Relationships: resources.EnvelopeRelationships{
				Ok: resources.Key{
					ID:   lorem.ULID(),
					Type: resources.OKS,
				}.AsRelation(),
			},
		},
		Included: resources.Included{},
	}
}
