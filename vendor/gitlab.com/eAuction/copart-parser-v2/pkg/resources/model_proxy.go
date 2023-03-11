/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Proxy struct {
	Key
	Attributes ProxyAttributes `json:"attributes"`
}
type ProxyResponse struct {
	Data     Proxy    `json:"data"`
	Included Included `json:"included"`
}

type ProxyListResponse struct {
	Data     []Proxy  `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustProxy - returns Proxy from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustProxy(key Key) *Proxy {
	var proxy Proxy
	if c.tryFindEntry(key, &proxy) {
		return &proxy
	}
	return nil
}
