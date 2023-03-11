package data

import "gitlab.com/eAuction/events/go/kafka"

type Identity struct {
	ID        string              `db:"id"`
	Platform  string              `db:"platform"`
	Type      kafka.Identity_Type `db:"type"`
	IsDeleted bool                `db:"is_deleted"`
}

type IdentityWithDeposiAmount struct {
	Identity
	DepositAmount int64 `db:"deposit_amount"`
}

type Identities interface {
	Store(identity Identity) error
	//GetForUpdate - returns identity with lock for update
	GetForUpdate(address string) (*Identity, error)
	MarkDeleted(address string) error
	//Get - returns identity
	Get(address string) (*Identity, error)
	SelectNeedRecharge(types []kafka.Identity_Type, targetBalance int64) ([]IdentityWithDeposiAmount, error)
}
