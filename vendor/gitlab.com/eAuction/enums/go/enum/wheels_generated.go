// Generated. DO NOT EDIT.
package enum

import (
	"gitlab.com/eAuction/enums/go/internal/alias"
	"gitlab.com/eAuction/enums/go/pkg/alphanum"
)

var wheels = []EnumValue{
	{Name: "Left", ReadableName: "Left"},
	{Name: "Right", ReadableName: "Right"},
	{Name: "Unknown", ReadableName: "Unknown"},
}

// WheelPosition parses and returns value of Wheel_Position enum.
func WheelPosition(s string) Wheel_Position {
	norm, ok := alias.WheelTypes[alphanum.Lower(s)]
	if !ok {
		return Wheel_POSITION_UNKNOWN
	}
	return mapAliasToWheelTypes[norm]
}

// Name returns string name of the enum value.
func (x Wheel_Position) Name() string {
	v, ok := mapWheelTypesToEnumValue[x]
	if !ok {
		return "UNKNOWN"
	}
	return v.Name
}

var mapAliasToWheelTypes = map[string]Wheel_Position{
	"Left":    Wheel_POSITION_LEFT,
	"Unknown": Wheel_POSITION_NULL,
	"Right":   Wheel_POSITION_RIGHT,
}

var mapWheelTypesToEnumValue = map[Wheel_Position]EnumValue{
	Wheel_POSITION_LEFT:  {Name: "Left", ReadableName: "Left"},
	Wheel_POSITION_NULL:  {Name: "Unknown", ReadableName: "Unknown"},
	Wheel_POSITION_RIGHT: {Name: "Right", ReadableName: "Right"},
}
