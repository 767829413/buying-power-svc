package enumer

import (
	"fmt"

	"gitlab.com/eAuction/enumer/pkg/state/static"
)

// ConditionFromInt creates condition from itneger
func ConditionFromDamageLevel(i DamageLevel) Condition {
	condition, ok := getState().DamageLevelIndicatives[int32(i)]
	if !ok {
		return Condition(static.ConditionOther)
	}
	return Condition(condition)
}

// IsGood returns true if condition contains "good" value.
func (c Condition) IsGood() bool {
	_, ok := getState().ConditionIndicatives.IsGood[int32(c)]
	return ok
}

// IsDamage returns true if condition contains "damage" value.
func (c Condition) IsDamage() bool {
	_, ok := getState().ConditionIndicatives.IsDamage[int32(c)]
	return ok
}

// IsMinorDamage returns true if condition contains "minor_damage" value.
func (c Condition) IsMinorDamage() bool {
	_, ok := getState().ConditionIndicatives.IsMinorDamage[int32(c)]
	return ok
}

// ToDamageLevel - converts vehicle condition to dmg level
func (c Condition) ToDamageLevel() []DamageLevel {
	if !IsKnownCondition(int32(c)) {
		panic(fmt.Errorf("unexpected condition: %v", c))
	}
	var result []DamageLevel
	for k, v := range getState().DamageLevelIndicatives {
		if v == int32(c) {
			result = append(result, DamageLevel(k))
		}
	}
	return result
}

// ConditionsToDamageLevel converts slice of conditions to slice of dmg levels
func ConditionsToDamageLevel(conditions []Condition) []DamageLevel {
	var result []DamageLevel
	for _, condition := range conditions {
		result = append(result, condition.ToDamageLevel()...)
	}

	return result
}
