/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Participation struct {
	Key
	Attributes    ParticipationAttributes    `json:"attributes"`
	Relationships ParticipationRelationships `json:"relationships"`
}
type ParticipationResponse struct {
	Data     Participation `json:"data"`
	Included Included      `json:"included"`
}

type ParticipationListResponse struct {
	Data     []Participation `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustParticipation - returns Participation from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustParticipation(key Key) *Participation {
	var participation Participation
	if c.tryFindEntry(key, &participation) {
		return &participation
	}
	return nil
}
