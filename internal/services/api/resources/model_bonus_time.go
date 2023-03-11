/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type BonusTime struct {
	Key
	Attributes    BonusTimeAttributes    `json:"attributes"`
	Relationships BonusTimeRelationships `json:"relationships"`
}
type BonusTimeResponse struct {
	Data     BonusTime `json:"data"`
	Included Included  `json:"included"`
}

type BonusTimeListResponse struct {
	Data     []BonusTime `json:"data"`
	Included Included    `json:"included"`
	Links    *Links      `json:"links"`
}

// MustBonusTime - returns BonusTime from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBonusTime(key Key) *BonusTime {
	var bonusTime BonusTime
	if c.tryFindEntry(key, &bonusTime) {
		return &bonusTime
	}
	return nil
}
