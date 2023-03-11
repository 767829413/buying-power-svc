package copart

import (
	"fmt"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

//auctionTokenData -.
type auctionTokenData struct {
	AuctionToken string `json:"auctionToken"`
}

//auctionTokenResponse -.
type auctionTokenResponse struct {
	g2AuctionResponseMeta
	Data auctionTokenData `json:"data"`
}

//GetAuctionToken - returns auction token for specified sale
func (c *connector) GetAuctionToken(saleID string) (string, error) {
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://g2auction.copart.com/g2/api/v1/buyer/login/%s/false?siteCode=&language=en&auctioneerMode=false", saleID), nil)

	if err != nil {
		panic(errors.Wrap(err, "failed to create request"))
	}

	var result auctionTokenResponse
	err = c.doJson(req, &result)
	if err != nil {
		return "", errors.Wrap(err, "failed to perform request")
	}

	return result.Data.AuctionToken, nil
}
