package copart

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type LostResponse struct {
	BidStatusMeta
	Data []Lost `json:"aaData"`
}
type Lost struct {
	BidderStatus              string    `json:"bidderStatus"`
	BidderNumber              string    `json:"bidderNumber"`
	LotNumber                 LotNumber `json:"lotNumber"`
	LotYear                   string    `json:"lotYear"`
	LotMake                   string    `json:"lotMake"`
	LotModel                  string    `json:"lotModel"`
	YardNumber                string    `json:"yardNumber"`
	Vin                       string    `json:"vin"`
	YardName                  string    `json:"yardName"`
	Currency                  Currency  `json:"currency"`
	SaleDate                  Date      `json:"saleDate"`
	SaleDateStr               string    `json:"saleDateStr"`
	DocType                   string    `json:"docType"`
	DmvState                  string    `json:"dmvState"`
	SaleTitleType             string    `json:"saleTitleType"`
	DamageType                string    `json:"damageType"`
	ImageURL                  string    `json:"imageUrl"`
	SwiftImageURL             string    `json:"swiftImageUrl"`
	OldImageURL               string    `json:"oldImageUrl"`
	OdometerReading           string    `json:"odometerReading"`
	ActualCostOfVehicle       string    `json:"actualCostOfVehicle"`
	BidStatus                 string    `json:"bidStatus"`
	MyBid                     string    `json:"myBid"`
	ExportActualCostOfVehicle string    `json:"exportActualCostOfVehicle"`
	ExportMyBid               string    `json:"exportMyBid"`
	SellerName                string    `json:"sellerName"`
	Brand                     string    `json:"brand"`
}

func (c *connector) GetLost(page, pageSize int64) (*LostResponse, error) {
	requestBody := c.getLostRequestBody(page, pageSize, time.Now().UTC())
	request, err := http.NewRequest(http.MethodPost, "https://www.copart.com/data/bidStatus/lotsLost", requestBody)
	if err != nil {
		panic(errors.Wrap(err, "failed to craft request"))
	}

	request.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")

	var result LostResponse
	err = c.doJson(request, &result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform request")
	}

	return &result, result.ToError()
}

func (c *connector) getLostRequestBody(page, pageSize int64, now time.Time) io.Reader {
	data := url.Values{}
	data.Set("draw", strconv.FormatInt(c.getNextDraw(&c.lostDraw), 10))
	data.Set("columns[0][data]", "imgp")
	data.Set("columns[0][name]", "")
	data.Set("columns[0][searchable]", "true")
	data.Set("columns[0][orderable]", "false")
	data.Set("columns[0][search][value]", "")
	data.Set("columns[0][search][regex]", "false")
	data.Set("columns[1][data]", "lotNumber")
	data.Set("columns[1][name]", "")
	data.Set("columns[1][searchable]", "true")
	data.Set("columns[1][orderable]", "true")
	data.Set("columns[1][search][value]", "")
	data.Set("columns[1][search][regex]", "false")
	data.Set("columns[2][data]", "lotYear")
	data.Set("columns[2][name]", "")
	data.Set("columns[2][searchable]", "true")
	data.Set("columns[2][orderable]", "true")
	data.Set("columns[2][search][value]", "")
	data.Set("columns[2][search][regex]", "false")
	data.Set("columns[3][data]", "lotMake")
	data.Set("columns[3][name]", "")
	data.Set("columns[3][searchable]", "true")
	data.Set("columns[3][orderable]", "true")
	data.Set("columns[3][search][value]", "")
	data.Set("columns[3][search][regex]", "false")
	data.Set("columns[4][data]", "lotModel")
	data.Set("columns[4][name]", "")
	data.Set("columns[4][searchable]", "true")
	data.Set("columns[4][orderable]", "true")
	data.Set("columns[4][search][value]", "")
	data.Set("columns[4][search][regex]", "false")
	data.Set("columns[5][data]", "yardName")
	data.Set("columns[5][name]", "")
	data.Set("columns[5][searchable]", "true")
	data.Set("columns[5][orderable]", "true")
	data.Set("columns[5][search][value]", "")
	data.Set("columns[5][search][regex]", "false")
	data.Set("columns[6][data]", "runLights")
	data.Set("columns[6][name]", "")
	data.Set("columns[6][searchable]", "false")
	data.Set("columns[6][orderable]", "false")
	data.Set("columns[6][search][value]", "")
	data.Set("columns[6][search][regex]", "false")
	data.Set("columns[7][data]", "sellerName")
	data.Set("columns[7][name]", "")
	data.Set("columns[7][searchable]", "false")
	data.Set("columns[7][orderable]", "true")
	data.Set("columns[7][search][value]", "")
	data.Set("columns[7][search][regex]", "false")
	data.Set("columns[8][data]", "crGrade")
	data.Set("columns[8][name]", "")
	data.Set("columns[8][searchable]", "false")
	data.Set("columns[8][orderable]", "true")
	data.Set("columns[8][search][value]", "")
	data.Set("columns[8][search][regex]", "false")
	data.Set("columns[9][data]", "vin")
	data.Set("columns[9][name]", "")
	data.Set("columns[9][searchable]", "false")
	data.Set("columns[9][orderable]", "true")
	data.Set("columns[9][search][value]", "")
	data.Set("columns[9][search][regex]", "false")
	data.Set("columns[10][data]", "saleDateStr")
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
	data.Set("columns[15][data]", "myBid")
	data.Set("columns[15][name]", "")
	data.Set("columns[15][searchable]", "true")
	data.Set("columns[15][orderable]", "false")
	data.Set("columns[15][search][value]", "")
	data.Set("columns[15][search][regex]", "false")
	data.Set("columns[16][data]", "")
	data.Set("columns[16][name]", "")
	data.Set("columns[16][searchable]", "true")
	data.Set("columns[16][orderable]", "false")
	data.Set("columns[16][search][value]", "")
	data.Set("columns[16][search][regex]", "false")
	data.Set("order[0][column]", "10")
	data.Set("order[0][dir]", "desc")
	data.Set("start", strconv.FormatInt(page*pageSize, 10))
	data.Set("length", strconv.FormatInt(pageSize, 10))
	data.Set("search[value]", "")
	data.Set("search[regex]", "false")
	data.Set("sortOrder", "desc")
	data.Set("sortByColumnIndex", "10")
	data.Set("sortByColumn", "saleDateStr")
	data.Set("filterCriteria", `{"filters":[]}`)
	data.Set("readCriteriaFromSession", "N")
	data.Set("searchText", "")
	const dateLayout = "20060102"
	data.Set("fromDate", now.Add(-30*24*time.Hour).Format(dateLayout))
	data.Set("toDate", now.Format(dateLayout))
	return bytes.NewReader([]byte(data.Encode()))
}
