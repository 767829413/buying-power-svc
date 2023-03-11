package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"gitlab.com/eAuction/events/go/kafka"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"

	"gitlab.com/eAuction/buying-power-svc/internal/data"
)

type lots struct {
	db *pgdb.DB
}

//Lots - returns new instance of lots
func (s *storage) Lots() data.Lots {
	return &lots{db: s.db}
}

//ByID - returns lot by ID, returns nil, nil if one does not exists
func (q *lots) ByID(id string) (*data.Lot, error) {
	return q.byID(id, false)
}

func (q *lots) byID(id string, forUpdate bool) (*data.Lot, error) {
	stmt := squirrel.Select("*").From("lots").Where(squirrel.Eq{"id": id})
	if forUpdate {
		stmt = stmt.Suffix("for update")
	}
	var result data.Lot
	err := q.db.Get(&result, stmt)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to select lots by id")
	}

	return &result, nil

}

func (q *lots) ByIDs(ids []string) (map[string]data.Lot, error) {
	stmt := squirrel.Select("*").From("lots").Where(squirrel.Eq{"id": ids})
	var rawResult []data.Lot
	err := q.db.Select(&rawResult, stmt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select lots by ids")
	}

	result := map[string]data.Lot{}
	for _, l := range rawResult {
		result[l.ID] = l
	}

	return result, nil
}

//ByIDForUpdate - select lot by id for update
func (q *lots) ByIDForUpdate(id string) (*data.Lot, error) {
	return q.byID(id, true)
}

//Exists - returns true if lot exists
func (q *lots) Exists(id string) (bool, error) {
	var result int
	err := q.db.Get(&result, squirrel.Select("1").From("lots").Where(squirrel.Eq{
		"id": id,
	}))
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return false, nil
		}

		return false, errors.Wrap(err, "failed to exec stmt")
	}

	return true, nil
}

//DeleteRedundant - removes lots that are no longer needed for this service
func (q *lots) DeleteRedundant() (bool, error) {
	const batchSize = 1000
	result, err := q.db.ExecWithResult(squirrel.Delete("lots").Where(fmt.Sprintf("id in (select id from lots where id not in "+
		"(select lot_id from deposits where lot_id is not null union select lot_id from participants) and (ends_at + interval '720 hours') < now() limit %d)", batchSize)))
	if err != nil {
		return false, errors.Wrap(err, "failed to exec statement")
	}

	count, err := result.RowsAffected()
	if err != nil {
		return false, errors.Wrap(err, "failed to count rows affected")
	}

	return count >= batchSize, nil
}

// SetHighestBid updates value of "highest_bid" for lot with specified ID.
func (q *lots) SetHighestBid(lotID string, bid int64) error {
	stmt := squirrel.
		Update("lots").
		Set("highest_bid", bid).
		Where("id = ? AND (highest_bid is null OR highest_bid < ?)", lotID, bid)

	_, err := q.db.ExecWithResult(stmt)
	if err != nil {
		return err
	}

	return nil
}

//Store - saves lot into db
func (q *lots) Store(lot data.Lot) error {
	stmt := squirrel.Insert("lots").SetMap(map[string]interface{}{
		"id":          lot.ID,
		"currency":    lot.Currency,
		"ends_at":     lot.EndsAt,
		"state":       lot.State,
		"details":     lot.Details,
		"highest_bid": lot.HighestBid,
		"external_id": lot.ExternalID,
		"lane":        lot.Lane,
		"item_num":    lot.ItemNum,
	}).Suffix(`
		on conflict
		on constraint lots_pkey
		do update
		set (
			currency,
			state,
			ends_at,
			details,
			external_id,
			lane,
			item_num
		) = (
			excluded.currency,
			excluded.state,
			excluded.ends_at,
			excluded.details,
			excluded.external_id,
			excluded.lane,
			excluded.item_num
		)
	`)

	return q.db.Exec(stmt)
}

//SetLiveBiddingDetails - set live bidding details
func (q *lots) SetLiveBiddingDetails(id, saleID, lane string, itemNum uint64) error {
	return q.db.Exec(squirrel.Update("lots").SetMap(map[string]interface{}{
		"lane": sql.NullString{
			String: lane,
			Valid:  lane != "",
		},
		"item_num": sql.NullInt64{
			Int64: int64(itemNum),
			Valid: itemNum != 0,
		},
		"sale_id": sql.NullString{
			String: saleID,
			Valid:  saleID != "",
		},
	}).Where(squirrel.Eq{
		"id": id,
	}))
}

// SetState updates value of "state" column for lot with specified ID.
func (q *lots) SetState(lotID string, state int32) error {
	stmt := squirrel.
		Update("lots").
		Set("state", state).
		Where("id = ?", lotID)

	result, err := q.db.ExecWithResult(stmt)
	if err != nil {
		return err
	}

	return ensureLotUpdated(result)
}

func ensureLotUpdated(result sql.Result) error {
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "failed to check how many rows were affected")
	}
	if rowsAffected == 0 {
		return data.ErrorLotNotFound
	}

	return nil
}

//ByExternalIDPrefix - returns lot by external ID prefix
func (q *lots) ByExternalIDPrefix(id string, isSimulation bool) (*data.Lot, error) {
	var result data.Lot
	now := time.Now().UTC()
	stmt := squirrel.Select("*").From("lots").
		Where(squirrel.Like{"external_id": id + "%"}).
		OrderBy("ends_at desc").
		Limit(1)
	if !isSimulation {
		stmt = stmt.Where("ends_at <= ? AND ends_at > ?", now, now.Add(-3*time.Hour))

	}
	err := q.db.Get(&result, stmt)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to select lots by external id")
	}

	return &result, nil
}

func (q *lots) SelectExternalIDForSale(branchID string, lane string) ([]string, error) {
	// this method is only used in simulation on dev, so even though everything that is written here is complete shit and
	//you should never do it again, this is fine
	stmt := squirrel.Select("sale_id").From("lots").
		Where(squirrel.Like{"sale_id": fmt.Sprintf("GAKEDLRHM7GQPUSQ3WUEDLI4GNZ3W32N5N5V2KXTUQ7PKF3XL5XVQGXQ-%s-%s-%%", branchID, lane)}).
		Where("sale_id is not null AND ends_at > now() - interval '3 hour'").
		OrderBy("ends_at").Limit(1)

	var saleID string
	err := q.db.Get(&saleID, stmt)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to select sale id")
	}

	var result []string
	stmt = squirrel.Select("external_id").From("lots").
		Where(squirrel.Eq{"sale_id": saleID}).
		Where("sale_id is not null").
		Where("state = ?", kafka.Lot_STATE_PENDING).
		OrderBy("item_num")
	err = q.db.Select(&result, stmt)
	return result, errors.Wrap(err, "failed to select external id fro sale")
}

func (q *lots) SelectEndedPending() ([]data.Lot, error) {
	query := squirrel.Select("l.*").From("lots l").
		LeftJoin("participants p on p.lot_id = l.id").
		Where("p.id is null").
		Where("l.state = ?", kafka.Lot_STATE_PENDING).
		Where("l.ends_at < (now() - interval '6 hours')")

	var lots []data.Lot
	err := q.db.Select(&lots, query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select lots")
	}
	return lots, nil
}
