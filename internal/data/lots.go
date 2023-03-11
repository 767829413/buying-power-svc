package data

import (
	"database/sql"
	"database/sql/driver"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"strings"
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/eAuction/events/go/kafka"
)

type Lots interface {
	ByID(string) (*Lot, error)
	ByIDs(ids []string) (map[string]Lot, error)
	ByIDForUpdate(id string) (*Lot, error)
	Store(Lot) error
	SetState(lotID string, state int32) error
	Exists(id string) (bool, error)
	// SetHighestBid updates value of "highest_bid" for lot with specified ID.
	SetHighestBid(lotID string, bid int64) error
	//DeleteRedundant - removes lots that are no longer needed for this service
	DeleteRedundant() (bool, error)
	ByExternalIDPrefix(id string, isSimulation bool) (*Lot, error)
	SelectExternalIDForSale(branchID string, lane string) ([]string, error)
	SetLiveBiddingDetails(id, saleID, lane string, itemNum uint64) error
	SelectEndedPending() ([]Lot, error)
}

//LotDetails - defines all the data we have for the lot. We need to store it for cases when design changes and we need to return additional data,
// but do not want to do reingest. DO NOT filter by the details!!
type LotDetails kafka.Lot

func (d LotDetails) Value() (driver.Value, error) {
	return pgdb.JSONValue(d)
}

func (d *LotDetails) Scan(raw interface{}) error {
	return pgdb.JSONScan(raw, d)
}

type Lot struct {
	ID         string         `db:"id"`
	Currency   string         `db:"currency"`
	EndsAt     time.Time      `db:"ends_at"`
	State      int32          `json:"state"`
	Details    LotDetails     `db:"details"`
	HighestBid sql.NullInt64  `db:"highest_bid"`
	ExternalID sql.NullString `db:"external_id"`
	Lane       sql.NullString `db:"lane"`
	ItemNum    sql.NullInt64  `db:"item_num"`
	SaleID     sql.NullString `db:"sale_id"`
}

//GetHighestBid - returns highest bid for the lot
func (l Lot) GetHighestBid() int64 {
	if l.Details.AuctionPrices == nil {
		return 0
	}
	highestBid := l.Details.AuctionPrices.Start
	if l.HighestBid.Valid && l.HighestBid.Int64 > highestBid {
		return l.HighestBid.Int64
	}

	return highestBid
}

//ExtractCopartSaleID - extracts copart sale ID from internal sale ID
func ExtractCopartSaleID(raw string) (string, error) {
	subs := strings.Split(raw, "-")
	if len(subs) != 4 {
		return "", errors.From(errors.New("expected sale id to have 4 parts"), logan.F{
			"sale_id": raw,
		})
	}

	branchID := subs[1]
	const expectedLength = 3
	if len(branchID) < expectedLength {
		branchID = strings.Repeat("0", expectedLength-len(branchID)) + branchID
	}

	return "COPART" + branchID + subs[2], nil
}
