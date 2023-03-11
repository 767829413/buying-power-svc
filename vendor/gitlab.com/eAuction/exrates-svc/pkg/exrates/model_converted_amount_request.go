/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package exrates

type ConvertedAmountRequest struct {
	Key
	Attributes    ConvertedAmountRequestAttributes    `json:"attributes"`
	Relationships ConvertedAmountRequestRelationships `json:"relationships"`
}
type ConvertedAmountRequestResponse struct {
	Data     ConvertedAmountRequest `json:"data"`
	Included Included               `json:"included"`
}

type ConvertedAmountRequestListResponse struct {
	Data     []ConvertedAmountRequest `json:"data"`
	Included Included                 `json:"included"`
	Links    *Links                   `json:"links"`
}

// MustConvertedAmountRequest - returns ConvertedAmountRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustConvertedAmountRequest(key Key) *ConvertedAmountRequest {
	var convertedAmountRequest ConvertedAmountRequest
	if c.tryFindEntry(key, &convertedAmountRequest) {
		return &convertedAmountRequest
	}
	return nil
}
