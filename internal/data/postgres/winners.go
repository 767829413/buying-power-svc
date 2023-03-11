package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
)

// winners is a helper struct for managing "winners" table.
type winners struct {
	db *pgdb.DB
}

func (s *storage) Winners() data.Winners {
	return &winners{db: s.db}
}

// Store saves winner into database. If winner with provided lot ID
// already exists, Store updates the existing one.
func (q *winners) Store(winner data.Winner) error {
	stmt := squirrel.Insert("winners").SetMap(map[string]interface{}{
		"lot_id":                        winner.LotID,
		"invoice_id":                    winner.InvoiceID,
		"invoice_created_at":            winner.InvoiceCreatedAt,
		"container_id":                  winner.ContainerID,
		"container_link":                winner.ContainerLink,
		"fee":                           winner.Fee,
		"fee_currency":                  winner.FeeCurrency,
		"transportation_price":          winner.TransportationPrice,
		"transportation_price_currency": winner.TransportationPriceCurrency,
		"state":                         winner.State,
		"version":                       0,
		"published":                     winner.Published,
		"delivered_at":                  winner.DeliveredAt,
		"is_from_event":                 winner.IsFromEvent,
		"lot_sold_published":            winner.LotSoldPublished,
	}).Suffix(`
		on conflict on constraint winners_pkey
		do update
		set (
			lot_id,
			invoice_id,
			invoice_created_at,
			container_id,
			container_link,
			fee,
			fee_currency,
			transportation_price,
			transportation_price_currency,
			state,
			version,
			published,
			delivered_at,
			is_from_event,
			lot_sold_published
		) = (
			excluded.lot_id,
			excluded.invoice_id,
			excluded.invoice_created_at,
			excluded.container_id,
			excluded.container_link,
			excluded.fee,
			excluded.fee_currency,
			excluded.transportation_price,
			excluded.transportation_price_currency,
			excluded.state,
			winners.version + 1,
			excluded.published,
			excluded.delivered_at,
			excluded.is_from_event,
			excluded.lot_sold_published
		)
	`)

	return q.db.Exec(stmt)
}

// GetByLotID - returns winner by lot ID
func (q *winners) GetByLotID(lotID string) (*data.Winner, error) {
	var result data.Winner
	err := q.db.Get(&result, squirrel.Select("*").From("winners").Where(squirrel.Eq{
		"lot_id": lotID,
	}))
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to exec statement")
	}

	return &result, nil
}

// SelectNotPublished returns all winners that is known to be not published.
func (q *winners) SelectNotPublished(limit uint64) ([]data.Winner, error) {
	query := squirrel.Select("*").From("winners").Where("not published and not is_from_event").Limit(limit).OrderBy("lot_id ASC")

	var winners []data.Winner
	err := q.db.Select(&winners, query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute select winners query")
	}
	return winners, nil
}

// MarkAsPublished updates winner by its lot-id and version
// (if consecutive updates will happen it will preserve kafka history correctness).
func (q *winners) MarkAsPublished(lotID string, version int32) error {
	query := squirrel.Update("winners").SetMap(map[string]interface{}{
		"published":          true,
		"lot_sold_published": true,
	}).Where("lot_id = ? and version = ?", lotID, version)

	err := q.db.Exec(query)
	if err != nil {
		return errors.Wrap(err, "failed to mark winner as published")
	}
	return nil
}
