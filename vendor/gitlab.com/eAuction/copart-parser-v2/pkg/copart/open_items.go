package copart

import (
	"bytes"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type OpenItemsResponse struct {
	BidStatusMeta
	Data []OpenItem `json:"aaData"`
}

type SmallImage struct {
	URL                  string `json:"url"`
	SwiftFlag            bool   `json:"swiftFlag"`
	FrameCount           int    `json:"frameCount"`
	Status               string `json:"status"`
	ImageTypeDescription string `json:"imageTypeDescription"`
	HighRes              bool   `json:"highRes"`
}

type Address struct {
	Line2   string `json:"line2"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type Yard struct {
	Number    int      `json:"number"`
	Name      string   `json:"name"`
	Address   Address  `json:"address"`
	Currency  Currency `json:"currency"`
	MileAway  int      `json:"mileAway"`
	NumberStr string   `json:"numberStr"`
	SaleType  string   `json:"saleType"`
}

type OpenItem struct {
	LotID                       LotNumber  `json:"lotId"`
	AuctionDateUtc              int64      `json:"auctionDateUtc"`
	BidderNumber                string     `json:"bidderNumber"`
	BidderStatus                string     `json:"bidderStatus"`
	BidderName                  string     `json:"bidderName"`
	CurrentBid                  float64    `json:"currentBid"`
	ActualCostOfVehicle         float64    `json:"actualCostOfVehicle"`
	DamageType                  string     `json:"damageType"`
	DamageTypeSecondary         string     `json:"damageTypeSecondary"`
	DmvState                    string     `json:"dmvState"`
	ItemNumber                  int        `json:"itemNumber"`
	LossType                    string     `json:"lossType"`
	MakeOfVehicle               string     `json:"makeOfVehicle"`
	ModelOfVehicle              string     `json:"modelOfVehicle"`
	OdometerReading             float64    `json:"odometerReading"`
	OdometerType                string     `json:"odometerType"`
	SaleDate                    Date       `json:"saleDate"`
	SaleTitleType               string     `json:"saleTitleType"`
	ImageURL                    string     `json:"imageUrl"`
	VehicleIdentificationNumber string     `json:"vehicleIdentificationNumber"`
	YearOfVehicle               int        `json:"yearOfVehicle"`
	AutoCheckFlag               string     `json:"autoCheckFlag"`
	MarketGuideFlag             string     `json:"marketGuideFlag"`
	LotIcons                    []string   `json:"lotIcons"`
	AuctionLane                 string     `json:"auctionLane"`
	BuyItNowPrice               float64    `json:"buyItNowPrice"`
	CounterBidStatus            string     `json:"counterBidStatus"`
	MyBid                       float64    `json:"myBid"`
	BidStatusStr                string     `json:"bidStatusStr"`
	DocType                     string     `json:"docType"`
	AuctionHostID               int        `json:"auctionHostId"`
	YardName                    string     `json:"yardName"`
	YardCurrencyCode            string     `json:"yardCurrencyCode"`
	PhysicalYardNumber          int        `json:"physicalYardNumber"`
	AuctionDateType             string     `json:"auctionDateType"`
	AwardedHighBid              float64    `json:"awardedHighBid"`
	FutureSale                  bool       `json:"futureSale"`
	IsWatchable                 bool       `json:"isWatchable"`
	SmallImage                  SmallImage `json:"smallImage"`
	Yard                        Yard       `json:"yard"`
	Brand                       string     `json:"brand"`
}

func (c *connector) GetOpenItems(page, pageSize int64) (*OpenItemsResponse, error) {
	requestBody := c.getOpenItemsRequestBody(page, pageSize, time.Now().UTC())
	buyerNumber := c.GetBuyerNumber()
	request, err := http.NewRequest(http.MethodPost,
		fmt.Sprintf("https://www.copart.com/data/lots/openItems/%d/0", buyerNumber.Primary), requestBody)
	if err != nil {
		panic(errors.Wrap(err, "failed to craft request"))
	}

	request.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")

	var result OpenItemsResponse
	err = c.doJson(request, &result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform request")
	}

	return &result, result.ToError()
}

func (c *connector) getOpenItemsRequestBody(page, pageSize int64, now time.Time) io.Reader {
	data := url.Values{}
	data.Set("draw", strconv.FormatInt(c.getNextDraw(&c.openItemsDraw), 10))
	data.Set("columns[0][data]", "imgp")
	data.Set("columns[0][name]", "")
	data.Set("columns[0][searchable]", "true")
	data.Set("columns[0][orderable]", "false")
	data.Set("columns[0][search][value]", "")
	data.Set("columns[0][search][regex]", "false")
	data.Set("columns[1][data]", "lotId")
	data.Set("columns[1][name]", "")
	data.Set("columns[1][searchable]", "true")
	data.Set("columns[1][orderable]", "true")
	data.Set("columns[1][search][value]", "")
	data.Set("columns[1][search][regex]", "false")
	data.Set("columns[2][data]", "yearOfVehicle")
	data.Set("columns[2][name]", "")
	data.Set("columns[2][searchable]", "true")
	data.Set("columns[2][orderable]", "true")
	data.Set("columns[2][search][value]", "")
	data.Set("columns[2][search][regex]", "false")
	data.Set("columns[3][data]", "makeOfVehicle")
	data.Set("columns[3][name]", "")
	data.Set("columns[3][searchable]", "true")
	data.Set("columns[3][orderable]", "true")
	data.Set("columns[3][search][value]", "")
	data.Set("columns[3][search][regex]", "false")
	data.Set("columns[4][data]", "modelOfVehicle")
	data.Set("columns[4][name]", "")
	data.Set("columns[4][searchable]", "true")
	data.Set("columns[4][orderable]", "true")
	data.Set("columns[4][search][value]", "")
	data.Set("columns[4][search][regex]", "false")
	data.Set("columns[5][data]", "itemNumber")
	data.Set("columns[5][name]", "")
	data.Set("columns[5][searchable]", "true")
	data.Set("columns[5][orderable]", "true")
	data.Set("columns[5][search][value]", "")
	data.Set("columns[5][search][regex]", "false")
	data.Set("columns[6][data]", "yardName")
	data.Set("columns[6][name]", "")
	data.Set("columns[6][searchable]", "true")
	data.Set("columns[6][orderable]", "true")
	data.Set("columns[6][search][value]", "")
	data.Set("columns[6][search][regex]", "false")
	data.Set("columns[7][data]", "runLights")
	data.Set("columns[7][name]", "")
	data.Set("columns[7][searchable]", "false")
	data.Set("columns[7][orderable]", "false")
	data.Set("columns[7][search][value]", "")
	data.Set("columns[7][search][regex]", "false")
	data.Set("columns[8][data]", "sellerName")
	data.Set("columns[8][name]", "")
	data.Set("columns[8][searchable]", "false")
	data.Set("columns[8][orderable]", "true")
	data.Set("columns[8][search][value]", "")
	data.Set("columns[8][search][regex]", "false")
	data.Set("columns[9][data]", "crGrade")
	data.Set("columns[9][name]", "")
	data.Set("columns[9][searchable]", "false")
	data.Set("columns[9][orderable]", "true")
	data.Set("columns[9][search][value]", "")
	data.Set("columns[9][search][regex]", "false")
	data.Set("columns[10][data]", "auctionDateUtc")
	data.Set("columns[10][name]", "")
	data.Set("columns[10][searchable]", "true")
	data.Set("columns[10][orderable]", "true")
	data.Set("columns[10][search][value]", "")
	data.Set("columns[10][search][regex]", "false")
	data.Set("columns[11][data]", "odometerReading")
	data.Set("columns[11][name]", "")
	data.Set("columns[11][searchable]", "true")
	data.Set("columns[11][orderable]", "true")
	data.Set("columns[11][search][value]", "")
	data.Set("columns[11][search][regex]", "false")
	data.Set("columns[12][data]", "docType")
	data.Set("columns[12][name]", "")
	data.Set("columns[12][searchable]", "true")
	data.Set("columns[12][orderable]", "false")
	data.Set("columns[12][search][value]", "")
	data.Set("columns[12][search][regex]", "false")
	data.Set("columns[13][data]", "damageType")
	data.Set("columns[13][name]", "")
	data.Set("columns[13][searchable]", "true")
	data.Set("columns[13][orderable]", "true")
	data.Set("columns[13][search][value]", "")
	data.Set("columns[13][search][regex]", "false")
	data.Set("columns[14][data]", "actualCostOfVehicle")
	data.Set("columns[14][name]", "")
	data.Set("columns[14][searchable]", "true")
	data.Set("columns[14][orderable]", "true")
	data.Set("columns[14][search][value]", "")
	data.Set("columns[14][search][regex]", "false")
	data.Set("columns[15][data]", "bidderNumber")
	data.Set("columns[15][name]", "")
	data.Set("columns[15][searchable]", "false")
	data.Set("columns[15][orderable]", "true")
	data.Set("columns[15][search][value]", "")
	data.Set("columns[15][search][regex]", "false")
	data.Set("columns[16][data]", "bidStatusStr")
	data.Set("columns[16][name]", "")
	data.Set("columns[16][searchable]", "true")
	data.Set("columns[16][orderable]", "true")
	data.Set("columns[16][search][value]", "")
	data.Set("columns[16][search][regex]", "false")
	data.Set("columns[17][data]", "maxBid")
	data.Set("columns[17][name]", "")
	data.Set("columns[17][searchable]", "true")
	data.Set("columns[17][orderable]", "true")
	data.Set("columns[17][search][value]", "")
	data.Set("columns[17][search][regex]", "false")
	data.Set("columns[18][data]", "")
	data.Set("columns[18][name]", "")
	data.Set("columns[18][searchable]", "true")
	data.Set("columns[18][orderable]", "false")
	data.Set("columns[18][search][value]", "")
	data.Set("columns[18][search][regex]", "false")
	data.Set("order[0][column]", "10")
	data.Set("order[0][dir]", "desc")
	data.Set("start", strconv.FormatInt(page*pageSize, 10))
	data.Set("length", strconv.FormatInt(pageSize, 10))
	data.Set("search[value]", "")
	data.Set("search[regex]", "false")
	data.Set("sortOrder", "desc")
	data.Set("sortByColumnIndex", "10")
	data.Set("sortByColumn", "auctionDateUtc")
	data.Set("filterCriteria", `{"filters":[]}`)
	data.Set("readCriteriaFromSession", "N")
	data.Set("searchText", "")
	return bytes.NewReader([]byte(data.Encode()))
}
