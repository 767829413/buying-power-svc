/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"time"

	"gitlab.com/eAuction/go/amount"
)

type PutWinnerAttributes struct {
	// ID of container.
	ContainerId *string `json:"container_id,omitempty"`
	// Link of container.
	ContainerLink *string    `json:"container_link,omitempty"`
	DeliveredAt   *time.Time `json:"delivered_at,omitempty"`
	// Fee amount to be paid (or already paid).
	Fee amount.Fiat `json:"fee"`
	// Currency of fee to be paid in.
	FeeCurrency string `json:"fee_currency"`
	// Time when the invoice was created.
	InvoiceCreatedAt *time.Time `json:"invoice_created_at,omitempty"`
	// ID of invoice to be paid (or already paid).
	InvoiceId *string `json:"invoice_id,omitempty"`
	// Current state of lot winner.
	State string `json:"state"`
	// Price for transporting the vehicle.
	TransportationPrice amount.Fiat `json:"transportation_price"`
	// Currency of transportation price.
	TransportationPriceCurrency string `json:"transportation_price_currency"`
}
