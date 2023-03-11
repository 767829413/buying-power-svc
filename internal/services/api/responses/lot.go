package responses

import (
	"time"

	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
	"gitlab.com/eAuction/events/go/kafka"
	"gitlab.com/eAuction/go/amount"
)

func PopulateEnvelopeWithLot(dbLot data.Lot) resources.EnvelopeResponse {
	env := resources.EnvelopeResponse{
		Data: resources.Envelope{
			Key: resources.Key{
				Type: resources.ENVELOPES,
			},
			Relationships: resources.EnvelopeRelationships{
				CurrentLot: resources.Key{
					ID:   dbLot.ID,
					Type: resources.LOTS,
				}.AsRelation(),
			},
		},
		Included: resources.Included{},
	}

	return env
}

// NewLot returns new instance of lot.
func NewLot(dbLot data.Lot) *resources.Lot {
	vehicleKey := resources.Key{
		ID:   dbLot.ID,
		Type: resources.VEHICLES,
	}.AsRelation()

	var itemNum *int64
	if dbLot.ItemNum.Valid {
		itemNum = &dbLot.ItemNum.Int64
	}
	l := &resources.Lot{
		Key: resources.Key{
			ID:   dbLot.ID,
			Type: resources.LOTS,
		},
		Attributes: resources.LotAttributes{
			CreatedAt:  time.Unix(dbLot.Details.CreatedAt, 0),
			Currency:   dbLot.Details.Currency,
			Duration:   int64(dbLot.Details.DurationSec),
			EndTime:    time.Unix(dbLot.Details.StartTime+int64(dbLot.Details.DurationSec), 0),
			HighestBid: amount.Fiat(dbLot.HighestBid.Int64),
			LotType:    dbLot.Details.Type.String(),
			LotTypeI:   int32(dbLot.Details.Type),
			MinBidStep: amount.Fiat(dbLot.Details.MinBidStep),
			MaxBidStep: amount.Fiat(dbLot.Details.MaxBidStep),
			StartTime:  time.Unix(dbLot.Details.StartTime, 0),
			State:      kafka.Lot_State(dbLot.State).String(),
			Name:       dbLot.Details.Name,
			ItemNumber: itemNum,
		},
		Relationships: resources.LotRelationships{
			Organizer: resources.Relation{
				Data: &resources.Key{
					ID:   dbLot.Details.OrganizerId,
					Type: resources.IDENTITIES,
				},
			},
			Platform: resources.Relation{
				Data: &resources.Key{
					ID:   dbLot.Details.PlatformId,
					Type: resources.IDENTITIES,
				},
			},
			Vehicle: *vehicleKey,
			Sale: resources.Key{
				ID:   dbLot.SaleID.String,
				Type: resources.SALES,
			}.AsRelation(),
		},
	}

	return l
}
