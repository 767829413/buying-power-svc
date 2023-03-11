package copart

import (
	"fmt"
	"net/http"
	"time"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

//SaleData - represents data of the sale
type SaleData struct {
	AuctionKey         string `json:"auctionKey"`
	AuctionID          int    `json:"auctionId"`
	Dbdate             string `json:"dbdate"`
	IsAdsSale          bool   `json:"isAdsSale"`
	YardNumber         int    `json:"yardNumber"`
	LaneCode           string `json:"laneCode"`
	SaleStartTimestamp string `json:"saleStartTimestamp"`
}

//salesResponse - presents response on sale request
type salesResponse struct {
	g2AuctionResponseMeta
	Data []SaleData `json:"data"`
}

func (c *connector) GetSales() ([]SaleData, error) {
	now := time.Now().UTC()
	from := now.Add(-24 * time.Hour).Truncate(time.Hour)
	to := now.Add(12 * time.Hour).Truncate(time.Hour)
	const layout = "2006-01-02:1504"
	u := fmt.Sprintf("https://g2auction.copart.com/g2/api/v1/sales?siteCode=&fromDateTimeUTC=%s&toDateTimeUTC=%s", from.Format(layout), to.Format(layout))
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to create request"))
	}

	var result salesResponse
	err = c.doJson(req, &result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform request")
	}

	return result.Data, nil
}

//GetSale - returns sale by it's id
func (c *connector) GetSale(saleID string) (*SaleData, error) {
	sales, err := c.GetSales()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sale data")
	}

	for _, s := range sales {
		if s.AuctionKey == saleID {
			return &s, nil
		}
	}

	return nil, nil
}
