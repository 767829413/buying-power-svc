package copart

import (
	"net/http"
	"strconv"
	"strings"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

//buyerNumberData - defines data of the buyer number request
type buyerNumberData struct {
	BuyerNumber string `json:"buyerNumber"`
}

//buyerNumberResponse - defines response for buyer number request
type buyerNumberResponse struct {
	g2AuctionResponseMeta
	Data buyerNumberData `json:"data"`
}

//BuyerNumber - represents buyer number
type BuyerNumber struct {
	Primary   int64
	Secondary string
}

func (c *connector) parseBuyerNumber(raw string) (BuyerNumber, error) {
	if c.isAnonymous {
		primary, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return BuyerNumber{}, errors.Wrap(err, "failed to parse buyer number for anonymous connector", logan.F{
				"buyer_number": raw,
			})
		}

		return BuyerNumber{
			Primary:   primary,
			Secondary: "",
		}, nil
	}
	rawBuyerNumber := strings.Split(raw, "-")
	if len(rawBuyerNumber) != 2 {
		return BuyerNumber{}, errors.From(errors.New("unexpected format of buyer number"), logan.F{
			"buyer_number": raw,
		})
	}

	primary, err := strconv.ParseInt(rawBuyerNumber[0], 10, 64)
	if err != nil {
		return BuyerNumber{}, errors.Wrap(err, "failed to parse buyers primary part", logan.F{
			"buyer_number": raw,
		})
	}

	return BuyerNumber{
		Primary:   primary,
		Secondary: rawBuyerNumber[1],
	}, nil
}

//GetBuyerNumber - returns buyer number
func (c *connector) GetBuyerNumber() BuyerNumber {
	return c.buyerNumber.Do(func() interface{} {
		defer func() {
			if rec := recover(); rec != nil {
				c.Close()
				panic(rec)
			}
		}()
		req, err := http.NewRequest(http.MethodGet, "https://g2auction.copart.com/g2/api/v1/buyerNumber/get", nil)
		if err != nil {
			panic(errors.Wrap(err, "failed to create request"))
		}

		req.Header.Add("Sec-Fetch-Site", "same-origin")
		req.Header.Add("Sec-Fetch-Mode", "cors")
		req.Header.Add("Sec-Fetch-Dest", "empty")
		req.Header.Add("Referer", "https://g2auction.copart.com/g2/")
		req.Header.Add("Accept-Encoding", "gzip, deflate, br")
		req.Header.Add("Accept-Language", "en-US,en;q=0.9")
		req.Header.Add("Accept", "application/json, text/plain, */*")

		var result buyerNumberResponse
		err = c.doJson(req, &result)
		if err != nil {
			panic(errors.Wrap(err, "failed to perform request"))
		}

		buyerNumber, err := c.parseBuyerNumber(result.Data.BuyerNumber)
		if err != nil {
			panic(errors.Wrap(err, "failed to parse buyer number"))
		}

		return buyerNumber
	}).(BuyerNumber)

}
