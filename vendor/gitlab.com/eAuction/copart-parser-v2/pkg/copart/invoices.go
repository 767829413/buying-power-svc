package copart

import (
	"bytes"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type InvoicesResponse struct {
	BidStatusMeta
	Data []Invoice `json:"aaData"`
}

type Invoice struct {
	SwiftImageURL     string    `json:"swiftImageUrl"`
	OldImageURL       string    `json:"oldImageUrl"`
	ID                int       `json:"id"`
	SecretID          string    `json:"secretId"`
	SaleDate          string    `json:"saleDate"`
	Description       string    `json:"description"`
	LotNumber         LotNumber `json:"lotNumber"`
	AmountDue         float64   `json:"amountDue"`
	InvoiceAmount     float64   `json:"invoiceAmount"`
	PastDue           bool      `json:"pastDue"`
	BidAmount         float64   `json:"bidAmount"`
	FundingStatus     string    `json:"fundingStatus"`
	ItemNumber        int       `json:"itemNumber"`
	LotYear           string    `json:"lotYear"`
	LotModel          string    `json:"lotModel"`
	LotMake           string    `json:"lotMake"`
	OwnerTitleType    string    `json:"ownerTitleType"`
	TransporterEmail  string    `json:"transporterEmail"`
	Pdf               bool      `json:"pdf"`
	Yn                string    `json:"yn"`
	Yname             string    `json:"yname"`
	Lfd               string    `json:"lfd"`
	Cf                float64   `json:"cf"`
	Sfda              float64   `json:"sfda"`
	Sfsia             float64   `json:"sfsia"`
	Cur               string    `json:"cur"`
	Efp               bool      `json:"efp"`
	Bid               string    `json:"bid"`
	Vin               string    `json:"vin"`
	Iba               bool      `json:"iba"`
	St                string    `json:"st"`
	Lori              string    `json:"lori"`
	Brand             string    `json:"brand"`
	BuyerPickupStatus string    `json:"buyerPickupStatus"`
	Plcflg            string    `json:"plcflg"`
	Ind               Date      `json:"ind"`
	Imgp              string    `json:"imgp,omitempty"`
}

func (c *connector) GetInvoices(page, pageSize int64) (*InvoicesResponse, error) {
	requestBody := c.getInvoicesRequestBody(page, pageSize)
	request, err := http.NewRequest(http.MethodPost, "https://www.copart.com/data/payments/getDueInvoiceList/O/USD", requestBody)
	if err != nil {
		panic(errors.Wrap(err, "failed to craft request"))
	}

	request.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")

	var result InvoicesResponse
	err = c.doJson(request, &result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform request")
	}

	return &result, result.ToError()
}

func (c *connector) DowloadInvoice(invoiceNumber, paymentType,currencyCode,invoiceType, lotNumber,secretID,bidderNumber string) (io.ReadCloser, error) {
	u, err := url.Parse("https://www.copart.com/invoice/pdf")
	if err != nil {
		panic(errors.Wrap(err, "failed to parse url"))
	}

	values := u.Query()
	values.Set("invoiceNumber", invoiceNumber)
	values.Set("paymentType", paymentType)
	values.Set("currencyCode", currencyCode)
	values.Set("invoiceType", invoiceType)
	values.Set("lotNumber", lotNumber)
	values.Set("outputType", "D")
	values.Set("secretId", secretID)
	values.Set("bidderNumber", bidderNumber)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to create request"))
	}

	result, err := c.do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform request")
	}

	return result, nil
}

func (c *connector) getInvoicesRequestBody(page, pageSize int64) io.Reader {
	data := url.Values{}
	data.Set("draw", strconv.FormatInt(c.getNextDraw(&c.invoicesDraw), 10))
	data.Set("columns[0][data]", "id")
	data.Set("columns[0][name]", "")
	data.Set("columns[0][searchable]", "false")
	data.Set("columns[0][orderable]", "false")
	data.Set("columns[0][search][value]", "")
	data.Set("columns[0][search][regex]", "false")
	data.Set("columns[1][data]", "imgp")
	data.Set("columns[1][name]", "imagePath")
	data.Set("columns[1][searchable]", "false")
	data.Set("columns[1][orderable]", "false")
	data.Set("columns[1][search][value]", "")
	data.Set("columns[1][search][regex]", "false")
	data.Set("columns[2][data]", "saleDate")
	data.Set("columns[2][name]", "")
	data.Set("columns[2][searchable]", "true")
	data.Set("columns[2][orderable]", "true")
	data.Set("columns[2][search][value]", "")
	data.Set("columns[2][search][regex]", "false")
	data.Set("columns[3][data]", "bid")
	data.Set("columns[3][name]", "")
	data.Set("columns[3][searchable]", "false")
	data.Set("columns[3][orderable]", "true")
	data.Set("columns[3][search][value]", "")
	data.Set("columns[3][search][regex]", "false")
	data.Set("columns[4][data]", "itemNumber")
	data.Set("columns[4][name]", "")
	data.Set("columns[4][searchable]", "true")
	data.Set("columns[4][orderable]", "true")
	data.Set("columns[4][search][value]", "")
	data.Set("columns[4][search][regex]", "false")
	data.Set("columns[5][data]", "lotNumber")
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
	data.Set("columns[7][data]", "st")
	data.Set("columns[7][name]", "")
	data.Set("columns[7][searchable]", "true")
	data.Set("columns[7][orderable]", "true")
	data.Set("columns[7][search][value]", "")
	data.Set("columns[7][search][regex]", "false")
	data.Set("columns[8][data]", "description")
	data.Set("columns[8][name]", "")
	data.Set("columns[8][searchable]", "true")
	data.Set("columns[8][orderable]", "true")
	data.Set("columns[8][search][value]", "")
	data.Set("columns[8][search][regex]", "false")
	data.Set("columns[9][data]", "vin")
	data.Set("columns[9][name]", "")
	data.Set("columns[9][searchable]", "true")
	data.Set("columns[9][orderable]", "true")
	data.Set("columns[9][search][value]", "")
	data.Set("columns[9][search][regex]", "false")
	data.Set("columns[10][data]", "ownerTitleType")
	data.Set("columns[10][name]", "")
	data.Set("columns[10][searchable]", "true")
	data.Set("columns[10][orderable]", "true")
	data.Set("columns[10][search][value]", "")
	data.Set("columns[10][search][regex]", "false")
	data.Set("columns[11][data]", "bidAmount")
	data.Set("columns[11][name]", "")
	data.Set("columns[11][searchable]", "true")
	data.Set("columns[11][orderable]", "true")
	data.Set("columns[11][search][value]", "")
	data.Set("columns[11][search][regex]", "false")
	data.Set("columns[12][data]", "invoiceAmount")
	data.Set("columns[12][name]", "")
	data.Set("columns[12][searchable]", "true")
	data.Set("columns[12][orderable]", "true")
	data.Set("columns[12][search][value]", "")
	data.Set("columns[12][search][regex]", "false")
	data.Set("columns[13][data]", "amountDue")
	data.Set("columns[13][name]", "amountDue")
	data.Set("columns[13][searchable]", "true")
	data.Set("columns[13][orderable]", "true")
	data.Set("columns[13][search][value]", "")
	data.Set("columns[13][search][regex]", "false")
	data.Set("columns[14][data]", "lfd")
	data.Set("columns[14][name]", "")
	data.Set("columns[14][searchable]", "true")
	data.Set("columns[14][orderable]", "true")
	data.Set("columns[14][search][value]", "")
	data.Set("columns[14][search][regex]", "false")
	data.Set("columns[15][data]", "lly")
	data.Set("columns[15][name]", "transport")
	data.Set("columns[15][searchable]", "false")
	data.Set("columns[15][orderable]", "false")
	data.Set("columns[15][search][value]", "")
	data.Set("columns[15][search][regex]", "false")
	data.Set("columns[16][data]", "buyerPickupStatus")
	data.Set("columns[16][name]", "")
	data.Set("columns[16][searchable]", "true")
	data.Set("columns[16][orderable]", "true")
	data.Set("columns[16][search][value]", "")
	data.Set("columns[16][search][regex]", "false")
	data.Set("columns[17][data]", "")
	data.Set("columns[17][name]", "")
	data.Set("columns[17][searchable]", "true")
	data.Set("columns[17][orderable]", "false")
	data.Set("columns[17][search][value]", "")
	data.Set("columns[17][search][regex]", "false")
	data.Set("start", strconv.FormatInt(page*pageSize, 10))
	data.Set("length", strconv.FormatInt(pageSize, 10))
	data.Set("search[value]", "")
	data.Set("search[regex]", "false")
	data.Set("filterCriteria", `{"filters":[]}`)
	data.Set("readCriteriaFromSession", "N")
	data.Set("searchText", "")
	return bytes.NewReader([]byte(data.Encode()))
}
