package enumer

// IsRejectedRepair returns true if DamageType represents "RejectedRepair" value.
func (t DamageType) IsRejectedRepair() bool {
	_, ok := getState().DamageTypeIndicatives.IsRejectedRepair[int32(t)]
	return ok
}

// IsElectrical returns true if DamageType represents "Electrical" value.
func (t DamageType) IsElectrical() bool {
	_, ok := getState().DamageTypeIndicatives.IsElectrical[int32(t)]
	return ok
}

// IsSide returns true if DamageType represents "Side" value.
func (t DamageType) IsSide() bool {
	_, ok := getState().DamageTypeIndicatives.IsSide[int32(t)]
	return ok
}

// IsTheft returns true if DamageType represents "Theft" value.
func (t DamageType) IsTheft() bool {
	_, ok := getState().DamageTypeIndicatives.IsTheft[int32(t)]
	return ok
}

// IsStrip returns true if DamageType represents "Strip" value.
func (t DamageType) IsStrip() bool {
	_, ok := getState().DamageTypeIndicatives.IsStrip[int32(t)]
	return ok
}

// IsRepossession returns true if DamageType represents "Repossession" value.
func (t DamageType) IsRepossession() bool {
	_, ok := getState().DamageTypeIndicatives.IsRepossession[int32(t)]
	return ok
}

// IsFrame returns true if DamageType represents "Frame" value.
func (t DamageType) IsFrame() bool {
	_, ok := getState().DamageTypeIndicatives.IsFrame[int32(t)]
	return ok
}

// IsNormalWear returns true if DamageType represents "NormalWear" value.
func (t DamageType) IsNormalWear() bool {
	_, ok := getState().DamageTypeIndicatives.IsNormalWear[int32(t)]
	return ok
}

// IsPartialRepair returns true if DamageType represents "PartialRepair" value.
func (t DamageType) IsPartialRepair() bool {
	_, ok := getState().DamageTypeIndicatives.IsPartialRepair[int32(t)]
	return ok
}

// IsFrontEnd returns true if DamageType represents "FrontEnd" value.
func (t DamageType) IsFrontEnd() bool {
	_, ok := getState().DamageTypeIndicatives.IsFrontEnd[int32(t)]
	return ok
}

// IsAllOver returns true if DamageType represents "AllOver" value.
func (t DamageType) IsAllOver() bool {
	_, ok := getState().DamageTypeIndicatives.IsAllOver[int32(t)]
	return ok
}

// IsMinorDentOrScratches returns true if DamageType represents "MinorDentOrScratches" value.
func (t DamageType) IsMinorDentOrScratches() bool {
	_, ok := getState().DamageTypeIndicatives.IsMinorDentOrScratches[int32(t)]
	return ok
}

// IsLeftRear returns true if DamageType represents "LeftRear" value.
func (t DamageType) IsLeftRear() bool {
	_, ok := getState().DamageTypeIndicatives.IsLeftRear[int32(t)]
	return ok
}

// IsTransmission returns true if DamageType represents "Transmission" value.
func (t DamageType) IsTransmission() bool {
	_, ok := getState().DamageTypeIndicatives.IsTransmission[int32(t)]
	return ok
}

// IsVandalism returns true if DamageType represents "Vandalism" value.
func (t DamageType) IsVandalism() bool {
	_, ok := getState().DamageTypeIndicatives.IsVandalism[int32(t)]
	return ok
}

// IsRightFront returns true if DamageType represents "RightFront" value.
func (t DamageType) IsRightFront() bool {
	_, ok := getState().DamageTypeIndicatives.IsRightFront[int32(t)]
	return ok
}

// IsSuspension returns true if DamageType represents "Suspension" value.
func (t DamageType) IsSuspension() bool {
	_, ok := getState().DamageTypeIndicatives.IsSuspension[int32(t)]
	return ok
}

// IsRearEnd returns true if DamageType represents "RearEnd" value.
func (t DamageType) IsRearEnd() bool {
	_, ok := getState().DamageTypeIndicatives.IsRearEnd[int32(t)]
	return ok
}

// IsRightSide returns true if DamageType represents "RightSide" value.
func (t DamageType) IsRightSide() bool {
	_, ok := getState().DamageTypeIndicatives.IsRightSide[int32(t)]
	return ok
}

// IsFlood returns true if DamageType represents "Flood" value.
func (t DamageType) IsFlood() bool {
	_, ok := getState().DamageTypeIndicatives.IsFlood[int32(t)]
	return ok
}

// IsFreshWater returns true if DamageType represents "FreshWater" value.
func (t DamageType) IsFreshWater() bool {
	_, ok := getState().DamageTypeIndicatives.IsFreshWater[int32(t)]
	return ok
}

// IsRollover returns true if DamageType represents "Rollover" value.
func (t DamageType) IsRollover() bool {
	_, ok := getState().DamageTypeIndicatives.IsRollover[int32(t)]
	return ok
}

// IsUndercarriage returns true if DamageType represents "Undercarriage" value.
func (t DamageType) IsUndercarriage() bool {
	_, ok := getState().DamageTypeIndicatives.IsUndercarriage[int32(t)]
	return ok
}

// IsRightRear returns true if DamageType represents "RightRear" value.
func (t DamageType) IsRightRear() bool {
	_, ok := getState().DamageTypeIndicatives.IsRightRear[int32(t)]
	return ok
}

// IsLeftFront returns true if DamageType represents "LeftFront" value.
func (t DamageType) IsLeftFront() bool {
	_, ok := getState().DamageTypeIndicatives.IsLeftFront[int32(t)]
	return ok
}

// IsInteriorBurn returns true if DamageType represents "InteriorBurn" value.
func (t DamageType) IsInteriorBurn() bool {
	_, ok := getState().DamageTypeIndicatives.IsInteriorBurn[int32(t)]
	return ok
}

// IsDamageHistory returns true if DamageType represents "DamageHistory" value.
func (t DamageType) IsDamageHistory() bool {
	_, ok := getState().DamageTypeIndicatives.IsDamageHistory[int32(t)]
	return ok
}

// IsRoof returns true if DamageType represents "Roof" value.
func (t DamageType) IsRoof() bool {
	_, ok := getState().DamageTypeIndicatives.IsRoof[int32(t)]
	return ok
}

// IsFrontAndRear returns true if DamageType represents "FrontAndRear" value.
func (t DamageType) IsFrontAndRear() bool {
	_, ok := getState().DamageTypeIndicatives.IsFrontAndRear[int32(t)]
	return ok
}

// IsHail returns true if DamageType represents "Hail" value.
func (t DamageType) IsHail() bool {
	_, ok := getState().DamageTypeIndicatives.IsHail[int32(t)]
	return ok
}

// IsMechanical returns true if DamageType represents "Mechanical" value.
func (t DamageType) IsMechanical() bool {
	_, ok := getState().DamageTypeIndicatives.IsMechanical[int32(t)]
	return ok
}

// IsSaltWater returns true if DamageType represents "SaltWater" value.
func (t DamageType) IsSaltWater() bool {
	_, ok := getState().DamageTypeIndicatives.IsSaltWater[int32(t)]
	return ok
}

// IsTotalBurn returns true if DamageType represents "TotalBurn" value.
func (t DamageType) IsTotalBurn() bool {
	_, ok := getState().DamageTypeIndicatives.IsTotalBurn[int32(t)]
	return ok
}

// IsBurn returns true if DamageType represents "Burn" value.
func (t DamageType) IsBurn() bool {
	_, ok := getState().DamageTypeIndicatives.IsBurn[int32(t)]
	return ok
}

// IsReplacedVin returns true if DamageType represents "ReplacedVin" value.
func (t DamageType) IsReplacedVin() bool {
	_, ok := getState().DamageTypeIndicatives.IsReplacedVin[int32(t)]
	return ok
}

// IsBiohazardOrChemical returns true if DamageType represents "BiohazardOrChemical" value.
func (t DamageType) IsBiohazardOrChemical() bool {
	_, ok := getState().DamageTypeIndicatives.IsBiohazardOrChemical[int32(t)]
	return ok
}

// IsLeftAndRight returns true if DamageType represents "LeftAndRight" value.
func (t DamageType) IsLeftAndRight() bool {
	_, ok := getState().DamageTypeIndicatives.IsLeftAndRight[int32(t)]
	return ok
}

// IsEngineBurn returns true if DamageType represents "EngineBurn" value.
func (t DamageType) IsEngineBurn() bool {
	_, ok := getState().DamageTypeIndicatives.IsEngineBurn[int32(t)]
	return ok
}

// IsLeftSide returns true if DamageType represents "LeftSide" value.
func (t DamageType) IsLeftSide() bool {
	_, ok := getState().DamageTypeIndicatives.IsLeftSide[int32(t)]
	return ok
}

// IsNeedsRepair returns true if DamageType represents "NeedsRepair" value.
func (t DamageType) IsNeedsRepair() bool {
	_, ok := getState().DamageTypeIndicatives.IsNeedsRepair[int32(t)]
	return ok
}

// IsStormDamage returns true if DamageType represents "StormDamage" value.
func (t DamageType) IsStormDamage() bool {
	_, ok := getState().DamageTypeIndicatives.IsStormDamage[int32(t)]
	return ok
}

// IsNone returns true if DamageType represents "None" value.
func (t DamageType) IsNone() bool {
	_, ok := getState().DamageTypeIndicatives.IsNone[int32(t)]
	return ok
}

// IsExteriorBurn returns true if DamageType represents "ExteriorBurn" value.
func (t DamageType) IsExteriorBurn() bool {
	_, ok := getState().DamageTypeIndicatives.IsExteriorBurn[int32(t)]
	return ok
}

// IsEngineDamage returns true if DamageType represents "EngineDamage" value.
func (t DamageType) IsEngineDamage() bool {
	_, ok := getState().DamageTypeIndicatives.IsEngineDamage[int32(t)]
	return ok
}

// IsMissingOrAlteredVin returns true if DamageType represents "MissingOrAlteredVin" value.
func (t DamageType) IsMissingOrAlteredVin() bool {
	_, ok := getState().DamageTypeIndicatives.IsMissingOrAlteredVin[int32(t)]
	return ok
}

// IsCashForClunkers returns true if DamageType represents "CashForClunkers" value.
func (t DamageType) IsCashForClunkers() bool {
	_, ok := getState().DamageTypeIndicatives.IsCashForClunkers[int32(t)]
	return ok
}
