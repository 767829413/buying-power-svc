package data

import (
	"time"
)

//Withdrawal -
type Withdrawal struct {
	ID        string    `db:"id"`
	State     int32     `db:"state"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Synced    bool      `db:"synced"`
	DepositID string    `db:"deposit_id"`
}

//WithdrawalsSelector - helper struct to define filters for withdrawal
type WithdrawalsSelector struct {
	States    []int32
	ForUpdate bool
	IsSynced  *bool
	Deposits  []string
}

//Withdrawals - defines helper interface to access stored withdrawals
type Withdrawals interface {
	Create(depositID string) (Withdrawal, error)
	GetForUpdate(id string) (*Withdrawal, error)
	SetState(id string, state int32) error
	Select(selector WithdrawalsSelector) ([]Withdrawal, error)
	MarkSynced(ids []string) error
}
