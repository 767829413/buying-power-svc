// Generated. DO NOT EDIT.
package enum

import (
	"gitlab.com/eAuction/enums/go/internal/alias"
	"gitlab.com/eAuction/enums/go/pkg/alphanum"
)

var transmissions = []EnumValue{
	{Name: "Automatic", ReadableName: "AUTOMATIC"},
	{Name: "Manual", ReadableName: "MANUAL_GEARBOX"},
	{Name: "Unknown", ReadableName: "Unknown"},
}

// TransmissionType parses and returns value of Transmission_Type enum.
func TransmissionType(s string) Transmission_Type {
	norm, ok := alias.Transmissions[alphanum.Lower(s)]
	if !ok {
		return Transmission_TYPE_UNKNOWN
	}
	return mapAliasToTransmissions[norm]
}

// Name returns string name of the enum value.
func (x Transmission_Type) Name() string {
	v, ok := mapTransmissionsToEnumValue[x]
	if !ok {
		return "UNKNOWN"
	}
	return v.Name
}

var mapAliasToTransmissions = map[string]Transmission_Type{
	"Automatic": Transmission_TYPE_AUTOMATIC,
	"Manual":    Transmission_TYPE_MANUAL,
	"Unknown":   Transmission_TYPE_NULL,
}

var mapTransmissionsToEnumValue = map[Transmission_Type]EnumValue{
	Transmission_TYPE_AUTOMATIC: {Name: "Automatic", ReadableName: "AUTOMATIC"},
	Transmission_TYPE_MANUAL:    {Name: "Manual", ReadableName: "MANUAL_GEARBOX"},
	Transmission_TYPE_NULL:      {Name: "Unknown", ReadableName: "Unknown"},
}
