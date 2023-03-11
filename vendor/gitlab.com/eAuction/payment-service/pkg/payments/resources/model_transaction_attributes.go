/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"time"

	"gitlab.com/eAuction/go/amount"
)

type TransactionAttributes struct {
	// Payee's account ID in Shelf.Network.
	AccountId   *string     `json:"account_id,omitempty"`
	Amount      amount.Fiat `json:"amount"`
	ConfirmedAt *time.Time  `json:"confirmed_at,omitempty"`
	CratedAt    *time.Time  `json:"crated_at,omitempty"`
	// Transaction currency [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code.
	Currency string `json:"currency"`
	// Optional transaction description. Up to 125 characters long.
	Description *string `json:"description,omitempty"`
	// Arbitrary invoice ID. Up to 255 characters long. Used for auto-processing.
	InvoiceId string `json:"invoice_id"`
	// [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) code of the desired UI language of the checkout form. *May not be supported by transaction processor.*
	Language *string `json:"language,omitempty"`
	// Additional info. E.g. charged amount for confirmed transactions, statuses.
	Meta *map[string]interface{} `json:"meta,omitempty"`
	// Transaction purpose.
	Purpose string `json:"purpose"`
	// Custom redirect URL template. The domain must be added to the whitelist in payment service config. `{{TRANSACTION}}` part in template will be replaced with the transaction ID.
	Redirect   *string    `json:"redirect,omitempty"`
	RefundedAt *time.Time `json:"refunded_at,omitempty"`
	// Transaction status.
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}
