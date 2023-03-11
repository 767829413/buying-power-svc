package enum

import (
	"gitlab.com/eAuction/enums/go/pkg/alphanum"
)

func (x Vehicle_State) ReadableString() (string, bool) {
	switch x {
	case Vehicle_STATE_RUNS_AND_DRIVES:
		return "RUNS AND DRIVES", true
	case Vehicle_STATE_ENGINE_START_PROGRAM:
		return "ENGINE START PROGRAM", true
	case Vehicle_STATE_ENHANCED:
		return "ENHANCED VEHICLES", true
	case Vehicle_STATE_NEW:
		return "NEW", true
	case Vehicle_STATE_NULL:
		return "", true
	default:
		return "", false
	}
}

// VehicleState builds Vehicle_State from its readable string representation
func VehicleState(s string) Vehicle_State {
	if s == "NEW" {
		return Vehicle_STATE_NEW
	}
	return CopartVehicleState(s)
}

// CopartVehicleState returns enum value for the provided copart "highlights" strings.
func CopartVehicleState(s string) Vehicle_State {
	switch alphanum.Lower(s) {
	case "enginestartprogram":
		return Vehicle_STATE_ENGINE_START_PROGRAM
	case "enhancedvehicles":
		return Vehicle_STATE_ENHANCED
	case "runsanddrives":
		return Vehicle_STATE_RUNS_AND_DRIVES
	case "":
		return Vehicle_STATE_NULL
	default:
		return Vehicle_STATE_UNKNOWN
	}
}

// IaaIVehicleState returns enum value for the provided iaai "run_and_drive" and "starts" strings.
func IaaIVehicleState(runsAndDrive, starts string) Vehicle_State {
	if runsAndDrive == "" && starts == "" {
		return Vehicle_STATE_NULL
	}

	if alphanum.Lower(runsAndDrive) == "yes" {
		return Vehicle_STATE_RUNS_AND_DRIVES
	}

	switch alphanum.Lower(starts) {
	case "yes", "vehiclestarts", "startswjump":
		return Vehicle_STATE_ENGINE_START_PROGRAM
	case "no", "wontstart", "didnttest", "canttest", "enginedamage":
		return Vehicle_STATE_ENHANCED
	default:
		return Vehicle_STATE_UNKNOWN
	}
}

func AutoPapaVehicleState(state string) Vehicle_State {
	switch alphanum.Lower(state) {
	case "new":
		return Vehicle_STATE_NEW
	case "good", "needsrepair":
		return Vehicle_STATE_RUNS_AND_DRIVES
	case "":
		return Vehicle_STATE_NULL
	default:
		return Vehicle_STATE_UNKNOWN
	}
}
