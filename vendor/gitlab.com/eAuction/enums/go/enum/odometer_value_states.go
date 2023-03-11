package enum

import (
	"gitlab.com/eAuction/enums/go/pkg/alphanum"
)

// OdometerValueState returns enum value for provided "odometer_value_state" string.
func OdometerValueState(s string) OdometerValue_State {
	switch alphanum.Lower(s) {
	case "actual":
		return OdometerValue_STATE_ACTUAL
	case "broken":
		return OdometerValue_STATE_BROKEN
	case "burnt":
		return OdometerValue_STATE_BURNT
	case "exceedsmechanicallimits":
		return OdometerValue_STATE_EXCEEDS_MECHANICAL_LIMITS
	case "exempt":
		return OdometerValue_STATE_EXEMPT
	case "inoperabledigitaldash":
		return OdometerValue_STATE_INOPERABLE_DIGITAL_DASH
	case "missing":
		return OdometerValue_STATE_MISSING
	case "notactual":
		return OdometerValue_STATE_NOT_ACTUAL
	case "notrequiredexempt":
		return OdometerValue_STATE_NOT_REQUIRED_OR_EXEMPT
	case "unknown", "":
		return OdometerValue_STATE_NULL
	default:
		return OdometerValue_STATE_UNKNOWN
	}
}
