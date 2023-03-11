// Generated. DO NOT EDIT.
package enum

import (
	"gitlab.com/eAuction/enums/go/internal/alias"
	"gitlab.com/eAuction/enums/go/pkg/alphanum"
)

var bodyStyles = []EnumValue{
	{Name: "Convertible", ReadableName: "CONVERTIBLE"},
	{Name: "Coupe", ReadableName: "COUPE"},
	{Name: "Hatchback", ReadableName: "HATCHBACK"},
	{Name: "Limousine", ReadableName: "LIMOUSINE"},
	{Name: "SUV", ReadableName: "SUV"},
	{Name: "Sedan", ReadableName: "SEDAN"},
	{Name: "Truck", ReadableName: "TRUCK"},
	{Name: "Van", ReadableName: "VAN"},
}

// BodyStyle parses and returns value of Body_Style enum.
func BodyStyle(s string) Body_Style {
	norm, ok := alias.BodyStyles[alphanum.Lower(s)]
	if !ok {
		return Body_STYLE_UNKNOWN
	}
	return mapAliasToBodyStyles[norm]
}

// Name returns string name of the enum value.
func (x Body_Style) Name() string {
	v, ok := mapBodyStylesToEnumValue[x]
	if !ok {
		return "UNKNOWN"
	}
	return v.Name
}

var mapAliasToBodyStyles = map[string]Body_Style{
	"Convertible": Body_STYLE_CONVERTIBLE,
	"Coupe":       Body_STYLE_COUPE,
	"Hatchback":   Body_STYLE_HATCHBACK,
	"Limousine":   Body_STYLE_LIMOUSINE,
	"SUV":         Body_STYLE_SUV,
	"Sedan":       Body_STYLE_SEDAN,
	"Truck":       Body_STYLE_TRUCK,
	"Van":         Body_STYLE_VAN,
}

var mapBodyStylesToEnumValue = map[Body_Style]EnumValue{
	Body_STYLE_CONVERTIBLE: {Name: "Convertible", ReadableName: "CONVERTIBLE"},
	Body_STYLE_COUPE:       {Name: "Coupe", ReadableName: "COUPE"},
	Body_STYLE_HATCHBACK:   {Name: "Hatchback", ReadableName: "HATCHBACK"},
	Body_STYLE_LIMOUSINE:   {Name: "Limousine", ReadableName: "LIMOUSINE"},
	Body_STYLE_SUV:         {Name: "SUV", ReadableName: "SUV"},
	Body_STYLE_SEDAN:       {Name: "Sedan", ReadableName: "SEDAN"},
	Body_STYLE_TRUCK:       {Name: "Truck", ReadableName: "TRUCK"},
	Body_STYLE_VAN:         {Name: "Van", ReadableName: "VAN"},
}
