package data

import "time"

//Sale - defines group of lots sold one by one
type Sale struct {
	ID         string    `db:"id"`
	StartsAt   time.Time `db:"starts_at"`
	ExternalID string    `db:"external_id"`
	Platform   string    `db:"platform"`
}

type Sales interface {
	Upsert(Sale) error
	GetSale(platform, externalID string, now time.Time) (*Sale, error)
	GetSaleWithinPeriod(platform string, from, to time.Time) (*Sale, error)
	GetSaleByID(id string) (*Sale, error)
}
