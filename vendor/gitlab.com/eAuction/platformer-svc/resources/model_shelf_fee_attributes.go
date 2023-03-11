/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "gitlab.com/eAuction/go/amount"

// Defines fee charged by shelf
type ShelfFeeAttributes struct {
	Currency string      `json:"currency"`
	Fixed    amount.Fiat `json:"fixed"`
}
