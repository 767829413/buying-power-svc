/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"time"

	"gitlab.com/eAuction/go/amount"
)

type DepositAttributes struct {
	// defines amount of the deposit in cents in the currency of the deposit.
	Amount    amount.Fiat `json:"amount"`
	CreatedAt time.Time   `json:"created_at"`
	// defines currency of the deposit in ISO4217
	Currency  string    `json:"currency"`
	State     string    `json:"state"`
	UpdatedAt time.Time `json:"updated_at"`
}
