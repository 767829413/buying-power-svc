/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type EnvelopeRequest struct {
	Key
	Relationships EnvelopeRequestRelationships `json:"relationships"`
}
type EnvelopeRequestResponse struct {
	Data     EnvelopeRequest `json:"data"`
	Included Included        `json:"included"`
}

type EnvelopeRequestListResponse struct {
	Data     []EnvelopeRequest `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
}

// MustEnvelopeRequest - returns EnvelopeRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustEnvelopeRequest(key Key) *EnvelopeRequest {
	var envelopeRequest EnvelopeRequest
	if c.tryFindEntry(key, &envelopeRequest) {
		return &envelopeRequest
	}
	return nil
}
