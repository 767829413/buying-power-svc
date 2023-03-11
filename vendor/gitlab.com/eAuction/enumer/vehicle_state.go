package enumer

func (t VehicleState) IsNew() bool {
	_, ok := getState().VehicleStateIndicatives.IsNew[int32(t)]
	return ok
}

func (t VehicleState) IsRunsAndDrives() bool {
	_, ok := getState().VehicleStateIndicatives.IsRunsAndDrives[int32(t)]
	return ok
}

func (t VehicleState) IsEnhanced() bool {
	_, ok := getState().VehicleStateIndicatives.IsEnhanced[int32(t)]
	return ok
}

func (t VehicleState) IsEngineStartProgram() bool {
	_, ok := getState().VehicleStateIndicatives.IsEngineStartProgram[int32(t)]
	return ok
}
