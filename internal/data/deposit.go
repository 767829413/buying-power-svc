package data

import (
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
)

//Deposit -
type Deposit struct {
	ID        string    `db:"id"`
	State     int32     `db:"state"`
	Amount    int64     `db:"amount"`
	Currency  string    `db:"currency"`
	Depositor string    `db:"depositor"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	LotID     *string   `db:"lot_id"`
	Synced    bool      `db:"synced"`
	// defines that deposits has been created by admin and
	IsArtificial bool `db:"is_artificial"`
}

type DepositsSelector struct {
	States       []int32
	HasLot       *bool
	Lots         []string
	Depositors   []string
	ForUpdate    bool
	MaxCreatedAt *time.Time
	MinCreatedAt *time.Time
	Currency     *string
	Page         *pgdb.OffsetPageParams
	IsSynced     *bool
	IsArtificial *bool
}

//Deposits - defines helper interface to access stored deposits
type Deposits interface {
	Create(amount int64, currency, depositor string) (Deposit, error)
	Get(id string, forUpdate bool) (*Deposit, error)
	GetForUpdate(id string) (*Deposit, error)
	SetState(id string, state int32) error
	SetLot(id string, lotID *string) error
	Select(selector DepositsSelector) ([]Deposit, error)
	MarkSynced(ids []string) error
	CreateLegacy(deposit Deposit) error
	//ManuallyApprove - marks deposit as paid and artificial
	// Note: DO NOT use directly - use deposits.ManuallyApprove
	ManuallyApprove(id string) error
	// ReturnDepositToBalance - returns deposit to the balance
	// NOTE: DO NOT use directly - use deposits.Unlock
	ReturnDepositToBalance(id string) error
	ScheduleReturnToCard(id string) error
}

//DepositsByCreatedAtAsc - alias that allows to sort deposits
type DepositsByCreatedAtAsc []Deposit

// Len is the number of elements in the collection.
func (d DepositsByCreatedAtAsc) Len() int {
	return len(d)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (d DepositsByCreatedAtAsc) Less(i, j int) bool {
	return d[i].CreatedAt.Before(d[j].CreatedAt)
}

// Swap swaps the elements with indexes i and j.
func (d DepositsByCreatedAtAsc) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
