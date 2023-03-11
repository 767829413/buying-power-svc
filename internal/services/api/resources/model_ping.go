/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Ping struct {
	Key
	Attributes PingAttributes `json:"attributes"`
}
type PingResponse struct {
	Data     Ping     `json:"data"`
	Included Included `json:"included"`
}

type PingListResponse struct {
	Data     []Ping   `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustPing - returns Ping from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPing(key Key) *Ping {
	var ping Ping
	if c.tryFindEntry(key, &ping) {
		return &ping
	}
	return nil
}
