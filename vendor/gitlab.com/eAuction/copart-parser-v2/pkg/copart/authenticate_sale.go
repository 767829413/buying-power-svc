package copart

import (
	"fmt"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

//AuthenticateSale - auths sale
func (c *connector) AuthenticateSale(saleID string) error {
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://g2auction.copart.com/g2/authenticate/api/v1/lot/%s/false", saleID), nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to create request"))
	}

	var result g2AuctionResponseMeta
	err = c.doJson(req, &result)
	if err != nil {
		return errors.Wrap(err, "failed to perform request")
	}

	return nil
}
