/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "gitlab.com/eAuction/go/amount"

type BidAttributes struct {
	// defines amount of bid
	Amount amount.Fiat `json:"amount"`
	// defines currency of the bid
	Currency string `json:"currency"`
	// defines next bid amount
	NextBidAmount amount.Fiat `json:"next_bid_amount"`
}
