/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package exrates

import (
	"time"
)

type ExchangeRateAttributes struct {
	// defines exchange rate for asset pair
	Price     float64   `json:"price"`
	UpdatedAt time.Time `json:"updated_at"`
}
