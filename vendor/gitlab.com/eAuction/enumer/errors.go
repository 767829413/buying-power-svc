package enumer

import (
	"fmt"

	"gitlab.com/distributed_lab/logan/v3"
)

type UnknownAliasError struct {
	Enum  string
	Value string
}

func (u UnknownAliasError) Error() string {
	return fmt.Sprintf("unknown alias %s for enum %s", u.Value, u.Enum)
}

func (u UnknownAliasError) GetFields() logan.F {
	return map[string]interface{}{"enum": u.Enum, "value": u.Value}
}
