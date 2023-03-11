/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type PutWinner struct {
	Key
	Attributes    PutWinnerAttributes    `json:"attributes"`
	Relationships PutWinnerRelationships `json:"relationships"`
}
type PutWinnerResponse struct {
	Data     PutWinner `json:"data"`
	Included Included  `json:"included"`
}

type PutWinnerListResponse struct {
	Data     []PutWinner `json:"data"`
	Included Included    `json:"included"`
	Links    *Links      `json:"links"`
}

// MustPutWinner - returns PutWinner from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPutWinner(key Key) *PutWinner {
	var putWinner PutWinner
	if c.tryFindEntry(key, &putWinner) {
		return &putWinner
	}
	return nil
}
