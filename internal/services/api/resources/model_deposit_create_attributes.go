/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "gitlab.com/eAuction/go/amount"

type DepositCreateAttributes struct {
	// defines amount of the deposit in cents in the currency of the deposit
	Amount amount.Fiat `json:"amount"`
	// defines currency of the deposit in ISO4217
	Currency string `json:"currency"`
}
