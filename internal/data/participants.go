package data

import (
	"database/sql"
	"time"

	"gitlab.com/eAuction/events/go/kafka"
)

//ParticipantsSelector - defines filters for participants selects
type ParticipantsSelector struct {
	State         []kafka.Participant_State
	Platforms     []string
	LotID         *string
	DepositStates []int32
	AccountID     *string
	UpdatedAfter  *time.Time
	ForUpdate     bool
	MinLotsEndsAt *time.Time
}

//Participants - helper interface to manage participant
type Participants interface {
	Create(identityID, lotID string) error
	Store(participant Participant) error
	// ByLotIDAccountAddress returns a participant by it's lot ID and account address.
	// Returns nil, nil if entry doesn't exist.
	ByLotIDAccountAddress(lotID, accountAddress string) (*Participant, error)
	GetForUpdate(lotID, accountAddress string) (*Participant, error)
	GetClosestForUpdate(lotID string, bid int64, platforms []string) (*Participant, error)
	Exists(lotID, accountAddress string) (bool, error)
	//Select - returns list of participants for specified selector
	Select(selector ParticipantsSelector) ([]Participant, error)
	SelectNotSynced() ([]Participant, error)
	MarkSynced(p Participant) error
	SetBidLimit(id, bidLimit int64) error
	SetAutoBidLimit(id, bidLimit int64) error
	SetBuyNow(id int64) error
	SetBid(id, bid int64) error
	SetState(id int64, state kafka.Participant_State) error
	SetRequestedBid(id, bid int64) error
	SetOutbidedWith(id, amount int64) error
}

// Participant represents a single record from "participants" storage.
type Participant struct {
	ID                int64         `db:"id"`
	AccountAddress    string        `db:"account_address"`
	LotID             string        `db:"lot_id"`
	State             int32         `db:"state"`
	RequestedBuyNow   bool          `db:"requested_buy_now"`
	CurrentBid        int64         `db:"current_bid"`
	BidLimit          int64         `db:"bid_limit"`
	AutoBidLimit      int64         `db:"auto_bid_limit"`
	CreatedAt         time.Time     `db:"created_at"`
	UpdatedAt         time.Time     `db:"updated_at"`
	IsLocal           bool          `db:"is_local"`
	IsSynced          bool          `db:"is_synced"`
	Version           int64         `db:"version"`
	PublishedCreation sql.NullBool  `db:"published_creation"`
	RequestedBid      int64         `db:"requested_bid"`
	OutbidedWith      sql.NullInt64 `db:"outbided_with"`
}
