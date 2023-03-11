package data

import (
	"time"
)

// Winners - helper interface manage winners.
type Winners interface {
	Store(Winner) error
	GetByLotID(lotID string) (*Winner, error)
	SelectNotPublished(limit uint64) ([]Winner, error)
	MarkAsPublished(id string, version int32) error
}

// Winner represents a single record from "winners" storage.
type Winner struct {
	LotID                       string     `db:"lot_id"`
	InvoiceID                   *string    `db:"invoice_id"`
	InvoiceCreatedAt            *time.Time `db:"invoice_created_at"`
	ContainerID                 *string    `db:"container_id"`
	ContainerLink               *string    `db:"container_link"`
	Fee                         int64      `db:"fee"`
	FeeCurrency                 string     `db:"fee_currency"`
	TransportationPrice         int64      `db:"transportation_price"`
	TransportationPriceCurrency string     `db:"transportation_price_currency"`
	State                       int16      `db:"state"`
	DeliveredAt                 *time.Time `db:"delivered_at"`
	Published                   bool       `db:"published"`
	Version                     int32      `db:"version"`
	IsFromEvent                 bool       `db:"is_from_event"`
	LotSoldPublished            bool       `db:"lot_sold_published"`
}
