/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "gitlab.com/eAuction/go/amount"

type DepositOptionAttributes struct {
	// defines amount of the deposit in cents
	Amount amount.Fiat `json:"amount"`
	// defines currency of the amount in ISO4217
	AmountCurrency string `json:"amount_currency"`
	// defines amount of the buying power
	BuyingPower amount.Fiat `json:"buying_power"`
	// defines currency of the buying power in ISO4217
	BuyingPowerCurrency string             `json:"buying_power_currency"`
	State               DepositOptionState `json:"state"`
}
