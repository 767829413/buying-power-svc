/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Vehicle struct {
	Key
	Attributes VehicleAttributes `json:"attributes"`
}
type VehicleResponse struct {
	Data     Vehicle  `json:"data"`
	Included Included `json:"included"`
}

type VehicleListResponse struct {
	Data     []Vehicle `json:"data"`
	Included Included  `json:"included"`
	Links    *Links    `json:"links"`
}

// MustVehicle - returns Vehicle from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustVehicle(key Key) *Vehicle {
	var vehicle Vehicle
	if c.tryFindEntry(key, &vehicle) {
		return &vehicle
	}
	return nil
}
