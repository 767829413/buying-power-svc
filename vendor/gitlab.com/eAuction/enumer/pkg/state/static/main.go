package static

import (
	"time"

	"gitlab.com/eAuction/enumer/pkg/state"
)

// NewGetter returns getter of static state ("backed" into library).
func NewGetter() state.Getter {
	stateInstance := newGenerated()
	stateInstance.Translations = Translations
	stateInstance.FuelTypeIndicatives = &state.FuelTypeIndicatives{
		IsDiesel: map[int32]struct{}{
			int32(FuelTypeDiesel): {},
		},
		IsElectric: map[int32]struct{}{
			int32(FuelTypeElectric): {},
		},
		IsHybrid: map[int32]struct{}{
			int32(FuelTypeHybrid): {},
		},
		IsGas: map[int32]struct{}{
			int32(FuelTypeGas): {},
		},
	}
	stateInstance.TransmissionTypeIndicatives = &state.TransmissionTypeIndicatives{
		IsManual: map[int32]struct{}{
			int32(TransmissionTypeManual): struct{}{},
		},
	}
	stateInstance.VehicleStateIndicatives = &state.VehicleStateIndicatives{
		IsRunsAndDrives: map[int32]struct{}{
			int32(VehicleStateRunsAndDrives): {},
		},
		IsEngineStartProgram: map[int32]struct{}{
			int32(VehicleStateEngineStartProgram): {},
		},
		IsEnhanced: map[int32]struct{}{
			int32(VehicleStateEnhanced): {},
		},
		IsNew: map[int32]struct{}{
			int32(VehicleStateNew): {},
		},
	}
	stateInstance.ConditionIndicatives = &state.ConditionIndicatives{
		IsGood: map[int32]struct{}{
			int32(ConditionGood): {},
		},
		IsDamage: map[int32]struct{}{
			int32(ConditionDamage): {},
		},
		IsMinorDamage: map[int32]struct{}{
			int32(ConditionMinorDamage): {},
		},
	}
	stateInstance.DamageLevelIndicatives = state.DamageLevelIndicatives{
		int32(DamageLevelLow):        int32(ConditionGood),
		int32(DamageLevelLowMedium):  int32(ConditionGood),
		int32(DamageLevelMedium):     int32(ConditionMinorDamage),
		int32(DamageLevelMediumHigh): int32(ConditionMinorDamage),
		int32(DamageLevelHigh):       int32(ConditionDamage),
	}
	stateInstance.DamageTypeIndicatives = &state.DamageTypeIndicatives{
		IsRejectedRepair: map[int32]struct{}{
			int32(DamageTypeRejectedRepair): {},
		},
		IsElectrical: map[int32]struct{}{
			int32(DamageTypeElectrical): {},
		},
		IsSide: map[int32]struct{}{
			int32(DamageTypeSide): {},
		},
		IsTheft: map[int32]struct{}{
			int32(DamageTypeTheft): {},
		},
		IsStrip: map[int32]struct{}{
			int32(DamageTypeStrip): {},
		},
		IsRepossession: map[int32]struct{}{
			int32(DamageTypeRepossession): {},
		},
		IsFrame: map[int32]struct{}{
			int32(DamageTypeFrame): {},
		},
		IsNormalWear: map[int32]struct{}{
			int32(DamageTypeNormalWear): {},
		},
		IsPartialRepair: map[int32]struct{}{
			int32(DamageTypePartialRepair): {},
		},
		IsFrontEnd: map[int32]struct{}{
			int32(DamageTypeFrontEnd): {},
		},
		IsAllOver: map[int32]struct{}{
			int32(DamageTypeAllOver): {},
		},
		IsMinorDentOrScratches: map[int32]struct{}{
			int32(DamageTypeMinorDentOrScratches): {},
		},
		IsLeftRear: map[int32]struct{}{
			int32(DamageTypeLeftRear): {},
		},
		IsTransmission: map[int32]struct{}{
			int32(DamageTypeTransmission): {},
		},
		IsVandalism: map[int32]struct{}{
			int32(DamageTypeVandalism): {},
		},
		IsRightFront: map[int32]struct{}{
			int32(DamageTypeRightFront): {},
		},
		IsSuspension: map[int32]struct{}{
			int32(DamageTypeSuspension): {},
		},
		IsRearEnd: map[int32]struct{}{
			int32(DamageTypeRearEnd): {},
		},
		IsRightSide: map[int32]struct{}{
			int32(DamageTypeRightSide): {},
		},
		IsFlood: map[int32]struct{}{
			int32(DamageTypeFlood): {},
		},
		IsFreshWater: map[int32]struct{}{
			int32(DamageTypeFreshWater): {},
		},
		IsRollover: map[int32]struct{}{
			int32(DamageTypeRollover): {},
		},
		IsUndercarriage: map[int32]struct{}{
			int32(DamageTypeUndercarriage): {},
		},
		IsRightRear: map[int32]struct{}{
			int32(DamageTypeRightRear): {},
		},
		IsLeftFront: map[int32]struct{}{
			int32(DamageTypeLeftFront): {},
		},
		IsInteriorBurn: map[int32]struct{}{
			int32(DamageTypeInteriorBurn): {},
		},
		IsDamageHistory: map[int32]struct{}{
			int32(DamageTypeDamageHistory): {},
		},
		IsRoof: map[int32]struct{}{
			int32(DamageTypeRoof): {},
		},
		IsFrontAndRear: map[int32]struct{}{
			int32(DamageTypeFrontAndRear): {},
		},
		IsHail: map[int32]struct{}{
			int32(DamageTypeHail): {},
		},
		IsMechanical: map[int32]struct{}{
			int32(DamageTypeMechanical): {},
		},
		IsSaltWater: map[int32]struct{}{
			int32(DamageTypeSaltWater): {},
		},
		IsTotalBurn: map[int32]struct{}{
			int32(DamageTypeTotalBurn): {},
		},
		IsBurn: map[int32]struct{}{
			int32(DamageTypeBurn): {},
		},
		IsReplacedVin: map[int32]struct{}{
			int32(DamageTypeReplacedVin): {},
		},
		IsBiohazardOrChemical: map[int32]struct{}{
			int32(DamageTypeBiohazardOrChemical): {},
		},
		IsLeftAndRight: map[int32]struct{}{
			int32(DamageTypeLeftAndRight): {},
		},
		IsEngineBurn: map[int32]struct{}{
			int32(DamageTypeEngineBurn): {},
		},
		IsLeftSide: map[int32]struct{}{
			int32(DamageTypeLeftSide): {},
		},
		IsNeedsRepair: map[int32]struct{}{
			int32(DamageTypeNeedsRepair): {},
		},
		IsStormDamage: map[int32]struct{}{
			int32(DamageTypeStormDamage): {},
		},
		IsNone: map[int32]struct{}{
			int32(DamageTypeNone): {},
		},
		IsExteriorBurn: map[int32]struct{}{
			int32(DamageTypeExteriorBurn): {},
		},
		IsEngineDamage: map[int32]struct{}{
			int32(DamageTypeEngineDamage): {},
		},
		IsMissingOrAlteredVin: map[int32]struct{}{
			int32(DamageTypeMissingOrAlteredVin): {},
		},
		IsCashForClunkers: map[int32]struct{}{
			int32(DamageTypeCashForClunkers): {},
		},
	}
	stateInstance.WheelPositionIndicatives = &state.WheelPositionIndicatives{
		IsRight: map[int32]struct{}{
			int32(WheelPositionRight): {},
		},
		IsLeft: map[int32]struct{}{
			int32(WheelPositionLeft): {},
		},
	}
	newVehicleAllLow := make(map[int32]int32, len(DamageTypeValues))
	for _, v := range DamageTypeValues {
		newVehicleAllLow[v] = int32(DamageLevelLow)
	}
	stateInstance.VehicleStateDamageTypeToDamageLevel = map[int32]map[int32]int32{
		1:                      {2: 6, 4: 2, 6: 2, 8: 6, 9: 2, 11: 4, 12: 4, 13: 2, 16: 2, 19: 2, 23: 6, 27: 4, 29: 4, 31: 4, 32: 6, 35: 6, 37: 4},
		2:                      {1: 4, 4: 3, 5: 3, 6: 3, 7: 3, 8: 4, 9: 3, 10: 4, 12: 4, 13: 3, 14: 4, 15: 4, 16: 3, 18: 4, 19: 3, 20: 4, 23: 4, 24: 4, 25: 4, 26: 4, 27: 4, 29: 4, 30: 4, 31: 4, 32: 4, 37: 3, 38: 4, 40: 4, 41: 4, 42: 4, 43: 3, 44: 4},
		3:                      {1: 4, 4: 4, 5: 4, 6: 4, 7: 4, 8: 4, 9: 4, 10: 4, 11: 4, 12: 4, 13: 4, 14: 4, 16: 4, 17: 4, 18: 4, 19: 4, 20: 4, 21: 4, 23: 4, 24: 4, 25: 4, 26: 4, 29: 4, 30: 4, 31: 4, 32: 4, 37: 4, 38: 4, 40: 4, 42: 4, 43: 4},
		4:                      {1: 5, 2: 6, 3: 6, 4: 5, 5: 4, 6: 5, 7: 4, 8: 6, 9: 4, 10: 5, 11: 5, 12: 5, 13: 4, 14: 5, 15: 6, 16: 5, 17: 5, 18: 6, 19: 5, 20: 5, 21: 6, 22: 6, 23: 6, 24: 6, 25: 5, 26: 5, 27: 6, 28: 6, 29: 5, 30: 5, 31: 4, 32: 6, 33: 6, 34: 6, 35: 6, 36: 6, 37: 4, 38: 5, 39: 6, 40: 5, 42: 6, 43: 4, 44: 6, 45: 6, 46: 6},
		int32(VehicleStateNew): newVehicleAllLow,
	}
	stateInstance.Timestamp = time.Now().UTC()
	return state.GetterFunc(func() (*state.State, error) {
		return stateInstance, nil
	})
}
