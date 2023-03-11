/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "gitlab.com/eAuction/go/amount"

type SoldAttributes struct {
	Amount amount.Fiat `json:"amount"`
	// sale is not finilized yet
	IsPending bool `json:"is_pending"`
}
