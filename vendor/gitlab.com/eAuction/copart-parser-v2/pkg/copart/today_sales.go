package copart

import (
	"encoding/json"
	"net/http"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

//Time - parses unix timestamp into time.Time
type Time time.Time

func (c *Time) UnmarshalJSON(data []byte) error {
	var raw int64
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal copart time", logan.F{
			"raw": string(data),
		})
	}

	*c = Time(time.Unix(raw/1000, 0))
	return nil
}

//Sale - defines copart sale
type Sale struct {
	AuctionKey      string `json:"auctionKey"`
	SaleUtcDateTime Time   `json:"saleUtcDateTime"`
}

type saleList struct {
	LiveSales  []Sale `json:"liveSales"`
	LaterSales []Sale `json:"laterSales"`
}

type todaysAuctionBody struct {
	SaleList saleList `json:"saleList"`
}

type todaysAuctionResponse struct {
	responseMeta
	Data todaysAuctionBody `json:"data"`
}

//GetUpcomingSales  returns list of upcoming sales
func (c *connector) GetUpcomingSales() ([]Sale, error) {
	const u = "https://www.copart.com/public/data/todaysAuctions?appId=g2&copartTimezonePref=%7B%22displayStr%22:%22EEST%22,%22offset%22:3,%22dst%22:true,%22windowsTz%22:%22Europe/Helsinki%22%7D"
	req, err := http.NewRequest(http.MethodGet, u, nil)

	if err != nil {
		return nil, errors.Wrap(err, "failed to get upcoming sales")
	}

	var response todaysAuctionResponse
	err = c.doJson(req, &response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get upcoming sales")
	}

	result := response.Data.SaleList.LiveSales
	return append(result, response.Data.SaleList.LaterSales...), nil

}
