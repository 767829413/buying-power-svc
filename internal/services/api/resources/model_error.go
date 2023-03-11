/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Error struct {
	Key
	Attributes ErrorAttributes `json:"attributes"`
}
type ErrorResponse struct {
	Data     Error    `json:"data"`
	Included Included `json:"included"`
}

type ErrorListResponse struct {
	Data     []Error  `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustError - returns Error from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustError(key Key) *Error {
	var error Error
	if c.tryFindEntry(key, &error) {
		return &error
	}
	return nil
}
