package responses

import (
	"gitlab.com/distributed_lab/lorem"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
)

//PopulateEnvelopWithParticipations - crafts response
func PopulateEnvelopWithParticipations(participations []data.Participant, lots map[string]data.Lot) resources.EnvelopeResponse {
	body := resources.EnvelopeResponse{
		Data: resources.Envelope{
			Key: resources.Key{
				ID:   lorem.ULID(),
				Type: resources.ENVELOPES,
			},
			Relationships: resources.EnvelopeRelationships{
				Participations: &resources.RelationCollection{},
			},
		},
		Included: resources.Included{},
	}

	for _, dbP :=range participations {
		p := NewParticipation(dbP)
		body.Data.Relationships.Participations.Data = append(body.Data.Relationships.Participations.Data, p.Key)
		body.Included.Add(p)
		lot := NewLot(lots[dbP.LotID])
		body.Included.Add(lot)
		vehicle := NewVehicle(lots[dbP.LotID])
		body.Included.Add(vehicle)
	}

	return body
}
