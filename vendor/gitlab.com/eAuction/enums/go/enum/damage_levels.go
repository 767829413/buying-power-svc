package enum

import "gitlab.com/eAuction/enums/go/pkg/alphanum"

func (x Damage_Level) ReadableString() (string, bool) {
	switch x {
	case Damage_LEVEL_LOW:
		return "low", true
	case Damage_LEVEL_LOW_MEDIUM:
		return "low-medium", true
	case Damage_LEVEL_MEDIUM:
		return "medium", true
	case Damage_LEVEL_MEDIUM_HIGH:
		return "medium-high", true
	case Damage_LEVEL_HIGH:
		return "high", true
	case Damage_LEVEL_NULL:
		return "", true
	default:
		return "", false
	}
}

func DamageLevel(s string) Damage_Level {
	switch alphanum.Lower(s) {
	case "low":
		return Damage_LEVEL_LOW
	case "lowmedium":
		return Damage_LEVEL_LOW_MEDIUM
	case "medium":
		return Damage_LEVEL_MEDIUM
	case "mediumhigh":
		return Damage_LEVEL_MEDIUM_HIGH
	case "high":
		return Damage_LEVEL_HIGH
	case "":
		return Damage_LEVEL_NULL
	default:
		return Damage_LEVEL_UNKNOWN
	}
}

func FindDamageLevel(state Vehicle_State, damageType string) (Damage_Level, bool) {
	if state == Vehicle_STATE_NEW {
		return Damage_LEVEL_LOW, true
	}

	for lvl, states := range damageLevels {
		for s, dmgs := range states {
			if VehicleState(s) != state {
				continue
			}
			for _, dmg := range dmgs {
				if alphanum.Lower(dmg) == alphanum.Lower(damageType) {
					return DamageLevel(lvl), true
				}
			}
		}
	}

	if state == Vehicle_STATE_RUNS_AND_DRIVES {
		return Damage_LEVEL_MEDIUM, true
	}

	return Damage_LEVEL_UNKNOWN, false
}
