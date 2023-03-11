// Generated. DO NOT EDIT.
package enum

import (
	"gitlab.com/eAuction/enums/go/internal/alias"
	"gitlab.com/eAuction/enums/go/pkg/alphanum"
)

var driveTypes = []EnumValue{
	{Name: "4X4 Drive", ReadableName: "4X4_DRIVE"},
	{Name: "All Drive", ReadableName: "ALL_DRIVE"},
	{Name: "Front Drive", ReadableName: "FRONT_DRIVE"},
	{Name: "Rear Drive", ReadableName: "REAR_DRIVE"},
}

// DriveType parses and returns value of Drive_Type enum.
func DriveType(s string) Drive_Type {
	norm, ok := alias.DriveTypes[alphanum.Lower(s)]
	if !ok {
		return Drive_TYPE_UNKNOWN
	}
	return mapAliasToDriveTypes[norm]
}

// Name returns string name of the enum value.
func (x Drive_Type) Name() string {
	v, ok := mapDriveTypesToEnumValue[x]
	if !ok {
		return "UNKNOWN"
	}
	return v.Name
}

var mapAliasToDriveTypes = map[string]Drive_Type{
	"4X4 Drive":   Drive_TYPE_4X4,
	"All Drive":   Drive_TYPE_ALL,
	"Front Drive": Drive_TYPE_FRONT,
	"Rear Drive":  Drive_TYPE_REAR,
}

var mapDriveTypesToEnumValue = map[Drive_Type]EnumValue{
	Drive_TYPE_4X4:   {Name: "4X4 Drive", ReadableName: "4X4_DRIVE"},
	Drive_TYPE_ALL:   {Name: "All Drive", ReadableName: "ALL_DRIVE"},
	Drive_TYPE_FRONT: {Name: "Front Drive", ReadableName: "FRONT_DRIVE"},
	Drive_TYPE_REAR:  {Name: "Rear Drive", ReadableName: "REAR_DRIVE"},
}
