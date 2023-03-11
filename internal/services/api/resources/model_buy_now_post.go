/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type BuyNowPost struct {
	Key
	Relationships BuyNowPostRelationships `json:"relationships"`
}
type BuyNowPostResponse struct {
	Data     BuyNowPost `json:"data"`
	Included Included   `json:"included"`
}

type BuyNowPostListResponse struct {
	Data     []BuyNowPost `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustBuyNowPost - returns BuyNowPost from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBuyNowPost(key Key) *BuyNowPost {
	var buyNowPost BuyNowPost
	if c.tryFindEntry(key, &buyNowPost) {
		return &buyNowPost
	}
	return nil
}
