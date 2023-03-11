package copart

import (
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type DynamicLotDetails struct {
	ErrorCode        string  `json:"errorCode"`
	BuyerNumber      int     `json:"buyerNumber"`
	Source           string  `json:"source"`
	BidIncrement     float64 `json:"bidIncrement"`
	BuyTodayFlag     string  `json:"buyTodayFlag"`
	BuyTodayBid      float64 `json:"buyTodayBid"`
	CurrentBid       float64   `json:"currentBid"`
	TotalAmountDue   float64 `json:"totalAmountDue"`
	SealedBid        bool    `json:"sealedBid"`
	MinBid           float64 `json:"minBid"`
	MaxBid           float64 `json:"maxBid"`
	CanBid           string  `json:"canBid"`
	FirstBid         bool    `json:"firstBid"`
	HasBid           bool    `json:"hasBid"`
	SellerReserveMet bool    `json:"sellerReserveMet"`
	LotSold          bool    `json:"lotSold"`
	BidStatus        string  `json:"bidStatus"`
	SaleStatus       string  `json:"saleStatus"`
	LotAuctionStatus string  `json:"lotAuctionStatus"`
	CounterBidStatus string  `json:"counterBidStatus"`
	RandomHash       string  `json:"randomHash"`
	MakeOfferEnabled bool    `json:"makeOfferEnabled"`
	BuyerHighBidder  bool    `json:"buyerHighBidder"`
	Anonymous        bool    `json:"anonymous"`
	NonSyncedBuyer   bool    `json:"nonSyncedBuyer"`
}

type dynamicLotDetailsData struct {
	LotDetails DynamicLotDetails `json:"lotDetails"`
}

type dynamicLotDetailsResponse struct {
	responseMeta
	Data dynamicLotDetailsData `json:"data"`
}

//GetDynamicLotDetails - returns details for specified lot ID. Returns nil, nil if not found
func (c *connector) GetDynamicLotDetails(lotID string) (*DynamicLotDetails, error) {
	r, err := http.NewRequest(http.MethodGet, "https://www.copart.com/data/lotdetails/dynamic/"+lotID, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to create request"))
	}

	var response dynamicLotDetailsResponse
	err = c.doJson(r, &response)
	if err != nil {
		if errors.Cause(err) == ErrNotFound {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to perform request")
	}

	return &response.Data.LotDetails, nil
}
