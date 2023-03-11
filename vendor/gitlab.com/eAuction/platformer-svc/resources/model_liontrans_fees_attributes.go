/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type LiontransFeesAttributes struct {
	ContainerFees      []ContainerFee `json:"container_fees"`
	Currency           string         `json:"currency"`
	DestinationCountry string         `json:"destination_country"`
	DestinationPort    string         `json:"destination_port"`
	ForwarderFee       float64        `json:"forwarder_fee"`
	InsuranceFee       float64        `json:"insurance_fee"`
	// defines fee charged by platform
	PlatformFee    float64 `json:"platform_fee"`
	TransactionFee float64 `json:"transaction_fee"`
}
