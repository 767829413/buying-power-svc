package postgres

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"time"
)

type sales struct {
	db *pgdb.DB
}

func (s *storage) Sales() data.Sales {
	return &sales{
		db: s.db,
	}
}

//Upsert - stores sale into db
func (q *sales) Upsert(sale data.Sale) error {
	stmt := squirrel.Insert("sales").SetMap(map[string]interface{}{
		"id":          sale.ID,
		"starts_at":   sale.StartsAt,
		"external_id": sale.ExternalID,
		"platform":    sale.Platform,
	}).Suffix(`
		on conflict
		on constraint sales_pkey
		do nothing
	`)

	return q.db.Exec(stmt)
}

//GetSale - returns sale by external id
func (q *sales) GetSale(platform, externalID string, now time.Time) (*data.Sale, error) {
	return q.getSale(squirrel.Select("*").From("sales").Where(
		"platform = ? AND external_id = ? and starts_at < ?", platform, externalID, now).OrderBy("starts_at desc").Limit(1))
}

func (q *sales) GetSaleWithinPeriod(platform string, from, to time.Time) (*data.Sale, error) {
	return q.getSale(squirrel.Select("*").From("sales").
		Where("platform = ? AND starts_at > ? AND starts_at < ?", platform, from, to).
		Limit(1))
}

func (q *sales) getSale(query squirrel.SelectBuilder) (*data.Sale, error) {
	var result data.Sale
	err := q.db.Get(&result, query)
	if err != nil {
		if sql.ErrNoRows == errors.Cause(err) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to select sale")
	}

	return &result, nil
}

func (q *sales) GetSaleByID(id string) (*data.Sale, error) {
	return q.getSale(squirrel.Select("*").From("sales").Where("id = ?", id))
}