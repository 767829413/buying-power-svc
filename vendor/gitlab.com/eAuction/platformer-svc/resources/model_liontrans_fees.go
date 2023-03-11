/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type LiontransFees struct {
	Key
	Attributes LiontransFeesAttributes `json:"attributes"`
}
type LiontransFeesResponse struct {
	Data     LiontransFees `json:"data"`
	Included Included      `json:"included"`
}

type LiontransFeesListResponse struct {
	Data     []LiontransFees `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustLiontransFees - returns LiontransFees from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustLiontransFees(key Key) *LiontransFees {
	var liontransFees LiontransFees
	if c.tryFindEntry(key, &liontransFees) {
		return &liontransFees
	}
	return nil
}
