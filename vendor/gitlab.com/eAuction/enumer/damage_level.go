package enumer

import (
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	// ErrUnknownDamageLevel is an error that indicates that damage level is unknown for given input.
	ErrUnknownDamageLevel = errors.New("unknown damage level")
)

// FindDamageLevel returns damage level based on vehicle state and damage values.
func FindDamageLevel(vehicleState VehicleState, damageType DamageType) (DamageLevel, error) {
	state := getState()
	typeToLevel := state.VehicleStateDamageTypeToDamageLevel[int32(vehicleState)]
	if typeToLevel == nil {
		return 0, errors.From(ErrUnknownDamageLevel, logan.F{"vehicle_state": vehicleState, "damage_type": damageType})
	}
	level, ok := typeToLevel[int32(damageType)]
	if !ok {
		return 0, errors.From(ErrUnknownDamageLevel, logan.F{"vehicle_state": vehicleState, "damage_type": damageType})
	}
	return DamageLevel(level), nil
}
