package responses

import (
	"gitlab.com/distributed_lab/lorem"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
)

//PopulateError - crafts error response
func PopulateError(requestID string, code int32, err error) *resources.EnvelopeResponse {
	errorKey := resources.Key{
		ID:   lorem.ULID(),
		Type: resources.ERRORS,
	}
	body := resources.EnvelopeResponse{
		Data: resources.Envelope{
			Key: resources.Key{
				ID:   requestID,
				Type: resources.ENVELOPES,
			},
			Relationships: resources.EnvelopeRelationships{
				Error: errorKey.AsRelation(),
			},
		},
		Included: resources.Included{},
	}

	body.Included.Add(&resources.Error{
		Key: errorKey,
		Attributes: resources.ErrorAttributes{
			Code:    code,
			Message: err.Error(),
		},
	})

	return &body
}
