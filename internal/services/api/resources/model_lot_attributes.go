/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"time"

	"gitlab.com/eAuction/go/amount"
)

type LotAttributes struct {
	// Fee to be paid for business when buying a car - already included in `business_total_price` (non-zero only for direct sales).
	BusinessFee amount.Fiat `json:"business_fee"`
	// Price of direct sale lot to be bought by business (non-zero only for direct sales).
	BusinessTotalPrice amount.Fiat `json:"business_total_price"`
	// Price to buy lot item without competition (if 0.00 - buy now option is turned off for the lot)
	BuyNowPrice amount.Fiat `json:"buy_now_price"`
	// Time when lot was created (may be different with `start_time`)
	CreatedAt time.Time `json:"created_at"`
	// Currency to calculate some of lot properties
	Currency string `json:"currency"`
	// Fee to be paid for customer when buying a car - already included in `customer_total_price` (non-zero only for direct sales).
	CustomerFee amount.Fiat `json:"customer_fee"`
	// Price of direct sale lot to be bought by customer (non-zero only for direct sales).
	CustomerTotalPrice amount.Fiat `json:"customer_total_price"`
	// Lot duration in seconds.
	Duration int64 `json:"duration"`
	// Time when lot finishes accepting bids.
	EndTime time.Time `json:"end_time"`
	// Time when lot was actually claimed as finished/sold.
	EndedAt *time.Time `json:"ended_at,omitempty"`
	// Highest bid for current time (currency is defined by `currency` field)
	HighestBid amount.Fiat `json:"highest_bid"`
	IsBidFinal bool        `json:"is_bid_final"`
	ItemNumber *int64      `json:"item_number,omitempty"`
	LotType    string      `json:"lot_type"`
	LotTypeI   int32       `json:"lot_type_i"`
	// Maximum step to update bid.
	MaxBidStep amount.Fiat `json:"max_bid_step"`
	// Minimum step to update bid.
	MinBidStep amount.Fiat `json:"min_bid_step"`
	Name       string      `json:"name"`
	// Price to start a bidding process from.
	StartPrice amount.Fiat `json:"start_price"`
	// Time lot becomes open for bids.
	StartTime time.Time `json:"start_time"`
	// Defines current state of the lot.
	State string `json:"state"`
	// Last time when lot state was updated.
	StateUpdatedAt time.Time `json:"state_updated_at"`
}
