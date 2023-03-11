package copart

import (
	"bytes"
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type BidStatusMeta struct {
	ITotalRecords        int             `json:"iTotalRecords"`
	ITotalDisplayRecords int             `json:"iTotalDisplayRecords"`
	SEcho                string          `json:"sEcho"`
	SColumns             string          `json:"sColumns"`
	Error                bool            `json:"error"`
	ErrorMsg             json.RawMessage `json:"errorMsg"`
}

func (b BidStatusMeta) ToError() error {
	if !b.Error {
		return nil
	}

	return errors.From(errors.New("received error from copart"), logan.F{
		"msg": string(b.ErrorMsg),
	})
}

//IsSuccess - returns true if request was successful
func (b BidStatusMeta) IsSuccess() bool {
	return !b.Error
}

type WonResponse struct {
	BidStatusMeta
	Data []Won `json:"aaData"`
}

type Currency struct {
	Code string `json:"code"`
}
type Date struct {
	DateAsInt int   `json:"dateAsInt"`
	UnixDate      int64 `json:"date"`
}

func (d Date) GetTime() time.Time {
	return time.Unix(d.UnixDate/1000, 0)
}

type Won struct {
	ImageURL        string    `json:"imageUrl"`
	LotNumber       LotNumber `json:"lotNumber"`
	LotYear         string    `json:"lotYear"`
	LotMake         string    `json:"lotMake"`
	LotModel        string    `json:"lotModel"`
	YardName        string    `json:"yardName"`
	Vin             string    `json:"vin"`
	SaleDate        Date      `json:"saleDate"`
	LotLeftYardDate Date      `json:"lotLeftYardDate,omitempty"`
	DocType         string    `json:"docType"`
	DamageType      string    `json:"damageType"`
	Currency        Currency  `json:"currency"`
	ExportTotalAmt  string    `json:"exportTotalAmt"`
	PaymentStatus   string    `json:"paymentStatus"`

	BidderStatus        string   `json:"bidderStatus"`
	BidderNumber        string   `json:"bidderNumber"`
	YardNumber          string   `json:"yardNumber"`
	SaleDateStr         string   `json:"saleDateStr"`
	SaleTitleType       string   `json:"saleTitleType"`
	SwiftImageURL       string   `json:"swiftImageUrl"`
	OdometerReading     string   `json:"odometerReading"`
	LotIcons            []string `json:"lotIcons"`
	ActualCostOfVehicle string   `json:"actualCostOfVehicle"`
	TotalAmt            string   `json:"totalAmt"`
	CrGrade             float64  `json:"crGrade"`
	AwardedHighBid      float64  `json:"awardedHighBid"`
	Brand               string   `json:"brand"`
}

func (c *connector) GetWon(page, pageSize int64) (*WonResponse, error) {
	requestBody := c.getWonRequestBody(page, pageSize, time.Now().UTC())
	request, err := http.NewRequest(http.MethodPost, "https://www.copart.com/data/bidStatus/lotsWon/0", requestBody)
	if err != nil {
		panic(errors.Wrap(err, "failed to craft request"))
	}

	request.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")

	var result WonResponse
	err = c.doJson(request, &result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform request")
	}

	return &result, result.ToError()
}

func (c *connector) getWonRequestBody(page, pageSize int64, now time.Time) io.Reader {
	data := url.Values{}
	data.Set("draw", strconv.FormatInt(c.getNextDraw(&c.wonDraw), 10))
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
	data.Set("columns[9][searchable]", "true")
	data.Set("columns[9][orderable]", "true")
	data.Set("columns[9][search][value]", "")
	data.Set("columns[9][search][regex]", "false")
	data.Set("columns[10][data]", "bidderNumber")
	data.Set("columns[10][name]", "")
	data.Set("columns[10][searchable]", "false")
	data.Set("columns[10][orderable]", "true")
	data.Set("columns[10][search][value]", "")
	data.Set("columns[10][search][regex]", "false")
	data.Set("columns[11][data]", "saleDateStr")
	data.Set("columns[11][name]", "")
	data.Set("columns[11][searchable]", "true")
	data.Set("columns[11][orderable]", "true")
	data.Set("columns[11][search][value]", "")
	data.Set("columns[11][search][regex]", "false")
	data.Set("columns[12][data]", "odometerReading")
	data.Set("columns[12][name]", "")
	data.Set("columns[12][searchable]", "false")
	data.Set("columns[12][orderable]", "true")
	data.Set("columns[12][search][value]", "")
	data.Set("columns[12][search][regex]", "false")
	data.Set("columns[13][data]", "lotLeftYardDate")
	data.Set("columns[13][name]", "")
	data.Set("columns[13][searchable]", "true")
	data.Set("columns[13][orderable]", "true")
	data.Set("columns[13][search][value]", "")
	data.Set("columns[13][search][regex]", "false")
	data.Set("columns[14][data]", "docType")
	data.Set("columns[14][name]", "")
	data.Set("columns[14][searchable]", "true")
	data.Set("columns[14][orderable]", "false")
	data.Set("columns[14][search][value]", "")
	data.Set("columns[14][search][regex]", "false")
	data.Set("columns[15][data]", "damageType")
	data.Set("columns[15][name]", "")
	data.Set("columns[15][searchable]", "true")
	data.Set("columns[15][orderable]", "true")
	data.Set("columns[15][search][value]", "")
	data.Set("columns[15][search][regex]", "false")
	data.Set("columns[16][data]", "actualCostOfVehicle")
	data.Set("columns[16][name]", "")
	data.Set("columns[16][searchable]", "false")
	data.Set("columns[16][orderable]", "true")
	data.Set("columns[16][search][value]", "")
	data.Set("columns[16][search][regex]", "false")
	data.Set("columns[17][data]", "totalAmt")
	data.Set("columns[17][name]", "")
	data.Set("columns[17][searchable]", "true")
	data.Set("columns[17][orderable]", "false")
	data.Set("columns[17][search][value]", "")
	data.Set("columns[17][search][regex]", "false")
	data.Set("columns[18][data]", "total")
	data.Set("columns[18][name]", "")
	data.Set("columns[18][searchable]", "false")
	data.Set("columns[18][orderable]", "false")
	data.Set("columns[18][search][value]", "")
	data.Set("columns[18][search][regex]", "false")
	data.Set("columns[19][data]", "totalAmt")
	data.Set("columns[19][name]", "")
	data.Set("columns[19][searchable]", "false")
	data.Set("columns[19][orderable]", "true")
	data.Set("columns[19][search][value]", "")
	data.Set("columns[19][search][regex]", "false")
	data.Set("columns[20][data]", "coeStatus")
	data.Set("columns[20][name]", "")
	data.Set("columns[20][searchable]", "false")
	data.Set("columns[20][orderable]", "false")
	data.Set("columns[20][search][value]", "")
	data.Set("columns[20][search][regex]", "false")
	data.Set("columns[21][data]", "dp")
	data.Set("columns[21][name]", "")
	data.Set("columns[21][searchable]", "true")
	data.Set("columns[21][orderable]", "false")
	data.Set("columns[21][search][value]", "")
	data.Set("columns[21][search][regex]", "false")
	data.Set("columns[22][data]", "")
	data.Set("columns[22][name]", "")
	data.Set("columns[22][searchable]", "true")
	data.Set("columns[22][orderable]", "false")
	data.Set("columns[22][search][value]", "")
	data.Set("columns[22][search][regex]", "false")
	data.Set("order[0][column]", "11")
	data.Set("order[0][dir]", "desc")
	data.Set("start", strconv.FormatInt(page*pageSize, 10))
	data.Set("length", strconv.FormatInt(pageSize, 10))
	data.Set("search[value]", "")
	data.Set("search[regex]", "false")
	data.Set("sortOrder", "desc")
	data.Set("sortByColumnIndex", "11")
	data.Set("sortByColumn", "saleDateStr")
	data.Set("filterCriteria", `{"filters":[]}`)
	data.Set("readCriteriaFromSession", "N")
	data.Set("searchText", "")
	const dateLayout = "20060102"
	data.Set("fromDate", now.Add(-30*24*time.Hour).Format(dateLayout))
	data.Set("toDate", now.Format(dateLayout))
	data.Set("dateRangeStringForSolr", "NOW-1MONTHS TO NOW")
	return bytes.NewReader([]byte(data.Encode()))
}
