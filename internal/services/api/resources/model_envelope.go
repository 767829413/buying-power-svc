/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Envelope struct {
	Key
	Relationships EnvelopeRelationships `json:"relationships"`
}
type EnvelopeResponse struct {
	Data     Envelope `json:"data"`
	Included Included `json:"included"`
}

type EnvelopeListResponse struct {
	Data     []Envelope `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustEnvelope - returns Envelope from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustEnvelope(key Key) *Envelope {
	var envelope Envelope
	if c.tryFindEntry(key, &envelope) {
		return &envelope
	}
	return nil
}
