/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ContainerFees struct {
	Key
	Attributes ContainerFeesAttributes `json:"attributes"`
}
type ContainerFeesResponse struct {
	Data     ContainerFees `json:"data"`
	Included Included      `json:"included"`
}

type ContainerFeesListResponse struct {
	Data     []ContainerFees `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustContainerFees - returns ContainerFees from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustContainerFees(key Key) *ContainerFees {
	var containerFees ContainerFees
	if c.tryFindEntry(key, &containerFees) {
		return &containerFees
	}
	return nil
}
