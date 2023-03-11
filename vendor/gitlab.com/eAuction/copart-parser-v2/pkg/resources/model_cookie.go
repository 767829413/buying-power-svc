/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Cookie struct {
	Key
	Attributes CookieAttributes `json:"attributes"`
}
type CookieResponse struct {
	Data     Cookie   `json:"data"`
	Included Included `json:"included"`
}

type CookieListResponse struct {
	Data     []Cookie `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustCookie - returns Cookie from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCookie(key Key) *Cookie {
	var cookie Cookie
	if c.tryFindEntry(key, &cookie) {
		return &cookie
	}
	return nil
}
