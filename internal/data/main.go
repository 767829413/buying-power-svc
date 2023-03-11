package data

import "gitlab.com/distributed_lab/kit/pgdb"

//Storage - allows to communicate with storage
type Storage interface {
	Clone() Storage
	Transaction(fn pgdb.TransactionFunc) error
	Deposits() Deposits
	Identities() Identities
	Lots() Lots
	Participants() Participants
	Winners() Winners
	Withdrawals() Withdrawals
	Movements() Movements
	IdentitiesToBrokers() IdentityToBrokers
	Sales() Sales
}

