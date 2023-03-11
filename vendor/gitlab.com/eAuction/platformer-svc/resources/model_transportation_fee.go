/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type TransportationFee struct {
	Key
	Attributes TransportationFeeAttributes `json:"attributes"`
}
type TransportationFeeResponse struct {
	Data     TransportationFee `json:"data"`
	Included Included          `json:"included"`
}

type TransportationFeeListResponse struct {
	Data     []TransportationFee `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
}

// MustTransportationFee - returns TransportationFee from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustTransportationFee(key Key) *TransportationFee {
	var transportationFee TransportationFee
	if c.tryFindEntry(key, &transportationFee) {
		return &transportationFee
	}
	return nil
}
