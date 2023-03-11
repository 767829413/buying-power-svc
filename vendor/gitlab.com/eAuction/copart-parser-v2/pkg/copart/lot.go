package copart

import (
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/enumer/pkg/normalization"
	"strconv"
	"strings"
	"time"
)

type Lot struct {
	ID                   int64   `csv:"Id"`
	YardNumber           int64   `csv:"Yard number"`
	YardName             string  `csv:"Yard name"`
	SaleDate             string  `csv:"Sale Date M/D/CY"`
	SaleTime             string  `csv:"Sale time (HHMM)"`
	ItemNumber           int64   `csv:"Item#"`
	LotNumber            int64   `csv:"Lot number"`
	HighestBid           string  `csv:"High Bid"`
	TimeZone             string  `csv:"Time Zone"`
	VIN                  string  `csv:"VIN"`
	EstimatedRetailValue float64 `csv:"Est. Retail Value"`
	Country              string  `csv:"Location country"`
	Maker                string  `csv:"Make"`
	ModelGroup           string  `csv:"Model Group"`
	Model                string  `csv:"Model Detail"`
	Year                 uint32  `csv:"Year"`
	DocIssuer            string  `csv:"Sale Title State"`
	DocType              string  `csv:"Sale Title Type"`
	Transmission         string  `csv:"Transmission"`
	BodyStyle            string  `csv:"Body Style"`
	Color                string  `csv:"Color"`
	FuelType             string  `csv:"Fuel Type"`
	Engine               string  `csv:"Engine"`
	MileageType          string  `csv:"Odometer Brand"`
	Mileage              float64 `csv:"Odometer"`
	PrimaryDamage        string  `csv:"Damage Description"`
	SecondaryDamage      string  `csv:"Secondary Damage"`
	State                string  `csv:"Location state"`
	City                 string  `csv:"Location city"`
	DriveType            string  `csv:"Drive"`
	BuyNowPrice          float32 `csv:"Buy-It-Now Price"`
	RetainValue          float32 `csv:"Est. Retail Value"`
	Highlights           string  `csv:"Runs/Drives"`
	Notes                string  `csv:"Special Note"`
	SaleStatus           string  `csv:"Sale Status"`
	Keys                 string  `csv:"Has Keys-Yes or No"`
	ImageLink            string  `csv:"Image URL"`
	Zip                  string  `csv:"Location ZIP"`
	SpecialNote          string  `csv:"Special Note"`
	ImageThumbnail       string  `csv:"Image Thumbnail"`
}

func (d Lot) ParseEngine() (float32, uint32) {
	// engine characteristics is stored as either "1.6L 4" or "4"
	// where 1.6 is volume and 4 is a cylinders num
	str := strings.Split(d.Engine, "L")
	if len(str) != 2 {
		return 0, 0
	}

	volume, err := strconv.ParseFloat(str[0], 32)
	if err != nil {
		return 0.0, 0
	}

	cylinders, err := strconv.ParseUint(str[1], 10, 64)
	if err != nil {
		return float32(volume), 0
	}
	return float32(volume), uint32(cylinders)
}

//GetHighestBid - parses highest bid
func (l Lot) GetHighestBid() int64 {
	result, err := strconv.ParseFloat(l.HighestBid, 64)
	if err != nil {
		return 0
	}

	return int64(result * 100)
}

func (l Lot) GetMilageType() string {
	switch strings.ToUpper(l.MileageType) {
	case "A":
		return "ACTUAL"
	case "E":
		return "EXCEEDS MECHANICAL LIMITS"
	case "X":
		return "EXEMPT"
	case "N":
		return "NOT ACTUAL"
	default:
		return ""
	}
}

//ErrNotScheduled - signals that lot's start time is not available yet
var ErrNotScheduled = errors.New("sale is not schedule yet")

var usTimeZonesOffsets = map[string]string{
	"AST":  "-0400",
	"EST":  "-0500",
	"EDT":  "-0400",
	"CST":  "-0600",
	"CDT":  "-0500",
	"MST":  "-0700",
	"MDT":  "-0600",
	"PST":  "-0800",
	"PDT":  "-0700",
	"AKST": "-0900",
	"AKDT": "-0800",
	"HST":  "-1000",
	"HAST": "-1000",
	"HADT": "-0900",
	"SST":  "-1100",
	"SDT":  "-1000",
	"CHST": "+1000",
	"ADT":  "-0300",
	"ASDT": "-0300",
}

//StartTime - parser sale start time
func (l Lot) StartTime() (time.Time, error) {
	if l.SaleDate == "0" || l.SaleDate == "" {
		return time.Time{}, ErrNotScheduled
	}

	fields := logan.F{
		"sale_date": l.SaleDate,
		"sale_time": l.SaleTime,
		"time_zone": l.TimeZone,
	}

	offset, ok := usTimeZonesOffsets[l.TimeZone]
	if !ok {
		return time.Time{}, errors.From(errors.New("to offset for specified timezone"), fields)
	}

	if len(l.SaleTime) == 3 {
		l.SaleTime = "0" + l.SaleTime
	}

	result, err := time.Parse("20060102 1504 -0700", strings.Join([]string{l.SaleDate, l.SaleTime, offset}, " "))
	if err != nil {
		return time.Time{}, errors.Wrap(err, "failed to parse time", fields)
	}

	return result.UTC(), nil
}

//IsSubLot - returns true if sublot
func (l Lot) IsSubLot() bool {
	return strings.Contains(normalization.Normalize(l.SpecialNote), "sublot")
}
