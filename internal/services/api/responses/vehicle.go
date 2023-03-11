package responses

import (
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
)

// NewVehicle returns new instance of Vehicle resource.
func NewVehicle(lot data.Lot) *resources.Vehicle {
	var img string
	if len(lot.Details.HdImgLinks) != 0 {
		img = lot.Details.HdImgLinks[0]
	} else if len(lot.Details.ImgLinks) != 0 {
		img = lot.Details.ImgLinks[0]
	}

	result := resources.Vehicle{
		Key: resources.Key{
			ID:   lot.ID,
			Type: resources.VEHICLES,
		},
		Attributes: resources.VehicleAttributes{
			ExtendedName: lot.Details.ExtendedName,
			Name:         lot.Details.Name,
			Images:       []string{img},
			OnApproval:   lot.Details.OnApproval,
		},
	}

	return &result
}
