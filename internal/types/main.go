package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

//UnmarshalSlice - tries to unmarshal slice of enums from raw slice of strings
func UnmarshalSlice(raw []string, res interface{}) error {
	asJSON := make([]string, len(raw))
	for i := range raw {
		asJSON[i] = fmt.Sprintf(`"%s"`, raw[i])
	}

	return json.Unmarshal([]byte(fmt.Sprintf(`[%s]`, strings.Join(asJSON, ","))), &res)
}
