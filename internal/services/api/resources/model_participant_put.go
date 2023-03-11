/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ParticipantPut struct {
	Key
	Relationships ParticipantPutRelationships `json:"relationships"`
}
type ParticipantPutResponse struct {
	Data     ParticipantPut `json:"data"`
	Included Included       `json:"included"`
}

type ParticipantPutListResponse struct {
	Data     []ParticipantPut `json:"data"`
	Included Included         `json:"included"`
	Links    *Links           `json:"links"`
}

// MustParticipantPut - returns ParticipantPut from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustParticipantPut(key Key) *ParticipantPut {
	var participantPut ParticipantPut
	if c.tryFindEntry(key, &participantPut) {
		return &participantPut
	}
	return nil
}
