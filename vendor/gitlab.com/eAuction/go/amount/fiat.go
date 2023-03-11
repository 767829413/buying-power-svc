package amount

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Fiat is a integer alias with custom marshaling/unmarshaling
// logic. Fiat is mostly needed for a consistent results in json responses
// and easy conversion from int values to string representations.
type Fiat int64

// FiatOne is the value of one whole unit of fiat currency.
const FiatOne uint = 100

// MarshalJSON marshals Fiat to json representation.
func (a Fiat) MarshalJSON() ([]byte, error) {
	return json.Marshal(StringWithOne(int64(a), FiatOne))
}

// UnmarshalJSON unmarshals Fiat from json representation.
func (a *Fiat) UnmarshalJSON(data []byte) error {
	var rawAmount string
	err := json.Unmarshal(data, &rawAmount)
	if err != nil {
		return errors.Wrap(err, "can't unmarshal amount")
	}

	rawA, err := ParseWithOne(rawAmount, FiatOne)
	if err != nil {
		return errors.Wrap(err, "can't parse amount")
	}

	*a = Fiat(rawA)
	return nil
}

func (a Fiat) String() string {
	return StringWithOne(int64(a), FiatOne)
}

// UnmarshalText parses decimal textual representation of Fiat.
func (a *Fiat) UnmarshalText(data []byte) error {
	rawA, err := ParseWithOne(string(data), FiatOne)
	if err != nil {
		return errors.Wrap(err, "can't parse amount")
	}

	*a = Fiat(rawA)
	return nil
}

// MarshalText codes decimal textual representation of Fiat.
func (a Fiat) MarshalText() ([]byte, error) {
	return []byte(a.String()), nil
}
