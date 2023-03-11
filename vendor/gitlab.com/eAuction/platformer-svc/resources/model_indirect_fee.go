/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type IndirectFee struct {
	// defines additional auction fees applied for the lot
	AuctionFee []RangeFee `json:"auction_fee"`
	// defines platform for whose lots fees are applied
	Platform string `json:"platform"`
	ShelfFee *Fee   `json:"shelf_fee,omitempty"`
}
