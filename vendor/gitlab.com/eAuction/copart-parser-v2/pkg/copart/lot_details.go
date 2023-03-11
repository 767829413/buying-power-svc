package copart

import (
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

//LotDetails - defines details of the lot
type LotDetails struct {
	LotNumber uint64 `json:"ln"`

	FinalPrice float32 `json:"awardedHighBid"`
	IsSold     bool    `json:"lotSold"`

	LotName    string `json:"ld"`
	Year       uint   `json:"lcy"`
	Maker      string `json:"mkn"`
	Model      string `json:"lm"`
	BodyStyle  string `json:"bstl"`
	Color      string `json:"clr"`
	DocType    string `json:"td"`
	DocIssuer  string `json:"ts"`
	Highlights string `json:"lcd"`
	SaleStatus string `json:"ess"`

	Currency    string  `json:"cuc"`
	BuyNowPrice float32 `json:"bnp"`
	CurrentBid  float32 `json:"hb"`
	RetainValue float32 `json:"la"`

	Mileage         float32 `json:"orr"`
	MileageType     string  `json:"ord"`
	EngineVolume    string  `json:"egn"` // like "1.6L 4" (for moto volume is not indicated)
	EngineCylinders string  `json:"cy"`
	FuelType        string  `json:"ft"`
	Transmission    string  `json:"tmtp"`
	DriveType       string  `json:"drv"`

	TimeZone   string `json:"tz"`
	StartTime  int64  `json:"ad"`
	YardNumber int64  `json:"ynumb"`
	Lane       string `json:"al"`
	ZIP        string `json:"zip"`
	SublotZIP  string `json:"slpc"`

	PrimaryDamage   string `json:"dd"`
	SecondaryDamage string `json:"sdd,omitempty"`
	ItemNum         int64  `json:"aan"`
	Country         string `json:"locCountry"`
	City            string `json:"locCity"`
	SublotCity      string `json:"slc"`
	State           string `json:"locState"`
	SublotState     string `json:"sls"`

	Keys  string `json:"hk"`
	Notes string `json:"ltnte,omitempty"`

	Seller     string `json:"scn,omitempty"`
	SellerType string `json:"std,omitempty"`
}

type lotDetailsData struct {
	Details *LotDetails `json:"lotDetails,omitempty"`
}

//LotDetailsResponse - defines response on lot details
type lotDetailsResponse struct {
	solrResponse
	Data lotDetailsData `json:"data,omitempty"`
}

//GetLotDetails - returns details for specified lot ID. Returns nil, nil if not found
func (c *connector) GetLotDetails(lotID string) (*LotDetails, error) {
	r, err := http.NewRequest(http.MethodGet, "https://www.copart.com/public/data/lotdetails/solr/"+lotID, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to create request"))
	}

	var response lotDetailsResponse
	err = c.doJson(r, &response)
	if err != nil {
		if errors.Cause(err) == ErrNotFound {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to perform request")
	}

	return response.Data.Details, nil
}
