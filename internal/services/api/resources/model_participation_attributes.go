/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"time"

	"gitlab.com/eAuction/go/amount"
)

type ParticipationAttributes struct {
	// Maximum bid to be set automatically (if user has enabled autobid option)
	AutoBidLimit amount.Fiat `json:"auto_bid_limit"`
	// Bid limit set to current participant (defined by admin, not by user).
	BidLimit amount.Fiat `json:"bid_limit"`
	// Time when participation was created
	CreatedAt time.Time `json:"created_at"`
	// Current bid of participant
	CurrentBid amount.Fiat `json:"current_bid"`
	// Defines whether participant has requested buy now option.
	RequestedBuyNow bool `json:"requested_buy_now"`
	// Current state of participation
	State string `json:"state"`
	// Time when participation was last updated (if not updated at all - is equal to `created_at`)
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
