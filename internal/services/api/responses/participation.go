package responses

import (
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
	"gitlab.com/eAuction/events/go/kafka"
	"gitlab.com/eAuction/go/amount"
)

// NewParticipation returns new instance of Participation.
func NewParticipation(dbParticipant data.Participant) *resources.Participation {

	p := resources.Participation{
		Key: resources.NewKeyInt64(dbParticipant.ID, resources.PARTICIPATIONS),
		Attributes: resources.ParticipationAttributes{
			AutoBidLimit:    amount.Fiat(dbParticipant.AutoBidLimit),
			BidLimit:        amount.Fiat(dbParticipant.BidLimit),
			CreatedAt:       dbParticipant.CreatedAt,
			CurrentBid:      amount.Fiat(dbParticipant.CurrentBid),
			RequestedBuyNow: dbParticipant.RequestedBuyNow,
			State:           kafka.Participant_State(dbParticipant.State).String(),
			UpdatedAt:       &dbParticipant.UpdatedAt,
		},
		Relationships: resources.ParticipationRelationships{
			Account: *resources.Key{
				ID:   dbParticipant.AccountAddress,
				Type: resources.IDENTITIES,
			}.AsRelation(),
			Lot: *resources.Key{
				ID:   dbParticipant.LotID,
				Type: resources.LOTS,
			}.AsRelation(),
		},
	}

	return &p
}
