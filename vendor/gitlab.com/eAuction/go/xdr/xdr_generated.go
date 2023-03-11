// Package xdr is generated from:
//
//  xdr/raw/Stellar-SCP.x
//  xdr/raw/Stellar-ledger-entries-account.x
//  xdr/raw/Stellar-ledger-entries-asset.x
//  xdr/raw/Stellar-ledger-entries-lot.x
//  xdr/raw/Stellar-ledger-entries-message.x
//  xdr/raw/Stellar-ledger-entries-migration.x
//  xdr/raw/Stellar-ledger-entries-offer.x
//  xdr/raw/Stellar-ledger-entries-participant.x
//  xdr/raw/Stellar-ledger-entries-reference.x
//  xdr/raw/Stellar-ledger-entries-reviewable-request.x
//  xdr/raw/Stellar-ledger-entries-trustline.x
//  xdr/raw/Stellar-ledger-entries.x
//  xdr/raw/Stellar-ledger.x
//  xdr/raw/Stellar-operation-allow-trust.x
//  xdr/raw/Stellar-operation-cancel-request.x
//  xdr/raw/Stellar-operation-change-trust.x
//  xdr/raw/Stellar-operation-close-auction.x
//  xdr/raw/Stellar-operation-create-account.x
//  xdr/raw/Stellar-operation-create-bid.x
//  xdr/raw/Stellar-operation-create-lot.x
//  xdr/raw/Stellar-operation-create-passive-offer.x
//  xdr/raw/Stellar-operation-manage-bid-limit.x
//  xdr/raw/Stellar-operation-manage-migration.x
//  xdr/raw/Stellar-operation-manage-offer.x
//  xdr/raw/Stellar-operation-manage-participant.x
//  xdr/raw/Stellar-operation-manage-permissions.x
//  xdr/raw/Stellar-operation-path-payment.x
//  xdr/raw/Stellar-operation-payment.x
//  xdr/raw/Stellar-operation-recover.x
//  xdr/raw/Stellar-operation-request-buy-now.x
//  xdr/raw/Stellar-operation-request-participation.x
//  xdr/raw/Stellar-operation-review-request.x
//  xdr/raw/Stellar-operation-send-message.x
//  xdr/raw/Stellar-operation-set-max-bid.x
//  xdr/raw/Stellar-operation-set-options.x
//  xdr/raw/Stellar-operation-update-lot.x
//  xdr/raw/Stellar-overlay.x
//  xdr/raw/Stellar-reviewable-request-create-bid.x
//  xdr/raw/Stellar-reviewable-request-lot.x
//  xdr/raw/Stellar-reviewable-request-participation.x
//  xdr/raw/Stellar-transaction.x
//  xdr/raw/Stellar-types.x
//
// DO NOT EDIT or your changes may be overwritten
package xdr

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/nullstyle/go-xdr/xdr3"
)

// Unmarshal reads an xdr element from `r` into `v`.
func Unmarshal(r io.Reader, v interface{}) (int, error) {
	// delegate to xdr package's Unmarshal
	return xdr.Unmarshal(r, v)
}

// Marshal writes an xdr element `v` into `w`.
func Marshal(w io.Writer, v interface{}) (int, error) {
	// delegate to xdr package's Marshal
	return xdr.Marshal(w, v)
}

// Value is an XDR Typedef defines as:
//
//   typedef opaque Value<>;
//
type Value []byte

// ScpBallot is an XDR Struct defines as:
//
//   struct SCPBallot
//    {
//        uint32 counter; // n
//        Value value;    // x
//    };
//
type ScpBallot struct {
	Counter Uint32 `json:"counter,omitempty"`
	Value   Value  `json:"value,omitempty"`
}

// ScpStatementType is an XDR Enum defines as:
//
//   enum SCPStatementType
//    {
//        SCP_ST_PREPARE = 0,
//        SCP_ST_CONFIRM = 1,
//        SCP_ST_EXTERNALIZE = 2,
//        SCP_ST_NOMINATE = 3
//    };
//
type ScpStatementType int32

const (
	ScpStatementTypeScpStPrepare     ScpStatementType = 0
	ScpStatementTypeScpStConfirm     ScpStatementType = 1
	ScpStatementTypeScpStExternalize ScpStatementType = 2
	ScpStatementTypeScpStNominate    ScpStatementType = 3
)

var ScpStatementTypeAll = []ScpStatementType{
	ScpStatementTypeScpStPrepare,
	ScpStatementTypeScpStConfirm,
	ScpStatementTypeScpStExternalize,
	ScpStatementTypeScpStNominate,
}

var scpStatementTypeMap = map[int32]string{
	0: "ScpStatementTypeScpStPrepare",
	1: "ScpStatementTypeScpStConfirm",
	2: "ScpStatementTypeScpStExternalize",
	3: "ScpStatementTypeScpStNominate",
}

var scpStatementTypeShortMap = map[int32]string{
	0: "scp_st_prepare",
	1: "scp_st_confirm",
	2: "scp_st_externalize",
	3: "scp_st_nominate",
}

var scpStatementTypeRevMap = map[string]int32{
	"ScpStatementTypeScpStPrepare":     0,
	"ScpStatementTypeScpStConfirm":     1,
	"ScpStatementTypeScpStExternalize": 2,
	"ScpStatementTypeScpStNominate":    3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ScpStatementType
func (e ScpStatementType) ValidEnum(v int32) bool {
	_, ok := scpStatementTypeMap[v]
	return ok
}
func (e ScpStatementType) isFlag() bool {
	for i := len(ScpStatementTypeAll) - 1; i >= 0; i-- {
		expected := ScpStatementType(2) << uint64(len(ScpStatementTypeAll)-1) >> uint64(len(ScpStatementTypeAll)-i)
		if expected != ScpStatementTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ScpStatementType) String() string {
	name, _ := scpStatementTypeMap[int32(e)]
	return name
}

func (e ScpStatementType) ShortString() string {
	name, _ := scpStatementTypeShortMap[int32(e)]
	return name
}

func (e ScpStatementType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ScpStatementTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ScpStatementType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ScpStatementType(t.Value)
	return nil
}

// ScpNomination is an XDR Struct defines as:
//
//   struct SCPNomination
//    {
//        Hash quorumSetHash; // D
//        Value votes<>;      // X
//        Value accepted<>;   // Y
//    };
//
type ScpNomination struct {
	QuorumSetHash Hash    `json:"quorumSetHash,omitempty"`
	Votes         []Value `json:"votes,omitempty"`
	Accepted      []Value `json:"accepted,omitempty"`
}

// ScpStatementPrepare is an XDR NestedStruct defines as:
//
//   struct
//            {
//                Hash quorumSetHash;       // D
//                SCPBallot ballot;         // b
//                SCPBallot* prepared;      // p
//                SCPBallot* preparedPrime; // p'
//                uint32 nC;                // c.n
//                uint32 nH;                // h.n
//            }
//
type ScpStatementPrepare struct {
	QuorumSetHash Hash       `json:"quorumSetHash,omitempty"`
	Ballot        ScpBallot  `json:"ballot,omitempty"`
	Prepared      *ScpBallot `json:"prepared,omitempty"`
	PreparedPrime *ScpBallot `json:"preparedPrime,omitempty"`
	NC            Uint32     `json:"nC,omitempty"`
	NH            Uint32     `json:"nH,omitempty"`
}

// ScpStatementConfirm is an XDR NestedStruct defines as:
//
//   struct
//            {
//                SCPBallot ballot;   // b
//                uint32 nPrepared;   // p.n
//                uint32 nCommit;     // c.n
//                uint32 nH;          // h.n
//                Hash quorumSetHash; // D
//            }
//
type ScpStatementConfirm struct {
	Ballot        ScpBallot `json:"ballot,omitempty"`
	NPrepared     Uint32    `json:"nPrepared,omitempty"`
	NCommit       Uint32    `json:"nCommit,omitempty"`
	NH            Uint32    `json:"nH,omitempty"`
	QuorumSetHash Hash      `json:"quorumSetHash,omitempty"`
}

// ScpStatementExternalize is an XDR NestedStruct defines as:
//
//   struct
//            {
//                SCPBallot commit;         // c
//                uint32 nH;                // h.n
//                Hash commitQuorumSetHash; // D used before EXTERNALIZE
//            }
//
type ScpStatementExternalize struct {
	Commit              ScpBallot `json:"commit,omitempty"`
	NH                  Uint32    `json:"nH,omitempty"`
	CommitQuorumSetHash Hash      `json:"commitQuorumSetHash,omitempty"`
}

// ScpStatementPledges is an XDR NestedUnion defines as:
//
//   union switch (SCPStatementType type)
//        {
//        case SCP_ST_PREPARE:
//            struct
//            {
//                Hash quorumSetHash;       // D
//                SCPBallot ballot;         // b
//                SCPBallot* prepared;      // p
//                SCPBallot* preparedPrime; // p'
//                uint32 nC;                // c.n
//                uint32 nH;                // h.n
//            } prepare;
//        case SCP_ST_CONFIRM:
//            struct
//            {
//                SCPBallot ballot;   // b
//                uint32 nPrepared;   // p.n
//                uint32 nCommit;     // c.n
//                uint32 nH;          // h.n
//                Hash quorumSetHash; // D
//            } confirm;
//        case SCP_ST_EXTERNALIZE:
//            struct
//            {
//                SCPBallot commit;         // c
//                uint32 nH;                // h.n
//                Hash commitQuorumSetHash; // D used before EXTERNALIZE
//            } externalize;
//        case SCP_ST_NOMINATE:
//            SCPNomination nominate;
//        }
//
type ScpStatementPledges struct {
	Type        ScpStatementType         `json:"type,omitempty"`
	Prepare     *ScpStatementPrepare     `json:"prepare,omitempty"`
	Confirm     *ScpStatementConfirm     `json:"confirm,omitempty"`
	Externalize *ScpStatementExternalize `json:"externalize,omitempty"`
	Nominate    *ScpNomination           `json:"nominate,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ScpStatementPledges) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ScpStatementPledges
func (u ScpStatementPledges) ArmForSwitch(sw int32) (string, bool) {
	switch ScpStatementType(sw) {
	case ScpStatementTypeScpStPrepare:
		return "Prepare", true
	case ScpStatementTypeScpStConfirm:
		return "Confirm", true
	case ScpStatementTypeScpStExternalize:
		return "Externalize", true
	case ScpStatementTypeScpStNominate:
		return "Nominate", true
	}
	return "-", false
}

// NewScpStatementPledges creates a new  ScpStatementPledges.
func NewScpStatementPledges(aType ScpStatementType, value interface{}) (result ScpStatementPledges, err error) {
	result.Type = aType
	switch ScpStatementType(aType) {
	case ScpStatementTypeScpStPrepare:
		tv, ok := value.(ScpStatementPrepare)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpStatementPrepare")
			return
		}
		result.Prepare = &tv
	case ScpStatementTypeScpStConfirm:
		tv, ok := value.(ScpStatementConfirm)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpStatementConfirm")
			return
		}
		result.Confirm = &tv
	case ScpStatementTypeScpStExternalize:
		tv, ok := value.(ScpStatementExternalize)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpStatementExternalize")
			return
		}
		result.Externalize = &tv
	case ScpStatementTypeScpStNominate:
		tv, ok := value.(ScpNomination)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpNomination")
			return
		}
		result.Nominate = &tv
	}
	return
}

// MustPrepare retrieves the Prepare value from the union,
// panicing if the value is not set.
func (u ScpStatementPledges) MustPrepare() ScpStatementPrepare {
	val, ok := u.GetPrepare()

	if !ok {
		panic("arm Prepare is not set")
	}

	return val
}

// GetPrepare retrieves the Prepare value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpStatementPledges) GetPrepare() (result ScpStatementPrepare, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Prepare" {
		result = *u.Prepare
		ok = true
	}

	return
}

// MustConfirm retrieves the Confirm value from the union,
// panicing if the value is not set.
func (u ScpStatementPledges) MustConfirm() ScpStatementConfirm {
	val, ok := u.GetConfirm()

	if !ok {
		panic("arm Confirm is not set")
	}

	return val
}

// GetConfirm retrieves the Confirm value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpStatementPledges) GetConfirm() (result ScpStatementConfirm, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Confirm" {
		result = *u.Confirm
		ok = true
	}

	return
}

// MustExternalize retrieves the Externalize value from the union,
// panicing if the value is not set.
func (u ScpStatementPledges) MustExternalize() ScpStatementExternalize {
	val, ok := u.GetExternalize()

	if !ok {
		panic("arm Externalize is not set")
	}

	return val
}

// GetExternalize retrieves the Externalize value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpStatementPledges) GetExternalize() (result ScpStatementExternalize, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Externalize" {
		result = *u.Externalize
		ok = true
	}

	return
}

// MustNominate retrieves the Nominate value from the union,
// panicing if the value is not set.
func (u ScpStatementPledges) MustNominate() ScpNomination {
	val, ok := u.GetNominate()

	if !ok {
		panic("arm Nominate is not set")
	}

	return val
}

// GetNominate retrieves the Nominate value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpStatementPledges) GetNominate() (result ScpNomination, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Nominate" {
		result = *u.Nominate
		ok = true
	}

	return
}

// ScpStatement is an XDR Struct defines as:
//
//   struct SCPStatement
//    {
//        NodeID nodeID;    // v
//        uint64 slotIndex; // i
//
//        union switch (SCPStatementType type)
//        {
//        case SCP_ST_PREPARE:
//            struct
//            {
//                Hash quorumSetHash;       // D
//                SCPBallot ballot;         // b
//                SCPBallot* prepared;      // p
//                SCPBallot* preparedPrime; // p'
//                uint32 nC;                // c.n
//                uint32 nH;                // h.n
//            } prepare;
//        case SCP_ST_CONFIRM:
//            struct
//            {
//                SCPBallot ballot;   // b
//                uint32 nPrepared;   // p.n
//                uint32 nCommit;     // c.n
//                uint32 nH;          // h.n
//                Hash quorumSetHash; // D
//            } confirm;
//        case SCP_ST_EXTERNALIZE:
//            struct
//            {
//                SCPBallot commit;         // c
//                uint32 nH;                // h.n
//                Hash commitQuorumSetHash; // D used before EXTERNALIZE
//            } externalize;
//        case SCP_ST_NOMINATE:
//            SCPNomination nominate;
//        }
//        pledges;
//    };
//
type ScpStatement struct {
	NodeId    NodeId              `json:"nodeID,omitempty"`
	SlotIndex Uint64              `json:"slotIndex,omitempty"`
	Pledges   ScpStatementPledges `json:"pledges,omitempty"`
}

// ScpEnvelope is an XDR Struct defines as:
//
//   struct SCPEnvelope
//    {
//        SCPStatement statement;
//        Signature signature;
//    };
//
type ScpEnvelope struct {
	Statement ScpStatement `json:"statement,omitempty"`
	Signature Signature    `json:"signature,omitempty"`
}

// ScpQuorumSet is an XDR Struct defines as:
//
//   struct SCPQuorumSet
//    {
//        uint32 threshold;
//    	uint32 thresholdPercent;
//    	bool isUpdated;
//        PublicKey trustValidators<>;
//        PublicKey validators<>;
//        SCPQuorumSet innerSets<>;
//    };
//
type ScpQuorumSet struct {
	Threshold        Uint32         `json:"threshold,omitempty"`
	ThresholdPercent Uint32         `json:"thresholdPercent,omitempty"`
	IsUpdated        bool           `json:"isUpdated,omitempty"`
	TrustValidators  []PublicKey    `json:"trustValidators,omitempty"`
	Validators       []PublicKey    `json:"validators,omitempty"`
	InnerSets        []ScpQuorumSet `json:"innerSets,omitempty"`
}

// SignerExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type SignerExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SignerExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SignerExt
func (u SignerExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSignerExt creates a new  SignerExt.
func NewSignerExt(v LedgerVersion, value interface{}) (result SignerExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// Signer is an XDR Struct defines as:
//
//   struct Signer
//        {
//            AccountID pubKey;
//            uint32 weight; // really only need 1byte
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        };
//
type Signer struct {
	PubKey AccountId `json:"pubKey,omitempty"`
	Weight Uint32    `json:"weight,omitempty"`
	Ext    SignerExt `json:"ext,omitempty"`
}

// AccountFlags is an XDR Enum defines as:
//
//   enum AccountFlags
//        {   // masks for each flag
//
//            // Flags set on issuer accounts
//            // TrustLines are created with authorized set to "false" requiring
//            // the issuer to set it for each TrustLine
//            AUTH_REQUIRED_FLAG = 0x1,
//            // If set, the authorized flag in TrustLines can be cleared
//            // otherwise, authorization cannot be revoked
//            AUTH_REVOCABLE_FLAG = 0x2,
//            // Once set, causes all AUTH_* flags to be read-only
//            AUTH_IMMUTABLE_FLAG = 0x4
//        };
//
type AccountFlags int32

const (
	AccountFlagsAuthRequiredFlag  AccountFlags = 1
	AccountFlagsAuthRevocableFlag AccountFlags = 2
	AccountFlagsAuthImmutableFlag AccountFlags = 4
)

var AccountFlagsAll = []AccountFlags{
	AccountFlagsAuthRequiredFlag,
	AccountFlagsAuthRevocableFlag,
	AccountFlagsAuthImmutableFlag,
}

var accountFlagsMap = map[int32]string{
	1: "AccountFlagsAuthRequiredFlag",
	2: "AccountFlagsAuthRevocableFlag",
	4: "AccountFlagsAuthImmutableFlag",
}

var accountFlagsShortMap = map[int32]string{
	1: "auth_required_flag",
	2: "auth_revocable_flag",
	4: "auth_immutable_flag",
}

var accountFlagsRevMap = map[string]int32{
	"AccountFlagsAuthRequiredFlag":  1,
	"AccountFlagsAuthRevocableFlag": 2,
	"AccountFlagsAuthImmutableFlag": 4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for AccountFlags
func (e AccountFlags) ValidEnum(v int32) bool {
	_, ok := accountFlagsMap[v]
	return ok
}
func (e AccountFlags) isFlag() bool {
	for i := len(AccountFlagsAll) - 1; i >= 0; i-- {
		expected := AccountFlags(2) << uint64(len(AccountFlagsAll)-1) >> uint64(len(AccountFlagsAll)-i)
		if expected != AccountFlagsAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e AccountFlags) String() string {
	name, _ := accountFlagsMap[int32(e)]
	return name
}

func (e AccountFlags) ShortString() string {
	name, _ := accountFlagsShortMap[int32(e)]
	return name
}

func (e AccountFlags) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range AccountFlagsAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *AccountFlags) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = AccountFlags(t.Value)
	return nil
}

// AccountType is an XDR Enum defines as:
//
//   enum AccountType
//        {
//            ACCOUNT_TYPE_MASTER = 1,
//            ACCOUNT_TYPE_PLATFORM = 2,
//            ACCOUNT_TYPE_DEALER = 3,
//            ACCOUNT_TYPE_USER = 4,
//            ACCOUNT_TYPE_ADMIN = 5,
//            ACCOUNT_TYPE_BROKER = 6,
//            ACCOUNT_TYPE_BUSINESS = 7
//        };
//
type AccountType int32

const (
	AccountTypeAccountTypeMaster   AccountType = 1
	AccountTypeAccountTypePlatform AccountType = 2
	AccountTypeAccountTypeDealer   AccountType = 3
	AccountTypeAccountTypeUser     AccountType = 4
	AccountTypeAccountTypeAdmin    AccountType = 5
	AccountTypeAccountTypeBroker   AccountType = 6
	AccountTypeAccountTypeBusiness AccountType = 7
)

var AccountTypeAll = []AccountType{
	AccountTypeAccountTypeMaster,
	AccountTypeAccountTypePlatform,
	AccountTypeAccountTypeDealer,
	AccountTypeAccountTypeUser,
	AccountTypeAccountTypeAdmin,
	AccountTypeAccountTypeBroker,
	AccountTypeAccountTypeBusiness,
}

var accountTypeMap = map[int32]string{
	1: "AccountTypeAccountTypeMaster",
	2: "AccountTypeAccountTypePlatform",
	3: "AccountTypeAccountTypeDealer",
	4: "AccountTypeAccountTypeUser",
	5: "AccountTypeAccountTypeAdmin",
	6: "AccountTypeAccountTypeBroker",
	7: "AccountTypeAccountTypeBusiness",
}

var accountTypeShortMap = map[int32]string{
	1: "account_type_master",
	2: "account_type_platform",
	3: "account_type_dealer",
	4: "account_type_user",
	5: "account_type_admin",
	6: "account_type_broker",
	7: "account_type_business",
}

var accountTypeRevMap = map[string]int32{
	"AccountTypeAccountTypeMaster":   1,
	"AccountTypeAccountTypePlatform": 2,
	"AccountTypeAccountTypeDealer":   3,
	"AccountTypeAccountTypeUser":     4,
	"AccountTypeAccountTypeAdmin":    5,
	"AccountTypeAccountTypeBroker":   6,
	"AccountTypeAccountTypeBusiness": 7,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for AccountType
func (e AccountType) ValidEnum(v int32) bool {
	_, ok := accountTypeMap[v]
	return ok
}
func (e AccountType) isFlag() bool {
	for i := len(AccountTypeAll) - 1; i >= 0; i-- {
		expected := AccountType(2) << uint64(len(AccountTypeAll)-1) >> uint64(len(AccountTypeAll)-i)
		if expected != AccountTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e AccountType) String() string {
	name, _ := accountTypeMap[int32(e)]
	return name
}

func (e AccountType) ShortString() string {
	name, _ := accountTypeShortMap[int32(e)]
	return name
}

func (e AccountType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range AccountTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *AccountType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = AccountType(t.Value)
	return nil
}

// AccountEntryTypeSpecificData is an XDR NestedUnion defines as:
//
//   union switch (AccountType t)
//                {
//                case ACCOUNT_TYPE_ADMIN:
//                    AccountID trustingPlatforms<>;
//                }
//
type AccountEntryTypeSpecificData struct {
	T                 AccountType  `json:"t,omitempty"`
	TrustingPlatforms *[]AccountId `json:"trustingPlatforms,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountEntryTypeSpecificData) SwitchFieldName() string {
	return "T"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AccountEntryTypeSpecificData
func (u AccountEntryTypeSpecificData) ArmForSwitch(sw int32) (string, bool) {
	switch AccountType(sw) {
	case AccountTypeAccountTypeAdmin:
		return "TrustingPlatforms", true
	}
	return "-", false
}

// NewAccountEntryTypeSpecificData creates a new  AccountEntryTypeSpecificData.
func NewAccountEntryTypeSpecificData(t AccountType, value interface{}) (result AccountEntryTypeSpecificData, err error) {
	result.T = t
	switch AccountType(t) {
	case AccountTypeAccountTypeAdmin:
		tv, ok := value.([]AccountId)
		if !ok {
			err = fmt.Errorf("invalid value, must be []AccountId")
			return
		}
		result.TrustingPlatforms = &tv
	}
	return
}

// MustTrustingPlatforms retrieves the TrustingPlatforms value from the union,
// panicing if the value is not set.
func (u AccountEntryTypeSpecificData) MustTrustingPlatforms() []AccountId {
	val, ok := u.GetTrustingPlatforms()

	if !ok {
		panic("arm TrustingPlatforms is not set")
	}

	return val
}

// GetTrustingPlatforms retrieves the TrustingPlatforms value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AccountEntryTypeSpecificData) GetTrustingPlatforms() (result []AccountId, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.T))

	if armName == "TrustingPlatforms" {
		result = *u.TrustingPlatforms
		ok = true
	}

	return
}

// AccountEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            case LEDGER_VERSION_ADMIN_PERMISSIONS:
//                union switch (AccountType t)
//                {
//                case ACCOUNT_TYPE_ADMIN:
//                    AccountID trustingPlatforms<>;
//                }
//                typeSpecificData;
//            }
//
type AccountEntryExt struct {
	V                LedgerVersion                 `json:"v,omitempty"`
	TypeSpecificData *AccountEntryTypeSpecificData `json:"typeSpecificData,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AccountEntryExt
func (u AccountEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionLedgerVersionAdminPermissions:
		return "TypeSpecificData", true
	}
	return "-", false
}

// NewAccountEntryExt creates a new  AccountEntryExt.
func NewAccountEntryExt(v LedgerVersion, value interface{}) (result AccountEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionLedgerVersionAdminPermissions:
		tv, ok := value.(AccountEntryTypeSpecificData)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountEntryTypeSpecificData")
			return
		}
		result.TypeSpecificData = &tv
	}
	return
}

// MustTypeSpecificData retrieves the TypeSpecificData value from the union,
// panicing if the value is not set.
func (u AccountEntryExt) MustTypeSpecificData() AccountEntryTypeSpecificData {
	val, ok := u.GetTypeSpecificData()

	if !ok {
		panic("arm TypeSpecificData is not set")
	}

	return val
}

// GetTypeSpecificData retrieves the TypeSpecificData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AccountEntryExt) GetTypeSpecificData() (result AccountEntryTypeSpecificData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "TypeSpecificData" {
		result = *u.TypeSpecificData
		ok = true
	}

	return
}

// AccountEntry is an XDR Struct defines as:
//
//   struct AccountEntry
//        {
//            AccountID accountID;
//            AccountType accountType;
//
//            AccountID platformID; // accountId of the platform this account belongs to
//
//            // fields used for signatures
//            // thresholds stores unsigned bytes: [weight of master|low|medium|high]
//            Thresholds thresholds;
//
//            Signer signers<20>; // possible signers for this account
//
//            // deprecated:
//            int64 balance;
//            uint32 flags;
//            uint32 numSubEntries;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            case LEDGER_VERSION_ADMIN_PERMISSIONS:
//                union switch (AccountType t)
//                {
//                case ACCOUNT_TYPE_ADMIN:
//                    AccountID trustingPlatforms<>;
//                }
//                typeSpecificData;
//            }
//            ext;
//        };
//
type AccountEntry struct {
	AccountId     AccountId       `json:"accountID,omitempty"`
	AccountType   AccountType     `json:"accountType,omitempty"`
	PlatformId    AccountId       `json:"platformID,omitempty"`
	Thresholds    Thresholds      `json:"thresholds,omitempty"`
	Signers       []Signer        `json:"signers,omitempty" xdrmaxsize:"20"`
	Balance       Int64           `json:"balance,omitempty"`
	Flags         Uint32          `json:"flags,omitempty"`
	NumSubEntries Uint32          `json:"numSubEntries,omitempty"`
	Ext           AccountEntryExt `json:"ext,omitempty"`
}

// AssetType is an XDR Enum defines as:
//
//   enum AssetType
//        {
//            ASSET_TYPE_NATIVE = 0,
//            ASSET_TYPE_CREDIT_ALPHANUM4 = 1,
//            ASSET_TYPE_CREDIT_ALPHANUM12 = 2
//        };
//
type AssetType int32

const (
	AssetTypeAssetTypeNative           AssetType = 0
	AssetTypeAssetTypeCreditAlphanum4  AssetType = 1
	AssetTypeAssetTypeCreditAlphanum12 AssetType = 2
)

var AssetTypeAll = []AssetType{
	AssetTypeAssetTypeNative,
	AssetTypeAssetTypeCreditAlphanum4,
	AssetTypeAssetTypeCreditAlphanum12,
}

var assetTypeMap = map[int32]string{
	0: "AssetTypeAssetTypeNative",
	1: "AssetTypeAssetTypeCreditAlphanum4",
	2: "AssetTypeAssetTypeCreditAlphanum12",
}

var assetTypeShortMap = map[int32]string{
	0: "asset_type_native",
	1: "asset_type_credit_alphanum4",
	2: "asset_type_credit_alphanum12",
}

var assetTypeRevMap = map[string]int32{
	"AssetTypeAssetTypeNative":           0,
	"AssetTypeAssetTypeCreditAlphanum4":  1,
	"AssetTypeAssetTypeCreditAlphanum12": 2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for AssetType
func (e AssetType) ValidEnum(v int32) bool {
	_, ok := assetTypeMap[v]
	return ok
}
func (e AssetType) isFlag() bool {
	for i := len(AssetTypeAll) - 1; i >= 0; i-- {
		expected := AssetType(2) << uint64(len(AssetTypeAll)-1) >> uint64(len(AssetTypeAll)-i)
		if expected != AssetTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e AssetType) String() string {
	name, _ := assetTypeMap[int32(e)]
	return name
}

func (e AssetType) ShortString() string {
	name, _ := assetTypeShortMap[int32(e)]
	return name
}

func (e AssetType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range AssetTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *AssetType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = AssetType(t.Value)
	return nil
}

// AssetAlphaNum4Ext is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                            {
//                            case EMPTY_VERSION:
//                                void;
//                            }
//
type AssetAlphaNum4Ext struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AssetAlphaNum4Ext) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AssetAlphaNum4Ext
func (u AssetAlphaNum4Ext) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAssetAlphaNum4Ext creates a new  AssetAlphaNum4Ext.
func NewAssetAlphaNum4Ext(v LedgerVersion, value interface{}) (result AssetAlphaNum4Ext, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AssetAlphaNum4 is an XDR NestedStruct defines as:
//
//   struct
//                {
//                    opaque assetCode[4]; // 1 to 4 characters
//                    AccountID issuer;
//                    union switch (LedgerVersion v)
//                            {
//                            case EMPTY_VERSION:
//                                void;
//                            }
//                            ext;
//                }
//
type AssetAlphaNum4 struct {
	AssetCode [4]byte           `json:"assetCode,omitempty"`
	Issuer    AccountId         `json:"issuer,omitempty"`
	Ext       AssetAlphaNum4Ext `json:"ext,omitempty"`
}

// AssetAlphaNum12Ext is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                            {
//                            case EMPTY_VERSION:
//                                void;
//                            }
//
type AssetAlphaNum12Ext struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AssetAlphaNum12Ext) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AssetAlphaNum12Ext
func (u AssetAlphaNum12Ext) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAssetAlphaNum12Ext creates a new  AssetAlphaNum12Ext.
func NewAssetAlphaNum12Ext(v LedgerVersion, value interface{}) (result AssetAlphaNum12Ext, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AssetAlphaNum12 is an XDR NestedStruct defines as:
//
//   struct
//                {
//                    opaque assetCode[12]; // 5 to 12 characters
//                    AccountID issuer;
//                    union switch (LedgerVersion v)
//                            {
//                            case EMPTY_VERSION:
//                                void;
//                            }
//                            ext;
//                }
//
type AssetAlphaNum12 struct {
	AssetCode [12]byte           `json:"assetCode,omitempty"`
	Issuer    AccountId          `json:"issuer,omitempty"`
	Ext       AssetAlphaNum12Ext `json:"ext,omitempty"`
}

// Asset is an XDR Union defines as:
//
//   union Asset switch (AssetType type)
//        {
//            case ASSET_TYPE_NATIVE: // Not credit
//                void;
//            case ASSET_TYPE_CREDIT_ALPHANUM4:
//                struct
//                {
//                    opaque assetCode[4]; // 1 to 4 characters
//                    AccountID issuer;
//                    union switch (LedgerVersion v)
//                            {
//                            case EMPTY_VERSION:
//                                void;
//                            }
//                            ext;
//                } alphaNum4;
//            case ASSET_TYPE_CREDIT_ALPHANUM12:
//                struct
//                {
//                    opaque assetCode[12]; // 5 to 12 characters
//                    AccountID issuer;
//                    union switch (LedgerVersion v)
//                            {
//                            case EMPTY_VERSION:
//                                void;
//                            }
//                            ext;
//                } alphaNum12;
//                // add other asset types here in the future
//        };
//
type Asset struct {
	Type       AssetType        `json:"type,omitempty"`
	AlphaNum4  *AssetAlphaNum4  `json:"alphaNum4,omitempty"`
	AlphaNum12 *AssetAlphaNum12 `json:"alphaNum12,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u Asset) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of Asset
func (u Asset) ArmForSwitch(sw int32) (string, bool) {
	switch AssetType(sw) {
	case AssetTypeAssetTypeNative:
		return "", true
	case AssetTypeAssetTypeCreditAlphanum4:
		return "AlphaNum4", true
	case AssetTypeAssetTypeCreditAlphanum12:
		return "AlphaNum12", true
	}
	return "-", false
}

// NewAsset creates a new  Asset.
func NewAsset(aType AssetType, value interface{}) (result Asset, err error) {
	result.Type = aType
	switch AssetType(aType) {
	case AssetTypeAssetTypeNative:
		// void
	case AssetTypeAssetTypeCreditAlphanum4:
		tv, ok := value.(AssetAlphaNum4)
		if !ok {
			err = fmt.Errorf("invalid value, must be AssetAlphaNum4")
			return
		}
		result.AlphaNum4 = &tv
	case AssetTypeAssetTypeCreditAlphanum12:
		tv, ok := value.(AssetAlphaNum12)
		if !ok {
			err = fmt.Errorf("invalid value, must be AssetAlphaNum12")
			return
		}
		result.AlphaNum12 = &tv
	}
	return
}

// MustAlphaNum4 retrieves the AlphaNum4 value from the union,
// panicing if the value is not set.
func (u Asset) MustAlphaNum4() AssetAlphaNum4 {
	val, ok := u.GetAlphaNum4()

	if !ok {
		panic("arm AlphaNum4 is not set")
	}

	return val
}

// GetAlphaNum4 retrieves the AlphaNum4 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Asset) GetAlphaNum4() (result AssetAlphaNum4, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AlphaNum4" {
		result = *u.AlphaNum4
		ok = true
	}

	return
}

// MustAlphaNum12 retrieves the AlphaNum12 value from the union,
// panicing if the value is not set.
func (u Asset) MustAlphaNum12() AssetAlphaNum12 {
	val, ok := u.GetAlphaNum12()

	if !ok {
		panic("arm AlphaNum12 is not set")
	}

	return val
}

// GetAlphaNum12 retrieves the AlphaNum12 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Asset) GetAlphaNum12() (result AssetAlphaNum12, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AlphaNum12" {
		result = *u.AlphaNum12
		ok = true
	}

	return
}

// PriceExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type PriceExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PriceExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PriceExt
func (u PriceExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPriceExt creates a new  PriceExt.
func NewPriceExt(v LedgerVersion, value interface{}) (result PriceExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// Price is an XDR Struct defines as:
//
//   struct Price
//        {
//            int32 n; // numerator
//            int32 d; // denominator
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        };
//
type Price struct {
	N   Int32    `json:"n,omitempty"`
	D   Int32    `json:"d,omitempty"`
	Ext PriceExt `json:"ext,omitempty"`
}

// LotType is an XDR Enum defines as:
//
//   enum LotType {
//            LOT_TYPE_NO_AUCTION = 0,
//            LOT_TYPE_ENG = 1,
//            LOT_TYPE_COPART = 2,
//            LOT_TYPE_IAAI = 3,
//            LOT_TYPE_BNP = 4
//        };
//
type LotType int32

const (
	LotTypeLotTypeNoAuction LotType = 0
	LotTypeLotTypeEng       LotType = 1
	LotTypeLotTypeCopart    LotType = 2
	LotTypeLotTypeIaai      LotType = 3
	LotTypeLotTypeBnp       LotType = 4
)

var LotTypeAll = []LotType{
	LotTypeLotTypeNoAuction,
	LotTypeLotTypeEng,
	LotTypeLotTypeCopart,
	LotTypeLotTypeIaai,
	LotTypeLotTypeBnp,
}

var lotTypeMap = map[int32]string{
	0: "LotTypeLotTypeNoAuction",
	1: "LotTypeLotTypeEng",
	2: "LotTypeLotTypeCopart",
	3: "LotTypeLotTypeIaai",
	4: "LotTypeLotTypeBnp",
}

var lotTypeShortMap = map[int32]string{
	0: "lot_type_no_auction",
	1: "lot_type_eng",
	2: "lot_type_copart",
	3: "lot_type_iaai",
	4: "lot_type_bnp",
}

var lotTypeRevMap = map[string]int32{
	"LotTypeLotTypeNoAuction": 0,
	"LotTypeLotTypeEng":       1,
	"LotTypeLotTypeCopart":    2,
	"LotTypeLotTypeIaai":      3,
	"LotTypeLotTypeBnp":       4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for LotType
func (e LotType) ValidEnum(v int32) bool {
	_, ok := lotTypeMap[v]
	return ok
}
func (e LotType) isFlag() bool {
	for i := len(LotTypeAll) - 1; i >= 0; i-- {
		expected := LotType(2) << uint64(len(LotTypeAll)-1) >> uint64(len(LotTypeAll)-i)
		if expected != LotTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e LotType) String() string {
	name, _ := lotTypeMap[int32(e)]
	return name
}

func (e LotType) ShortString() string {
	name, _ := lotTypeShortMap[int32(e)]
	return name
}

func (e LotType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range LotTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *LotType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = LotType(t.Value)
	return nil
}

// LotState is an XDR Enum defines as:
//
//   enum LotState {
//            LOT_STATE_PENDING = 1,
//            LOT_STATE_SOLD = 2,
//            LOT_STATE_FINISHED = 3,
//            LOT_STATE_NEGOTIATIONS = 4
//        };
//
type LotState int32

const (
	LotStateLotStatePending      LotState = 1
	LotStateLotStateSold         LotState = 2
	LotStateLotStateFinished     LotState = 3
	LotStateLotStateNegotiations LotState = 4
)

var LotStateAll = []LotState{
	LotStateLotStatePending,
	LotStateLotStateSold,
	LotStateLotStateFinished,
	LotStateLotStateNegotiations,
}

var lotStateMap = map[int32]string{
	1: "LotStateLotStatePending",
	2: "LotStateLotStateSold",
	3: "LotStateLotStateFinished",
	4: "LotStateLotStateNegotiations",
}

var lotStateShortMap = map[int32]string{
	1: "lot_state_pending",
	2: "lot_state_sold",
	3: "lot_state_finished",
	4: "lot_state_negotiations",
}

var lotStateRevMap = map[string]int32{
	"LotStateLotStatePending":      1,
	"LotStateLotStateSold":         2,
	"LotStateLotStateFinished":     3,
	"LotStateLotStateNegotiations": 4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for LotState
func (e LotState) ValidEnum(v int32) bool {
	_, ok := lotStateMap[v]
	return ok
}
func (e LotState) isFlag() bool {
	for i := len(LotStateAll) - 1; i >= 0; i-- {
		expected := LotState(2) << uint64(len(LotStateAll)-1) >> uint64(len(LotStateAll)-i)
		if expected != LotStateAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e LotState) String() string {
	name, _ := lotStateMap[int32(e)]
	return name
}

func (e LotState) ShortString() string {
	name, _ := lotStateShortMap[int32(e)]
	return name
}

func (e LotState) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range LotStateAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *LotState) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = LotState(t.Value)
	return nil
}

// LotPricesExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LotPricesExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LotPricesExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LotPricesExt
func (u LotPricesExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLotPricesExt creates a new  LotPricesExt.
func NewLotPricesExt(v LedgerVersion, value interface{}) (result LotPricesExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LotPrices is an XDR Struct defines as:
//
//   struct LotPrices {
//            uint64 b2cTotalPrice;
//            uint64 b2cVehiclePrice;
//
//            uint64 b2bVehiclePrice;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            } ext;
//        };
//
type LotPrices struct {
	B2cTotalPrice   Uint64       `json:"b2cTotalPrice,omitempty"`
	B2cVehiclePrice Uint64       `json:"b2cVehiclePrice,omitempty"`
	B2bVehiclePrice Uint64       `json:"b2bVehiclePrice,omitempty"`
	Ext             LotPricesExt `json:"ext,omitempty"`
}

// LotEntryData is an XDR NestedUnion defines as:
//
//   union switch (LotType lt)
//                {
//                case LOT_TYPE_NO_AUCTION:
//                    LotPrices lotPrices;
//                }
//
type LotEntryData struct {
	Lt        LotType    `json:"lt,omitempty"`
	LotPrices *LotPrices `json:"lotPrices,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LotEntryData) SwitchFieldName() string {
	return "Lt"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LotEntryData
func (u LotEntryData) ArmForSwitch(sw int32) (string, bool) {
	switch LotType(sw) {
	case LotTypeLotTypeNoAuction:
		return "LotPrices", true
	}
	return "-", false
}

// NewLotEntryData creates a new  LotEntryData.
func NewLotEntryData(lt LotType, value interface{}) (result LotEntryData, err error) {
	result.Lt = lt
	switch LotType(lt) {
	case LotTypeLotTypeNoAuction:
		tv, ok := value.(LotPrices)
		if !ok {
			err = fmt.Errorf("invalid value, must be LotPrices")
			return
		}
		result.LotPrices = &tv
	}
	return
}

// MustLotPrices retrieves the LotPrices value from the union,
// panicing if the value is not set.
func (u LotEntryData) MustLotPrices() LotPrices {
	val, ok := u.GetLotPrices()

	if !ok {
		panic("arm LotPrices is not set")
	}

	return val
}

// GetLotPrices retrieves the LotPrices value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LotEntryData) GetLotPrices() (result LotPrices, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Lt))

	if armName == "LotPrices" {
		result = *u.LotPrices
		ok = true
	}

	return
}

// LotEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            case LEDGER_VERSION_B2B_B2C_PRICES:
//                union switch (LotType lt)
//                {
//                case LOT_TYPE_NO_AUCTION:
//                    LotPrices lotPrices;
//                } data;
//            }
//
type LotEntryExt struct {
	V    LedgerVersion `json:"v,omitempty"`
	Data *LotEntryData `json:"data,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LotEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LotEntryExt
func (u LotEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionLedgerVersionB2BB2CPrices:
		return "Data", true
	}
	return "-", false
}

// NewLotEntryExt creates a new  LotEntryExt.
func NewLotEntryExt(v LedgerVersion, value interface{}) (result LotEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionLedgerVersionB2BB2CPrices:
		tv, ok := value.(LotEntryData)
		if !ok {
			err = fmt.Errorf("invalid value, must be LotEntryData")
			return
		}
		result.Data = &tv
	}
	return
}

// MustData retrieves the Data value from the union,
// panicing if the value is not set.
func (u LotEntryExt) MustData() LotEntryData {
	val, ok := u.GetData()

	if !ok {
		panic("arm Data is not set")
	}

	return val
}

// GetData retrieves the Data value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LotEntryExt) GetData() (result LotEntryData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "Data" {
		result = *u.Data
		ok = true
	}

	return
}

// LotEntry is an XDR Struct defines as:
//
//   struct LotEntry
//        {
//            uint64 lotID;
//            AccountID organizerID;  // lot organizer
//            AccountID platformID;   // platform
//
//            AccountID* winnerID;
//            LotState state;
//
//            LotType lotType;
//            bool buyNowSupport;
//            string details<>;       // JSON string with full lot description
//
//            int64 startTime;        // auction start time in Unix epoch format
//            uint64 duration;        // duration in seconds
//
//            string currency<3>;     // need to keep this here in order to be sure
//                                    // that all the lots have proper currency
//            uint64 startPrice;      // all prices are indicated in cents
//            uint64 deposit;         // minimal deposit to pay for participating
//            uint64 buyNowPrice;
//            uint64 minStep;
//
//            // DEPRECATED
//            uint64 maxStep;
//
//            int64 createdAt;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            case LEDGER_VERSION_B2B_B2C_PRICES:
//                union switch (LotType lt)
//                {
//                case LOT_TYPE_NO_AUCTION:
//                    LotPrices lotPrices;
//                } data;
//            }
//            ext;
//        };
//
type LotEntry struct {
	LotId         Uint64      `json:"lotID,omitempty"`
	OrganizerId   AccountId   `json:"organizerID,omitempty"`
	PlatformId    AccountId   `json:"platformID,omitempty"`
	WinnerId      *AccountId  `json:"winnerID,omitempty"`
	State         LotState    `json:"state,omitempty"`
	LotType       LotType     `json:"lotType,omitempty"`
	BuyNowSupport bool        `json:"buyNowSupport,omitempty"`
	Details       string      `json:"details,omitempty"`
	StartTime     Int64       `json:"startTime,omitempty"`
	Duration      Uint64      `json:"duration,omitempty"`
	Currency      string      `json:"currency,omitempty" xdrmaxsize:"3"`
	StartPrice    Uint64      `json:"startPrice,omitempty"`
	Deposit       Uint64      `json:"deposit,omitempty"`
	BuyNowPrice   Uint64      `json:"buyNowPrice,omitempty"`
	MinStep       Uint64      `json:"minStep,omitempty"`
	MaxStep       Uint64      `json:"maxStep,omitempty"`
	CreatedAt     Int64       `json:"createdAt,omitempty"`
	Ext           LotEntryExt `json:"ext,omitempty"`
}

// MessageEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//            void;
//            }
//
type MessageEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u MessageEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of MessageEntryExt
func (u MessageEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewMessageEntryExt creates a new  MessageEntryExt.
func NewMessageEntryExt(v LedgerVersion, value interface{}) (result MessageEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// MessageEntry is an XDR Struct defines as:
//
//   struct MessageEntry
//        {
//            uint64 messageID;
//
//            AccountID senderID;
//            AccountID receiverID;
//            uint64 lotID;
//
//            string text<1000>;
//
//            int64 createdAt;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//            void;
//            }
//            ext;
//        };
//
type MessageEntry struct {
	MessageId  Uint64          `json:"messageID,omitempty"`
	SenderId   AccountId       `json:"senderID,omitempty"`
	ReceiverId AccountId       `json:"receiverID,omitempty"`
	LotId      Uint64          `json:"lotID,omitempty"`
	Text       string          `json:"text,omitempty" xdrmaxsize:"1000"`
	CreatedAt  Int64           `json:"createdAt,omitempty"`
	Ext        MessageEntryExt `json:"ext,omitempty"`
}

// MigrationVersion is an XDR Enum defines as:
//
//   enum MigrationVersion {
//            MIGRATION_MIGRATE_PART_AND_BID_REQUESTS = 1,
//            MIGRATION_MIGRATE_TO_LIVE_BID = 2
//        };
//
type MigrationVersion int32

const (
	MigrationVersionMigrationMigratePartAndBidRequests MigrationVersion = 1
	MigrationVersionMigrationMigrateToLiveBid          MigrationVersion = 2
)

var MigrationVersionAll = []MigrationVersion{
	MigrationVersionMigrationMigratePartAndBidRequests,
	MigrationVersionMigrationMigrateToLiveBid,
}

var migrationVersionMap = map[int32]string{
	1: "MigrationVersionMigrationMigratePartAndBidRequests",
	2: "MigrationVersionMigrationMigrateToLiveBid",
}

var migrationVersionShortMap = map[int32]string{
	1: "migration_migrate_part_and_bid_requests",
	2: "migration_migrate_to_live_bid",
}

var migrationVersionRevMap = map[string]int32{
	"MigrationVersionMigrationMigratePartAndBidRequests": 1,
	"MigrationVersionMigrationMigrateToLiveBid":          2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for MigrationVersion
func (e MigrationVersion) ValidEnum(v int32) bool {
	_, ok := migrationVersionMap[v]
	return ok
}
func (e MigrationVersion) isFlag() bool {
	for i := len(MigrationVersionAll) - 1; i >= 0; i-- {
		expected := MigrationVersion(2) << uint64(len(MigrationVersionAll)-1) >> uint64(len(MigrationVersionAll)-i)
		if expected != MigrationVersionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e MigrationVersion) String() string {
	name, _ := migrationVersionMap[int32(e)]
	return name
}

func (e MigrationVersion) ShortString() string {
	name, _ := migrationVersionShortMap[int32(e)]
	return name
}

func (e MigrationVersion) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range MigrationVersionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *MigrationVersion) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = MigrationVersion(t.Value)
	return nil
}

// MigrationEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                    case EMPTY_VERSION:
//                        void;
//                }
//
type MigrationEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u MigrationEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of MigrationEntryExt
func (u MigrationEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewMigrationEntryExt creates a new  MigrationEntryExt.
func NewMigrationEntryExt(v LedgerVersion, value interface{}) (result MigrationEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// MigrationEntry is an XDR Struct defines as:
//
//   struct MigrationEntry {
//            MigrationVersion migrationVersion;
//
//            union switch (LedgerVersion v)
//                {
//                    case EMPTY_VERSION:
//                        void;
//                }
//            ext;
//        };
//
type MigrationEntry struct {
	MigrationVersion MigrationVersion  `json:"migrationVersion,omitempty"`
	Ext              MigrationEntryExt `json:"ext,omitempty"`
}

// OfferEntryFlags is an XDR Enum defines as:
//
//   enum OfferEntryFlags
//        {
//            // issuer has authorized account to perform transactions with its credit
//            PASSIVE_FLAG = 1
//        };
//
type OfferEntryFlags int32

const (
	OfferEntryFlagsPassiveFlag OfferEntryFlags = 1
)

var OfferEntryFlagsAll = []OfferEntryFlags{
	OfferEntryFlagsPassiveFlag,
}

var offerEntryFlagsMap = map[int32]string{
	1: "OfferEntryFlagsPassiveFlag",
}

var offerEntryFlagsShortMap = map[int32]string{
	1: "passive_flag",
}

var offerEntryFlagsRevMap = map[string]int32{
	"OfferEntryFlagsPassiveFlag": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for OfferEntryFlags
func (e OfferEntryFlags) ValidEnum(v int32) bool {
	_, ok := offerEntryFlagsMap[v]
	return ok
}
func (e OfferEntryFlags) isFlag() bool {
	for i := len(OfferEntryFlagsAll) - 1; i >= 0; i-- {
		expected := OfferEntryFlags(2) << uint64(len(OfferEntryFlagsAll)-1) >> uint64(len(OfferEntryFlagsAll)-i)
		if expected != OfferEntryFlagsAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e OfferEntryFlags) String() string {
	name, _ := offerEntryFlagsMap[int32(e)]
	return name
}

func (e OfferEntryFlags) ShortString() string {
	name, _ := offerEntryFlagsShortMap[int32(e)]
	return name
}

func (e OfferEntryFlags) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range OfferEntryFlagsAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *OfferEntryFlags) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = OfferEntryFlags(t.Value)
	return nil
}

// OfferEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type OfferEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u OfferEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of OfferEntryExt
func (u OfferEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewOfferEntryExt creates a new  OfferEntryExt.
func NewOfferEntryExt(v LedgerVersion, value interface{}) (result OfferEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// OfferEntry is an XDR Struct defines as:
//
//   struct OfferEntry
//        {
//            AccountID sellerID;
//            uint64 offerID;
//            Asset selling; // A
//            Asset buying;  // B
//            int64 amount;  // amount of A
//
//            /* price for this offer:
//                price of A in terms of B
//                price=AmountB/AmountA=priceNumerator/priceDenominator
//                price is after fees
//            */
//            Price price;
//            uint32 flags; // see OfferEntryFlags
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type OfferEntry struct {
	SellerId AccountId     `json:"sellerID,omitempty"`
	OfferId  Uint64        `json:"offerID,omitempty"`
	Selling  Asset         `json:"selling,omitempty"`
	Buying   Asset         `json:"buying,omitempty"`
	Amount   Int64         `json:"amount,omitempty"`
	Price    Price         `json:"price,omitempty"`
	Flags    Uint32        `json:"flags,omitempty"`
	Ext      OfferEntryExt `json:"ext,omitempty"`
}

// ParticipantState is an XDR Enum defines as:
//
//   enum ParticipantState {
//            PARTICIPANT_STATE_PENDING = 0,
//            PARTICIPANT_STATE_REJECTED = 1,
//            PARTICIPANT_STATE_WINNER = 2,
//            PARTICIPANT_STATE_BUY_NOW_WINNER = 3,
//            PARTICIPANT_STATE_DEPOSIT_PENDING = 4
//        };
//
type ParticipantState int32

const (
	ParticipantStateParticipantStatePending        ParticipantState = 0
	ParticipantStateParticipantStateRejected       ParticipantState = 1
	ParticipantStateParticipantStateWinner         ParticipantState = 2
	ParticipantStateParticipantStateBuyNowWinner   ParticipantState = 3
	ParticipantStateParticipantStateDepositPending ParticipantState = 4
)

var ParticipantStateAll = []ParticipantState{
	ParticipantStateParticipantStatePending,
	ParticipantStateParticipantStateRejected,
	ParticipantStateParticipantStateWinner,
	ParticipantStateParticipantStateBuyNowWinner,
	ParticipantStateParticipantStateDepositPending,
}

var participantStateMap = map[int32]string{
	0: "ParticipantStateParticipantStatePending",
	1: "ParticipantStateParticipantStateRejected",
	2: "ParticipantStateParticipantStateWinner",
	3: "ParticipantStateParticipantStateBuyNowWinner",
	4: "ParticipantStateParticipantStateDepositPending",
}

var participantStateShortMap = map[int32]string{
	0: "participant_state_pending",
	1: "participant_state_rejected",
	2: "participant_state_winner",
	3: "participant_state_buy_now_winner",
	4: "participant_state_deposit_pending",
}

var participantStateRevMap = map[string]int32{
	"ParticipantStateParticipantStatePending":        0,
	"ParticipantStateParticipantStateRejected":       1,
	"ParticipantStateParticipantStateWinner":         2,
	"ParticipantStateParticipantStateBuyNowWinner":   3,
	"ParticipantStateParticipantStateDepositPending": 4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ParticipantState
func (e ParticipantState) ValidEnum(v int32) bool {
	_, ok := participantStateMap[v]
	return ok
}
func (e ParticipantState) isFlag() bool {
	for i := len(ParticipantStateAll) - 1; i >= 0; i-- {
		expected := ParticipantState(2) << uint64(len(ParticipantStateAll)-1) >> uint64(len(ParticipantStateAll)-i)
		if expected != ParticipantStateAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ParticipantState) String() string {
	name, _ := participantStateMap[int32(e)]
	return name
}

func (e ParticipantState) ShortString() string {
	name, _ := participantStateShortMap[int32(e)]
	return name
}

func (e ParticipantState) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ParticipantStateAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ParticipantState) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ParticipantState(t.Value)
	return nil
}

// ParticipantLiveBidData is an XDR Struct defines as:
//
//   struct ParticipantLiveBidData
//        {
//            uint64 limit;
//            uint64 maxBid;
//        };
//
type ParticipantLiveBidData struct {
	Limit  Uint64 `json:"limit,omitempty"`
	MaxBid Uint64 `json:"maxBid,omitempty"`
}

// ParticipantEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            case LEDGER_VERSION_LIVE_BID:
//                ParticipantLiveBidData liveBidData;
//            }
//
type ParticipantEntryExt struct {
	V           LedgerVersion           `json:"v,omitempty"`
	LiveBidData *ParticipantLiveBidData `json:"liveBidData,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ParticipantEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ParticipantEntryExt
func (u ParticipantEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionLedgerVersionLiveBid:
		return "LiveBidData", true
	}
	return "-", false
}

// NewParticipantEntryExt creates a new  ParticipantEntryExt.
func NewParticipantEntryExt(v LedgerVersion, value interface{}) (result ParticipantEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionLedgerVersionLiveBid:
		tv, ok := value.(ParticipantLiveBidData)
		if !ok {
			err = fmt.Errorf("invalid value, must be ParticipantLiveBidData")
			return
		}
		result.LiveBidData = &tv
	}
	return
}

// MustLiveBidData retrieves the LiveBidData value from the union,
// panicing if the value is not set.
func (u ParticipantEntryExt) MustLiveBidData() ParticipantLiveBidData {
	val, ok := u.GetLiveBidData()

	if !ok {
		panic("arm LiveBidData is not set")
	}

	return val
}

// GetLiveBidData retrieves the LiveBidData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ParticipantEntryExt) GetLiveBidData() (result ParticipantLiveBidData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "LiveBidData" {
		result = *u.LiveBidData
		ok = true
	}

	return
}

// ParticipantEntry is an XDR Struct defines as:
//
//   struct ParticipantEntry
//        {
//            AccountID accountID;
//            uint64 lotID;
//
//            ParticipantState state;
//            bool buyNow;
//
//            uint64 bestBid;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            case LEDGER_VERSION_LIVE_BID:
//                ParticipantLiveBidData liveBidData;
//            }
//            ext;
//        };
//
type ParticipantEntry struct {
	AccountId AccountId           `json:"accountID,omitempty"`
	LotId     Uint64              `json:"lotID,omitempty"`
	State     ParticipantState    `json:"state,omitempty"`
	BuyNow    bool                `json:"buyNow,omitempty"`
	BestBid   Uint64              `json:"bestBid,omitempty"`
	Ext       ParticipantEntryExt `json:"ext,omitempty"`
}

// ReferenceEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ReferenceEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReferenceEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReferenceEntryExt
func (u ReferenceEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewReferenceEntryExt creates a new  ReferenceEntryExt.
func NewReferenceEntryExt(v LedgerVersion, value interface{}) (result ReferenceEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ReferenceEntry is an XDR Struct defines as:
//
//   struct ReferenceEntry
//    {
//    	AccountID sender;
//        string64 reference;
//
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ReferenceEntry struct {
	Sender    AccountId         `json:"sender,omitempty"`
	Reference String64          `json:"reference,omitempty"`
	Ext       ReferenceEntryExt `json:"ext,omitempty"`
}

// ReviewableRequestType is an XDR Enum defines as:
//
//   enum ReviewableRequestType
//    {
//        REQUEST_TYPE_CREATE_LOT = 1,            // creates new lot
//        REQUEST_TYPE_REQUEST_PARTICIPATION = 2,
//        REQUEST_TYPE_CREATE_BID = 3
//    };
//
type ReviewableRequestType int32

const (
	ReviewableRequestTypeRequestTypeCreateLot            ReviewableRequestType = 1
	ReviewableRequestTypeRequestTypeRequestParticipation ReviewableRequestType = 2
	ReviewableRequestTypeRequestTypeCreateBid            ReviewableRequestType = 3
)

var ReviewableRequestTypeAll = []ReviewableRequestType{
	ReviewableRequestTypeRequestTypeCreateLot,
	ReviewableRequestTypeRequestTypeRequestParticipation,
	ReviewableRequestTypeRequestTypeCreateBid,
}

var reviewableRequestTypeMap = map[int32]string{
	1: "ReviewableRequestTypeRequestTypeCreateLot",
	2: "ReviewableRequestTypeRequestTypeRequestParticipation",
	3: "ReviewableRequestTypeRequestTypeCreateBid",
}

var reviewableRequestTypeShortMap = map[int32]string{
	1: "request_type_create_lot",
	2: "request_type_request_participation",
	3: "request_type_create_bid",
}

var reviewableRequestTypeRevMap = map[string]int32{
	"ReviewableRequestTypeRequestTypeCreateLot":            1,
	"ReviewableRequestTypeRequestTypeRequestParticipation": 2,
	"ReviewableRequestTypeRequestTypeCreateBid":            3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ReviewableRequestType
func (e ReviewableRequestType) ValidEnum(v int32) bool {
	_, ok := reviewableRequestTypeMap[v]
	return ok
}
func (e ReviewableRequestType) isFlag() bool {
	for i := len(ReviewableRequestTypeAll) - 1; i >= 0; i-- {
		expected := ReviewableRequestType(2) << uint64(len(ReviewableRequestTypeAll)-1) >> uint64(len(ReviewableRequestTypeAll)-i)
		if expected != ReviewableRequestTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ReviewableRequestType) String() string {
	name, _ := reviewableRequestTypeMap[int32(e)]
	return name
}

func (e ReviewableRequestType) ShortString() string {
	name, _ := reviewableRequestTypeShortMap[int32(e)]
	return name
}

func (e ReviewableRequestType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ReviewableRequestTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ReviewableRequestType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ReviewableRequestType(t.Value)
	return nil
}

// ReviewableRequestEntryBody is an XDR NestedUnion defines as:
//
//   union switch(ReviewableRequestType type)
//        {
//         case REQUEST_TYPE_CREATE_LOT:
//            LotCreationRequest lotCreationRequest;
//         case REQUEST_TYPE_CREATE_BID:
//            BidRequest bidRequest;
//         case REQUEST_TYPE_REQUEST_PARTICIPATION:
//            ParticipationRequest participationRequest;
//        }
//
type ReviewableRequestEntryBody struct {
	Type                 ReviewableRequestType `json:"type,omitempty"`
	LotCreationRequest   *LotCreationRequest   `json:"lotCreationRequest,omitempty"`
	BidRequest           *BidRequest           `json:"bidRequest,omitempty"`
	ParticipationRequest *ParticipationRequest `json:"participationRequest,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewableRequestEntryBody) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewableRequestEntryBody
func (u ReviewableRequestEntryBody) ArmForSwitch(sw int32) (string, bool) {
	switch ReviewableRequestType(sw) {
	case ReviewableRequestTypeRequestTypeCreateLot:
		return "LotCreationRequest", true
	case ReviewableRequestTypeRequestTypeCreateBid:
		return "BidRequest", true
	case ReviewableRequestTypeRequestTypeRequestParticipation:
		return "ParticipationRequest", true
	}
	return "-", false
}

// NewReviewableRequestEntryBody creates a new  ReviewableRequestEntryBody.
func NewReviewableRequestEntryBody(aType ReviewableRequestType, value interface{}) (result ReviewableRequestEntryBody, err error) {
	result.Type = aType
	switch ReviewableRequestType(aType) {
	case ReviewableRequestTypeRequestTypeCreateLot:
		tv, ok := value.(LotCreationRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be LotCreationRequest")
			return
		}
		result.LotCreationRequest = &tv
	case ReviewableRequestTypeRequestTypeCreateBid:
		tv, ok := value.(BidRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be BidRequest")
			return
		}
		result.BidRequest = &tv
	case ReviewableRequestTypeRequestTypeRequestParticipation:
		tv, ok := value.(ParticipationRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be ParticipationRequest")
			return
		}
		result.ParticipationRequest = &tv
	}
	return
}

// MustLotCreationRequest retrieves the LotCreationRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustLotCreationRequest() LotCreationRequest {
	val, ok := u.GetLotCreationRequest()

	if !ok {
		panic("arm LotCreationRequest is not set")
	}

	return val
}

// GetLotCreationRequest retrieves the LotCreationRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetLotCreationRequest() (result LotCreationRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "LotCreationRequest" {
		result = *u.LotCreationRequest
		ok = true
	}

	return
}

// MustBidRequest retrieves the BidRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustBidRequest() BidRequest {
	val, ok := u.GetBidRequest()

	if !ok {
		panic("arm BidRequest is not set")
	}

	return val
}

// GetBidRequest retrieves the BidRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetBidRequest() (result BidRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "BidRequest" {
		result = *u.BidRequest
		ok = true
	}

	return
}

// MustParticipationRequest retrieves the ParticipationRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustParticipationRequest() ParticipationRequest {
	val, ok := u.GetParticipationRequest()

	if !ok {
		panic("arm ParticipationRequest is not set")
	}

	return val
}

// GetParticipationRequest retrieves the ParticipationRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetParticipationRequest() (result ParticipationRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ParticipationRequest" {
		result = *u.ParticipationRequest
		ok = true
	}

	return
}

// ReviewableRequestEntryExt is an XDR NestedUnion defines as:
//
//   union switch(LedgerVersion v)
//        {
//         case EMPTY_VERSION:
//            void;
//        }
//
type ReviewableRequestEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewableRequestEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewableRequestEntryExt
func (u ReviewableRequestEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewReviewableRequestEntryExt creates a new  ReviewableRequestEntryExt.
func NewReviewableRequestEntryExt(v LedgerVersion, value interface{}) (result ReviewableRequestEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ReviewableRequestEntry is an XDR Struct defines as:
//
//   struct ReviewableRequestEntry
//    {
//        uint64 requestID;
//        AccountID requestor;
//        AccountID reviewer;
//        Hash hash;
//        string256 rejectReason;
//        string64* reference;
//        int64 createdAt;
//
//        // body of the request
//        union switch(ReviewableRequestType type)
//        {
//         case REQUEST_TYPE_CREATE_LOT:
//            LotCreationRequest lotCreationRequest;
//         case REQUEST_TYPE_CREATE_BID:
//            BidRequest bidRequest;
//         case REQUEST_TYPE_REQUEST_PARTICIPATION:
//            ParticipationRequest participationRequest;
//        } body;
//
//        //reserved for future use
//        union switch(LedgerVersion v)
//        {
//         case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type ReviewableRequestEntry struct {
	RequestId    Uint64                     `json:"requestID,omitempty"`
	Requestor    AccountId                  `json:"requestor,omitempty"`
	Reviewer     AccountId                  `json:"reviewer,omitempty"`
	Hash         Hash                       `json:"hash,omitempty"`
	RejectReason String256                  `json:"rejectReason,omitempty"`
	Reference    *String64                  `json:"reference,omitempty"`
	CreatedAt    Int64                      `json:"createdAt,omitempty"`
	Body         ReviewableRequestEntryBody `json:"body,omitempty"`
	Ext          ReviewableRequestEntryExt  `json:"ext,omitempty"`
}

// TrustLineFlags is an XDR Enum defines as:
//
//   enum TrustLineFlags
//        {
//            // issuer has authorized account to perform transactions with its credit
//            AUTHORIZED_FLAG = 1
//        };
//
type TrustLineFlags int32

const (
	TrustLineFlagsAuthorizedFlag TrustLineFlags = 1
)

var TrustLineFlagsAll = []TrustLineFlags{
	TrustLineFlagsAuthorizedFlag,
}

var trustLineFlagsMap = map[int32]string{
	1: "TrustLineFlagsAuthorizedFlag",
}

var trustLineFlagsShortMap = map[int32]string{
	1: "authorized_flag",
}

var trustLineFlagsRevMap = map[string]int32{
	"TrustLineFlagsAuthorizedFlag": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for TrustLineFlags
func (e TrustLineFlags) ValidEnum(v int32) bool {
	_, ok := trustLineFlagsMap[v]
	return ok
}
func (e TrustLineFlags) isFlag() bool {
	for i := len(TrustLineFlagsAll) - 1; i >= 0; i-- {
		expected := TrustLineFlags(2) << uint64(len(TrustLineFlagsAll)-1) >> uint64(len(TrustLineFlagsAll)-i)
		if expected != TrustLineFlagsAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e TrustLineFlags) String() string {
	name, _ := trustLineFlagsMap[int32(e)]
	return name
}

func (e TrustLineFlags) ShortString() string {
	name, _ := trustLineFlagsShortMap[int32(e)]
	return name
}

func (e TrustLineFlags) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range TrustLineFlagsAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *TrustLineFlags) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = TrustLineFlags(t.Value)
	return nil
}

// TrustLineEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type TrustLineEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TrustLineEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TrustLineEntryExt
func (u TrustLineEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTrustLineEntryExt creates a new  TrustLineEntryExt.
func NewTrustLineEntryExt(v LedgerVersion, value interface{}) (result TrustLineEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TrustLineEntry is an XDR Struct defines as:
//
//   struct TrustLineEntry
//        {
//            AccountID accountID; // account this trustline belongs to
//            Asset asset;         // type of asset (with issuer)
//            int64 balance;       // how much of this asset the user has.
//                                 // Asset defines the unit for this;
//
//            int64 limit;  // balance cannot be above this
//            uint32 flags; // see TrustLineFlags
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type TrustLineEntry struct {
	AccountId AccountId         `json:"accountID,omitempty"`
	Asset     Asset             `json:"asset,omitempty"`
	Balance   Int64             `json:"balance,omitempty"`
	Limit     Int64             `json:"limit,omitempty"`
	Flags     Uint32            `json:"flags,omitempty"`
	Ext       TrustLineEntryExt `json:"ext,omitempty"`
}

// ThresholdIndexes is an XDR Enum defines as:
//
//   enum ThresholdIndexes
//        {
//            THRESHOLD_MASTER_WEIGHT = 0,
//            THRESHOLD_LOW = 1,
//            THRESHOLD_MED = 2,
//            THRESHOLD_HIGH = 3
//        };
//
type ThresholdIndexes int32

const (
	ThresholdIndexesThresholdMasterWeight ThresholdIndexes = 0
	ThresholdIndexesThresholdLow          ThresholdIndexes = 1
	ThresholdIndexesThresholdMed          ThresholdIndexes = 2
	ThresholdIndexesThresholdHigh         ThresholdIndexes = 3
)

var ThresholdIndexesAll = []ThresholdIndexes{
	ThresholdIndexesThresholdMasterWeight,
	ThresholdIndexesThresholdLow,
	ThresholdIndexesThresholdMed,
	ThresholdIndexesThresholdHigh,
}

var thresholdIndexesMap = map[int32]string{
	0: "ThresholdIndexesThresholdMasterWeight",
	1: "ThresholdIndexesThresholdLow",
	2: "ThresholdIndexesThresholdMed",
	3: "ThresholdIndexesThresholdHigh",
}

var thresholdIndexesShortMap = map[int32]string{
	0: "threshold_master_weight",
	1: "threshold_low",
	2: "threshold_med",
	3: "threshold_high",
}

var thresholdIndexesRevMap = map[string]int32{
	"ThresholdIndexesThresholdMasterWeight": 0,
	"ThresholdIndexesThresholdLow":          1,
	"ThresholdIndexesThresholdMed":          2,
	"ThresholdIndexesThresholdHigh":         3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ThresholdIndexes
func (e ThresholdIndexes) ValidEnum(v int32) bool {
	_, ok := thresholdIndexesMap[v]
	return ok
}
func (e ThresholdIndexes) isFlag() bool {
	for i := len(ThresholdIndexesAll) - 1; i >= 0; i-- {
		expected := ThresholdIndexes(2) << uint64(len(ThresholdIndexesAll)-1) >> uint64(len(ThresholdIndexesAll)-i)
		if expected != ThresholdIndexesAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ThresholdIndexes) String() string {
	name, _ := thresholdIndexesMap[int32(e)]
	return name
}

func (e ThresholdIndexes) ShortString() string {
	name, _ := thresholdIndexesShortMap[int32(e)]
	return name
}

func (e ThresholdIndexes) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ThresholdIndexesAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ThresholdIndexes) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ThresholdIndexes(t.Value)
	return nil
}

// LedgerEntryType is an XDR Enum defines as:
//
//   enum LedgerEntryType
//        {
//            ACCOUNT = 0,
//            TRUSTLINE = 1,
//            OFFER = 2,
//            LOT = 3,
//            PARTICIPANT = 4,
//            MESSAGE = 5,
//            MIGRATION =6,
//            REVIEWABLE_REQUEST = 7,
//            REFERENCE = 8
//        };
//
type LedgerEntryType int32

const (
	LedgerEntryTypeAccount           LedgerEntryType = 0
	LedgerEntryTypeTrustline         LedgerEntryType = 1
	LedgerEntryTypeOffer             LedgerEntryType = 2
	LedgerEntryTypeLot               LedgerEntryType = 3
	LedgerEntryTypeParticipant       LedgerEntryType = 4
	LedgerEntryTypeMessage           LedgerEntryType = 5
	LedgerEntryTypeMigration         LedgerEntryType = 6
	LedgerEntryTypeReviewableRequest LedgerEntryType = 7
	LedgerEntryTypeReference         LedgerEntryType = 8
)

var LedgerEntryTypeAll = []LedgerEntryType{
	LedgerEntryTypeAccount,
	LedgerEntryTypeTrustline,
	LedgerEntryTypeOffer,
	LedgerEntryTypeLot,
	LedgerEntryTypeParticipant,
	LedgerEntryTypeMessage,
	LedgerEntryTypeMigration,
	LedgerEntryTypeReviewableRequest,
	LedgerEntryTypeReference,
}

var ledgerEntryTypeMap = map[int32]string{
	0: "LedgerEntryTypeAccount",
	1: "LedgerEntryTypeTrustline",
	2: "LedgerEntryTypeOffer",
	3: "LedgerEntryTypeLot",
	4: "LedgerEntryTypeParticipant",
	5: "LedgerEntryTypeMessage",
	6: "LedgerEntryTypeMigration",
	7: "LedgerEntryTypeReviewableRequest",
	8: "LedgerEntryTypeReference",
}

var ledgerEntryTypeShortMap = map[int32]string{
	0: "account",
	1: "trustline",
	2: "offer",
	3: "lot",
	4: "participant",
	5: "message",
	6: "migration",
	7: "reviewable_request",
	8: "reference",
}

var ledgerEntryTypeRevMap = map[string]int32{
	"LedgerEntryTypeAccount":           0,
	"LedgerEntryTypeTrustline":         1,
	"LedgerEntryTypeOffer":             2,
	"LedgerEntryTypeLot":               3,
	"LedgerEntryTypeParticipant":       4,
	"LedgerEntryTypeMessage":           5,
	"LedgerEntryTypeMigration":         6,
	"LedgerEntryTypeReviewableRequest": 7,
	"LedgerEntryTypeReference":         8,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for LedgerEntryType
func (e LedgerEntryType) ValidEnum(v int32) bool {
	_, ok := ledgerEntryTypeMap[v]
	return ok
}
func (e LedgerEntryType) isFlag() bool {
	for i := len(LedgerEntryTypeAll) - 1; i >= 0; i-- {
		expected := LedgerEntryType(2) << uint64(len(LedgerEntryTypeAll)-1) >> uint64(len(LedgerEntryTypeAll)-i)
		if expected != LedgerEntryTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e LedgerEntryType) String() string {
	name, _ := ledgerEntryTypeMap[int32(e)]
	return name
}

func (e LedgerEntryType) ShortString() string {
	name, _ := ledgerEntryTypeShortMap[int32(e)]
	return name
}

func (e LedgerEntryType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range LedgerEntryTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *LedgerEntryType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = LedgerEntryType(t.Value)
	return nil
}

// LedgerEntryData is an XDR NestedUnion defines as:
//
//   union switch (LedgerEntryType type)
//            {
//            case ACCOUNT:
//                AccountEntry account;
//            case TRUSTLINE:
//                TrustLineEntry trustLine;
//            case OFFER:
//                OfferEntry offer;
//            case LOT:
//                LotEntry lot;
//            case PARTICIPANT:
//                ParticipantEntry participant;
//            case MESSAGE:
//                MessageEntry message;
//            case MIGRATION:
//                MigrationEntry migration;
//            case REVIEWABLE_REQUEST:
//                ReviewableRequestEntry reviewableRequest;
//            case REFERENCE:
//                ReferenceEntry reference;
//            }
//
type LedgerEntryData struct {
	Type              LedgerEntryType         `json:"type,omitempty"`
	Account           *AccountEntry           `json:"account,omitempty"`
	TrustLine         *TrustLineEntry         `json:"trustLine,omitempty"`
	Offer             *OfferEntry             `json:"offer,omitempty"`
	Lot               *LotEntry               `json:"lot,omitempty"`
	Participant       *ParticipantEntry       `json:"participant,omitempty"`
	Message           *MessageEntry           `json:"message,omitempty"`
	Migration         *MigrationEntry         `json:"migration,omitempty"`
	ReviewableRequest *ReviewableRequestEntry `json:"reviewableRequest,omitempty"`
	Reference         *ReferenceEntry         `json:"reference,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerEntryData) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerEntryData
func (u LedgerEntryData) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerEntryType(sw) {
	case LedgerEntryTypeAccount:
		return "Account", true
	case LedgerEntryTypeTrustline:
		return "TrustLine", true
	case LedgerEntryTypeOffer:
		return "Offer", true
	case LedgerEntryTypeLot:
		return "Lot", true
	case LedgerEntryTypeParticipant:
		return "Participant", true
	case LedgerEntryTypeMessage:
		return "Message", true
	case LedgerEntryTypeMigration:
		return "Migration", true
	case LedgerEntryTypeReviewableRequest:
		return "ReviewableRequest", true
	case LedgerEntryTypeReference:
		return "Reference", true
	}
	return "-", false
}

// NewLedgerEntryData creates a new  LedgerEntryData.
func NewLedgerEntryData(aType LedgerEntryType, value interface{}) (result LedgerEntryData, err error) {
	result.Type = aType
	switch LedgerEntryType(aType) {
	case LedgerEntryTypeAccount:
		tv, ok := value.(AccountEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountEntry")
			return
		}
		result.Account = &tv
	case LedgerEntryTypeTrustline:
		tv, ok := value.(TrustLineEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be TrustLineEntry")
			return
		}
		result.TrustLine = &tv
	case LedgerEntryTypeOffer:
		tv, ok := value.(OfferEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be OfferEntry")
			return
		}
		result.Offer = &tv
	case LedgerEntryTypeLot:
		tv, ok := value.(LotEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LotEntry")
			return
		}
		result.Lot = &tv
	case LedgerEntryTypeParticipant:
		tv, ok := value.(ParticipantEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be ParticipantEntry")
			return
		}
		result.Participant = &tv
	case LedgerEntryTypeMessage:
		tv, ok := value.(MessageEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be MessageEntry")
			return
		}
		result.Message = &tv
	case LedgerEntryTypeMigration:
		tv, ok := value.(MigrationEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be MigrationEntry")
			return
		}
		result.Migration = &tv
	case LedgerEntryTypeReviewableRequest:
		tv, ok := value.(ReviewableRequestEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReviewableRequestEntry")
			return
		}
		result.ReviewableRequest = &tv
	case LedgerEntryTypeReference:
		tv, ok := value.(ReferenceEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReferenceEntry")
			return
		}
		result.Reference = &tv
	}
	return
}

// MustAccount retrieves the Account value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustAccount() AccountEntry {
	val, ok := u.GetAccount()

	if !ok {
		panic("arm Account is not set")
	}

	return val
}

// GetAccount retrieves the Account value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetAccount() (result AccountEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Account" {
		result = *u.Account
		ok = true
	}

	return
}

// MustTrustLine retrieves the TrustLine value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustTrustLine() TrustLineEntry {
	val, ok := u.GetTrustLine()

	if !ok {
		panic("arm TrustLine is not set")
	}

	return val
}

// GetTrustLine retrieves the TrustLine value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetTrustLine() (result TrustLineEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "TrustLine" {
		result = *u.TrustLine
		ok = true
	}

	return
}

// MustOffer retrieves the Offer value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustOffer() OfferEntry {
	val, ok := u.GetOffer()

	if !ok {
		panic("arm Offer is not set")
	}

	return val
}

// GetOffer retrieves the Offer value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetOffer() (result OfferEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Offer" {
		result = *u.Offer
		ok = true
	}

	return
}

// MustLot retrieves the Lot value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustLot() LotEntry {
	val, ok := u.GetLot()

	if !ok {
		panic("arm Lot is not set")
	}

	return val
}

// GetLot retrieves the Lot value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetLot() (result LotEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Lot" {
		result = *u.Lot
		ok = true
	}

	return
}

// MustParticipant retrieves the Participant value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustParticipant() ParticipantEntry {
	val, ok := u.GetParticipant()

	if !ok {
		panic("arm Participant is not set")
	}

	return val
}

// GetParticipant retrieves the Participant value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetParticipant() (result ParticipantEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Participant" {
		result = *u.Participant
		ok = true
	}

	return
}

// MustMessage retrieves the Message value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustMessage() MessageEntry {
	val, ok := u.GetMessage()

	if !ok {
		panic("arm Message is not set")
	}

	return val
}

// GetMessage retrieves the Message value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetMessage() (result MessageEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Message" {
		result = *u.Message
		ok = true
	}

	return
}

// MustMigration retrieves the Migration value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustMigration() MigrationEntry {
	val, ok := u.GetMigration()

	if !ok {
		panic("arm Migration is not set")
	}

	return val
}

// GetMigration retrieves the Migration value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetMigration() (result MigrationEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Migration" {
		result = *u.Migration
		ok = true
	}

	return
}

// MustReviewableRequest retrieves the ReviewableRequest value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustReviewableRequest() ReviewableRequestEntry {
	val, ok := u.GetReviewableRequest()

	if !ok {
		panic("arm ReviewableRequest is not set")
	}

	return val
}

// GetReviewableRequest retrieves the ReviewableRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetReviewableRequest() (result ReviewableRequestEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ReviewableRequest" {
		result = *u.ReviewableRequest
		ok = true
	}

	return
}

// MustReference retrieves the Reference value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustReference() ReferenceEntry {
	val, ok := u.GetReference()

	if !ok {
		panic("arm Reference is not set")
	}

	return val
}

// GetReference retrieves the Reference value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetReference() (result ReferenceEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Reference" {
		result = *u.Reference
		ok = true
	}

	return
}

// LedgerEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerEntryExt
func (u LedgerEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerEntryExt creates a new  LedgerEntryExt.
func NewLedgerEntryExt(v LedgerVersion, value interface{}) (result LedgerEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerEntry is an XDR Struct defines as:
//
//   struct LedgerEntry
//        {
//            uint32 lastModifiedLedgerSeq; // ledger the LedgerEntry was last changed
//
//            union switch (LedgerEntryType type)
//            {
//            case ACCOUNT:
//                AccountEntry account;
//            case TRUSTLINE:
//                TrustLineEntry trustLine;
//            case OFFER:
//                OfferEntry offer;
//            case LOT:
//                LotEntry lot;
//            case PARTICIPANT:
//                ParticipantEntry participant;
//            case MESSAGE:
//                MessageEntry message;
//            case MIGRATION:
//                MigrationEntry migration;
//            case REVIEWABLE_REQUEST:
//                ReviewableRequestEntry reviewableRequest;
//            case REFERENCE:
//                ReferenceEntry reference;
//            }
//            data;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type LedgerEntry struct {
	LastModifiedLedgerSeq Uint32          `json:"lastModifiedLedgerSeq,omitempty"`
	Data                  LedgerEntryData `json:"data,omitempty"`
	Ext                   LedgerEntryExt  `json:"ext,omitempty"`
}

// EnvelopeType is an XDR Enum defines as:
//
//   enum EnvelopeType
//        {
//            ENVELOPE_TYPE_SCP = 1,
//            ENVELOPE_TYPE_TX = 2,
//            ENVELOPE_TYPE_AUTH = 3
//        };
//
type EnvelopeType int32

const (
	EnvelopeTypeEnvelopeTypeScp  EnvelopeType = 1
	EnvelopeTypeEnvelopeTypeTx   EnvelopeType = 2
	EnvelopeTypeEnvelopeTypeAuth EnvelopeType = 3
)

var EnvelopeTypeAll = []EnvelopeType{
	EnvelopeTypeEnvelopeTypeScp,
	EnvelopeTypeEnvelopeTypeTx,
	EnvelopeTypeEnvelopeTypeAuth,
}

var envelopeTypeMap = map[int32]string{
	1: "EnvelopeTypeEnvelopeTypeScp",
	2: "EnvelopeTypeEnvelopeTypeTx",
	3: "EnvelopeTypeEnvelopeTypeAuth",
}

var envelopeTypeShortMap = map[int32]string{
	1: "envelope_type_scp",
	2: "envelope_type_tx",
	3: "envelope_type_auth",
}

var envelopeTypeRevMap = map[string]int32{
	"EnvelopeTypeEnvelopeTypeScp":  1,
	"EnvelopeTypeEnvelopeTypeTx":   2,
	"EnvelopeTypeEnvelopeTypeAuth": 3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for EnvelopeType
func (e EnvelopeType) ValidEnum(v int32) bool {
	_, ok := envelopeTypeMap[v]
	return ok
}
func (e EnvelopeType) isFlag() bool {
	for i := len(EnvelopeTypeAll) - 1; i >= 0; i-- {
		expected := EnvelopeType(2) << uint64(len(EnvelopeTypeAll)-1) >> uint64(len(EnvelopeTypeAll)-i)
		if expected != EnvelopeTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e EnvelopeType) String() string {
	name, _ := envelopeTypeMap[int32(e)]
	return name
}

func (e EnvelopeType) ShortString() string {
	name, _ := envelopeTypeShortMap[int32(e)]
	return name
}

func (e EnvelopeType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range EnvelopeTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *EnvelopeType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = EnvelopeType(t.Value)
	return nil
}

// UpgradeType is an XDR Typedef defines as:
//
//   typedef opaque UpgradeType<128>;
//
type UpgradeType []byte

// StellarValueExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type StellarValueExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u StellarValueExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of StellarValueExt
func (u StellarValueExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewStellarValueExt creates a new  StellarValueExt.
func NewStellarValueExt(v LedgerVersion, value interface{}) (result StellarValueExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// StellarValue is an XDR Struct defines as:
//
//   struct StellarValue
//    {
//        Hash txSetHash;   // transaction set to apply to previous ledger
//        uint64 closeTime; // network close time
//
//        // upgrades to apply to the previous ledger (usually empty)
//        // this is a vector of encoded 'LedgerUpgrade' so that nodes can drop
//        // unknown steps during consensus if needed.
//        // see notes below on 'LedgerUpgrade' for more detail
//        // max size is dictated by number of upgrade types (+ room for future)
//        UpgradeType upgrades<6>;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type StellarValue struct {
	TxSetHash Hash            `json:"txSetHash,omitempty"`
	CloseTime Uint64          `json:"closeTime,omitempty"`
	Upgrades  []UpgradeType   `json:"upgrades,omitempty" xdrmaxsize:"6"`
	Ext       StellarValueExt `json:"ext,omitempty"`
}

// LedgerHeaderExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type LedgerHeaderExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerHeaderExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerHeaderExt
func (u LedgerHeaderExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerHeaderExt creates a new  LedgerHeaderExt.
func NewLedgerHeaderExt(v LedgerVersion, value interface{}) (result LedgerHeaderExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerHeader is an XDR Struct defines as:
//
//   struct LedgerHeader
//    {
//        uint32 ledgerVersion;    // the protocol version of the ledger
//        Hash previousLedgerHash; // hash of the previous ledger header
//        StellarValue scpValue;   // what consensus agreed to
//        Hash txSetResultHash;    // the TransactionResultSet that led to this ledger
//        Hash bucketListHash;     // hash of the ledger state
//
//        uint32 ledgerSeq; // sequence number of this ledger
//
//        int64 totalCoins; // total number of stroops in existence.
//                          // 10,000,000 stroops in 1 XLM
//
//        int64 feePool;       // fees burned since last inflation run
//        uint32 inflationSeq; // inflation sequence number
//
//        uint64 idPool; // last used global ID, used for generating objects
//
//        uint32 baseFee;     // base fee per operation in stroops
//        uint32 baseReserve; // account base reserve in stroops
//
//        uint32 maxTxSetSize; // maximum size a transaction set can be
//
//        Hash skipList[4]; // hashes of ledgers in the past. allows you to jump back
//                          // in time without walking the chain back ledger by ledger
//                          // each slot contains the oldest ledger that is mod of
//                          // either 50  5000  50000 or 500000 depending on index
//                          // skipList[0] mod(50), skipList[1] mod(5000), etc
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type LedgerHeader struct {
	LedgerVersion      Uint32          `json:"ledgerVersion,omitempty"`
	PreviousLedgerHash Hash            `json:"previousLedgerHash,omitempty"`
	ScpValue           StellarValue    `json:"scpValue,omitempty"`
	TxSetResultHash    Hash            `json:"txSetResultHash,omitempty"`
	BucketListHash     Hash            `json:"bucketListHash,omitempty"`
	LedgerSeq          Uint32          `json:"ledgerSeq,omitempty"`
	TotalCoins         Int64           `json:"totalCoins,omitempty"`
	FeePool            Int64           `json:"feePool,omitempty"`
	InflationSeq       Uint32          `json:"inflationSeq,omitempty"`
	IdPool             Uint64          `json:"idPool,omitempty"`
	BaseFee            Uint32          `json:"baseFee,omitempty"`
	BaseReserve        Uint32          `json:"baseReserve,omitempty"`
	MaxTxSetSize       Uint32          `json:"maxTxSetSize,omitempty"`
	SkipList           [4]Hash         `json:"skipList,omitempty"`
	Ext                LedgerHeaderExt `json:"ext,omitempty"`
}

// LedgerUpgradeType is an XDR Enum defines as:
//
//   enum LedgerUpgradeType
//    {
//        LEDGER_UPGRADE_VERSION = 1,
//        LEDGER_UPGRADE_BASE_FEE = 2,
//        LEDGER_UPGRADE_MAX_TX_SET_SIZE = 3
//    };
//
type LedgerUpgradeType int32

const (
	LedgerUpgradeTypeLedgerUpgradeVersion      LedgerUpgradeType = 1
	LedgerUpgradeTypeLedgerUpgradeBaseFee      LedgerUpgradeType = 2
	LedgerUpgradeTypeLedgerUpgradeMaxTxSetSize LedgerUpgradeType = 3
)

var LedgerUpgradeTypeAll = []LedgerUpgradeType{
	LedgerUpgradeTypeLedgerUpgradeVersion,
	LedgerUpgradeTypeLedgerUpgradeBaseFee,
	LedgerUpgradeTypeLedgerUpgradeMaxTxSetSize,
}

var ledgerUpgradeTypeMap = map[int32]string{
	1: "LedgerUpgradeTypeLedgerUpgradeVersion",
	2: "LedgerUpgradeTypeLedgerUpgradeBaseFee",
	3: "LedgerUpgradeTypeLedgerUpgradeMaxTxSetSize",
}

var ledgerUpgradeTypeShortMap = map[int32]string{
	1: "ledger_upgrade_version",
	2: "ledger_upgrade_base_fee",
	3: "ledger_upgrade_max_tx_set_size",
}

var ledgerUpgradeTypeRevMap = map[string]int32{
	"LedgerUpgradeTypeLedgerUpgradeVersion":      1,
	"LedgerUpgradeTypeLedgerUpgradeBaseFee":      2,
	"LedgerUpgradeTypeLedgerUpgradeMaxTxSetSize": 3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for LedgerUpgradeType
func (e LedgerUpgradeType) ValidEnum(v int32) bool {
	_, ok := ledgerUpgradeTypeMap[v]
	return ok
}
func (e LedgerUpgradeType) isFlag() bool {
	for i := len(LedgerUpgradeTypeAll) - 1; i >= 0; i-- {
		expected := LedgerUpgradeType(2) << uint64(len(LedgerUpgradeTypeAll)-1) >> uint64(len(LedgerUpgradeTypeAll)-i)
		if expected != LedgerUpgradeTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e LedgerUpgradeType) String() string {
	name, _ := ledgerUpgradeTypeMap[int32(e)]
	return name
}

func (e LedgerUpgradeType) ShortString() string {
	name, _ := ledgerUpgradeTypeShortMap[int32(e)]
	return name
}

func (e LedgerUpgradeType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range LedgerUpgradeTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *LedgerUpgradeType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = LedgerUpgradeType(t.Value)
	return nil
}

// LedgerUpgrade is an XDR Union defines as:
//
//   union LedgerUpgrade switch (LedgerUpgradeType type)
//    {
//    case LEDGER_UPGRADE_VERSION:
//        uint32 newLedgerVersion; // update ledgerVersion
//    case LEDGER_UPGRADE_BASE_FEE:
//        uint32 newBaseFee; // update baseFee
//    case LEDGER_UPGRADE_MAX_TX_SET_SIZE:
//        uint32 newMaxTxSetSize; // update maxTxSetSize
//    };
//
type LedgerUpgrade struct {
	Type             LedgerUpgradeType `json:"type,omitempty"`
	NewLedgerVersion *Uint32           `json:"newLedgerVersion,omitempty"`
	NewBaseFee       *Uint32           `json:"newBaseFee,omitempty"`
	NewMaxTxSetSize  *Uint32           `json:"newMaxTxSetSize,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerUpgrade) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerUpgrade
func (u LedgerUpgrade) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerUpgradeType(sw) {
	case LedgerUpgradeTypeLedgerUpgradeVersion:
		return "NewLedgerVersion", true
	case LedgerUpgradeTypeLedgerUpgradeBaseFee:
		return "NewBaseFee", true
	case LedgerUpgradeTypeLedgerUpgradeMaxTxSetSize:
		return "NewMaxTxSetSize", true
	}
	return "-", false
}

// NewLedgerUpgrade creates a new  LedgerUpgrade.
func NewLedgerUpgrade(aType LedgerUpgradeType, value interface{}) (result LedgerUpgrade, err error) {
	result.Type = aType
	switch LedgerUpgradeType(aType) {
	case LedgerUpgradeTypeLedgerUpgradeVersion:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.NewLedgerVersion = &tv
	case LedgerUpgradeTypeLedgerUpgradeBaseFee:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.NewBaseFee = &tv
	case LedgerUpgradeTypeLedgerUpgradeMaxTxSetSize:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.NewMaxTxSetSize = &tv
	}
	return
}

// MustNewLedgerVersion retrieves the NewLedgerVersion value from the union,
// panicing if the value is not set.
func (u LedgerUpgrade) MustNewLedgerVersion() Uint32 {
	val, ok := u.GetNewLedgerVersion()

	if !ok {
		panic("arm NewLedgerVersion is not set")
	}

	return val
}

// GetNewLedgerVersion retrieves the NewLedgerVersion value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerUpgrade) GetNewLedgerVersion() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "NewLedgerVersion" {
		result = *u.NewLedgerVersion
		ok = true
	}

	return
}

// MustNewBaseFee retrieves the NewBaseFee value from the union,
// panicing if the value is not set.
func (u LedgerUpgrade) MustNewBaseFee() Uint32 {
	val, ok := u.GetNewBaseFee()

	if !ok {
		panic("arm NewBaseFee is not set")
	}

	return val
}

// GetNewBaseFee retrieves the NewBaseFee value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerUpgrade) GetNewBaseFee() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "NewBaseFee" {
		result = *u.NewBaseFee
		ok = true
	}

	return
}

// MustNewMaxTxSetSize retrieves the NewMaxTxSetSize value from the union,
// panicing if the value is not set.
func (u LedgerUpgrade) MustNewMaxTxSetSize() Uint32 {
	val, ok := u.GetNewMaxTxSetSize()

	if !ok {
		panic("arm NewMaxTxSetSize is not set")
	}

	return val
}

// GetNewMaxTxSetSize retrieves the NewMaxTxSetSize value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerUpgrade) GetNewMaxTxSetSize() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "NewMaxTxSetSize" {
		result = *u.NewMaxTxSetSize
		ok = true
	}

	return
}

// LedgerKeyAccountExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type LedgerKeyAccountExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyAccountExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyAccountExt
func (u LedgerKeyAccountExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyAccountExt creates a new  LedgerKeyAccountExt.
func NewLedgerKeyAccountExt(v LedgerVersion, value interface{}) (result LedgerKeyAccountExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyAccount is an XDR NestedStruct defines as:
//
//   struct
//        {
//            AccountID accountID;
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        }
//
type LedgerKeyAccount struct {
	AccountId AccountId           `json:"accountID,omitempty"`
	Ext       LedgerKeyAccountExt `json:"ext,omitempty"`
}

// LedgerKeyTrustLineExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type LedgerKeyTrustLineExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyTrustLineExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyTrustLineExt
func (u LedgerKeyTrustLineExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyTrustLineExt creates a new  LedgerKeyTrustLineExt.
func NewLedgerKeyTrustLineExt(v LedgerVersion, value interface{}) (result LedgerKeyTrustLineExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyTrustLine is an XDR NestedStruct defines as:
//
//   struct
//        {
//            AccountID accountID;
//            Asset asset;
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        }
//
type LedgerKeyTrustLine struct {
	AccountId AccountId             `json:"accountID,omitempty"`
	Asset     Asset                 `json:"asset,omitempty"`
	Ext       LedgerKeyTrustLineExt `json:"ext,omitempty"`
}

// LedgerKeyOfferExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type LedgerKeyOfferExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyOfferExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyOfferExt
func (u LedgerKeyOfferExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyOfferExt creates a new  LedgerKeyOfferExt.
func NewLedgerKeyOfferExt(v LedgerVersion, value interface{}) (result LedgerKeyOfferExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyOffer is an XDR NestedStruct defines as:
//
//   struct
//        {
//            AccountID sellerID;
//            uint64 offerID;
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        }
//
type LedgerKeyOffer struct {
	SellerId AccountId         `json:"sellerID,omitempty"`
	OfferId  Uint64            `json:"offerID,omitempty"`
	Ext      LedgerKeyOfferExt `json:"ext,omitempty"`
}

// LedgerKeyLotExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type LedgerKeyLotExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyLotExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyLotExt
func (u LedgerKeyLotExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyLotExt creates a new  LedgerKeyLotExt.
func NewLedgerKeyLotExt(v LedgerVersion, value interface{}) (result LedgerKeyLotExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyLot is an XDR NestedStruct defines as:
//
//   struct
//        {
//            uint64 lotID;
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        }
//
type LedgerKeyLot struct {
	LotId Uint64          `json:"lotID,omitempty"`
	Ext   LedgerKeyLotExt `json:"ext,omitempty"`
}

// LedgerKeyParticipantExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type LedgerKeyParticipantExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyParticipantExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyParticipantExt
func (u LedgerKeyParticipantExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyParticipantExt creates a new  LedgerKeyParticipantExt.
func NewLedgerKeyParticipantExt(v LedgerVersion, value interface{}) (result LedgerKeyParticipantExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyParticipant is an XDR NestedStruct defines as:
//
//   struct
//        {
//            uint64 lotID;
//            AccountID accountID;
//
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        }
//
type LedgerKeyParticipant struct {
	LotId     Uint64                  `json:"lotID,omitempty"`
	AccountId AccountId               `json:"accountID,omitempty"`
	Ext       LedgerKeyParticipantExt `json:"ext,omitempty"`
}

// LedgerKeyMessageExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeyMessageExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyMessageExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyMessageExt
func (u LedgerKeyMessageExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyMessageExt creates a new  LedgerKeyMessageExt.
func NewLedgerKeyMessageExt(v LedgerVersion, value interface{}) (result LedgerKeyMessageExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyMessage is an XDR NestedStruct defines as:
//
//   struct
//        {
//        uint64 messageID;
//        union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        }
//
type LedgerKeyMessage struct {
	MessageId Uint64              `json:"messageID,omitempty"`
	Ext       LedgerKeyMessageExt `json:"ext,omitempty"`
}

// LedgerKeyMigrationExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//
type LedgerKeyMigrationExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyMigrationExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyMigrationExt
func (u LedgerKeyMigrationExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyMigrationExt creates a new  LedgerKeyMigrationExt.
func NewLedgerKeyMigrationExt(v LedgerVersion, value interface{}) (result LedgerKeyMigrationExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyMigration is an XDR NestedStruct defines as:
//
//   struct
//        {
//            MigrationVersion migrationVersion;
//            union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//                ext;
//        }
//
type LedgerKeyMigration struct {
	MigrationVersion MigrationVersion      `json:"migrationVersion,omitempty"`
	Ext              LedgerKeyMigrationExt `json:"ext,omitempty"`
}

// LedgerKeyReviewableRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//
type LedgerKeyReviewableRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyReviewableRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyReviewableRequestExt
func (u LedgerKeyReviewableRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyReviewableRequestExt creates a new  LedgerKeyReviewableRequestExt.
func NewLedgerKeyReviewableRequestExt(v LedgerVersion, value interface{}) (result LedgerKeyReviewableRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyReviewableRequest is an XDR NestedStruct defines as:
//
//   struct
//        {
//            uint64 requestID;
//            union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//            ext;
//        }
//
type LedgerKeyReviewableRequest struct {
	RequestId Uint64                        `json:"requestID,omitempty"`
	Ext       LedgerKeyReviewableRequestExt `json:"ext,omitempty"`
}

// LedgerKeyReferenceExt is an XDR NestedUnion defines as:
//
//   union switch(LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//
type LedgerKeyReferenceExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyReferenceExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyReferenceExt
func (u LedgerKeyReferenceExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyReferenceExt creates a new  LedgerKeyReferenceExt.
func NewLedgerKeyReferenceExt(v LedgerVersion, value interface{}) (result LedgerKeyReferenceExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyReference is an XDR NestedStruct defines as:
//
//   struct
//        {
//            AccountID sender;
//            string64 reference;
//            union switch(LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//            ext;
//        }
//
type LedgerKeyReference struct {
	Sender    AccountId             `json:"sender,omitempty"`
	Reference String64              `json:"reference,omitempty"`
	Ext       LedgerKeyReferenceExt `json:"ext,omitempty"`
}

// LedgerKey is an XDR Union defines as:
//
//   union LedgerKey switch (LedgerEntryType type)
//    {
//    case ACCOUNT:
//        struct
//        {
//            AccountID accountID;
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        } account;
//
//    case TRUSTLINE:
//        struct
//        {
//            AccountID accountID;
//            Asset asset;
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        } trustLine;
//
//    case OFFER:
//        struct
//        {
//            AccountID sellerID;
//            uint64 offerID;
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        } offer;
//    case LOT:
//        struct
//        {
//            uint64 lotID;
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        } lot;
//    case PARTICIPANT:
//        struct
//        {
//            uint64 lotID;
//            AccountID accountID;
//
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        } participant;
//    case MESSAGE:
//        struct
//        {
//        uint64 messageID;
//        union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } message;
//    case MIGRATION:
//        struct
//        {
//            MigrationVersion migrationVersion;
//            union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//                ext;
//        } migration;
//    case REVIEWABLE_REQUEST:
//        struct
//        {
//            uint64 requestID;
//            union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//            ext;
//        } reviewableRequest;
//    case REFERENCE:
//        struct
//        {
//            AccountID sender;
//            string64 reference;
//            union switch(LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//            ext;
//        } reference;
//    };
//
type LedgerKey struct {
	Type              LedgerEntryType             `json:"type,omitempty"`
	Account           *LedgerKeyAccount           `json:"account,omitempty"`
	TrustLine         *LedgerKeyTrustLine         `json:"trustLine,omitempty"`
	Offer             *LedgerKeyOffer             `json:"offer,omitempty"`
	Lot               *LedgerKeyLot               `json:"lot,omitempty"`
	Participant       *LedgerKeyParticipant       `json:"participant,omitempty"`
	Message           *LedgerKeyMessage           `json:"message,omitempty"`
	Migration         *LedgerKeyMigration         `json:"migration,omitempty"`
	ReviewableRequest *LedgerKeyReviewableRequest `json:"reviewableRequest,omitempty"`
	Reference         *LedgerKeyReference         `json:"reference,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKey) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKey
func (u LedgerKey) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerEntryType(sw) {
	case LedgerEntryTypeAccount:
		return "Account", true
	case LedgerEntryTypeTrustline:
		return "TrustLine", true
	case LedgerEntryTypeOffer:
		return "Offer", true
	case LedgerEntryTypeLot:
		return "Lot", true
	case LedgerEntryTypeParticipant:
		return "Participant", true
	case LedgerEntryTypeMessage:
		return "Message", true
	case LedgerEntryTypeMigration:
		return "Migration", true
	case LedgerEntryTypeReviewableRequest:
		return "ReviewableRequest", true
	case LedgerEntryTypeReference:
		return "Reference", true
	}
	return "-", false
}

// NewLedgerKey creates a new  LedgerKey.
func NewLedgerKey(aType LedgerEntryType, value interface{}) (result LedgerKey, err error) {
	result.Type = aType
	switch LedgerEntryType(aType) {
	case LedgerEntryTypeAccount:
		tv, ok := value.(LedgerKeyAccount)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyAccount")
			return
		}
		result.Account = &tv
	case LedgerEntryTypeTrustline:
		tv, ok := value.(LedgerKeyTrustLine)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyTrustLine")
			return
		}
		result.TrustLine = &tv
	case LedgerEntryTypeOffer:
		tv, ok := value.(LedgerKeyOffer)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyOffer")
			return
		}
		result.Offer = &tv
	case LedgerEntryTypeLot:
		tv, ok := value.(LedgerKeyLot)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyLot")
			return
		}
		result.Lot = &tv
	case LedgerEntryTypeParticipant:
		tv, ok := value.(LedgerKeyParticipant)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyParticipant")
			return
		}
		result.Participant = &tv
	case LedgerEntryTypeMessage:
		tv, ok := value.(LedgerKeyMessage)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyMessage")
			return
		}
		result.Message = &tv
	case LedgerEntryTypeMigration:
		tv, ok := value.(LedgerKeyMigration)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyMigration")
			return
		}
		result.Migration = &tv
	case LedgerEntryTypeReviewableRequest:
		tv, ok := value.(LedgerKeyReviewableRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyReviewableRequest")
			return
		}
		result.ReviewableRequest = &tv
	case LedgerEntryTypeReference:
		tv, ok := value.(LedgerKeyReference)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyReference")
			return
		}
		result.Reference = &tv
	}
	return
}

// MustAccount retrieves the Account value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustAccount() LedgerKeyAccount {
	val, ok := u.GetAccount()

	if !ok {
		panic("arm Account is not set")
	}

	return val
}

// GetAccount retrieves the Account value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetAccount() (result LedgerKeyAccount, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Account" {
		result = *u.Account
		ok = true
	}

	return
}

// MustTrustLine retrieves the TrustLine value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustTrustLine() LedgerKeyTrustLine {
	val, ok := u.GetTrustLine()

	if !ok {
		panic("arm TrustLine is not set")
	}

	return val
}

// GetTrustLine retrieves the TrustLine value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetTrustLine() (result LedgerKeyTrustLine, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "TrustLine" {
		result = *u.TrustLine
		ok = true
	}

	return
}

// MustOffer retrieves the Offer value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustOffer() LedgerKeyOffer {
	val, ok := u.GetOffer()

	if !ok {
		panic("arm Offer is not set")
	}

	return val
}

// GetOffer retrieves the Offer value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetOffer() (result LedgerKeyOffer, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Offer" {
		result = *u.Offer
		ok = true
	}

	return
}

// MustLot retrieves the Lot value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustLot() LedgerKeyLot {
	val, ok := u.GetLot()

	if !ok {
		panic("arm Lot is not set")
	}

	return val
}

// GetLot retrieves the Lot value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetLot() (result LedgerKeyLot, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Lot" {
		result = *u.Lot
		ok = true
	}

	return
}

// MustParticipant retrieves the Participant value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustParticipant() LedgerKeyParticipant {
	val, ok := u.GetParticipant()

	if !ok {
		panic("arm Participant is not set")
	}

	return val
}

// GetParticipant retrieves the Participant value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetParticipant() (result LedgerKeyParticipant, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Participant" {
		result = *u.Participant
		ok = true
	}

	return
}

// MustMessage retrieves the Message value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustMessage() LedgerKeyMessage {
	val, ok := u.GetMessage()

	if !ok {
		panic("arm Message is not set")
	}

	return val
}

// GetMessage retrieves the Message value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetMessage() (result LedgerKeyMessage, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Message" {
		result = *u.Message
		ok = true
	}

	return
}

// MustMigration retrieves the Migration value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustMigration() LedgerKeyMigration {
	val, ok := u.GetMigration()

	if !ok {
		panic("arm Migration is not set")
	}

	return val
}

// GetMigration retrieves the Migration value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetMigration() (result LedgerKeyMigration, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Migration" {
		result = *u.Migration
		ok = true
	}

	return
}

// MustReviewableRequest retrieves the ReviewableRequest value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustReviewableRequest() LedgerKeyReviewableRequest {
	val, ok := u.GetReviewableRequest()

	if !ok {
		panic("arm ReviewableRequest is not set")
	}

	return val
}

// GetReviewableRequest retrieves the ReviewableRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetReviewableRequest() (result LedgerKeyReviewableRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ReviewableRequest" {
		result = *u.ReviewableRequest
		ok = true
	}

	return
}

// MustReference retrieves the Reference value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustReference() LedgerKeyReference {
	val, ok := u.GetReference()

	if !ok {
		panic("arm Reference is not set")
	}

	return val
}

// GetReference retrieves the Reference value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetReference() (result LedgerKeyReference, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Reference" {
		result = *u.Reference
		ok = true
	}

	return
}

// BucketEntryType is an XDR Enum defines as:
//
//   enum BucketEntryType
//    {
//        LIVEENTRY = 0,
//        DEADENTRY = 1
//    };
//
type BucketEntryType int32

const (
	BucketEntryTypeLiveentry BucketEntryType = 0
	BucketEntryTypeDeadentry BucketEntryType = 1
)

var BucketEntryTypeAll = []BucketEntryType{
	BucketEntryTypeLiveentry,
	BucketEntryTypeDeadentry,
}

var bucketEntryTypeMap = map[int32]string{
	0: "BucketEntryTypeLiveentry",
	1: "BucketEntryTypeDeadentry",
}

var bucketEntryTypeShortMap = map[int32]string{
	0: "liveentry",
	1: "deadentry",
}

var bucketEntryTypeRevMap = map[string]int32{
	"BucketEntryTypeLiveentry": 0,
	"BucketEntryTypeDeadentry": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for BucketEntryType
func (e BucketEntryType) ValidEnum(v int32) bool {
	_, ok := bucketEntryTypeMap[v]
	return ok
}
func (e BucketEntryType) isFlag() bool {
	for i := len(BucketEntryTypeAll) - 1; i >= 0; i-- {
		expected := BucketEntryType(2) << uint64(len(BucketEntryTypeAll)-1) >> uint64(len(BucketEntryTypeAll)-i)
		if expected != BucketEntryTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e BucketEntryType) String() string {
	name, _ := bucketEntryTypeMap[int32(e)]
	return name
}

func (e BucketEntryType) ShortString() string {
	name, _ := bucketEntryTypeShortMap[int32(e)]
	return name
}

func (e BucketEntryType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range BucketEntryTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *BucketEntryType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = BucketEntryType(t.Value)
	return nil
}

// BucketEntry is an XDR Union defines as:
//
//   union BucketEntry switch (BucketEntryType type)
//    {
//    case LIVEENTRY:
//        LedgerEntry liveEntry;
//
//    case DEADENTRY:
//        LedgerKey deadEntry;
//    };
//
type BucketEntry struct {
	Type      BucketEntryType `json:"type,omitempty"`
	LiveEntry *LedgerEntry    `json:"liveEntry,omitempty"`
	DeadEntry *LedgerKey      `json:"deadEntry,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u BucketEntry) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of BucketEntry
func (u BucketEntry) ArmForSwitch(sw int32) (string, bool) {
	switch BucketEntryType(sw) {
	case BucketEntryTypeLiveentry:
		return "LiveEntry", true
	case BucketEntryTypeDeadentry:
		return "DeadEntry", true
	}
	return "-", false
}

// NewBucketEntry creates a new  BucketEntry.
func NewBucketEntry(aType BucketEntryType, value interface{}) (result BucketEntry, err error) {
	result.Type = aType
	switch BucketEntryType(aType) {
	case BucketEntryTypeLiveentry:
		tv, ok := value.(LedgerEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntry")
			return
		}
		result.LiveEntry = &tv
	case BucketEntryTypeDeadentry:
		tv, ok := value.(LedgerKey)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKey")
			return
		}
		result.DeadEntry = &tv
	}
	return
}

// MustLiveEntry retrieves the LiveEntry value from the union,
// panicing if the value is not set.
func (u BucketEntry) MustLiveEntry() LedgerEntry {
	val, ok := u.GetLiveEntry()

	if !ok {
		panic("arm LiveEntry is not set")
	}

	return val
}

// GetLiveEntry retrieves the LiveEntry value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u BucketEntry) GetLiveEntry() (result LedgerEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "LiveEntry" {
		result = *u.LiveEntry
		ok = true
	}

	return
}

// MustDeadEntry retrieves the DeadEntry value from the union,
// panicing if the value is not set.
func (u BucketEntry) MustDeadEntry() LedgerKey {
	val, ok := u.GetDeadEntry()

	if !ok {
		panic("arm DeadEntry is not set")
	}

	return val
}

// GetDeadEntry retrieves the DeadEntry value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u BucketEntry) GetDeadEntry() (result LedgerKey, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "DeadEntry" {
		result = *u.DeadEntry
		ok = true
	}

	return
}

// MaxTxPerLedger is an XDR Const defines as:
//
//   const MAX_TX_PER_LEDGER = 5000;
//
const MaxTxPerLedger = 5000

// TransactionSetExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type TransactionSetExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionSetExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionSetExt
func (u TransactionSetExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTransactionSetExt creates a new  TransactionSetExt.
func NewTransactionSetExt(v LedgerVersion, value interface{}) (result TransactionSetExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TransactionSet is an XDR Struct defines as:
//
//   struct TransactionSet
//    {
//        Hash previousLedgerHash;
//        TransactionEnvelope txs<MAX_TX_PER_LEDGER>;
//        union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//    };
//
type TransactionSet struct {
	PreviousLedgerHash Hash                  `json:"previousLedgerHash,omitempty"`
	Txs                []TransactionEnvelope `json:"txs,omitempty" xdrmaxsize:"5000"`
	Ext                TransactionSetExt     `json:"ext,omitempty"`
}

// TransactionResultPairExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type TransactionResultPairExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionResultPairExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionResultPairExt
func (u TransactionResultPairExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTransactionResultPairExt creates a new  TransactionResultPairExt.
func NewTransactionResultPairExt(v LedgerVersion, value interface{}) (result TransactionResultPairExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TransactionResultPair is an XDR Struct defines as:
//
//   struct TransactionResultPair
//    {
//        Hash transactionHash;
//        TransactionResult result; // result for the transaction
//        union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//    };
//
type TransactionResultPair struct {
	TransactionHash Hash                     `json:"transactionHash,omitempty"`
	Result          TransactionResult        `json:"result,omitempty"`
	Ext             TransactionResultPairExt `json:"ext,omitempty"`
}

// TransactionResultSetExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type TransactionResultSetExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionResultSetExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionResultSetExt
func (u TransactionResultSetExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTransactionResultSetExt creates a new  TransactionResultSetExt.
func NewTransactionResultSetExt(v LedgerVersion, value interface{}) (result TransactionResultSetExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TransactionResultSet is an XDR Struct defines as:
//
//   struct TransactionResultSet
//    {
//        TransactionResultPair results<MAX_TX_PER_LEDGER>;
//        union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//    };
//
type TransactionResultSet struct {
	Results []TransactionResultPair `json:"results,omitempty" xdrmaxsize:"5000"`
	Ext     TransactionResultSetExt `json:"ext,omitempty"`
}

// TransactionHistoryEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type TransactionHistoryEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionHistoryEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionHistoryEntryExt
func (u TransactionHistoryEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTransactionHistoryEntryExt creates a new  TransactionHistoryEntryExt.
func NewTransactionHistoryEntryExt(v LedgerVersion, value interface{}) (result TransactionHistoryEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TransactionHistoryEntry is an XDR Struct defines as:
//
//   struct TransactionHistoryEntry
//    {
//        uint32 ledgerSeq;
//        TransactionSet txSet;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type TransactionHistoryEntry struct {
	LedgerSeq Uint32                     `json:"ledgerSeq,omitempty"`
	TxSet     TransactionSet             `json:"txSet,omitempty"`
	Ext       TransactionHistoryEntryExt `json:"ext,omitempty"`
}

// TransactionHistoryResultEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type TransactionHistoryResultEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionHistoryResultEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionHistoryResultEntryExt
func (u TransactionHistoryResultEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTransactionHistoryResultEntryExt creates a new  TransactionHistoryResultEntryExt.
func NewTransactionHistoryResultEntryExt(v LedgerVersion, value interface{}) (result TransactionHistoryResultEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TransactionHistoryResultEntry is an XDR Struct defines as:
//
//   struct TransactionHistoryResultEntry
//    {
//        uint32 ledgerSeq;
//        TransactionResultSet txResultSet;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type TransactionHistoryResultEntry struct {
	LedgerSeq   Uint32                           `json:"ledgerSeq,omitempty"`
	TxResultSet TransactionResultSet             `json:"txResultSet,omitempty"`
	Ext         TransactionHistoryResultEntryExt `json:"ext,omitempty"`
}

// LedgerHeaderHistoryEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type LedgerHeaderHistoryEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerHeaderHistoryEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerHeaderHistoryEntryExt
func (u LedgerHeaderHistoryEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerHeaderHistoryEntryExt creates a new  LedgerHeaderHistoryEntryExt.
func NewLedgerHeaderHistoryEntryExt(v LedgerVersion, value interface{}) (result LedgerHeaderHistoryEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerHeaderHistoryEntry is an XDR Struct defines as:
//
//   struct LedgerHeaderHistoryEntry
//    {
//        Hash hash;
//        LedgerHeader header;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type LedgerHeaderHistoryEntry struct {
	Hash   Hash                        `json:"hash,omitempty"`
	Header LedgerHeader                `json:"header,omitempty"`
	Ext    LedgerHeaderHistoryEntryExt `json:"ext,omitempty"`
}

// LedgerScpMessages is an XDR Struct defines as:
//
//   struct LedgerSCPMessages
//    {
//        uint32 ledgerSeq;
//        SCPEnvelope messages<>;
//    };
//
type LedgerScpMessages struct {
	LedgerSeq Uint32        `json:"ledgerSeq,omitempty"`
	Messages  []ScpEnvelope `json:"messages,omitempty"`
}

// ScpHistoryEntryV0 is an XDR Struct defines as:
//
//   struct SCPHistoryEntryV0
//    {
//        SCPQuorumSet quorumSets<>; // additional quorum sets used by ledgerMessages
//        LedgerSCPMessages ledgerMessages;
//    };
//
type ScpHistoryEntryV0 struct {
	QuorumSets     []ScpQuorumSet    `json:"quorumSets,omitempty"`
	LedgerMessages LedgerScpMessages `json:"ledgerMessages,omitempty"`
}

// ScpHistoryEntry is an XDR Union defines as:
//
//   union SCPHistoryEntry switch (LedgerVersion v)
//    {
//    case EMPTY_VERSION:
//        SCPHistoryEntryV0 v0;
//    };
//
type ScpHistoryEntry struct {
	V  LedgerVersion      `json:"v,omitempty"`
	V0 *ScpHistoryEntryV0 `json:"v0,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ScpHistoryEntry) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ScpHistoryEntry
func (u ScpHistoryEntry) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "V0", true
	}
	return "-", false
}

// NewScpHistoryEntry creates a new  ScpHistoryEntry.
func NewScpHistoryEntry(v LedgerVersion, value interface{}) (result ScpHistoryEntry, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		tv, ok := value.(ScpHistoryEntryV0)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpHistoryEntryV0")
			return
		}
		result.V0 = &tv
	}
	return
}

// MustV0 retrieves the V0 value from the union,
// panicing if the value is not set.
func (u ScpHistoryEntry) MustV0() ScpHistoryEntryV0 {
	val, ok := u.GetV0()

	if !ok {
		panic("arm V0 is not set")
	}

	return val
}

// GetV0 retrieves the V0 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpHistoryEntry) GetV0() (result ScpHistoryEntryV0, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "V0" {
		result = *u.V0
		ok = true
	}

	return
}

// LedgerEntryChangeType is an XDR Enum defines as:
//
//   enum LedgerEntryChangeType
//    {
//        LEDGER_ENTRY_CREATED = 0, // entry was added to the ledger
//        LEDGER_ENTRY_UPDATED = 1, // entry was modified in the ledger
//        LEDGER_ENTRY_REMOVED = 2, // entry was removed from the ledger
//        LEDGER_ENTRY_STATE = 3    // value of the entry
//    };
//
type LedgerEntryChangeType int32

const (
	LedgerEntryChangeTypeLedgerEntryCreated LedgerEntryChangeType = 0
	LedgerEntryChangeTypeLedgerEntryUpdated LedgerEntryChangeType = 1
	LedgerEntryChangeTypeLedgerEntryRemoved LedgerEntryChangeType = 2
	LedgerEntryChangeTypeLedgerEntryState   LedgerEntryChangeType = 3
)

var LedgerEntryChangeTypeAll = []LedgerEntryChangeType{
	LedgerEntryChangeTypeLedgerEntryCreated,
	LedgerEntryChangeTypeLedgerEntryUpdated,
	LedgerEntryChangeTypeLedgerEntryRemoved,
	LedgerEntryChangeTypeLedgerEntryState,
}

var ledgerEntryChangeTypeMap = map[int32]string{
	0: "LedgerEntryChangeTypeLedgerEntryCreated",
	1: "LedgerEntryChangeTypeLedgerEntryUpdated",
	2: "LedgerEntryChangeTypeLedgerEntryRemoved",
	3: "LedgerEntryChangeTypeLedgerEntryState",
}

var ledgerEntryChangeTypeShortMap = map[int32]string{
	0: "ledger_entry_created",
	1: "ledger_entry_updated",
	2: "ledger_entry_removed",
	3: "ledger_entry_state",
}

var ledgerEntryChangeTypeRevMap = map[string]int32{
	"LedgerEntryChangeTypeLedgerEntryCreated": 0,
	"LedgerEntryChangeTypeLedgerEntryUpdated": 1,
	"LedgerEntryChangeTypeLedgerEntryRemoved": 2,
	"LedgerEntryChangeTypeLedgerEntryState":   3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for LedgerEntryChangeType
func (e LedgerEntryChangeType) ValidEnum(v int32) bool {
	_, ok := ledgerEntryChangeTypeMap[v]
	return ok
}
func (e LedgerEntryChangeType) isFlag() bool {
	for i := len(LedgerEntryChangeTypeAll) - 1; i >= 0; i-- {
		expected := LedgerEntryChangeType(2) << uint64(len(LedgerEntryChangeTypeAll)-1) >> uint64(len(LedgerEntryChangeTypeAll)-i)
		if expected != LedgerEntryChangeTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e LedgerEntryChangeType) String() string {
	name, _ := ledgerEntryChangeTypeMap[int32(e)]
	return name
}

func (e LedgerEntryChangeType) ShortString() string {
	name, _ := ledgerEntryChangeTypeShortMap[int32(e)]
	return name
}

func (e LedgerEntryChangeType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range LedgerEntryChangeTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *LedgerEntryChangeType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = LedgerEntryChangeType(t.Value)
	return nil
}

// LedgerEntryChange is an XDR Union defines as:
//
//   union LedgerEntryChange switch (LedgerEntryChangeType type)
//    {
//    case LEDGER_ENTRY_CREATED:
//        LedgerEntry created;
//    case LEDGER_ENTRY_UPDATED:
//        LedgerEntry updated;
//    case LEDGER_ENTRY_REMOVED:
//        LedgerKey removed;
//    case LEDGER_ENTRY_STATE:
//        LedgerEntry state;
//    };
//
type LedgerEntryChange struct {
	Type    LedgerEntryChangeType `json:"type,omitempty"`
	Created *LedgerEntry          `json:"created,omitempty"`
	Updated *LedgerEntry          `json:"updated,omitempty"`
	Removed *LedgerKey            `json:"removed,omitempty"`
	State   *LedgerEntry          `json:"state,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerEntryChange) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerEntryChange
func (u LedgerEntryChange) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerEntryChangeType(sw) {
	case LedgerEntryChangeTypeLedgerEntryCreated:
		return "Created", true
	case LedgerEntryChangeTypeLedgerEntryUpdated:
		return "Updated", true
	case LedgerEntryChangeTypeLedgerEntryRemoved:
		return "Removed", true
	case LedgerEntryChangeTypeLedgerEntryState:
		return "State", true
	}
	return "-", false
}

// NewLedgerEntryChange creates a new  LedgerEntryChange.
func NewLedgerEntryChange(aType LedgerEntryChangeType, value interface{}) (result LedgerEntryChange, err error) {
	result.Type = aType
	switch LedgerEntryChangeType(aType) {
	case LedgerEntryChangeTypeLedgerEntryCreated:
		tv, ok := value.(LedgerEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntry")
			return
		}
		result.Created = &tv
	case LedgerEntryChangeTypeLedgerEntryUpdated:
		tv, ok := value.(LedgerEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntry")
			return
		}
		result.Updated = &tv
	case LedgerEntryChangeTypeLedgerEntryRemoved:
		tv, ok := value.(LedgerKey)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKey")
			return
		}
		result.Removed = &tv
	case LedgerEntryChangeTypeLedgerEntryState:
		tv, ok := value.(LedgerEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntry")
			return
		}
		result.State = &tv
	}
	return
}

// MustCreated retrieves the Created value from the union,
// panicing if the value is not set.
func (u LedgerEntryChange) MustCreated() LedgerEntry {
	val, ok := u.GetCreated()

	if !ok {
		panic("arm Created is not set")
	}

	return val
}

// GetCreated retrieves the Created value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryChange) GetCreated() (result LedgerEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Created" {
		result = *u.Created
		ok = true
	}

	return
}

// MustUpdated retrieves the Updated value from the union,
// panicing if the value is not set.
func (u LedgerEntryChange) MustUpdated() LedgerEntry {
	val, ok := u.GetUpdated()

	if !ok {
		panic("arm Updated is not set")
	}

	return val
}

// GetUpdated retrieves the Updated value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryChange) GetUpdated() (result LedgerEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Updated" {
		result = *u.Updated
		ok = true
	}

	return
}

// MustRemoved retrieves the Removed value from the union,
// panicing if the value is not set.
func (u LedgerEntryChange) MustRemoved() LedgerKey {
	val, ok := u.GetRemoved()

	if !ok {
		panic("arm Removed is not set")
	}

	return val
}

// GetRemoved retrieves the Removed value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryChange) GetRemoved() (result LedgerKey, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Removed" {
		result = *u.Removed
		ok = true
	}

	return
}

// MustState retrieves the State value from the union,
// panicing if the value is not set.
func (u LedgerEntryChange) MustState() LedgerEntry {
	val, ok := u.GetState()

	if !ok {
		panic("arm State is not set")
	}

	return val
}

// GetState retrieves the State value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryChange) GetState() (result LedgerEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "State" {
		result = *u.State
		ok = true
	}

	return
}

// LedgerEntryChanges is an XDR Typedef defines as:
//
//   typedef LedgerEntryChange LedgerEntryChanges<>;
//
type LedgerEntryChanges []LedgerEntryChange

// OperationMeta is an XDR Struct defines as:
//
//   struct OperationMeta
//    {
//        LedgerEntryChanges changes;
//    };
//
type OperationMeta struct {
	Changes LedgerEntryChanges `json:"changes,omitempty"`
}

// TransactionMeta is an XDR Union defines as:
//
//   union TransactionMeta switch (LedgerVersion v)
//    {
//    case EMPTY_VERSION:
//        OperationMeta operations<>;
//    };
//
type TransactionMeta struct {
	V          LedgerVersion    `json:"v,omitempty"`
	Operations *[]OperationMeta `json:"operations,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionMeta) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionMeta
func (u TransactionMeta) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "Operations", true
	}
	return "-", false
}

// NewTransactionMeta creates a new  TransactionMeta.
func NewTransactionMeta(v LedgerVersion, value interface{}) (result TransactionMeta, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		tv, ok := value.([]OperationMeta)
		if !ok {
			err = fmt.Errorf("invalid value, must be []OperationMeta")
			return
		}
		result.Operations = &tv
	}
	return
}

// MustOperations retrieves the Operations value from the union,
// panicing if the value is not set.
func (u TransactionMeta) MustOperations() []OperationMeta {
	val, ok := u.GetOperations()

	if !ok {
		panic("arm Operations is not set")
	}

	return val
}

// GetOperations retrieves the Operations value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionMeta) GetOperations() (result []OperationMeta, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "Operations" {
		result = *u.Operations
		ok = true
	}

	return
}

// AllowTrustOpAsset is an XDR NestedUnion defines as:
//
//   union switch (AssetType type)
//            {
//            // ASSET_TYPE_NATIVE is not allowed
//            case ASSET_TYPE_CREDIT_ALPHANUM4:
//                opaque assetCode4[4];
//
//            case ASSET_TYPE_CREDIT_ALPHANUM12:
//                opaque assetCode12[12];
//
//                // add other asset types here in the future
//            }
//
type AllowTrustOpAsset struct {
	Type        AssetType `json:"type,omitempty"`
	AssetCode4  *[4]byte  `json:"assetCode4,omitempty"`
	AssetCode12 *[12]byte `json:"assetCode12,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AllowTrustOpAsset) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AllowTrustOpAsset
func (u AllowTrustOpAsset) ArmForSwitch(sw int32) (string, bool) {
	switch AssetType(sw) {
	case AssetTypeAssetTypeCreditAlphanum4:
		return "AssetCode4", true
	case AssetTypeAssetTypeCreditAlphanum12:
		return "AssetCode12", true
	}
	return "-", false
}

// NewAllowTrustOpAsset creates a new  AllowTrustOpAsset.
func NewAllowTrustOpAsset(aType AssetType, value interface{}) (result AllowTrustOpAsset, err error) {
	result.Type = aType
	switch AssetType(aType) {
	case AssetTypeAssetTypeCreditAlphanum4:
		tv, ok := value.([4]byte)
		if !ok {
			err = fmt.Errorf("invalid value, must be [4]byte")
			return
		}
		result.AssetCode4 = &tv
	case AssetTypeAssetTypeCreditAlphanum12:
		tv, ok := value.([12]byte)
		if !ok {
			err = fmt.Errorf("invalid value, must be [12]byte")
			return
		}
		result.AssetCode12 = &tv
	}
	return
}

// MustAssetCode4 retrieves the AssetCode4 value from the union,
// panicing if the value is not set.
func (u AllowTrustOpAsset) MustAssetCode4() [4]byte {
	val, ok := u.GetAssetCode4()

	if !ok {
		panic("arm AssetCode4 is not set")
	}

	return val
}

// GetAssetCode4 retrieves the AssetCode4 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AllowTrustOpAsset) GetAssetCode4() (result [4]byte, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AssetCode4" {
		result = *u.AssetCode4
		ok = true
	}

	return
}

// MustAssetCode12 retrieves the AssetCode12 value from the union,
// panicing if the value is not set.
func (u AllowTrustOpAsset) MustAssetCode12() [12]byte {
	val, ok := u.GetAssetCode12()

	if !ok {
		panic("arm AssetCode12 is not set")
	}

	return val
}

// GetAssetCode12 retrieves the AssetCode12 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AllowTrustOpAsset) GetAssetCode12() (result [12]byte, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AssetCode12" {
		result = *u.AssetCode12
		ok = true
	}

	return
}

// AllowTrustOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type AllowTrustOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AllowTrustOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AllowTrustOpExt
func (u AllowTrustOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAllowTrustOpExt creates a new  AllowTrustOpExt.
func NewAllowTrustOpExt(v LedgerVersion, value interface{}) (result AllowTrustOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AllowTrustOp is an XDR Struct defines as:
//
//   struct AllowTrustOp
//        {
//            AccountID trustor;
//            union switch (AssetType type)
//            {
//            // ASSET_TYPE_NATIVE is not allowed
//            case ASSET_TYPE_CREDIT_ALPHANUM4:
//                opaque assetCode4[4];
//
//            case ASSET_TYPE_CREDIT_ALPHANUM12:
//                opaque assetCode12[12];
//
//                // add other asset types here in the future
//            }
//            asset;
//
//            bool authorize;
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        };
//
type AllowTrustOp struct {
	Trustor   AccountId         `json:"trustor,omitempty"`
	Asset     AllowTrustOpAsset `json:"asset,omitempty"`
	Authorize bool              `json:"authorize,omitempty"`
	Ext       AllowTrustOpExt   `json:"ext,omitempty"`
}

// AllowTrustResultCode is an XDR Enum defines as:
//
//   enum AllowTrustResultCode
//        {
//            // codes considered as "success" for the operation
//            ALLOW_TRUST_SUCCESS = 0,
//            // codes considered as "failure" for the operation
//            ALLOW_TRUST_MALFORMED = -1,     // asset is not ASSET_TYPE_ALPHANUM
//            ALLOW_TRUST_NO_TRUST_LINE = -2, // trustor does not have a trustline
//                                            // source account does not require trust
//            ALLOW_TRUST_TRUST_NOT_REQUIRED = -3,
//            ALLOW_TRUST_CANT_REVOKE = -4 // source account can't revoke trust
//        };
//
type AllowTrustResultCode int32

const (
	AllowTrustResultCodeAllowTrustSuccess          AllowTrustResultCode = 0
	AllowTrustResultCodeAllowTrustMalformed        AllowTrustResultCode = -1
	AllowTrustResultCodeAllowTrustNoTrustLine      AllowTrustResultCode = -2
	AllowTrustResultCodeAllowTrustTrustNotRequired AllowTrustResultCode = -3
	AllowTrustResultCodeAllowTrustCantRevoke       AllowTrustResultCode = -4
)

var AllowTrustResultCodeAll = []AllowTrustResultCode{
	AllowTrustResultCodeAllowTrustSuccess,
	AllowTrustResultCodeAllowTrustMalformed,
	AllowTrustResultCodeAllowTrustNoTrustLine,
	AllowTrustResultCodeAllowTrustTrustNotRequired,
	AllowTrustResultCodeAllowTrustCantRevoke,
}

var allowTrustResultCodeMap = map[int32]string{
	0:  "AllowTrustResultCodeAllowTrustSuccess",
	-1: "AllowTrustResultCodeAllowTrustMalformed",
	-2: "AllowTrustResultCodeAllowTrustNoTrustLine",
	-3: "AllowTrustResultCodeAllowTrustTrustNotRequired",
	-4: "AllowTrustResultCodeAllowTrustCantRevoke",
}

var allowTrustResultCodeShortMap = map[int32]string{
	0:  "allow_trust_success",
	-1: "allow_trust_malformed",
	-2: "allow_trust_no_trust_line",
	-3: "allow_trust_trust_not_required",
	-4: "allow_trust_cant_revoke",
}

var allowTrustResultCodeRevMap = map[string]int32{
	"AllowTrustResultCodeAllowTrustSuccess":          0,
	"AllowTrustResultCodeAllowTrustMalformed":        -1,
	"AllowTrustResultCodeAllowTrustNoTrustLine":      -2,
	"AllowTrustResultCodeAllowTrustTrustNotRequired": -3,
	"AllowTrustResultCodeAllowTrustCantRevoke":       -4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for AllowTrustResultCode
func (e AllowTrustResultCode) ValidEnum(v int32) bool {
	_, ok := allowTrustResultCodeMap[v]
	return ok
}
func (e AllowTrustResultCode) isFlag() bool {
	for i := len(AllowTrustResultCodeAll) - 1; i >= 0; i-- {
		expected := AllowTrustResultCode(2) << uint64(len(AllowTrustResultCodeAll)-1) >> uint64(len(AllowTrustResultCodeAll)-i)
		if expected != AllowTrustResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e AllowTrustResultCode) String() string {
	name, _ := allowTrustResultCodeMap[int32(e)]
	return name
}

func (e AllowTrustResultCode) ShortString() string {
	name, _ := allowTrustResultCodeShortMap[int32(e)]
	return name
}

func (e AllowTrustResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range AllowTrustResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *AllowTrustResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = AllowTrustResultCode(t.Value)
	return nil
}

// AllowTrustResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                        {
//                        case EMPTY_VERSION:
//                          void;
//                        }
//
type AllowTrustResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AllowTrustResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AllowTrustResultSuccessExt
func (u AllowTrustResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAllowTrustResultSuccessExt creates a new  AllowTrustResultSuccessExt.
func NewAllowTrustResultSuccessExt(v LedgerVersion, value interface{}) (result AllowTrustResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AllowTrustResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//                        // reserved for future use
//                        union switch (LedgerVersion v)
//                        {
//                        case EMPTY_VERSION:
//                          void;
//                        }
//                        ext;
//                      }
//
type AllowTrustResultSuccess struct {
	Ext AllowTrustResultSuccessExt `json:"ext,omitempty"`
}

// AllowTrustResult is an XDR Union defines as:
//
//   union AllowTrustResult switch (AllowTrustResultCode code)
//        {
//        case ALLOW_TRUST_SUCCESS:
//            struct {
//                        // reserved for future use
//                        union switch (LedgerVersion v)
//                        {
//                        case EMPTY_VERSION:
//                          void;
//                        }
//                        ext;
//                      } success;
//        default:
//            void;
//        };
//
type AllowTrustResult struct {
	Code    AllowTrustResultCode     `json:"code,omitempty"`
	Success *AllowTrustResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AllowTrustResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AllowTrustResult
func (u AllowTrustResult) ArmForSwitch(sw int32) (string, bool) {
	switch AllowTrustResultCode(sw) {
	case AllowTrustResultCodeAllowTrustSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewAllowTrustResult creates a new  AllowTrustResult.
func NewAllowTrustResult(code AllowTrustResultCode, value interface{}) (result AllowTrustResult, err error) {
	result.Code = code
	switch AllowTrustResultCode(code) {
	case AllowTrustResultCodeAllowTrustSuccess:
		tv, ok := value.(AllowTrustResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be AllowTrustResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u AllowTrustResult) MustSuccess() AllowTrustResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AllowTrustResult) GetSuccess() (result AllowTrustResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// CancelRequestOpExt is an XDR NestedUnion defines as:
//
//   union switch(LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//
type CancelRequestOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CancelRequestOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CancelRequestOpExt
func (u CancelRequestOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCancelRequestOpExt creates a new  CancelRequestOpExt.
func NewCancelRequestOpExt(v LedgerVersion, value interface{}) (result CancelRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CancelRequestOp is an XDR Struct defines as:
//
//   struct CancelRequestOp
//        {
//            uint64 requestID;
//
//            union switch(LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            } ext;
//        };
//
type CancelRequestOp struct {
	RequestId Uint64             `json:"requestID,omitempty"`
	Ext       CancelRequestOpExt `json:"ext,omitempty"`
}

// CancelRequestResultCode is an XDR Enum defines as:
//
//   enum CancelRequestResultCode
//        {
//            CANCEL_REQUEST_SUCCESS = 0,
//
//            CANCEL_REQUEST_NOT_FOUND = -1, // request not found
//            CANCEL_REQUEST_NOT_ALLOWED = -2, // only requestor can cancel his request
//            CANCEL_REQUEST_TYPE_NOT_ALLOWED = -3 // can't cancel this type of request
//        };
//
type CancelRequestResultCode int32

const (
	CancelRequestResultCodeCancelRequestSuccess        CancelRequestResultCode = 0
	CancelRequestResultCodeCancelRequestNotFound       CancelRequestResultCode = -1
	CancelRequestResultCodeCancelRequestNotAllowed     CancelRequestResultCode = -2
	CancelRequestResultCodeCancelRequestTypeNotAllowed CancelRequestResultCode = -3
)

var CancelRequestResultCodeAll = []CancelRequestResultCode{
	CancelRequestResultCodeCancelRequestSuccess,
	CancelRequestResultCodeCancelRequestNotFound,
	CancelRequestResultCodeCancelRequestNotAllowed,
	CancelRequestResultCodeCancelRequestTypeNotAllowed,
}

var cancelRequestResultCodeMap = map[int32]string{
	0:  "CancelRequestResultCodeCancelRequestSuccess",
	-1: "CancelRequestResultCodeCancelRequestNotFound",
	-2: "CancelRequestResultCodeCancelRequestNotAllowed",
	-3: "CancelRequestResultCodeCancelRequestTypeNotAllowed",
}

var cancelRequestResultCodeShortMap = map[int32]string{
	0:  "cancel_request_success",
	-1: "cancel_request_not_found",
	-2: "cancel_request_not_allowed",
	-3: "cancel_request_type_not_allowed",
}

var cancelRequestResultCodeRevMap = map[string]int32{
	"CancelRequestResultCodeCancelRequestSuccess":        0,
	"CancelRequestResultCodeCancelRequestNotFound":       -1,
	"CancelRequestResultCodeCancelRequestNotAllowed":     -2,
	"CancelRequestResultCodeCancelRequestTypeNotAllowed": -3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CancelRequestResultCode
func (e CancelRequestResultCode) ValidEnum(v int32) bool {
	_, ok := cancelRequestResultCodeMap[v]
	return ok
}
func (e CancelRequestResultCode) isFlag() bool {
	for i := len(CancelRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := CancelRequestResultCode(2) << uint64(len(CancelRequestResultCodeAll)-1) >> uint64(len(CancelRequestResultCodeAll)-i)
		if expected != CancelRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CancelRequestResultCode) String() string {
	name, _ := cancelRequestResultCodeMap[int32(e)]
	return name
}

func (e CancelRequestResultCode) ShortString() string {
	name, _ := cancelRequestResultCodeShortMap[int32(e)]
	return name
}

func (e CancelRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range CancelRequestResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CancelRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CancelRequestResultCode(t.Value)
	return nil
}

// CancelRequestSuccessExt is an XDR NestedUnion defines as:
//
//   union switch(LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//
type CancelRequestSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CancelRequestSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CancelRequestSuccessExt
func (u CancelRequestSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCancelRequestSuccessExt creates a new  CancelRequestSuccessExt.
func NewCancelRequestSuccessExt(v LedgerVersion, value interface{}) (result CancelRequestSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CancelRequestSuccess is an XDR Struct defines as:
//
//   struct CancelRequestSuccess
//        {
//
//            // reserved for future use
//            union switch(LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            } ext;
//        };
//
type CancelRequestSuccess struct {
	Ext CancelRequestSuccessExt `json:"ext,omitempty"`
}

// CancelRequestResult is an XDR Union defines as:
//
//   union CancelRequestResult switch (CancelRequestResultCode code)
//        {
//            case CANCEL_REQUEST_SUCCESS:
//                CancelRequestSuccess success;
//            default:
//                void;
//        };
//
type CancelRequestResult struct {
	Code    CancelRequestResultCode `json:"code,omitempty"`
	Success *CancelRequestSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CancelRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CancelRequestResult
func (u CancelRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch CancelRequestResultCode(sw) {
	case CancelRequestResultCodeCancelRequestSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCancelRequestResult creates a new  CancelRequestResult.
func NewCancelRequestResult(code CancelRequestResultCode, value interface{}) (result CancelRequestResult, err error) {
	result.Code = code
	switch CancelRequestResultCode(code) {
	case CancelRequestResultCodeCancelRequestSuccess:
		tv, ok := value.(CancelRequestSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CancelRequestSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CancelRequestResult) MustSuccess() CancelRequestSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CancelRequestResult) GetSuccess() (result CancelRequestSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ChangeTrustOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type ChangeTrustOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ChangeTrustOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ChangeTrustOpExt
func (u ChangeTrustOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewChangeTrustOpExt creates a new  ChangeTrustOpExt.
func NewChangeTrustOpExt(v LedgerVersion, value interface{}) (result ChangeTrustOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ChangeTrustOp is an XDR Struct defines as:
//
//   struct ChangeTrustOp
//        {
//            Asset line;
//
//            // if limit is set to 0, deletes the trust line
//            int64 limit;
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        };
//
type ChangeTrustOp struct {
	Line  Asset            `json:"line,omitempty"`
	Limit Int64            `json:"limit,omitempty"`
	Ext   ChangeTrustOpExt `json:"ext,omitempty"`
}

// ChangeTrustResultCode is an XDR Enum defines as:
//
//   enum ChangeTrustResultCode
//        {
//            // codes considered as "success" for the operation
//            CHANGE_TRUST_SUCCESS = 0,
//            // codes considered as "failure" for the operation
//            CHANGE_TRUST_MALFORMED = -1,     // bad input
//            CHANGE_TRUST_NO_ISSUER = -2,     // could not find issuer
//            CHANGE_TRUST_INVALID_LIMIT = -3, // cannot drop limit below balance
//                                             // cannot create with a limit of 0
//            CHANGE_TRUST_LOW_RESERVE = -4 // not enough funds to create a new trust line
//        };
//
type ChangeTrustResultCode int32

const (
	ChangeTrustResultCodeChangeTrustSuccess      ChangeTrustResultCode = 0
	ChangeTrustResultCodeChangeTrustMalformed    ChangeTrustResultCode = -1
	ChangeTrustResultCodeChangeTrustNoIssuer     ChangeTrustResultCode = -2
	ChangeTrustResultCodeChangeTrustInvalidLimit ChangeTrustResultCode = -3
	ChangeTrustResultCodeChangeTrustLowReserve   ChangeTrustResultCode = -4
)

var ChangeTrustResultCodeAll = []ChangeTrustResultCode{
	ChangeTrustResultCodeChangeTrustSuccess,
	ChangeTrustResultCodeChangeTrustMalformed,
	ChangeTrustResultCodeChangeTrustNoIssuer,
	ChangeTrustResultCodeChangeTrustInvalidLimit,
	ChangeTrustResultCodeChangeTrustLowReserve,
}

var changeTrustResultCodeMap = map[int32]string{
	0:  "ChangeTrustResultCodeChangeTrustSuccess",
	-1: "ChangeTrustResultCodeChangeTrustMalformed",
	-2: "ChangeTrustResultCodeChangeTrustNoIssuer",
	-3: "ChangeTrustResultCodeChangeTrustInvalidLimit",
	-4: "ChangeTrustResultCodeChangeTrustLowReserve",
}

var changeTrustResultCodeShortMap = map[int32]string{
	0:  "change_trust_success",
	-1: "change_trust_malformed",
	-2: "change_trust_no_issuer",
	-3: "change_trust_invalid_limit",
	-4: "change_trust_low_reserve",
}

var changeTrustResultCodeRevMap = map[string]int32{
	"ChangeTrustResultCodeChangeTrustSuccess":      0,
	"ChangeTrustResultCodeChangeTrustMalformed":    -1,
	"ChangeTrustResultCodeChangeTrustNoIssuer":     -2,
	"ChangeTrustResultCodeChangeTrustInvalidLimit": -3,
	"ChangeTrustResultCodeChangeTrustLowReserve":   -4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ChangeTrustResultCode
func (e ChangeTrustResultCode) ValidEnum(v int32) bool {
	_, ok := changeTrustResultCodeMap[v]
	return ok
}
func (e ChangeTrustResultCode) isFlag() bool {
	for i := len(ChangeTrustResultCodeAll) - 1; i >= 0; i-- {
		expected := ChangeTrustResultCode(2) << uint64(len(ChangeTrustResultCodeAll)-1) >> uint64(len(ChangeTrustResultCodeAll)-i)
		if expected != ChangeTrustResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ChangeTrustResultCode) String() string {
	name, _ := changeTrustResultCodeMap[int32(e)]
	return name
}

func (e ChangeTrustResultCode) ShortString() string {
	name, _ := changeTrustResultCodeShortMap[int32(e)]
	return name
}

func (e ChangeTrustResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ChangeTrustResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ChangeTrustResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ChangeTrustResultCode(t.Value)
	return nil
}

// ChangeTrustResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                        {
//                        case EMPTY_VERSION:
//                          void;
//                        }
//
type ChangeTrustResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ChangeTrustResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ChangeTrustResultSuccessExt
func (u ChangeTrustResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewChangeTrustResultSuccessExt creates a new  ChangeTrustResultSuccessExt.
func NewChangeTrustResultSuccessExt(v LedgerVersion, value interface{}) (result ChangeTrustResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ChangeTrustResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//                        // reserved for future use
//                        union switch (LedgerVersion v)
//                        {
//                        case EMPTY_VERSION:
//                          void;
//                        }
//                        ext;
//                      }
//
type ChangeTrustResultSuccess struct {
	Ext ChangeTrustResultSuccessExt `json:"ext,omitempty"`
}

// ChangeTrustResult is an XDR Union defines as:
//
//   union ChangeTrustResult switch (ChangeTrustResultCode code)
//        {
//        case CHANGE_TRUST_SUCCESS:
//            struct {
//                        // reserved for future use
//                        union switch (LedgerVersion v)
//                        {
//                        case EMPTY_VERSION:
//                          void;
//                        }
//                        ext;
//                      } success;
//        default:
//            void;
//        };
//
type ChangeTrustResult struct {
	Code    ChangeTrustResultCode     `json:"code,omitempty"`
	Success *ChangeTrustResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ChangeTrustResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ChangeTrustResult
func (u ChangeTrustResult) ArmForSwitch(sw int32) (string, bool) {
	switch ChangeTrustResultCode(sw) {
	case ChangeTrustResultCodeChangeTrustSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewChangeTrustResult creates a new  ChangeTrustResult.
func NewChangeTrustResult(code ChangeTrustResultCode, value interface{}) (result ChangeTrustResult, err error) {
	result.Code = code
	switch ChangeTrustResultCode(code) {
	case ChangeTrustResultCodeChangeTrustSuccess:
		tv, ok := value.(ChangeTrustResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ChangeTrustResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ChangeTrustResult) MustSuccess() ChangeTrustResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ChangeTrustResult) GetSuccess() (result ChangeTrustResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// CloseAuctionAction is an XDR Enum defines as:
//
//   enum CloseAuctionAction
//        {
//            CLOSE_AUCTION_ACTION_MARK_SOLD = 1,
//            CLOSE_AUCTION_ACTION_MARK_FINISHED = 2
//        };
//
type CloseAuctionAction int32

const (
	CloseAuctionActionCloseAuctionActionMarkSold     CloseAuctionAction = 1
	CloseAuctionActionCloseAuctionActionMarkFinished CloseAuctionAction = 2
)

var CloseAuctionActionAll = []CloseAuctionAction{
	CloseAuctionActionCloseAuctionActionMarkSold,
	CloseAuctionActionCloseAuctionActionMarkFinished,
}

var closeAuctionActionMap = map[int32]string{
	1: "CloseAuctionActionCloseAuctionActionMarkSold",
	2: "CloseAuctionActionCloseAuctionActionMarkFinished",
}

var closeAuctionActionShortMap = map[int32]string{
	1: "close_auction_action_mark_sold",
	2: "close_auction_action_mark_finished",
}

var closeAuctionActionRevMap = map[string]int32{
	"CloseAuctionActionCloseAuctionActionMarkSold":     1,
	"CloseAuctionActionCloseAuctionActionMarkFinished": 2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CloseAuctionAction
func (e CloseAuctionAction) ValidEnum(v int32) bool {
	_, ok := closeAuctionActionMap[v]
	return ok
}
func (e CloseAuctionAction) isFlag() bool {
	for i := len(CloseAuctionActionAll) - 1; i >= 0; i-- {
		expected := CloseAuctionAction(2) << uint64(len(CloseAuctionActionAll)-1) >> uint64(len(CloseAuctionActionAll)-i)
		if expected != CloseAuctionActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CloseAuctionAction) String() string {
	name, _ := closeAuctionActionMap[int32(e)]
	return name
}

func (e CloseAuctionAction) ShortString() string {
	name, _ := closeAuctionActionShortMap[int32(e)]
	return name
}

func (e CloseAuctionAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range CloseAuctionActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CloseAuctionAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CloseAuctionAction(t.Value)
	return nil
}

// CloseAuctionOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type CloseAuctionOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CloseAuctionOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CloseAuctionOpExt
func (u CloseAuctionOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCloseAuctionOpExt creates a new  CloseAuctionOpExt.
func NewCloseAuctionOpExt(v LedgerVersion, value interface{}) (result CloseAuctionOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CloseAuctionOp is an XDR Struct defines as:
//
//   struct CloseAuctionOp
//        {
//            CloseAuctionAction action;
//            uint64 lotID;
//
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type CloseAuctionOp struct {
	Action CloseAuctionAction `json:"action,omitempty"`
	LotId  Uint64             `json:"lotID,omitempty"`
	Ext    CloseAuctionOpExt  `json:"ext,omitempty"`
}

// CloseAuctionResultCode is an XDR Enum defines as:
//
//   enum CloseAuctionResultCode
//        {
//            CLOSE_AUCTION_SUCCESS = 0,
//
//            CLOSE_AUCTION_NOT_FOUND = -1,        // lot not found
//            CLOSE_AUCTION_TOO_EARLY = -2,        // can't close auction before
//            CLOSE_AUCTION_NO_WINNERS = -3,       // can't mark as winner since there are no participants with state winner
//            CLOSE_AUCTION_ALREADY_CLOSED = -4,   // lot is already closed
//            CLOSE_AUCTION_HAS_WINNER = -5,       // can't mark as finished since lot has winner participant
//            CLOSE_AUCTION_NOT_ALLOWED = -6       // only platform can resolve auction
//        };
//
type CloseAuctionResultCode int32

const (
	CloseAuctionResultCodeCloseAuctionSuccess       CloseAuctionResultCode = 0
	CloseAuctionResultCodeCloseAuctionNotFound      CloseAuctionResultCode = -1
	CloseAuctionResultCodeCloseAuctionTooEarly      CloseAuctionResultCode = -2
	CloseAuctionResultCodeCloseAuctionNoWinners     CloseAuctionResultCode = -3
	CloseAuctionResultCodeCloseAuctionAlreadyClosed CloseAuctionResultCode = -4
	CloseAuctionResultCodeCloseAuctionHasWinner     CloseAuctionResultCode = -5
	CloseAuctionResultCodeCloseAuctionNotAllowed    CloseAuctionResultCode = -6
)

var CloseAuctionResultCodeAll = []CloseAuctionResultCode{
	CloseAuctionResultCodeCloseAuctionSuccess,
	CloseAuctionResultCodeCloseAuctionNotFound,
	CloseAuctionResultCodeCloseAuctionTooEarly,
	CloseAuctionResultCodeCloseAuctionNoWinners,
	CloseAuctionResultCodeCloseAuctionAlreadyClosed,
	CloseAuctionResultCodeCloseAuctionHasWinner,
	CloseAuctionResultCodeCloseAuctionNotAllowed,
}

var closeAuctionResultCodeMap = map[int32]string{
	0:  "CloseAuctionResultCodeCloseAuctionSuccess",
	-1: "CloseAuctionResultCodeCloseAuctionNotFound",
	-2: "CloseAuctionResultCodeCloseAuctionTooEarly",
	-3: "CloseAuctionResultCodeCloseAuctionNoWinners",
	-4: "CloseAuctionResultCodeCloseAuctionAlreadyClosed",
	-5: "CloseAuctionResultCodeCloseAuctionHasWinner",
	-6: "CloseAuctionResultCodeCloseAuctionNotAllowed",
}

var closeAuctionResultCodeShortMap = map[int32]string{
	0:  "close_auction_success",
	-1: "close_auction_not_found",
	-2: "close_auction_too_early",
	-3: "close_auction_no_winners",
	-4: "close_auction_already_closed",
	-5: "close_auction_has_winner",
	-6: "close_auction_not_allowed",
}

var closeAuctionResultCodeRevMap = map[string]int32{
	"CloseAuctionResultCodeCloseAuctionSuccess":       0,
	"CloseAuctionResultCodeCloseAuctionNotFound":      -1,
	"CloseAuctionResultCodeCloseAuctionTooEarly":      -2,
	"CloseAuctionResultCodeCloseAuctionNoWinners":     -3,
	"CloseAuctionResultCodeCloseAuctionAlreadyClosed": -4,
	"CloseAuctionResultCodeCloseAuctionHasWinner":     -5,
	"CloseAuctionResultCodeCloseAuctionNotAllowed":    -6,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CloseAuctionResultCode
func (e CloseAuctionResultCode) ValidEnum(v int32) bool {
	_, ok := closeAuctionResultCodeMap[v]
	return ok
}
func (e CloseAuctionResultCode) isFlag() bool {
	for i := len(CloseAuctionResultCodeAll) - 1; i >= 0; i-- {
		expected := CloseAuctionResultCode(2) << uint64(len(CloseAuctionResultCodeAll)-1) >> uint64(len(CloseAuctionResultCodeAll)-i)
		if expected != CloseAuctionResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CloseAuctionResultCode) String() string {
	name, _ := closeAuctionResultCodeMap[int32(e)]
	return name
}

func (e CloseAuctionResultCode) ShortString() string {
	name, _ := closeAuctionResultCodeShortMap[int32(e)]
	return name
}

func (e CloseAuctionResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range CloseAuctionResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CloseAuctionResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CloseAuctionResultCode(t.Value)
	return nil
}

// CloseAuctionSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type CloseAuctionSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CloseAuctionSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CloseAuctionSuccessExt
func (u CloseAuctionSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCloseAuctionSuccessExt creates a new  CloseAuctionSuccessExt.
func NewCloseAuctionSuccessExt(v LedgerVersion, value interface{}) (result CloseAuctionSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CloseAuctionSuccess is an XDR Struct defines as:
//
//   struct CloseAuctionSuccess
//        {
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type CloseAuctionSuccess struct {
	Ext CloseAuctionSuccessExt `json:"ext,omitempty"`
}

// CloseAuctionResult is an XDR Union defines as:
//
//   union CloseAuctionResult switch (CloseAuctionResultCode code)
//        {
//        case CLOSE_AUCTION_SUCCESS:
//            CloseAuctionSuccess success;
//        default:
//            void;
//        };
//
type CloseAuctionResult struct {
	Code    CloseAuctionResultCode `json:"code,omitempty"`
	Success *CloseAuctionSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CloseAuctionResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CloseAuctionResult
func (u CloseAuctionResult) ArmForSwitch(sw int32) (string, bool) {
	switch CloseAuctionResultCode(sw) {
	case CloseAuctionResultCodeCloseAuctionSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCloseAuctionResult creates a new  CloseAuctionResult.
func NewCloseAuctionResult(code CloseAuctionResultCode, value interface{}) (result CloseAuctionResult, err error) {
	result.Code = code
	switch CloseAuctionResultCode(code) {
	case CloseAuctionResultCodeCloseAuctionSuccess:
		tv, ok := value.(CloseAuctionSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CloseAuctionSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CloseAuctionResult) MustSuccess() CloseAuctionSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CloseAuctionResult) GetSuccess() (result CloseAuctionSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// CreateAccountOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//
type CreateAccountOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateAccountOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateAccountOpExt
func (u CreateAccountOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateAccountOpExt creates a new  CreateAccountOpExt.
func NewCreateAccountOpExt(v LedgerVersion, value interface{}) (result CreateAccountOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateAccountOp is an XDR Struct defines as:
//
//   struct CreateAccountOp
//        {
//            AccountID destination;     // account to create
//            AccountID platformID;
//            AccountType accountType;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//            ext;
//        };
//
type CreateAccountOp struct {
	Destination AccountId          `json:"destination,omitempty"`
	PlatformId  AccountId          `json:"platformID,omitempty"`
	AccountType AccountType        `json:"accountType,omitempty"`
	Ext         CreateAccountOpExt `json:"ext,omitempty"`
}

// CreateAccountResultCode is an XDR Enum defines as:
//
//   enum CreateAccountResultCode
//        {
//            // codes considered as "success" for the operation
//            CREATE_ACCOUNT_SUCCESS = 0, // account was created
//
//            // codes considered as "failure" for the operation
//            CREATE_ACCOUNT_MALFORMED = -1,   // general problems
//            CREATE_ACCOUNT_TYPE_NOT_ALLOWED = -2,
//            CREATE_ACCOUNT_ALREADY_EXISTS = -3,
//            CREATE_ACCOUNT_PLATFORM_NOT_FOUND = -4,
//            CREATE_ACCOUNT_PLATFORM_ID_MISMATCHED = -5, // platform should belong to itself
//            CREATE_ACCOUNT_INVALID_VERSION = -6
//        };
//
type CreateAccountResultCode int32

const (
	CreateAccountResultCodeCreateAccountSuccess              CreateAccountResultCode = 0
	CreateAccountResultCodeCreateAccountMalformed            CreateAccountResultCode = -1
	CreateAccountResultCodeCreateAccountTypeNotAllowed       CreateAccountResultCode = -2
	CreateAccountResultCodeCreateAccountAlreadyExists        CreateAccountResultCode = -3
	CreateAccountResultCodeCreateAccountPlatformNotFound     CreateAccountResultCode = -4
	CreateAccountResultCodeCreateAccountPlatformIdMismatched CreateAccountResultCode = -5
	CreateAccountResultCodeCreateAccountInvalidVersion       CreateAccountResultCode = -6
)

var CreateAccountResultCodeAll = []CreateAccountResultCode{
	CreateAccountResultCodeCreateAccountSuccess,
	CreateAccountResultCodeCreateAccountMalformed,
	CreateAccountResultCodeCreateAccountTypeNotAllowed,
	CreateAccountResultCodeCreateAccountAlreadyExists,
	CreateAccountResultCodeCreateAccountPlatformNotFound,
	CreateAccountResultCodeCreateAccountPlatformIdMismatched,
	CreateAccountResultCodeCreateAccountInvalidVersion,
}

var createAccountResultCodeMap = map[int32]string{
	0:  "CreateAccountResultCodeCreateAccountSuccess",
	-1: "CreateAccountResultCodeCreateAccountMalformed",
	-2: "CreateAccountResultCodeCreateAccountTypeNotAllowed",
	-3: "CreateAccountResultCodeCreateAccountAlreadyExists",
	-4: "CreateAccountResultCodeCreateAccountPlatformNotFound",
	-5: "CreateAccountResultCodeCreateAccountPlatformIdMismatched",
	-6: "CreateAccountResultCodeCreateAccountInvalidVersion",
}

var createAccountResultCodeShortMap = map[int32]string{
	0:  "create_account_success",
	-1: "create_account_malformed",
	-2: "create_account_type_not_allowed",
	-3: "create_account_already_exists",
	-4: "create_account_platform_not_found",
	-5: "create_account_platform_id_mismatched",
	-6: "create_account_invalid_version",
}

var createAccountResultCodeRevMap = map[string]int32{
	"CreateAccountResultCodeCreateAccountSuccess":              0,
	"CreateAccountResultCodeCreateAccountMalformed":            -1,
	"CreateAccountResultCodeCreateAccountTypeNotAllowed":       -2,
	"CreateAccountResultCodeCreateAccountAlreadyExists":        -3,
	"CreateAccountResultCodeCreateAccountPlatformNotFound":     -4,
	"CreateAccountResultCodeCreateAccountPlatformIdMismatched": -5,
	"CreateAccountResultCodeCreateAccountInvalidVersion":       -6,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateAccountResultCode
func (e CreateAccountResultCode) ValidEnum(v int32) bool {
	_, ok := createAccountResultCodeMap[v]
	return ok
}
func (e CreateAccountResultCode) isFlag() bool {
	for i := len(CreateAccountResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateAccountResultCode(2) << uint64(len(CreateAccountResultCodeAll)-1) >> uint64(len(CreateAccountResultCodeAll)-i)
		if expected != CreateAccountResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateAccountResultCode) String() string {
	name, _ := createAccountResultCodeMap[int32(e)]
	return name
}

func (e CreateAccountResultCode) ShortString() string {
	name, _ := createAccountResultCodeShortMap[int32(e)]
	return name
}

func (e CreateAccountResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range CreateAccountResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CreateAccountResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateAccountResultCode(t.Value)
	return nil
}

// CreateAccountResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                  void;
//                }
//
type CreateAccountResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateAccountResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateAccountResultSuccessExt
func (u CreateAccountResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateAccountResultSuccessExt creates a new  CreateAccountResultSuccessExt.
func NewCreateAccountResultSuccessExt(v LedgerVersion, value interface{}) (result CreateAccountResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateAccountResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//                // reserved for future use
//                union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                  void;
//                }
//                ext;
//              }
//
type CreateAccountResultSuccess struct {
	Ext CreateAccountResultSuccessExt `json:"ext,omitempty"`
}

// CreateAccountResult is an XDR Union defines as:
//
//   union CreateAccountResult switch (CreateAccountResultCode code)
//        {
//        case CREATE_ACCOUNT_SUCCESS:
//            struct {
//                // reserved for future use
//                union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                  void;
//                }
//                ext;
//              } success;
//        default:
//            void;
//        };
//
type CreateAccountResult struct {
	Code    CreateAccountResultCode     `json:"code,omitempty"`
	Success *CreateAccountResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateAccountResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateAccountResult
func (u CreateAccountResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateAccountResultCode(sw) {
	case CreateAccountResultCodeCreateAccountSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCreateAccountResult creates a new  CreateAccountResult.
func NewCreateAccountResult(code CreateAccountResultCode, value interface{}) (result CreateAccountResult, err error) {
	result.Code = code
	switch CreateAccountResultCode(code) {
	case CreateAccountResultCodeCreateAccountSuccess:
		tv, ok := value.(CreateAccountResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAccountResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CreateAccountResult) MustSuccess() CreateAccountResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateAccountResult) GetSuccess() (result CreateAccountResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// CreateBidOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type CreateBidOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateBidOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateBidOpExt
func (u CreateBidOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateBidOpExt creates a new  CreateBidOpExt.
func NewCreateBidOpExt(v LedgerVersion, value interface{}) (result CreateBidOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateBidOp is an XDR Struct defines as:
//
//   struct CreateBidOp
//        {
//            uint64 lotID;
//            uint64 amount;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type CreateBidOp struct {
	LotId  Uint64         `json:"lotID,omitempty"`
	Amount Uint64         `json:"amount,omitempty"`
	Ext    CreateBidOpExt `json:"ext,omitempty"`
}

// CreateBidResultCode is an XDR Enum defines as:
//
//   enum CreateBidResultCode
//        {
//            CREATE_BID_SUCCESS = 0,
//            CREATE_BID_NO_LOT = -1,
//            CREATE_BID_NO_PARTICIPANT = -2,
//            CREATE_BID_TOO_EARLY = -3,
//            CREATE_BID_TOO_LATE = -4,
//            CREATE_BID_LESS_THAN_START_PRICE = -5,
//            CREATE_BID_LESS_THAN_PREV_BID = -6,
//            CREATE_BID_INVALID_STEP = -7,
//            CREATE_BID_MALFORMED = -8,
//            CREATE_BID_NO_AUCTION = -9,                 // can't bid on no-auction lot type
//            CREATE_BID_NOT_ALLOWED = -10,               // participant is not in pending state
//            CREATE_BID_COPART_ALREADY_HAS_BID = -15,
//            CREATE_BID_COPART_ALREADY_REQUESTED = -16,
//            CREATE_BID_COPART_NOT_ALLOWED = -17,
//            CREATE_BID_LIMIT_EXCEEDED = -18,            // participant exceeded bid limit
//            CREATE_BID_MAX_BID_EXCEEDED = -19           // bid exceeds max bid value
//        };
//
type CreateBidResultCode int32

const (
	CreateBidResultCodeCreateBidSuccess                CreateBidResultCode = 0
	CreateBidResultCodeCreateBidNoLot                  CreateBidResultCode = -1
	CreateBidResultCodeCreateBidNoParticipant          CreateBidResultCode = -2
	CreateBidResultCodeCreateBidTooEarly               CreateBidResultCode = -3
	CreateBidResultCodeCreateBidTooLate                CreateBidResultCode = -4
	CreateBidResultCodeCreateBidLessThanStartPrice     CreateBidResultCode = -5
	CreateBidResultCodeCreateBidLessThanPrevBid        CreateBidResultCode = -6
	CreateBidResultCodeCreateBidInvalidStep            CreateBidResultCode = -7
	CreateBidResultCodeCreateBidMalformed              CreateBidResultCode = -8
	CreateBidResultCodeCreateBidNoAuction              CreateBidResultCode = -9
	CreateBidResultCodeCreateBidNotAllowed             CreateBidResultCode = -10
	CreateBidResultCodeCreateBidCopartAlreadyHasBid    CreateBidResultCode = -15
	CreateBidResultCodeCreateBidCopartAlreadyRequested CreateBidResultCode = -16
	CreateBidResultCodeCreateBidCopartNotAllowed       CreateBidResultCode = -17
	CreateBidResultCodeCreateBidLimitExceeded          CreateBidResultCode = -18
	CreateBidResultCodeCreateBidMaxBidExceeded         CreateBidResultCode = -19
)

var CreateBidResultCodeAll = []CreateBidResultCode{
	CreateBidResultCodeCreateBidSuccess,
	CreateBidResultCodeCreateBidNoLot,
	CreateBidResultCodeCreateBidNoParticipant,
	CreateBidResultCodeCreateBidTooEarly,
	CreateBidResultCodeCreateBidTooLate,
	CreateBidResultCodeCreateBidLessThanStartPrice,
	CreateBidResultCodeCreateBidLessThanPrevBid,
	CreateBidResultCodeCreateBidInvalidStep,
	CreateBidResultCodeCreateBidMalformed,
	CreateBidResultCodeCreateBidNoAuction,
	CreateBidResultCodeCreateBidNotAllowed,
	CreateBidResultCodeCreateBidCopartAlreadyHasBid,
	CreateBidResultCodeCreateBidCopartAlreadyRequested,
	CreateBidResultCodeCreateBidCopartNotAllowed,
	CreateBidResultCodeCreateBidLimitExceeded,
	CreateBidResultCodeCreateBidMaxBidExceeded,
}

var createBidResultCodeMap = map[int32]string{
	0:   "CreateBidResultCodeCreateBidSuccess",
	-1:  "CreateBidResultCodeCreateBidNoLot",
	-2:  "CreateBidResultCodeCreateBidNoParticipant",
	-3:  "CreateBidResultCodeCreateBidTooEarly",
	-4:  "CreateBidResultCodeCreateBidTooLate",
	-5:  "CreateBidResultCodeCreateBidLessThanStartPrice",
	-6:  "CreateBidResultCodeCreateBidLessThanPrevBid",
	-7:  "CreateBidResultCodeCreateBidInvalidStep",
	-8:  "CreateBidResultCodeCreateBidMalformed",
	-9:  "CreateBidResultCodeCreateBidNoAuction",
	-10: "CreateBidResultCodeCreateBidNotAllowed",
	-15: "CreateBidResultCodeCreateBidCopartAlreadyHasBid",
	-16: "CreateBidResultCodeCreateBidCopartAlreadyRequested",
	-17: "CreateBidResultCodeCreateBidCopartNotAllowed",
	-18: "CreateBidResultCodeCreateBidLimitExceeded",
	-19: "CreateBidResultCodeCreateBidMaxBidExceeded",
}

var createBidResultCodeShortMap = map[int32]string{
	0:   "create_bid_success",
	-1:  "create_bid_no_lot",
	-2:  "create_bid_no_participant",
	-3:  "create_bid_too_early",
	-4:  "create_bid_too_late",
	-5:  "create_bid_less_than_start_price",
	-6:  "create_bid_less_than_prev_bid",
	-7:  "create_bid_invalid_step",
	-8:  "create_bid_malformed",
	-9:  "create_bid_no_auction",
	-10: "create_bid_not_allowed",
	-15: "create_bid_copart_already_has_bid",
	-16: "create_bid_copart_already_requested",
	-17: "create_bid_copart_not_allowed",
	-18: "create_bid_limit_exceeded",
	-19: "create_bid_max_bid_exceeded",
}

var createBidResultCodeRevMap = map[string]int32{
	"CreateBidResultCodeCreateBidSuccess":                0,
	"CreateBidResultCodeCreateBidNoLot":                  -1,
	"CreateBidResultCodeCreateBidNoParticipant":          -2,
	"CreateBidResultCodeCreateBidTooEarly":               -3,
	"CreateBidResultCodeCreateBidTooLate":                -4,
	"CreateBidResultCodeCreateBidLessThanStartPrice":     -5,
	"CreateBidResultCodeCreateBidLessThanPrevBid":        -6,
	"CreateBidResultCodeCreateBidInvalidStep":            -7,
	"CreateBidResultCodeCreateBidMalformed":              -8,
	"CreateBidResultCodeCreateBidNoAuction":              -9,
	"CreateBidResultCodeCreateBidNotAllowed":             -10,
	"CreateBidResultCodeCreateBidCopartAlreadyHasBid":    -15,
	"CreateBidResultCodeCreateBidCopartAlreadyRequested": -16,
	"CreateBidResultCodeCreateBidCopartNotAllowed":       -17,
	"CreateBidResultCodeCreateBidLimitExceeded":          -18,
	"CreateBidResultCodeCreateBidMaxBidExceeded":         -19,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateBidResultCode
func (e CreateBidResultCode) ValidEnum(v int32) bool {
	_, ok := createBidResultCodeMap[v]
	return ok
}
func (e CreateBidResultCode) isFlag() bool {
	for i := len(CreateBidResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateBidResultCode(2) << uint64(len(CreateBidResultCodeAll)-1) >> uint64(len(CreateBidResultCodeAll)-i)
		if expected != CreateBidResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateBidResultCode) String() string {
	name, _ := createBidResultCodeMap[int32(e)]
	return name
}

func (e CreateBidResultCode) ShortString() string {
	name, _ := createBidResultCodeShortMap[int32(e)]
	return name
}

func (e CreateBidResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range CreateBidResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CreateBidResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateBidResultCode(t.Value)
	return nil
}

// CreateBidSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type CreateBidSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateBidSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateBidSuccessExt
func (u CreateBidSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateBidSuccessExt creates a new  CreateBidSuccessExt.
func NewCreateBidSuccessExt(v LedgerVersion, value interface{}) (result CreateBidSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateBidSuccess is an XDR Struct defines as:
//
//   struct CreateBidSuccess
//        {
//            bool fulfilled;
//            uint64 requestID;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type CreateBidSuccess struct {
	Fulfilled bool                `json:"fulfilled,omitempty"`
	RequestId Uint64              `json:"requestID,omitempty"`
	Ext       CreateBidSuccessExt `json:"ext,omitempty"`
}

// CreateBidResult is an XDR Union defines as:
//
//   union CreateBidResult switch (CreateBidResultCode code)
//        {
//        case CREATE_BID_SUCCESS:
//           CreateBidSuccess success;
//        default:
//            void;
//        };
//
type CreateBidResult struct {
	Code    CreateBidResultCode `json:"code,omitempty"`
	Success *CreateBidSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateBidResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateBidResult
func (u CreateBidResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateBidResultCode(sw) {
	case CreateBidResultCodeCreateBidSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCreateBidResult creates a new  CreateBidResult.
func NewCreateBidResult(code CreateBidResultCode, value interface{}) (result CreateBidResult, err error) {
	result.Code = code
	switch CreateBidResultCode(code) {
	case CreateBidResultCodeCreateBidSuccess:
		tv, ok := value.(CreateBidSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateBidSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CreateBidResult) MustSuccess() CreateBidSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateBidResult) GetSuccess() (result CreateBidSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// LotTypeSpecificData is an XDR Union defines as:
//
//   union LotTypeSpecificData switch (LotType lt)
//        {
//        case LOT_TYPE_NO_AUCTION:
//            LotPrices lotPrices;
//        };
//
type LotTypeSpecificData struct {
	Lt        LotType    `json:"lt,omitempty"`
	LotPrices *LotPrices `json:"lotPrices,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LotTypeSpecificData) SwitchFieldName() string {
	return "Lt"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LotTypeSpecificData
func (u LotTypeSpecificData) ArmForSwitch(sw int32) (string, bool) {
	switch LotType(sw) {
	case LotTypeLotTypeNoAuction:
		return "LotPrices", true
	}
	return "-", false
}

// NewLotTypeSpecificData creates a new  LotTypeSpecificData.
func NewLotTypeSpecificData(lt LotType, value interface{}) (result LotTypeSpecificData, err error) {
	result.Lt = lt
	switch LotType(lt) {
	case LotTypeLotTypeNoAuction:
		tv, ok := value.(LotPrices)
		if !ok {
			err = fmt.Errorf("invalid value, must be LotPrices")
			return
		}
		result.LotPrices = &tv
	}
	return
}

// MustLotPrices retrieves the LotPrices value from the union,
// panicing if the value is not set.
func (u LotTypeSpecificData) MustLotPrices() LotPrices {
	val, ok := u.GetLotPrices()

	if !ok {
		panic("arm LotPrices is not set")
	}

	return val
}

// GetLotPrices retrieves the LotPrices value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LotTypeSpecificData) GetLotPrices() (result LotPrices, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Lt))

	if armName == "LotPrices" {
		result = *u.LotPrices
		ok = true
	}

	return
}

// CreateLotPlatformPartExt is an XDR Struct defines as:
//
//   struct CreateLotPlatformPartExt
//        {
//            LotTypeSpecificData data;
//            bool createPlatformPart;
//        };
//
type CreateLotPlatformPartExt struct {
	Data               LotTypeSpecificData `json:"data,omitempty"`
	CreatePlatformPart bool                `json:"createPlatformPart,omitempty"`
}

// CreateLotOrganizerExt is an XDR Struct defines as:
//
//   struct CreateLotOrganizerExt
//        {
//            AccountID organizerID;
//            CreateLotPlatformPartExt createLotPlatformPartExt;
//        };
//
type CreateLotOrganizerExt struct {
	OrganizerId              AccountId                `json:"organizerID,omitempty"`
	CreateLotPlatformPartExt CreateLotPlatformPartExt `json:"createLotPlatformPartExt,omitempty"`
}

// CreateLotOpData is an XDR NestedUnion defines as:
//
//   union switch (LotType lt)
//                {
//                case LOT_TYPE_NO_AUCTION:
//                    LotPrices lotPrices;
//                }
//
type CreateLotOpData struct {
	Lt        LotType    `json:"lt,omitempty"`
	LotPrices *LotPrices `json:"lotPrices,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateLotOpData) SwitchFieldName() string {
	return "Lt"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateLotOpData
func (u CreateLotOpData) ArmForSwitch(sw int32) (string, bool) {
	switch LotType(sw) {
	case LotTypeLotTypeNoAuction:
		return "LotPrices", true
	}
	return "-", false
}

// NewCreateLotOpData creates a new  CreateLotOpData.
func NewCreateLotOpData(lt LotType, value interface{}) (result CreateLotOpData, err error) {
	result.Lt = lt
	switch LotType(lt) {
	case LotTypeLotTypeNoAuction:
		tv, ok := value.(LotPrices)
		if !ok {
			err = fmt.Errorf("invalid value, must be LotPrices")
			return
		}
		result.LotPrices = &tv
	}
	return
}

// MustLotPrices retrieves the LotPrices value from the union,
// panicing if the value is not set.
func (u CreateLotOpData) MustLotPrices() LotPrices {
	val, ok := u.GetLotPrices()

	if !ok {
		panic("arm LotPrices is not set")
	}

	return val
}

// GetLotPrices retrieves the LotPrices value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateLotOpData) GetLotPrices() (result LotPrices, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Lt))

	if armName == "LotPrices" {
		result = *u.LotPrices
		ok = true
	}

	return
}

// CreateLotOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            case LEDGER_VERSION_B2B_B2C_PRICES:
//                union switch (LotType lt)
//                {
//                case LOT_TYPE_NO_AUCTION:
//                    LotPrices lotPrices;
//                } data;
//            case LEDGER_VERSION_PLATFORM_PARTICIPANT:
//                CreateLotPlatformPartExt platformPartExt;
//            case LEDGER_VERSION_ORGANIZER_ID:
//                CreateLotOrganizerExt organizerExt;
//            }
//
type CreateLotOpExt struct {
	V               LedgerVersion             `json:"v,omitempty"`
	Data            *CreateLotOpData          `json:"data,omitempty"`
	PlatformPartExt *CreateLotPlatformPartExt `json:"platformPartExt,omitempty"`
	OrganizerExt    *CreateLotOrganizerExt    `json:"organizerExt,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateLotOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateLotOpExt
func (u CreateLotOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionLedgerVersionB2BB2CPrices:
		return "Data", true
	case LedgerVersionLedgerVersionPlatformParticipant:
		return "PlatformPartExt", true
	case LedgerVersionLedgerVersionOrganizerId:
		return "OrganizerExt", true
	}
	return "-", false
}

// NewCreateLotOpExt creates a new  CreateLotOpExt.
func NewCreateLotOpExt(v LedgerVersion, value interface{}) (result CreateLotOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionLedgerVersionB2BB2CPrices:
		tv, ok := value.(CreateLotOpData)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateLotOpData")
			return
		}
		result.Data = &tv
	case LedgerVersionLedgerVersionPlatformParticipant:
		tv, ok := value.(CreateLotPlatformPartExt)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateLotPlatformPartExt")
			return
		}
		result.PlatformPartExt = &tv
	case LedgerVersionLedgerVersionOrganizerId:
		tv, ok := value.(CreateLotOrganizerExt)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateLotOrganizerExt")
			return
		}
		result.OrganizerExt = &tv
	}
	return
}

// MustData retrieves the Data value from the union,
// panicing if the value is not set.
func (u CreateLotOpExt) MustData() CreateLotOpData {
	val, ok := u.GetData()

	if !ok {
		panic("arm Data is not set")
	}

	return val
}

// GetData retrieves the Data value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateLotOpExt) GetData() (result CreateLotOpData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "Data" {
		result = *u.Data
		ok = true
	}

	return
}

// MustPlatformPartExt retrieves the PlatformPartExt value from the union,
// panicing if the value is not set.
func (u CreateLotOpExt) MustPlatformPartExt() CreateLotPlatformPartExt {
	val, ok := u.GetPlatformPartExt()

	if !ok {
		panic("arm PlatformPartExt is not set")
	}

	return val
}

// GetPlatformPartExt retrieves the PlatformPartExt value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateLotOpExt) GetPlatformPartExt() (result CreateLotPlatformPartExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "PlatformPartExt" {
		result = *u.PlatformPartExt
		ok = true
	}

	return
}

// MustOrganizerExt retrieves the OrganizerExt value from the union,
// panicing if the value is not set.
func (u CreateLotOpExt) MustOrganizerExt() CreateLotOrganizerExt {
	val, ok := u.GetOrganizerExt()

	if !ok {
		panic("arm OrganizerExt is not set")
	}

	return val
}

// GetOrganizerExt retrieves the OrganizerExt value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateLotOpExt) GetOrganizerExt() (result CreateLotOrganizerExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "OrganizerExt" {
		result = *u.OrganizerExt
		ok = true
	}

	return
}

// CreateLotOp is an XDR Struct defines as:
//
//   struct CreateLotOp
//        {
//            uint64 requestID;       // 0 to create new, otherwise update existing
//
//            LotType lotType;
//            bool buyNowSupport;     // does lot support buy now
//
//            string details<>;       // JSON string with full lot description
//
//            string currency<3>;     // currency of start price and buy now price
//            uint64 startPrice;      // all prices are indicated in cents
//            uint64 deposit;         // minimal deposit to pay for participating
//            uint64 buyNowPrice;
//            uint64 minStep;
//            uint64 maxStep;
//
//            int64 startTime;        // auction start time in Unix epoch format
//            uint64 duration;        // duration in seconds (optional for certain type of lot)
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            case LEDGER_VERSION_B2B_B2C_PRICES:
//                union switch (LotType lt)
//                {
//                case LOT_TYPE_NO_AUCTION:
//                    LotPrices lotPrices;
//                } data;
//            case LEDGER_VERSION_PLATFORM_PARTICIPANT:
//                CreateLotPlatformPartExt platformPartExt;
//            case LEDGER_VERSION_ORGANIZER_ID:
//                CreateLotOrganizerExt organizerExt;
//            }
//            ext;
//        };
//
type CreateLotOp struct {
	RequestId     Uint64         `json:"requestID,omitempty"`
	LotType       LotType        `json:"lotType,omitempty"`
	BuyNowSupport bool           `json:"buyNowSupport,omitempty"`
	Details       string         `json:"details,omitempty"`
	Currency      string         `json:"currency,omitempty" xdrmaxsize:"3"`
	StartPrice    Uint64         `json:"startPrice,omitempty"`
	Deposit       Uint64         `json:"deposit,omitempty"`
	BuyNowPrice   Uint64         `json:"buyNowPrice,omitempty"`
	MinStep       Uint64         `json:"minStep,omitempty"`
	MaxStep       Uint64         `json:"maxStep,omitempty"`
	StartTime     Int64          `json:"startTime,omitempty"`
	Duration      Uint64         `json:"duration,omitempty"`
	Ext           CreateLotOpExt `json:"ext,omitempty"`
}

// CreateLotSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type CreateLotSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateLotSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateLotSuccessExt
func (u CreateLotSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateLotSuccessExt creates a new  CreateLotSuccessExt.
func NewCreateLotSuccessExt(v LedgerVersion, value interface{}) (result CreateLotSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateLotSuccess is an XDR Struct defines as:
//
//   struct CreateLotSuccess {
//            uint64 requestID;
//            uint64 lotID;
//            bool fulfilled;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type CreateLotSuccess struct {
	RequestId Uint64              `json:"requestID,omitempty"`
	LotId     Uint64              `json:"lotID,omitempty"`
	Fulfilled bool                `json:"fulfilled,omitempty"`
	Ext       CreateLotSuccessExt `json:"ext,omitempty"`
}

// CreateLotResultCode is an XDR Enum defines as:
//
//   enum CreateLotResultCode
//        {
//            CREATE_LOT_SUCCESS = 0, // lot was created
//
//            CREATE_LOT_ALREADY_EXIST = -1,              // for future use
//            CREATE_LOT_MALFORMED = -2,
//            CREATE_LOT_MALFORMED_START_TIME = -3,       // attempt to create lot in past
//            CREATE_LOT_TOO_LOW_MIN_STEP = -4,           // should be >= 1
//            CREATE_LOT_TOO_LOW_MAX_STEP = -5,           // should be > min step
//            CREATE_LOT_INVALID_JSON = -6,
//            CREATE_LOT_MALFORMED_CURRENCY = -7,
//            CREATE_LOT_MALFORMED_PRICE = -8,            // should be buyNow > startPrice (if supported)
//            CREATE_LOT_INVALID_VERSION = -9,
//            CREATE_LOT_REQUEST_NOT_FOUND = -10,
//            CREATE_LOT_REFERENCE_DUPLICATION = -11,
//            CREATE_LOT_LOT_TYPE_NOT_ALLOWED = -12,      // source is not allowed to create lots of this type
//            CREATE_LOT_NO_AUCTION_MALFORMED = -13,      // lot must support buy now and have zero start price and steps
//            CREATE_LOT_COPART_MALFORMED = -14,          // copart lots must have zero steps and duration
//            CREATE_LOT_IAAI_MALFORMED = -15,
//            CREATE_LOT_DURATION_MALFORMED = -16
//        };
//
type CreateLotResultCode int32

const (
	CreateLotResultCodeCreateLotSuccess              CreateLotResultCode = 0
	CreateLotResultCodeCreateLotAlreadyExist         CreateLotResultCode = -1
	CreateLotResultCodeCreateLotMalformed            CreateLotResultCode = -2
	CreateLotResultCodeCreateLotMalformedStartTime   CreateLotResultCode = -3
	CreateLotResultCodeCreateLotTooLowMinStep        CreateLotResultCode = -4
	CreateLotResultCodeCreateLotTooLowMaxStep        CreateLotResultCode = -5
	CreateLotResultCodeCreateLotInvalidJson          CreateLotResultCode = -6
	CreateLotResultCodeCreateLotMalformedCurrency    CreateLotResultCode = -7
	CreateLotResultCodeCreateLotMalformedPrice       CreateLotResultCode = -8
	CreateLotResultCodeCreateLotInvalidVersion       CreateLotResultCode = -9
	CreateLotResultCodeCreateLotRequestNotFound      CreateLotResultCode = -10
	CreateLotResultCodeCreateLotReferenceDuplication CreateLotResultCode = -11
	CreateLotResultCodeCreateLotLotTypeNotAllowed    CreateLotResultCode = -12
	CreateLotResultCodeCreateLotNoAuctionMalformed   CreateLotResultCode = -13
	CreateLotResultCodeCreateLotCopartMalformed      CreateLotResultCode = -14
	CreateLotResultCodeCreateLotIaaiMalformed        CreateLotResultCode = -15
	CreateLotResultCodeCreateLotDurationMalformed    CreateLotResultCode = -16
)

var CreateLotResultCodeAll = []CreateLotResultCode{
	CreateLotResultCodeCreateLotSuccess,
	CreateLotResultCodeCreateLotAlreadyExist,
	CreateLotResultCodeCreateLotMalformed,
	CreateLotResultCodeCreateLotMalformedStartTime,
	CreateLotResultCodeCreateLotTooLowMinStep,
	CreateLotResultCodeCreateLotTooLowMaxStep,
	CreateLotResultCodeCreateLotInvalidJson,
	CreateLotResultCodeCreateLotMalformedCurrency,
	CreateLotResultCodeCreateLotMalformedPrice,
	CreateLotResultCodeCreateLotInvalidVersion,
	CreateLotResultCodeCreateLotRequestNotFound,
	CreateLotResultCodeCreateLotReferenceDuplication,
	CreateLotResultCodeCreateLotLotTypeNotAllowed,
	CreateLotResultCodeCreateLotNoAuctionMalformed,
	CreateLotResultCodeCreateLotCopartMalformed,
	CreateLotResultCodeCreateLotIaaiMalformed,
	CreateLotResultCodeCreateLotDurationMalformed,
}

var createLotResultCodeMap = map[int32]string{
	0:   "CreateLotResultCodeCreateLotSuccess",
	-1:  "CreateLotResultCodeCreateLotAlreadyExist",
	-2:  "CreateLotResultCodeCreateLotMalformed",
	-3:  "CreateLotResultCodeCreateLotMalformedStartTime",
	-4:  "CreateLotResultCodeCreateLotTooLowMinStep",
	-5:  "CreateLotResultCodeCreateLotTooLowMaxStep",
	-6:  "CreateLotResultCodeCreateLotInvalidJson",
	-7:  "CreateLotResultCodeCreateLotMalformedCurrency",
	-8:  "CreateLotResultCodeCreateLotMalformedPrice",
	-9:  "CreateLotResultCodeCreateLotInvalidVersion",
	-10: "CreateLotResultCodeCreateLotRequestNotFound",
	-11: "CreateLotResultCodeCreateLotReferenceDuplication",
	-12: "CreateLotResultCodeCreateLotLotTypeNotAllowed",
	-13: "CreateLotResultCodeCreateLotNoAuctionMalformed",
	-14: "CreateLotResultCodeCreateLotCopartMalformed",
	-15: "CreateLotResultCodeCreateLotIaaiMalformed",
	-16: "CreateLotResultCodeCreateLotDurationMalformed",
}

var createLotResultCodeShortMap = map[int32]string{
	0:   "create_lot_success",
	-1:  "create_lot_already_exist",
	-2:  "create_lot_malformed",
	-3:  "create_lot_malformed_start_time",
	-4:  "create_lot_too_low_min_step",
	-5:  "create_lot_too_low_max_step",
	-6:  "create_lot_invalid_json",
	-7:  "create_lot_malformed_currency",
	-8:  "create_lot_malformed_price",
	-9:  "create_lot_invalid_version",
	-10: "create_lot_request_not_found",
	-11: "create_lot_reference_duplication",
	-12: "create_lot_lot_type_not_allowed",
	-13: "create_lot_no_auction_malformed",
	-14: "create_lot_copart_malformed",
	-15: "create_lot_iaai_malformed",
	-16: "create_lot_duration_malformed",
}

var createLotResultCodeRevMap = map[string]int32{
	"CreateLotResultCodeCreateLotSuccess":              0,
	"CreateLotResultCodeCreateLotAlreadyExist":         -1,
	"CreateLotResultCodeCreateLotMalformed":            -2,
	"CreateLotResultCodeCreateLotMalformedStartTime":   -3,
	"CreateLotResultCodeCreateLotTooLowMinStep":        -4,
	"CreateLotResultCodeCreateLotTooLowMaxStep":        -5,
	"CreateLotResultCodeCreateLotInvalidJson":          -6,
	"CreateLotResultCodeCreateLotMalformedCurrency":    -7,
	"CreateLotResultCodeCreateLotMalformedPrice":       -8,
	"CreateLotResultCodeCreateLotInvalidVersion":       -9,
	"CreateLotResultCodeCreateLotRequestNotFound":      -10,
	"CreateLotResultCodeCreateLotReferenceDuplication": -11,
	"CreateLotResultCodeCreateLotLotTypeNotAllowed":    -12,
	"CreateLotResultCodeCreateLotNoAuctionMalformed":   -13,
	"CreateLotResultCodeCreateLotCopartMalformed":      -14,
	"CreateLotResultCodeCreateLotIaaiMalformed":        -15,
	"CreateLotResultCodeCreateLotDurationMalformed":    -16,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateLotResultCode
func (e CreateLotResultCode) ValidEnum(v int32) bool {
	_, ok := createLotResultCodeMap[v]
	return ok
}
func (e CreateLotResultCode) isFlag() bool {
	for i := len(CreateLotResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateLotResultCode(2) << uint64(len(CreateLotResultCodeAll)-1) >> uint64(len(CreateLotResultCodeAll)-i)
		if expected != CreateLotResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateLotResultCode) String() string {
	name, _ := createLotResultCodeMap[int32(e)]
	return name
}

func (e CreateLotResultCode) ShortString() string {
	name, _ := createLotResultCodeShortMap[int32(e)]
	return name
}

func (e CreateLotResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range CreateLotResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CreateLotResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateLotResultCode(t.Value)
	return nil
}

// CreateLotResult is an XDR Union defines as:
//
//   union CreateLotResult switch (CreateLotResultCode code)
//        {
//        case CREATE_LOT_SUCCESS:
//            CreateLotSuccess success;
//        default:
//            void;
//        };
//
type CreateLotResult struct {
	Code    CreateLotResultCode `json:"code,omitempty"`
	Success *CreateLotSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateLotResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateLotResult
func (u CreateLotResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateLotResultCode(sw) {
	case CreateLotResultCodeCreateLotSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCreateLotResult creates a new  CreateLotResult.
func NewCreateLotResult(code CreateLotResultCode, value interface{}) (result CreateLotResult, err error) {
	result.Code = code
	switch CreateLotResultCode(code) {
	case CreateLotResultCodeCreateLotSuccess:
		tv, ok := value.(CreateLotSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateLotSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CreateLotResult) MustSuccess() CreateLotSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateLotResult) GetSuccess() (result CreateLotSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// CreatePassiveOfferOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type CreatePassiveOfferOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreatePassiveOfferOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreatePassiveOfferOpExt
func (u CreatePassiveOfferOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreatePassiveOfferOpExt creates a new  CreatePassiveOfferOpExt.
func NewCreatePassiveOfferOpExt(v LedgerVersion, value interface{}) (result CreatePassiveOfferOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreatePassiveOfferOp is an XDR Struct defines as:
//
//   struct CreatePassiveOfferOp
//        {
//            Asset selling; // A
//            Asset buying;  // B
//            int64 amount;  // amount taker gets. if set to 0, delete the offer
//            Price price;   // cost of A in terms of B
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        };
//
type CreatePassiveOfferOp struct {
	Selling Asset                   `json:"selling,omitempty"`
	Buying  Asset                   `json:"buying,omitempty"`
	Amount  Int64                   `json:"amount,omitempty"`
	Price   Price                   `json:"price,omitempty"`
	Ext     CreatePassiveOfferOpExt `json:"ext,omitempty"`
}

// ManageBidLimitOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type ManageBidLimitOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageBidLimitOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageBidLimitOpExt
func (u ManageBidLimitOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageBidLimitOpExt creates a new  ManageBidLimitOpExt.
func NewManageBidLimitOpExt(v LedgerVersion, value interface{}) (result ManageBidLimitOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageBidLimitOp is an XDR Struct defines as:
//
//   struct ManageBidLimitOp
//        {
//            uint64 lotID;
//            AccountID accountID;
//
//            uint64 limit;
//
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type ManageBidLimitOp struct {
	LotId     Uint64              `json:"lotID,omitempty"`
	AccountId AccountId           `json:"accountID,omitempty"`
	Limit     Uint64              `json:"limit,omitempty"`
	Ext       ManageBidLimitOpExt `json:"ext,omitempty"`
}

// ManageBidLimitResultCode is an XDR Enum defines as:
//
//   enum ManageBidLimitResultCode
//        {
//            MANAGE_BID_LIMIT_SUCCESS = 0,
//
//            MANAGE_BID_LIMIT_NOT_ALLOWED = -1,                  // source is not allowed to manage limit of this participant
//            MANAGE_BID_LIMIT_NOT_FOUND = -2,                    // participant not found
//            MANAGE_BID_LIMIT_LOT_IS_NOT_ACTIVE = -3,            // lot is not active (closed or finished)
//            MANAGE_BID_LIMIT_PARTICIPANT_IS_WINNER = -4,        // participant is already winner
//            MANAGE_BID_LIMIT_PARTICIPANT_IS_REJECTED = -5,      // participant is rejected
//            MANAGE_BID_LIMIT_PARTICIPANT_IS_NOT_APPROVED = -6,  // participant is in state 'deposit pending'
//            MANAGE_BID_LIMIT_INVALID_VERSION = -7               // current ledger version does not support this op
//        };
//
type ManageBidLimitResultCode int32

const (
	ManageBidLimitResultCodeManageBidLimitSuccess                  ManageBidLimitResultCode = 0
	ManageBidLimitResultCodeManageBidLimitNotAllowed               ManageBidLimitResultCode = -1
	ManageBidLimitResultCodeManageBidLimitNotFound                 ManageBidLimitResultCode = -2
	ManageBidLimitResultCodeManageBidLimitLotIsNotActive           ManageBidLimitResultCode = -3
	ManageBidLimitResultCodeManageBidLimitParticipantIsWinner      ManageBidLimitResultCode = -4
	ManageBidLimitResultCodeManageBidLimitParticipantIsRejected    ManageBidLimitResultCode = -5
	ManageBidLimitResultCodeManageBidLimitParticipantIsNotApproved ManageBidLimitResultCode = -6
	ManageBidLimitResultCodeManageBidLimitInvalidVersion           ManageBidLimitResultCode = -7
)

var ManageBidLimitResultCodeAll = []ManageBidLimitResultCode{
	ManageBidLimitResultCodeManageBidLimitSuccess,
	ManageBidLimitResultCodeManageBidLimitNotAllowed,
	ManageBidLimitResultCodeManageBidLimitNotFound,
	ManageBidLimitResultCodeManageBidLimitLotIsNotActive,
	ManageBidLimitResultCodeManageBidLimitParticipantIsWinner,
	ManageBidLimitResultCodeManageBidLimitParticipantIsRejected,
	ManageBidLimitResultCodeManageBidLimitParticipantIsNotApproved,
	ManageBidLimitResultCodeManageBidLimitInvalidVersion,
}

var manageBidLimitResultCodeMap = map[int32]string{
	0:  "ManageBidLimitResultCodeManageBidLimitSuccess",
	-1: "ManageBidLimitResultCodeManageBidLimitNotAllowed",
	-2: "ManageBidLimitResultCodeManageBidLimitNotFound",
	-3: "ManageBidLimitResultCodeManageBidLimitLotIsNotActive",
	-4: "ManageBidLimitResultCodeManageBidLimitParticipantIsWinner",
	-5: "ManageBidLimitResultCodeManageBidLimitParticipantIsRejected",
	-6: "ManageBidLimitResultCodeManageBidLimitParticipantIsNotApproved",
	-7: "ManageBidLimitResultCodeManageBidLimitInvalidVersion",
}

var manageBidLimitResultCodeShortMap = map[int32]string{
	0:  "manage_bid_limit_success",
	-1: "manage_bid_limit_not_allowed",
	-2: "manage_bid_limit_not_found",
	-3: "manage_bid_limit_lot_is_not_active",
	-4: "manage_bid_limit_participant_is_winner",
	-5: "manage_bid_limit_participant_is_rejected",
	-6: "manage_bid_limit_participant_is_not_approved",
	-7: "manage_bid_limit_invalid_version",
}

var manageBidLimitResultCodeRevMap = map[string]int32{
	"ManageBidLimitResultCodeManageBidLimitSuccess":                  0,
	"ManageBidLimitResultCodeManageBidLimitNotAllowed":               -1,
	"ManageBidLimitResultCodeManageBidLimitNotFound":                 -2,
	"ManageBidLimitResultCodeManageBidLimitLotIsNotActive":           -3,
	"ManageBidLimitResultCodeManageBidLimitParticipantIsWinner":      -4,
	"ManageBidLimitResultCodeManageBidLimitParticipantIsRejected":    -5,
	"ManageBidLimitResultCodeManageBidLimitParticipantIsNotApproved": -6,
	"ManageBidLimitResultCodeManageBidLimitInvalidVersion":           -7,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageBidLimitResultCode
func (e ManageBidLimitResultCode) ValidEnum(v int32) bool {
	_, ok := manageBidLimitResultCodeMap[v]
	return ok
}
func (e ManageBidLimitResultCode) isFlag() bool {
	for i := len(ManageBidLimitResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageBidLimitResultCode(2) << uint64(len(ManageBidLimitResultCodeAll)-1) >> uint64(len(ManageBidLimitResultCodeAll)-i)
		if expected != ManageBidLimitResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageBidLimitResultCode) String() string {
	name, _ := manageBidLimitResultCodeMap[int32(e)]
	return name
}

func (e ManageBidLimitResultCode) ShortString() string {
	name, _ := manageBidLimitResultCodeShortMap[int32(e)]
	return name
}

func (e ManageBidLimitResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageBidLimitResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageBidLimitResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageBidLimitResultCode(t.Value)
	return nil
}

// ManageBidLimitSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type ManageBidLimitSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageBidLimitSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageBidLimitSuccessExt
func (u ManageBidLimitSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageBidLimitSuccessExt creates a new  ManageBidLimitSuccessExt.
func NewManageBidLimitSuccessExt(v LedgerVersion, value interface{}) (result ManageBidLimitSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageBidLimitSuccess is an XDR Struct defines as:
//
//   struct ManageBidLimitSuccess
//        {
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type ManageBidLimitSuccess struct {
	Ext ManageBidLimitSuccessExt `json:"ext,omitempty"`
}

// ManageBidLimitResult is an XDR Union defines as:
//
//   union ManageBidLimitResult switch (ManageBidLimitResultCode code)
//        {
//        case MANAGE_BID_LIMIT_SUCCESS:
//            ManageBidLimitSuccess success;
//        default:
//            void;
//        };
//
type ManageBidLimitResult struct {
	Code    ManageBidLimitResultCode `json:"code,omitempty"`
	Success *ManageBidLimitSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageBidLimitResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageBidLimitResult
func (u ManageBidLimitResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageBidLimitResultCode(sw) {
	case ManageBidLimitResultCodeManageBidLimitSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManageBidLimitResult creates a new  ManageBidLimitResult.
func NewManageBidLimitResult(code ManageBidLimitResultCode, value interface{}) (result ManageBidLimitResult, err error) {
	result.Code = code
	switch ManageBidLimitResultCode(code) {
	case ManageBidLimitResultCodeManageBidLimitSuccess:
		tv, ok := value.(ManageBidLimitSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageBidLimitSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageBidLimitResult) MustSuccess() ManageBidLimitSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageBidLimitResult) GetSuccess() (result ManageBidLimitSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ManageMigrationOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//
type ManageMigrationOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageMigrationOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageMigrationOpExt
func (u ManageMigrationOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageMigrationOpExt creates a new  ManageMigrationOpExt.
func NewManageMigrationOpExt(v LedgerVersion, value interface{}) (result ManageMigrationOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageMigrationOp is an XDR Struct defines as:
//
//   struct ManageMigrationOp {
//            MigrationVersion migrationVersion;
//
//            union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//            ext;
//        };
//
type ManageMigrationOp struct {
	MigrationVersion MigrationVersion     `json:"migrationVersion,omitempty"`
	Ext              ManageMigrationOpExt `json:"ext,omitempty"`
}

// ManageMigrationResultCode is an XDR Enum defines as:
//
//   enum ManageMigrationResultCode
//        {
//            MANAGE_MIGRATION_SUCCESS = 0, // lot was created
//            MANAGE_MIGRATION_NOT_ROOT = -1,
//            MANAGE_MIGRATION_MALFORMED = -2,
//            MANAGE_MIGRATION_ALREADY_MIGRATE = -3,
//            MANAGE_MIGRATION_INVALID_MIGRATE_VERSION = -4,
//            MANAGE_MIGRATION_FAILED = -5
//        };
//
type ManageMigrationResultCode int32

const (
	ManageMigrationResultCodeManageMigrationSuccess               ManageMigrationResultCode = 0
	ManageMigrationResultCodeManageMigrationNotRoot               ManageMigrationResultCode = -1
	ManageMigrationResultCodeManageMigrationMalformed             ManageMigrationResultCode = -2
	ManageMigrationResultCodeManageMigrationAlreadyMigrate        ManageMigrationResultCode = -3
	ManageMigrationResultCodeManageMigrationInvalidMigrateVersion ManageMigrationResultCode = -4
	ManageMigrationResultCodeManageMigrationFailed                ManageMigrationResultCode = -5
)

var ManageMigrationResultCodeAll = []ManageMigrationResultCode{
	ManageMigrationResultCodeManageMigrationSuccess,
	ManageMigrationResultCodeManageMigrationNotRoot,
	ManageMigrationResultCodeManageMigrationMalformed,
	ManageMigrationResultCodeManageMigrationAlreadyMigrate,
	ManageMigrationResultCodeManageMigrationInvalidMigrateVersion,
	ManageMigrationResultCodeManageMigrationFailed,
}

var manageMigrationResultCodeMap = map[int32]string{
	0:  "ManageMigrationResultCodeManageMigrationSuccess",
	-1: "ManageMigrationResultCodeManageMigrationNotRoot",
	-2: "ManageMigrationResultCodeManageMigrationMalformed",
	-3: "ManageMigrationResultCodeManageMigrationAlreadyMigrate",
	-4: "ManageMigrationResultCodeManageMigrationInvalidMigrateVersion",
	-5: "ManageMigrationResultCodeManageMigrationFailed",
}

var manageMigrationResultCodeShortMap = map[int32]string{
	0:  "manage_migration_success",
	-1: "manage_migration_not_root",
	-2: "manage_migration_malformed",
	-3: "manage_migration_already_migrate",
	-4: "manage_migration_invalid_migrate_version",
	-5: "manage_migration_failed",
}

var manageMigrationResultCodeRevMap = map[string]int32{
	"ManageMigrationResultCodeManageMigrationSuccess":               0,
	"ManageMigrationResultCodeManageMigrationNotRoot":               -1,
	"ManageMigrationResultCodeManageMigrationMalformed":             -2,
	"ManageMigrationResultCodeManageMigrationAlreadyMigrate":        -3,
	"ManageMigrationResultCodeManageMigrationInvalidMigrateVersion": -4,
	"ManageMigrationResultCodeManageMigrationFailed":                -5,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageMigrationResultCode
func (e ManageMigrationResultCode) ValidEnum(v int32) bool {
	_, ok := manageMigrationResultCodeMap[v]
	return ok
}
func (e ManageMigrationResultCode) isFlag() bool {
	for i := len(ManageMigrationResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageMigrationResultCode(2) << uint64(len(ManageMigrationResultCodeAll)-1) >> uint64(len(ManageMigrationResultCodeAll)-i)
		if expected != ManageMigrationResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageMigrationResultCode) String() string {
	name, _ := manageMigrationResultCodeMap[int32(e)]
	return name
}

func (e ManageMigrationResultCode) ShortString() string {
	name, _ := manageMigrationResultCodeShortMap[int32(e)]
	return name
}

func (e ManageMigrationResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageMigrationResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageMigrationResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageMigrationResultCode(t.Value)
	return nil
}

// ManageMigrationResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                    {
//                    case EMPTY_VERSION:
//                      void;
//                    }
//
type ManageMigrationResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageMigrationResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageMigrationResultSuccessExt
func (u ManageMigrationResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageMigrationResultSuccessExt creates a new  ManageMigrationResultSuccessExt.
func NewManageMigrationResultSuccessExt(v LedgerVersion, value interface{}) (result ManageMigrationResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageMigrationResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//                    // reserved for future use
//                    union switch (LedgerVersion v)
//                    {
//                    case EMPTY_VERSION:
//                      void;
//                    }
//                    ext;
//                  }
//
type ManageMigrationResultSuccess struct {
	Ext ManageMigrationResultSuccessExt `json:"ext,omitempty"`
}

// ManageMigrationResult is an XDR Union defines as:
//
//   union ManageMigrationResult switch (ManageMigrationResultCode code)
//        {
//        case MANAGE_MIGRATION_SUCCESS :
//            struct {
//                    // reserved for future use
//                    union switch (LedgerVersion v)
//                    {
//                    case EMPTY_VERSION:
//                      void;
//                    }
//                    ext;
//                  } success;
//        default:
//            void;
//        };
//
type ManageMigrationResult struct {
	Code    ManageMigrationResultCode     `json:"code,omitempty"`
	Success *ManageMigrationResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageMigrationResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageMigrationResult
func (u ManageMigrationResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageMigrationResultCode(sw) {
	case ManageMigrationResultCodeManageMigrationSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManageMigrationResult creates a new  ManageMigrationResult.
func NewManageMigrationResult(code ManageMigrationResultCode, value interface{}) (result ManageMigrationResult, err error) {
	result.Code = code
	switch ManageMigrationResultCode(code) {
	case ManageMigrationResultCodeManageMigrationSuccess:
		tv, ok := value.(ManageMigrationResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageMigrationResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageMigrationResult) MustSuccess() ManageMigrationResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageMigrationResult) GetSuccess() (result ManageMigrationResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ManageOfferOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type ManageOfferOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageOfferOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageOfferOpExt
func (u ManageOfferOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageOfferOpExt creates a new  ManageOfferOpExt.
func NewManageOfferOpExt(v LedgerVersion, value interface{}) (result ManageOfferOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageOfferOp is an XDR Struct defines as:
//
//   struct ManageOfferOp
//        {
//            Asset selling;
//            Asset buying;
//            int64 amount; // amount being sold. if set to 0, delete the offer
//            Price price;  // price of thing being sold in terms of what you are buying
//
//            // 0=create a new offer, otherwise edit an existing offer
//            uint64 offerID;
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        };
//
type ManageOfferOp struct {
	Selling Asset            `json:"selling,omitempty"`
	Buying  Asset            `json:"buying,omitempty"`
	Amount  Int64            `json:"amount,omitempty"`
	Price   Price            `json:"price,omitempty"`
	OfferId Uint64           `json:"offerID,omitempty"`
	Ext     ManageOfferOpExt `json:"ext,omitempty"`
}

// ManageOfferResultCode is an XDR Enum defines as:
//
//   enum ManageOfferResultCode
//        {
//            // codes considered as "success" for the operation
//            MANAGE_OFFER_SUCCESS = 0,
//
//            // codes considered as "failure" for the operation
//            MANAGE_OFFER_MALFORMED = -1,     // generated offer would be invalid
//            MANAGE_OFFER_SELL_NO_TRUST = -2, // no trust line for what we're selling
//            MANAGE_OFFER_BUY_NO_TRUST = -3,  // no trust line for what we're buying
//            MANAGE_OFFER_SELL_NOT_AUTHORIZED = -4, // not authorized to sell
//            MANAGE_OFFER_BUY_NOT_AUTHORIZED = -5,  // not authorized to buy
//            MANAGE_OFFER_LINE_FULL = -6,      // can't receive more of what it's buying
//            MANAGE_OFFER_UNDERFUNDED = -7,    // doesn't hold what it's trying to sell
//            MANAGE_OFFER_CROSS_SELF = -8,     // would cross an offer from the same user
//            MANAGE_OFFER_SELL_NO_ISSUER = -9, // no issuer for what we're selling
//            MANAGE_OFFER_BUY_NO_ISSUER = -10, // no issuer for what we're buying
//
//            // update errors
//            MANAGE_OFFER_NOT_FOUND = -11, // offerID does not match an existing offer
//
//            MANAGE_OFFER_LOW_RESERVE = -12 // not enough funds to create a new Offer
//        };
//
type ManageOfferResultCode int32

const (
	ManageOfferResultCodeManageOfferSuccess           ManageOfferResultCode = 0
	ManageOfferResultCodeManageOfferMalformed         ManageOfferResultCode = -1
	ManageOfferResultCodeManageOfferSellNoTrust       ManageOfferResultCode = -2
	ManageOfferResultCodeManageOfferBuyNoTrust        ManageOfferResultCode = -3
	ManageOfferResultCodeManageOfferSellNotAuthorized ManageOfferResultCode = -4
	ManageOfferResultCodeManageOfferBuyNotAuthorized  ManageOfferResultCode = -5
	ManageOfferResultCodeManageOfferLineFull          ManageOfferResultCode = -6
	ManageOfferResultCodeManageOfferUnderfunded       ManageOfferResultCode = -7
	ManageOfferResultCodeManageOfferCrossSelf         ManageOfferResultCode = -8
	ManageOfferResultCodeManageOfferSellNoIssuer      ManageOfferResultCode = -9
	ManageOfferResultCodeManageOfferBuyNoIssuer       ManageOfferResultCode = -10
	ManageOfferResultCodeManageOfferNotFound          ManageOfferResultCode = -11
	ManageOfferResultCodeManageOfferLowReserve        ManageOfferResultCode = -12
)

var ManageOfferResultCodeAll = []ManageOfferResultCode{
	ManageOfferResultCodeManageOfferSuccess,
	ManageOfferResultCodeManageOfferMalformed,
	ManageOfferResultCodeManageOfferSellNoTrust,
	ManageOfferResultCodeManageOfferBuyNoTrust,
	ManageOfferResultCodeManageOfferSellNotAuthorized,
	ManageOfferResultCodeManageOfferBuyNotAuthorized,
	ManageOfferResultCodeManageOfferLineFull,
	ManageOfferResultCodeManageOfferUnderfunded,
	ManageOfferResultCodeManageOfferCrossSelf,
	ManageOfferResultCodeManageOfferSellNoIssuer,
	ManageOfferResultCodeManageOfferBuyNoIssuer,
	ManageOfferResultCodeManageOfferNotFound,
	ManageOfferResultCodeManageOfferLowReserve,
}

var manageOfferResultCodeMap = map[int32]string{
	0:   "ManageOfferResultCodeManageOfferSuccess",
	-1:  "ManageOfferResultCodeManageOfferMalformed",
	-2:  "ManageOfferResultCodeManageOfferSellNoTrust",
	-3:  "ManageOfferResultCodeManageOfferBuyNoTrust",
	-4:  "ManageOfferResultCodeManageOfferSellNotAuthorized",
	-5:  "ManageOfferResultCodeManageOfferBuyNotAuthorized",
	-6:  "ManageOfferResultCodeManageOfferLineFull",
	-7:  "ManageOfferResultCodeManageOfferUnderfunded",
	-8:  "ManageOfferResultCodeManageOfferCrossSelf",
	-9:  "ManageOfferResultCodeManageOfferSellNoIssuer",
	-10: "ManageOfferResultCodeManageOfferBuyNoIssuer",
	-11: "ManageOfferResultCodeManageOfferNotFound",
	-12: "ManageOfferResultCodeManageOfferLowReserve",
}

var manageOfferResultCodeShortMap = map[int32]string{
	0:   "manage_offer_success",
	-1:  "manage_offer_malformed",
	-2:  "manage_offer_sell_no_trust",
	-3:  "manage_offer_buy_no_trust",
	-4:  "manage_offer_sell_not_authorized",
	-5:  "manage_offer_buy_not_authorized",
	-6:  "manage_offer_line_full",
	-7:  "manage_offer_underfunded",
	-8:  "manage_offer_cross_self",
	-9:  "manage_offer_sell_no_issuer",
	-10: "manage_offer_buy_no_issuer",
	-11: "manage_offer_not_found",
	-12: "manage_offer_low_reserve",
}

var manageOfferResultCodeRevMap = map[string]int32{
	"ManageOfferResultCodeManageOfferSuccess":           0,
	"ManageOfferResultCodeManageOfferMalformed":         -1,
	"ManageOfferResultCodeManageOfferSellNoTrust":       -2,
	"ManageOfferResultCodeManageOfferBuyNoTrust":        -3,
	"ManageOfferResultCodeManageOfferSellNotAuthorized": -4,
	"ManageOfferResultCodeManageOfferBuyNotAuthorized":  -5,
	"ManageOfferResultCodeManageOfferLineFull":          -6,
	"ManageOfferResultCodeManageOfferUnderfunded":       -7,
	"ManageOfferResultCodeManageOfferCrossSelf":         -8,
	"ManageOfferResultCodeManageOfferSellNoIssuer":      -9,
	"ManageOfferResultCodeManageOfferBuyNoIssuer":       -10,
	"ManageOfferResultCodeManageOfferNotFound":          -11,
	"ManageOfferResultCodeManageOfferLowReserve":        -12,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageOfferResultCode
func (e ManageOfferResultCode) ValidEnum(v int32) bool {
	_, ok := manageOfferResultCodeMap[v]
	return ok
}
func (e ManageOfferResultCode) isFlag() bool {
	for i := len(ManageOfferResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageOfferResultCode(2) << uint64(len(ManageOfferResultCodeAll)-1) >> uint64(len(ManageOfferResultCodeAll)-i)
		if expected != ManageOfferResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageOfferResultCode) String() string {
	name, _ := manageOfferResultCodeMap[int32(e)]
	return name
}

func (e ManageOfferResultCode) ShortString() string {
	name, _ := manageOfferResultCodeShortMap[int32(e)]
	return name
}

func (e ManageOfferResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageOfferResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageOfferResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageOfferResultCode(t.Value)
	return nil
}

// ManageOfferEffect is an XDR Enum defines as:
//
//   enum ManageOfferEffect
//        {
//            MANAGE_OFFER_CREATED = 0,
//            MANAGE_OFFER_UPDATED = 1,
//            MANAGE_OFFER_DELETED = 2
//        };
//
type ManageOfferEffect int32

const (
	ManageOfferEffectManageOfferCreated ManageOfferEffect = 0
	ManageOfferEffectManageOfferUpdated ManageOfferEffect = 1
	ManageOfferEffectManageOfferDeleted ManageOfferEffect = 2
)

var ManageOfferEffectAll = []ManageOfferEffect{
	ManageOfferEffectManageOfferCreated,
	ManageOfferEffectManageOfferUpdated,
	ManageOfferEffectManageOfferDeleted,
}

var manageOfferEffectMap = map[int32]string{
	0: "ManageOfferEffectManageOfferCreated",
	1: "ManageOfferEffectManageOfferUpdated",
	2: "ManageOfferEffectManageOfferDeleted",
}

var manageOfferEffectShortMap = map[int32]string{
	0: "manage_offer_created",
	1: "manage_offer_updated",
	2: "manage_offer_deleted",
}

var manageOfferEffectRevMap = map[string]int32{
	"ManageOfferEffectManageOfferCreated": 0,
	"ManageOfferEffectManageOfferUpdated": 1,
	"ManageOfferEffectManageOfferDeleted": 2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageOfferEffect
func (e ManageOfferEffect) ValidEnum(v int32) bool {
	_, ok := manageOfferEffectMap[v]
	return ok
}
func (e ManageOfferEffect) isFlag() bool {
	for i := len(ManageOfferEffectAll) - 1; i >= 0; i-- {
		expected := ManageOfferEffect(2) << uint64(len(ManageOfferEffectAll)-1) >> uint64(len(ManageOfferEffectAll)-i)
		if expected != ManageOfferEffectAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageOfferEffect) String() string {
	name, _ := manageOfferEffectMap[int32(e)]
	return name
}

func (e ManageOfferEffect) ShortString() string {
	name, _ := manageOfferEffectShortMap[int32(e)]
	return name
}

func (e ManageOfferEffect) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageOfferEffectAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageOfferEffect) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageOfferEffect(t.Value)
	return nil
}

// ClaimOfferAtomExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type ClaimOfferAtomExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ClaimOfferAtomExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ClaimOfferAtomExt
func (u ClaimOfferAtomExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewClaimOfferAtomExt creates a new  ClaimOfferAtomExt.
func NewClaimOfferAtomExt(v LedgerVersion, value interface{}) (result ClaimOfferAtomExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ClaimOfferAtom is an XDR Struct defines as:
//
//   struct ClaimOfferAtom
//        {
//            // emitted to identify the offer
//            AccountID sellerID; // Account that owns the offer
//            uint64 offerID;
//
//            // amount and asset taken from the owner
//            Asset assetSold;
//            int64 amountSold;
//
//            // amount and asset sent to the owner
//            Asset assetBought;
//            int64 amountBought;
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        };
//
type ClaimOfferAtom struct {
	SellerId     AccountId         `json:"sellerID,omitempty"`
	OfferId      Uint64            `json:"offerID,omitempty"`
	AssetSold    Asset             `json:"assetSold,omitempty"`
	AmountSold   Int64             `json:"amountSold,omitempty"`
	AssetBought  Asset             `json:"assetBought,omitempty"`
	AmountBought Int64             `json:"amountBought,omitempty"`
	Ext          ClaimOfferAtomExt `json:"ext,omitempty"`
}

// ManageOfferSuccessResultOffer is an XDR NestedUnion defines as:
//
//   union switch (ManageOfferEffect effect)
//            {
//            case MANAGE_OFFER_CREATED:
//            case MANAGE_OFFER_UPDATED:
//                OfferEntry offer;
//            default:
//                void;
//            }
//
type ManageOfferSuccessResultOffer struct {
	Effect ManageOfferEffect `json:"effect,omitempty"`
	Offer  *OfferEntry       `json:"offer,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageOfferSuccessResultOffer) SwitchFieldName() string {
	return "Effect"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageOfferSuccessResultOffer
func (u ManageOfferSuccessResultOffer) ArmForSwitch(sw int32) (string, bool) {
	switch ManageOfferEffect(sw) {
	case ManageOfferEffectManageOfferCreated:
		return "Offer", true
	case ManageOfferEffectManageOfferUpdated:
		return "Offer", true
	default:
		return "", true
	}
}

// NewManageOfferSuccessResultOffer creates a new  ManageOfferSuccessResultOffer.
func NewManageOfferSuccessResultOffer(effect ManageOfferEffect, value interface{}) (result ManageOfferSuccessResultOffer, err error) {
	result.Effect = effect
	switch ManageOfferEffect(effect) {
	case ManageOfferEffectManageOfferCreated:
		tv, ok := value.(OfferEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be OfferEntry")
			return
		}
		result.Offer = &tv
	case ManageOfferEffectManageOfferUpdated:
		tv, ok := value.(OfferEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be OfferEntry")
			return
		}
		result.Offer = &tv
	default:
		// void
	}
	return
}

// MustOffer retrieves the Offer value from the union,
// panicing if the value is not set.
func (u ManageOfferSuccessResultOffer) MustOffer() OfferEntry {
	val, ok := u.GetOffer()

	if !ok {
		panic("arm Offer is not set")
	}

	return val
}

// GetOffer retrieves the Offer value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageOfferSuccessResultOffer) GetOffer() (result OfferEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Effect))

	if armName == "Offer" {
		result = *u.Offer
		ok = true
	}

	return
}

// ManageOfferSuccessResultExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type ManageOfferSuccessResultExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageOfferSuccessResultExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageOfferSuccessResultExt
func (u ManageOfferSuccessResultExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageOfferSuccessResultExt creates a new  ManageOfferSuccessResultExt.
func NewManageOfferSuccessResultExt(v LedgerVersion, value interface{}) (result ManageOfferSuccessResultExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageOfferSuccessResult is an XDR Struct defines as:
//
//   struct ManageOfferSuccessResult
//        {
//            // offers that got claimed while creating this offer
//            ClaimOfferAtom offersClaimed<>;
//
//            union switch (ManageOfferEffect effect)
//            {
//            case MANAGE_OFFER_CREATED:
//            case MANAGE_OFFER_UPDATED:
//                OfferEntry offer;
//            default:
//                void;
//            }
//            offer;
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        };
//
type ManageOfferSuccessResult struct {
	OffersClaimed []ClaimOfferAtom              `json:"offersClaimed,omitempty"`
	Offer         ManageOfferSuccessResultOffer `json:"offer,omitempty"`
	Ext           ManageOfferSuccessResultExt   `json:"ext,omitempty"`
}

// ManageOfferResult is an XDR Union defines as:
//
//   union ManageOfferResult switch (ManageOfferResultCode code)
//        {
//        case MANAGE_OFFER_SUCCESS:
//            ManageOfferSuccessResult success;
//        default:
//            void;
//        };
//
type ManageOfferResult struct {
	Code    ManageOfferResultCode     `json:"code,omitempty"`
	Success *ManageOfferSuccessResult `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageOfferResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageOfferResult
func (u ManageOfferResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageOfferResultCode(sw) {
	case ManageOfferResultCodeManageOfferSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManageOfferResult creates a new  ManageOfferResult.
func NewManageOfferResult(code ManageOfferResultCode, value interface{}) (result ManageOfferResult, err error) {
	result.Code = code
	switch ManageOfferResultCode(code) {
	case ManageOfferResultCodeManageOfferSuccess:
		tv, ok := value.(ManageOfferSuccessResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageOfferSuccessResult")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageOfferResult) MustSuccess() ManageOfferSuccessResult {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageOfferResult) GetSuccess() (result ManageOfferSuccessResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ManageParticipantAction is an XDR Enum defines as:
//
//   enum ManageParticipantAction
//        {
//            MANAGE_PARTICIPANT_ACTION_CHOOSE_WINNER = 0,
//            MANAGE_PARTICIPANT_ACTION_REJECT = 1,
//            MANAGE_PARTICIPANT_ACTION_APPROVE_BUY_NOW = 2,
//            MANAGE_PARTICIPANT_ACTION_APPROVE_DEPOSIT = 3
//        };
//
type ManageParticipantAction int32

const (
	ManageParticipantActionManageParticipantActionChooseWinner   ManageParticipantAction = 0
	ManageParticipantActionManageParticipantActionReject         ManageParticipantAction = 1
	ManageParticipantActionManageParticipantActionApproveBuyNow  ManageParticipantAction = 2
	ManageParticipantActionManageParticipantActionApproveDeposit ManageParticipantAction = 3
)

var ManageParticipantActionAll = []ManageParticipantAction{
	ManageParticipantActionManageParticipantActionChooseWinner,
	ManageParticipantActionManageParticipantActionReject,
	ManageParticipantActionManageParticipantActionApproveBuyNow,
	ManageParticipantActionManageParticipantActionApproveDeposit,
}

var manageParticipantActionMap = map[int32]string{
	0: "ManageParticipantActionManageParticipantActionChooseWinner",
	1: "ManageParticipantActionManageParticipantActionReject",
	2: "ManageParticipantActionManageParticipantActionApproveBuyNow",
	3: "ManageParticipantActionManageParticipantActionApproveDeposit",
}

var manageParticipantActionShortMap = map[int32]string{
	0: "manage_participant_action_choose_winner",
	1: "manage_participant_action_reject",
	2: "manage_participant_action_approve_buy_now",
	3: "manage_participant_action_approve_deposit",
}

var manageParticipantActionRevMap = map[string]int32{
	"ManageParticipantActionManageParticipantActionChooseWinner":   0,
	"ManageParticipantActionManageParticipantActionReject":         1,
	"ManageParticipantActionManageParticipantActionApproveBuyNow":  2,
	"ManageParticipantActionManageParticipantActionApproveDeposit": 3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageParticipantAction
func (e ManageParticipantAction) ValidEnum(v int32) bool {
	_, ok := manageParticipantActionMap[v]
	return ok
}
func (e ManageParticipantAction) isFlag() bool {
	for i := len(ManageParticipantActionAll) - 1; i >= 0; i-- {
		expected := ManageParticipantAction(2) << uint64(len(ManageParticipantActionAll)-1) >> uint64(len(ManageParticipantActionAll)-i)
		if expected != ManageParticipantActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageParticipantAction) String() string {
	name, _ := manageParticipantActionMap[int32(e)]
	return name
}

func (e ManageParticipantAction) ShortString() string {
	name, _ := manageParticipantActionShortMap[int32(e)]
	return name
}

func (e ManageParticipantAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageParticipantActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageParticipantAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageParticipantAction(t.Value)
	return nil
}

// ManageParticipantOpData is an XDR NestedUnion defines as:
//
//   union switch (ManageParticipantAction action)
//                {
//                case MANAGE_PARTICIPANT_ACTION_APPROVE_DEPOSIT:
//                    uint64 limit;
//                default:
//                    void;
//                }
//
type ManageParticipantOpData struct {
	Action ManageParticipantAction `json:"action,omitempty"`
	Limit  *Uint64                 `json:"limit,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageParticipantOpData) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageParticipantOpData
func (u ManageParticipantOpData) ArmForSwitch(sw int32) (string, bool) {
	switch ManageParticipantAction(sw) {
	case ManageParticipantActionManageParticipantActionApproveDeposit:
		return "Limit", true
	default:
		return "", true
	}
}

// NewManageParticipantOpData creates a new  ManageParticipantOpData.
func NewManageParticipantOpData(action ManageParticipantAction, value interface{}) (result ManageParticipantOpData, err error) {
	result.Action = action
	switch ManageParticipantAction(action) {
	case ManageParticipantActionManageParticipantActionApproveDeposit:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.Limit = &tv
	default:
		// void
	}
	return
}

// MustLimit retrieves the Limit value from the union,
// panicing if the value is not set.
func (u ManageParticipantOpData) MustLimit() Uint64 {
	val, ok := u.GetLimit()

	if !ok {
		panic("arm Limit is not set")
	}

	return val
}

// GetLimit retrieves the Limit value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageParticipantOpData) GetLimit() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "Limit" {
		result = *u.Limit
		ok = true
	}

	return
}

// ManageParticipantOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            case LEDGER_VERSION_LIVE_BID:
//                union switch (ManageParticipantAction action)
//                {
//                case MANAGE_PARTICIPANT_ACTION_APPROVE_DEPOSIT:
//                    uint64 limit;
//                default:
//                    void;
//                } data;
//            }
//
type ManageParticipantOpExt struct {
	V    LedgerVersion            `json:"v,omitempty"`
	Data *ManageParticipantOpData `json:"data,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageParticipantOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageParticipantOpExt
func (u ManageParticipantOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionLedgerVersionLiveBid:
		return "Data", true
	}
	return "-", false
}

// NewManageParticipantOpExt creates a new  ManageParticipantOpExt.
func NewManageParticipantOpExt(v LedgerVersion, value interface{}) (result ManageParticipantOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionLedgerVersionLiveBid:
		tv, ok := value.(ManageParticipantOpData)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageParticipantOpData")
			return
		}
		result.Data = &tv
	}
	return
}

// MustData retrieves the Data value from the union,
// panicing if the value is not set.
func (u ManageParticipantOpExt) MustData() ManageParticipantOpData {
	val, ok := u.GetData()

	if !ok {
		panic("arm Data is not set")
	}

	return val
}

// GetData retrieves the Data value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageParticipantOpExt) GetData() (result ManageParticipantOpData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "Data" {
		result = *u.Data
		ok = true
	}

	return
}

// ManageParticipantOp is an XDR Struct defines as:
//
//   struct ManageParticipantOp
//        {
//            ManageParticipantAction action;
//
//            uint64 lotID;
//            AccountID accountID;
//
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            case LEDGER_VERSION_LIVE_BID:
//                union switch (ManageParticipantAction action)
//                {
//                case MANAGE_PARTICIPANT_ACTION_APPROVE_DEPOSIT:
//                    uint64 limit;
//                default:
//                    void;
//                } data;
//            }
//            ext;
//        };
//
type ManageParticipantOp struct {
	Action    ManageParticipantAction `json:"action,omitempty"`
	LotId     Uint64                  `json:"lotID,omitempty"`
	AccountId AccountId               `json:"accountID,omitempty"`
	Ext       ManageParticipantOpExt  `json:"ext,omitempty"`
}

// ManageParticipantResultCode is an XDR Enum defines as:
//
//   enum ManageParticipantResultCode
//        {
//            MANAGE_PARTICIPANT_SUCCESS = 0,
//
//            MANAGE_PARTICIPANT_BIDDER_NOT_FOUND = -1,
//            MANAGE_PARTICIPANT_TOO_EARLY = -2,                  // auction hasn't started or currently active (not for buy now)
//            MANAGE_PARTICIPANT_AUCTION_CLOSED = -3,             // auction is closed
//            MANAGE_PARTICIPANT_ALREADY_HAS_WINNER = -4,         // lot has participant with state winner or buy now winner
//            MANAGE_PARTICIPANT_ALREADY_REJECTED = -5,           // lot has participant with state winner or buy now winner
//            MANAGE_PARTICIPANT_BAD_CHOICE = -6,                 // lot has better bidders, reject them first
//            MANAGE_PARTICIPANT_NOT_ALLOWED = -7,                // only platform can resolve auction
//            MANAGE_PARTICIPANT_DEPOSIT_IS_ALREADY_APPROVED = -8,// participant is already in state pending
//            MANAGE_PARTICIPANT_INVALID_VERSION = -9,            // ledger version is 'live bid' but field limit is missing
//            MANAGE_PARTICIPANT_MALFORMED = -10                  // ext is malformed
//        };
//
type ManageParticipantResultCode int32

const (
	ManageParticipantResultCodeManageParticipantSuccess                  ManageParticipantResultCode = 0
	ManageParticipantResultCodeManageParticipantBidderNotFound           ManageParticipantResultCode = -1
	ManageParticipantResultCodeManageParticipantTooEarly                 ManageParticipantResultCode = -2
	ManageParticipantResultCodeManageParticipantAuctionClosed            ManageParticipantResultCode = -3
	ManageParticipantResultCodeManageParticipantAlreadyHasWinner         ManageParticipantResultCode = -4
	ManageParticipantResultCodeManageParticipantAlreadyRejected          ManageParticipantResultCode = -5
	ManageParticipantResultCodeManageParticipantBadChoice                ManageParticipantResultCode = -6
	ManageParticipantResultCodeManageParticipantNotAllowed               ManageParticipantResultCode = -7
	ManageParticipantResultCodeManageParticipantDepositIsAlreadyApproved ManageParticipantResultCode = -8
	ManageParticipantResultCodeManageParticipantInvalidVersion           ManageParticipantResultCode = -9
	ManageParticipantResultCodeManageParticipantMalformed                ManageParticipantResultCode = -10
)

var ManageParticipantResultCodeAll = []ManageParticipantResultCode{
	ManageParticipantResultCodeManageParticipantSuccess,
	ManageParticipantResultCodeManageParticipantBidderNotFound,
	ManageParticipantResultCodeManageParticipantTooEarly,
	ManageParticipantResultCodeManageParticipantAuctionClosed,
	ManageParticipantResultCodeManageParticipantAlreadyHasWinner,
	ManageParticipantResultCodeManageParticipantAlreadyRejected,
	ManageParticipantResultCodeManageParticipantBadChoice,
	ManageParticipantResultCodeManageParticipantNotAllowed,
	ManageParticipantResultCodeManageParticipantDepositIsAlreadyApproved,
	ManageParticipantResultCodeManageParticipantInvalidVersion,
	ManageParticipantResultCodeManageParticipantMalformed,
}

var manageParticipantResultCodeMap = map[int32]string{
	0:   "ManageParticipantResultCodeManageParticipantSuccess",
	-1:  "ManageParticipantResultCodeManageParticipantBidderNotFound",
	-2:  "ManageParticipantResultCodeManageParticipantTooEarly",
	-3:  "ManageParticipantResultCodeManageParticipantAuctionClosed",
	-4:  "ManageParticipantResultCodeManageParticipantAlreadyHasWinner",
	-5:  "ManageParticipantResultCodeManageParticipantAlreadyRejected",
	-6:  "ManageParticipantResultCodeManageParticipantBadChoice",
	-7:  "ManageParticipantResultCodeManageParticipantNotAllowed",
	-8:  "ManageParticipantResultCodeManageParticipantDepositIsAlreadyApproved",
	-9:  "ManageParticipantResultCodeManageParticipantInvalidVersion",
	-10: "ManageParticipantResultCodeManageParticipantMalformed",
}

var manageParticipantResultCodeShortMap = map[int32]string{
	0:   "manage_participant_success",
	-1:  "manage_participant_bidder_not_found",
	-2:  "manage_participant_too_early",
	-3:  "manage_participant_auction_closed",
	-4:  "manage_participant_already_has_winner",
	-5:  "manage_participant_already_rejected",
	-6:  "manage_participant_bad_choice",
	-7:  "manage_participant_not_allowed",
	-8:  "manage_participant_deposit_is_already_approved",
	-9:  "manage_participant_invalid_version",
	-10: "manage_participant_malformed",
}

var manageParticipantResultCodeRevMap = map[string]int32{
	"ManageParticipantResultCodeManageParticipantSuccess":                  0,
	"ManageParticipantResultCodeManageParticipantBidderNotFound":           -1,
	"ManageParticipantResultCodeManageParticipantTooEarly":                 -2,
	"ManageParticipantResultCodeManageParticipantAuctionClosed":            -3,
	"ManageParticipantResultCodeManageParticipantAlreadyHasWinner":         -4,
	"ManageParticipantResultCodeManageParticipantAlreadyRejected":          -5,
	"ManageParticipantResultCodeManageParticipantBadChoice":                -6,
	"ManageParticipantResultCodeManageParticipantNotAllowed":               -7,
	"ManageParticipantResultCodeManageParticipantDepositIsAlreadyApproved": -8,
	"ManageParticipantResultCodeManageParticipantInvalidVersion":           -9,
	"ManageParticipantResultCodeManageParticipantMalformed":                -10,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageParticipantResultCode
func (e ManageParticipantResultCode) ValidEnum(v int32) bool {
	_, ok := manageParticipantResultCodeMap[v]
	return ok
}
func (e ManageParticipantResultCode) isFlag() bool {
	for i := len(ManageParticipantResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageParticipantResultCode(2) << uint64(len(ManageParticipantResultCodeAll)-1) >> uint64(len(ManageParticipantResultCodeAll)-i)
		if expected != ManageParticipantResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageParticipantResultCode) String() string {
	name, _ := manageParticipantResultCodeMap[int32(e)]
	return name
}

func (e ManageParticipantResultCode) ShortString() string {
	name, _ := manageParticipantResultCodeShortMap[int32(e)]
	return name
}

func (e ManageParticipantResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageParticipantResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageParticipantResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageParticipantResultCode(t.Value)
	return nil
}

// ManageParticipantSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type ManageParticipantSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageParticipantSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageParticipantSuccessExt
func (u ManageParticipantSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageParticipantSuccessExt creates a new  ManageParticipantSuccessExt.
func NewManageParticipantSuccessExt(v LedgerVersion, value interface{}) (result ManageParticipantSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageParticipantSuccess is an XDR Struct defines as:
//
//   struct ManageParticipantSuccess
//        {
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type ManageParticipantSuccess struct {
	Ext ManageParticipantSuccessExt `json:"ext,omitempty"`
}

// ManageParticipantResult is an XDR Union defines as:
//
//   union ManageParticipantResult switch (ManageParticipantResultCode code)
//        {
//        case MANAGE_PARTICIPANT_SUCCESS:
//            ManageParticipantSuccess success;
//        default:
//            void;
//        };
//
type ManageParticipantResult struct {
	Code    ManageParticipantResultCode `json:"code,omitempty"`
	Success *ManageParticipantSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageParticipantResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageParticipantResult
func (u ManageParticipantResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageParticipantResultCode(sw) {
	case ManageParticipantResultCodeManageParticipantSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManageParticipantResult creates a new  ManageParticipantResult.
func NewManageParticipantResult(code ManageParticipantResultCode, value interface{}) (result ManageParticipantResult, err error) {
	result.Code = code
	switch ManageParticipantResultCode(code) {
	case ManageParticipantResultCodeManageParticipantSuccess:
		tv, ok := value.(ManageParticipantSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageParticipantSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageParticipantResult) MustSuccess() ManageParticipantSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageParticipantResult) GetSuccess() (result ManageParticipantSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ManagePermissionsAction is an XDR Enum defines as:
//
//   enum ManagePermissionsAction
//        {
//            MANAGE_PERMISSIONS_ACTION_ADD_ADMIN = 0,
//            MANAGE_PERMISSIONS_ACTION_REMOVE_ADMIN = 1
//        };
//
type ManagePermissionsAction int32

const (
	ManagePermissionsActionManagePermissionsActionAddAdmin    ManagePermissionsAction = 0
	ManagePermissionsActionManagePermissionsActionRemoveAdmin ManagePermissionsAction = 1
)

var ManagePermissionsActionAll = []ManagePermissionsAction{
	ManagePermissionsActionManagePermissionsActionAddAdmin,
	ManagePermissionsActionManagePermissionsActionRemoveAdmin,
}

var managePermissionsActionMap = map[int32]string{
	0: "ManagePermissionsActionManagePermissionsActionAddAdmin",
	1: "ManagePermissionsActionManagePermissionsActionRemoveAdmin",
}

var managePermissionsActionShortMap = map[int32]string{
	0: "manage_permissions_action_add_admin",
	1: "manage_permissions_action_remove_admin",
}

var managePermissionsActionRevMap = map[string]int32{
	"ManagePermissionsActionManagePermissionsActionAddAdmin":    0,
	"ManagePermissionsActionManagePermissionsActionRemoveAdmin": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManagePermissionsAction
func (e ManagePermissionsAction) ValidEnum(v int32) bool {
	_, ok := managePermissionsActionMap[v]
	return ok
}
func (e ManagePermissionsAction) isFlag() bool {
	for i := len(ManagePermissionsActionAll) - 1; i >= 0; i-- {
		expected := ManagePermissionsAction(2) << uint64(len(ManagePermissionsActionAll)-1) >> uint64(len(ManagePermissionsActionAll)-i)
		if expected != ManagePermissionsActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManagePermissionsAction) String() string {
	name, _ := managePermissionsActionMap[int32(e)]
	return name
}

func (e ManagePermissionsAction) ShortString() string {
	name, _ := managePermissionsActionShortMap[int32(e)]
	return name
}

func (e ManagePermissionsAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManagePermissionsActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManagePermissionsAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManagePermissionsAction(t.Value)
	return nil
}

// ManagePermissionsOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type ManagePermissionsOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManagePermissionsOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManagePermissionsOpExt
func (u ManagePermissionsOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManagePermissionsOpExt creates a new  ManagePermissionsOpExt.
func NewManagePermissionsOpExt(v LedgerVersion, value interface{}) (result ManagePermissionsOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManagePermissionsOp is an XDR Struct defines as:
//
//   struct ManagePermissionsOp
//        {
//            AccountID target;                   // target admin account id
//            ManagePermissionsAction action;
//
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type ManagePermissionsOp struct {
	Target AccountId               `json:"target,omitempty"`
	Action ManagePermissionsAction `json:"action,omitempty"`
	Ext    ManagePermissionsOpExt  `json:"ext,omitempty"`
}

// ManagePermissionsResultCode is an XDR Enum defines as:
//
//   enum ManagePermissionsResultCode
//        {
//            MANAGE_PERMISSIONS_SUCCESS = 0,
//
//            MANAGE_PERMISSIONS_TARGET_NOT_FOUND = -1,
//            MANAGE_PERMISSIONS_ADMIN_IS_ALREADY_ADDED = -2,        // admin is already trusted
//            MANAGE_PERMISSIONS_NO_ADMIN_TO_REMOVE = -3             // admin has not permissions from source platform
//        };
//
type ManagePermissionsResultCode int32

const (
	ManagePermissionsResultCodeManagePermissionsSuccess             ManagePermissionsResultCode = 0
	ManagePermissionsResultCodeManagePermissionsTargetNotFound      ManagePermissionsResultCode = -1
	ManagePermissionsResultCodeManagePermissionsAdminIsAlreadyAdded ManagePermissionsResultCode = -2
	ManagePermissionsResultCodeManagePermissionsNoAdminToRemove     ManagePermissionsResultCode = -3
)

var ManagePermissionsResultCodeAll = []ManagePermissionsResultCode{
	ManagePermissionsResultCodeManagePermissionsSuccess,
	ManagePermissionsResultCodeManagePermissionsTargetNotFound,
	ManagePermissionsResultCodeManagePermissionsAdminIsAlreadyAdded,
	ManagePermissionsResultCodeManagePermissionsNoAdminToRemove,
}

var managePermissionsResultCodeMap = map[int32]string{
	0:  "ManagePermissionsResultCodeManagePermissionsSuccess",
	-1: "ManagePermissionsResultCodeManagePermissionsTargetNotFound",
	-2: "ManagePermissionsResultCodeManagePermissionsAdminIsAlreadyAdded",
	-3: "ManagePermissionsResultCodeManagePermissionsNoAdminToRemove",
}

var managePermissionsResultCodeShortMap = map[int32]string{
	0:  "manage_permissions_success",
	-1: "manage_permissions_target_not_found",
	-2: "manage_permissions_admin_is_already_added",
	-3: "manage_permissions_no_admin_to_remove",
}

var managePermissionsResultCodeRevMap = map[string]int32{
	"ManagePermissionsResultCodeManagePermissionsSuccess":             0,
	"ManagePermissionsResultCodeManagePermissionsTargetNotFound":      -1,
	"ManagePermissionsResultCodeManagePermissionsAdminIsAlreadyAdded": -2,
	"ManagePermissionsResultCodeManagePermissionsNoAdminToRemove":     -3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManagePermissionsResultCode
func (e ManagePermissionsResultCode) ValidEnum(v int32) bool {
	_, ok := managePermissionsResultCodeMap[v]
	return ok
}
func (e ManagePermissionsResultCode) isFlag() bool {
	for i := len(ManagePermissionsResultCodeAll) - 1; i >= 0; i-- {
		expected := ManagePermissionsResultCode(2) << uint64(len(ManagePermissionsResultCodeAll)-1) >> uint64(len(ManagePermissionsResultCodeAll)-i)
		if expected != ManagePermissionsResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManagePermissionsResultCode) String() string {
	name, _ := managePermissionsResultCodeMap[int32(e)]
	return name
}

func (e ManagePermissionsResultCode) ShortString() string {
	name, _ := managePermissionsResultCodeShortMap[int32(e)]
	return name
}

func (e ManagePermissionsResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManagePermissionsResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManagePermissionsResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManagePermissionsResultCode(t.Value)
	return nil
}

// ManagePermissionsSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type ManagePermissionsSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManagePermissionsSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManagePermissionsSuccessExt
func (u ManagePermissionsSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManagePermissionsSuccessExt creates a new  ManagePermissionsSuccessExt.
func NewManagePermissionsSuccessExt(v LedgerVersion, value interface{}) (result ManagePermissionsSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManagePermissionsSuccess is an XDR Struct defines as:
//
//   struct ManagePermissionsSuccess
//        {
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type ManagePermissionsSuccess struct {
	Ext ManagePermissionsSuccessExt `json:"ext,omitempty"`
}

// ManagePermissionsResult is an XDR Union defines as:
//
//   union ManagePermissionsResult switch (ManagePermissionsResultCode code)
//        {
//        case MANAGE_PERMISSIONS_SUCCESS:
//            ManagePermissionsSuccess success;
//        default:
//            void;
//        };
//
type ManagePermissionsResult struct {
	Code    ManagePermissionsResultCode `json:"code,omitempty"`
	Success *ManagePermissionsSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManagePermissionsResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManagePermissionsResult
func (u ManagePermissionsResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManagePermissionsResultCode(sw) {
	case ManagePermissionsResultCodeManagePermissionsSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManagePermissionsResult creates a new  ManagePermissionsResult.
func NewManagePermissionsResult(code ManagePermissionsResultCode, value interface{}) (result ManagePermissionsResult, err error) {
	result.Code = code
	switch ManagePermissionsResultCode(code) {
	case ManagePermissionsResultCodeManagePermissionsSuccess:
		tv, ok := value.(ManagePermissionsSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManagePermissionsSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManagePermissionsResult) MustSuccess() ManagePermissionsSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManagePermissionsResult) GetSuccess() (result ManagePermissionsSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// PathPaymentOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type PathPaymentOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PathPaymentOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PathPaymentOpExt
func (u PathPaymentOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPathPaymentOpExt creates a new  PathPaymentOpExt.
func NewPathPaymentOpExt(v LedgerVersion, value interface{}) (result PathPaymentOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PathPaymentOp is an XDR Struct defines as:
//
//   struct PathPaymentOp
//        {
//            Asset sendAsset; // asset we pay with
//            int64 sendMax;   // the maximum amount of sendAsset to
//                             // send (excluding fees).
//                             // The operation will fail if can't be met
//
//            AccountID destination; // recipient of the payment
//            Asset destAsset;       // what they end up with
//            int64 destAmount;      // amount they end up with
//
//            Asset path<5>; // additional hops it must go through to get there
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        };
//
type PathPaymentOp struct {
	SendAsset   Asset            `json:"sendAsset,omitempty"`
	SendMax     Int64            `json:"sendMax,omitempty"`
	Destination AccountId        `json:"destination,omitempty"`
	DestAsset   Asset            `json:"destAsset,omitempty"`
	DestAmount  Int64            `json:"destAmount,omitempty"`
	Path        []Asset          `json:"path,omitempty" xdrmaxsize:"5"`
	Ext         PathPaymentOpExt `json:"ext,omitempty"`
}

// PathPaymentResultCode is an XDR Enum defines as:
//
//   enum PathPaymentResultCode
//        {
//            // codes considered as "success" for the operation
//            PATH_PAYMENT_SUCCESS = 0, // success
//
//            // codes considered as "failure" for the operation
//            PATH_PAYMENT_MALFORMED = -1,          // bad input
//            PATH_PAYMENT_UNDERFUNDED = -2,        // not enough funds in source account
//            PATH_PAYMENT_SRC_NO_TRUST = -3,       // no trust line on source account
//            PATH_PAYMENT_SRC_NOT_AUTHORIZED = -4, // source not authorized to transfer
//            PATH_PAYMENT_NO_DESTINATION = -5,     // destination account does not exist
//            PATH_PAYMENT_NO_TRUST = -6,           // dest missing a trust line for asset
//            PATH_PAYMENT_NOT_AUTHORIZED = -7,     // dest not authorized to hold asset
//            PATH_PAYMENT_LINE_FULL = -8,          // dest would go above their limit
//            PATH_PAYMENT_NO_ISSUER = -9,          // missing issuer on one asset
//            PATH_PAYMENT_TOO_FEW_OFFERS = -10,    // not enough offers to satisfy path
//            PATH_PAYMENT_OFFER_CROSS_SELF = -11,  // would cross one of its own offers
//            PATH_PAYMENT_OVER_SENDMAX = -12       // could not satisfy sendmax
//        };
//
type PathPaymentResultCode int32

const (
	PathPaymentResultCodePathPaymentSuccess          PathPaymentResultCode = 0
	PathPaymentResultCodePathPaymentMalformed        PathPaymentResultCode = -1
	PathPaymentResultCodePathPaymentUnderfunded      PathPaymentResultCode = -2
	PathPaymentResultCodePathPaymentSrcNoTrust       PathPaymentResultCode = -3
	PathPaymentResultCodePathPaymentSrcNotAuthorized PathPaymentResultCode = -4
	PathPaymentResultCodePathPaymentNoDestination    PathPaymentResultCode = -5
	PathPaymentResultCodePathPaymentNoTrust          PathPaymentResultCode = -6
	PathPaymentResultCodePathPaymentNotAuthorized    PathPaymentResultCode = -7
	PathPaymentResultCodePathPaymentLineFull         PathPaymentResultCode = -8
	PathPaymentResultCodePathPaymentNoIssuer         PathPaymentResultCode = -9
	PathPaymentResultCodePathPaymentTooFewOffers     PathPaymentResultCode = -10
	PathPaymentResultCodePathPaymentOfferCrossSelf   PathPaymentResultCode = -11
	PathPaymentResultCodePathPaymentOverSendmax      PathPaymentResultCode = -12
)

var PathPaymentResultCodeAll = []PathPaymentResultCode{
	PathPaymentResultCodePathPaymentSuccess,
	PathPaymentResultCodePathPaymentMalformed,
	PathPaymentResultCodePathPaymentUnderfunded,
	PathPaymentResultCodePathPaymentSrcNoTrust,
	PathPaymentResultCodePathPaymentSrcNotAuthorized,
	PathPaymentResultCodePathPaymentNoDestination,
	PathPaymentResultCodePathPaymentNoTrust,
	PathPaymentResultCodePathPaymentNotAuthorized,
	PathPaymentResultCodePathPaymentLineFull,
	PathPaymentResultCodePathPaymentNoIssuer,
	PathPaymentResultCodePathPaymentTooFewOffers,
	PathPaymentResultCodePathPaymentOfferCrossSelf,
	PathPaymentResultCodePathPaymentOverSendmax,
}

var pathPaymentResultCodeMap = map[int32]string{
	0:   "PathPaymentResultCodePathPaymentSuccess",
	-1:  "PathPaymentResultCodePathPaymentMalformed",
	-2:  "PathPaymentResultCodePathPaymentUnderfunded",
	-3:  "PathPaymentResultCodePathPaymentSrcNoTrust",
	-4:  "PathPaymentResultCodePathPaymentSrcNotAuthorized",
	-5:  "PathPaymentResultCodePathPaymentNoDestination",
	-6:  "PathPaymentResultCodePathPaymentNoTrust",
	-7:  "PathPaymentResultCodePathPaymentNotAuthorized",
	-8:  "PathPaymentResultCodePathPaymentLineFull",
	-9:  "PathPaymentResultCodePathPaymentNoIssuer",
	-10: "PathPaymentResultCodePathPaymentTooFewOffers",
	-11: "PathPaymentResultCodePathPaymentOfferCrossSelf",
	-12: "PathPaymentResultCodePathPaymentOverSendmax",
}

var pathPaymentResultCodeShortMap = map[int32]string{
	0:   "path_payment_success",
	-1:  "path_payment_malformed",
	-2:  "path_payment_underfunded",
	-3:  "path_payment_src_no_trust",
	-4:  "path_payment_src_not_authorized",
	-5:  "path_payment_no_destination",
	-6:  "path_payment_no_trust",
	-7:  "path_payment_not_authorized",
	-8:  "path_payment_line_full",
	-9:  "path_payment_no_issuer",
	-10: "path_payment_too_few_offers",
	-11: "path_payment_offer_cross_self",
	-12: "path_payment_over_sendmax",
}

var pathPaymentResultCodeRevMap = map[string]int32{
	"PathPaymentResultCodePathPaymentSuccess":          0,
	"PathPaymentResultCodePathPaymentMalformed":        -1,
	"PathPaymentResultCodePathPaymentUnderfunded":      -2,
	"PathPaymentResultCodePathPaymentSrcNoTrust":       -3,
	"PathPaymentResultCodePathPaymentSrcNotAuthorized": -4,
	"PathPaymentResultCodePathPaymentNoDestination":    -5,
	"PathPaymentResultCodePathPaymentNoTrust":          -6,
	"PathPaymentResultCodePathPaymentNotAuthorized":    -7,
	"PathPaymentResultCodePathPaymentLineFull":         -8,
	"PathPaymentResultCodePathPaymentNoIssuer":         -9,
	"PathPaymentResultCodePathPaymentTooFewOffers":     -10,
	"PathPaymentResultCodePathPaymentOfferCrossSelf":   -11,
	"PathPaymentResultCodePathPaymentOverSendmax":      -12,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for PathPaymentResultCode
func (e PathPaymentResultCode) ValidEnum(v int32) bool {
	_, ok := pathPaymentResultCodeMap[v]
	return ok
}
func (e PathPaymentResultCode) isFlag() bool {
	for i := len(PathPaymentResultCodeAll) - 1; i >= 0; i-- {
		expected := PathPaymentResultCode(2) << uint64(len(PathPaymentResultCodeAll)-1) >> uint64(len(PathPaymentResultCodeAll)-i)
		if expected != PathPaymentResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e PathPaymentResultCode) String() string {
	name, _ := pathPaymentResultCodeMap[int32(e)]
	return name
}

func (e PathPaymentResultCode) ShortString() string {
	name, _ := pathPaymentResultCodeShortMap[int32(e)]
	return name
}

func (e PathPaymentResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range PathPaymentResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *PathPaymentResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = PathPaymentResultCode(t.Value)
	return nil
}

// SimplePaymentResultExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type SimplePaymentResultExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SimplePaymentResultExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SimplePaymentResultExt
func (u SimplePaymentResultExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSimplePaymentResultExt creates a new  SimplePaymentResultExt.
func NewSimplePaymentResultExt(v LedgerVersion, value interface{}) (result SimplePaymentResultExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SimplePaymentResult is an XDR Struct defines as:
//
//   struct SimplePaymentResult
//        {
//            AccountID destination;
//            Asset asset;
//            int64 amount;
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        };
//
type SimplePaymentResult struct {
	Destination AccountId              `json:"destination,omitempty"`
	Asset       Asset                  `json:"asset,omitempty"`
	Amount      Int64                  `json:"amount,omitempty"`
	Ext         SimplePaymentResultExt `json:"ext,omitempty"`
}

// PathPaymentResultSuccess is an XDR NestedStruct defines as:
//
//   struct
//            {
//                ClaimOfferAtom offers<>;
//                SimplePaymentResult last;
//            }
//
type PathPaymentResultSuccess struct {
	Offers []ClaimOfferAtom    `json:"offers,omitempty"`
	Last   SimplePaymentResult `json:"last,omitempty"`
}

// PathPaymentResult is an XDR Union defines as:
//
//   union PathPaymentResult switch (PathPaymentResultCode code)
//        {
//        case PATH_PAYMENT_SUCCESS:
//            struct
//            {
//                ClaimOfferAtom offers<>;
//                SimplePaymentResult last;
//            } success;
//        case PATH_PAYMENT_NO_ISSUER:
//            Asset noIssuer; // the asset that caused the error
//        default:
//            void;
//        };
//
type PathPaymentResult struct {
	Code     PathPaymentResultCode     `json:"code,omitempty"`
	Success  *PathPaymentResultSuccess `json:"success,omitempty"`
	NoIssuer *Asset                    `json:"noIssuer,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PathPaymentResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PathPaymentResult
func (u PathPaymentResult) ArmForSwitch(sw int32) (string, bool) {
	switch PathPaymentResultCode(sw) {
	case PathPaymentResultCodePathPaymentSuccess:
		return "Success", true
	case PathPaymentResultCodePathPaymentNoIssuer:
		return "NoIssuer", true
	default:
		return "", true
	}
}

// NewPathPaymentResult creates a new  PathPaymentResult.
func NewPathPaymentResult(code PathPaymentResultCode, value interface{}) (result PathPaymentResult, err error) {
	result.Code = code
	switch PathPaymentResultCode(code) {
	case PathPaymentResultCodePathPaymentSuccess:
		tv, ok := value.(PathPaymentResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be PathPaymentResultSuccess")
			return
		}
		result.Success = &tv
	case PathPaymentResultCodePathPaymentNoIssuer:
		tv, ok := value.(Asset)
		if !ok {
			err = fmt.Errorf("invalid value, must be Asset")
			return
		}
		result.NoIssuer = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u PathPaymentResult) MustSuccess() PathPaymentResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PathPaymentResult) GetSuccess() (result PathPaymentResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// MustNoIssuer retrieves the NoIssuer value from the union,
// panicing if the value is not set.
func (u PathPaymentResult) MustNoIssuer() Asset {
	val, ok := u.GetNoIssuer()

	if !ok {
		panic("arm NoIssuer is not set")
	}

	return val
}

// GetNoIssuer retrieves the NoIssuer value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PathPaymentResult) GetNoIssuer() (result Asset, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "NoIssuer" {
		result = *u.NoIssuer
		ok = true
	}

	return
}

// PaymentOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type PaymentOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PaymentOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PaymentOpExt
func (u PaymentOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPaymentOpExt creates a new  PaymentOpExt.
func NewPaymentOpExt(v LedgerVersion, value interface{}) (result PaymentOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PaymentOp is an XDR Struct defines as:
//
//   struct PaymentOp
//        {
//            AccountID destination; // recipient of the payment
//            Asset asset;           // what they end up with
//            int64 amount;          // amount they end up with
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        };
//
type PaymentOp struct {
	Destination AccountId    `json:"destination,omitempty"`
	Asset       Asset        `json:"asset,omitempty"`
	Amount      Int64        `json:"amount,omitempty"`
	Ext         PaymentOpExt `json:"ext,omitempty"`
}

// PaymentResultCode is an XDR Enum defines as:
//
//   enum PaymentResultCode
//        {
//            // codes considered as "success" for the operation
//            PAYMENT_SUCCESS = 0, // payment successfuly completed
//
//            // codes considered as "failure" for the operation
//            PAYMENT_MALFORMED = -1,          // bad input
//            PAYMENT_UNDERFUNDED = -2,        // not enough funds in source account
//            PAYMENT_SRC_NO_TRUST = -3,       // no trust line on source account
//            PAYMENT_SRC_NOT_AUTHORIZED = -4, // source not authorized to transfer
//            PAYMENT_NO_DESTINATION = -5,     // destination account does not exist
//            PAYMENT_NO_TRUST = -6,       // destination missing a trust line for asset
//            PAYMENT_NOT_AUTHORIZED = -7, // destination not authorized to hold asset
//            PAYMENT_LINE_FULL = -8,      // destination would go above their limit
//            PAYMENT_NO_ISSUER = -9       // missing issuer on asset
//        };
//
type PaymentResultCode int32

const (
	PaymentResultCodePaymentSuccess          PaymentResultCode = 0
	PaymentResultCodePaymentMalformed        PaymentResultCode = -1
	PaymentResultCodePaymentUnderfunded      PaymentResultCode = -2
	PaymentResultCodePaymentSrcNoTrust       PaymentResultCode = -3
	PaymentResultCodePaymentSrcNotAuthorized PaymentResultCode = -4
	PaymentResultCodePaymentNoDestination    PaymentResultCode = -5
	PaymentResultCodePaymentNoTrust          PaymentResultCode = -6
	PaymentResultCodePaymentNotAuthorized    PaymentResultCode = -7
	PaymentResultCodePaymentLineFull         PaymentResultCode = -8
	PaymentResultCodePaymentNoIssuer         PaymentResultCode = -9
)

var PaymentResultCodeAll = []PaymentResultCode{
	PaymentResultCodePaymentSuccess,
	PaymentResultCodePaymentMalformed,
	PaymentResultCodePaymentUnderfunded,
	PaymentResultCodePaymentSrcNoTrust,
	PaymentResultCodePaymentSrcNotAuthorized,
	PaymentResultCodePaymentNoDestination,
	PaymentResultCodePaymentNoTrust,
	PaymentResultCodePaymentNotAuthorized,
	PaymentResultCodePaymentLineFull,
	PaymentResultCodePaymentNoIssuer,
}

var paymentResultCodeMap = map[int32]string{
	0:  "PaymentResultCodePaymentSuccess",
	-1: "PaymentResultCodePaymentMalformed",
	-2: "PaymentResultCodePaymentUnderfunded",
	-3: "PaymentResultCodePaymentSrcNoTrust",
	-4: "PaymentResultCodePaymentSrcNotAuthorized",
	-5: "PaymentResultCodePaymentNoDestination",
	-6: "PaymentResultCodePaymentNoTrust",
	-7: "PaymentResultCodePaymentNotAuthorized",
	-8: "PaymentResultCodePaymentLineFull",
	-9: "PaymentResultCodePaymentNoIssuer",
}

var paymentResultCodeShortMap = map[int32]string{
	0:  "payment_success",
	-1: "payment_malformed",
	-2: "payment_underfunded",
	-3: "payment_src_no_trust",
	-4: "payment_src_not_authorized",
	-5: "payment_no_destination",
	-6: "payment_no_trust",
	-7: "payment_not_authorized",
	-8: "payment_line_full",
	-9: "payment_no_issuer",
}

var paymentResultCodeRevMap = map[string]int32{
	"PaymentResultCodePaymentSuccess":          0,
	"PaymentResultCodePaymentMalformed":        -1,
	"PaymentResultCodePaymentUnderfunded":      -2,
	"PaymentResultCodePaymentSrcNoTrust":       -3,
	"PaymentResultCodePaymentSrcNotAuthorized": -4,
	"PaymentResultCodePaymentNoDestination":    -5,
	"PaymentResultCodePaymentNoTrust":          -6,
	"PaymentResultCodePaymentNotAuthorized":    -7,
	"PaymentResultCodePaymentLineFull":         -8,
	"PaymentResultCodePaymentNoIssuer":         -9,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for PaymentResultCode
func (e PaymentResultCode) ValidEnum(v int32) bool {
	_, ok := paymentResultCodeMap[v]
	return ok
}
func (e PaymentResultCode) isFlag() bool {
	for i := len(PaymentResultCodeAll) - 1; i >= 0; i-- {
		expected := PaymentResultCode(2) << uint64(len(PaymentResultCodeAll)-1) >> uint64(len(PaymentResultCodeAll)-i)
		if expected != PaymentResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e PaymentResultCode) String() string {
	name, _ := paymentResultCodeMap[int32(e)]
	return name
}

func (e PaymentResultCode) ShortString() string {
	name, _ := paymentResultCodeShortMap[int32(e)]
	return name
}

func (e PaymentResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range PaymentResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *PaymentResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = PaymentResultCode(t.Value)
	return nil
}

// PaymentResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                        {
//                        case EMPTY_VERSION:
//                          void;
//                        }
//
type PaymentResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PaymentResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PaymentResultSuccessExt
func (u PaymentResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPaymentResultSuccessExt creates a new  PaymentResultSuccessExt.
func NewPaymentResultSuccessExt(v LedgerVersion, value interface{}) (result PaymentResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PaymentResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//                        // reserved for future use
//                        union switch (LedgerVersion v)
//                        {
//                        case EMPTY_VERSION:
//                          void;
//                        }
//                        ext;
//                      }
//
type PaymentResultSuccess struct {
	Ext PaymentResultSuccessExt `json:"ext,omitempty"`
}

// PaymentResult is an XDR Union defines as:
//
//   union PaymentResult switch (PaymentResultCode code)
//        {
//        case PAYMENT_SUCCESS:
//            struct {
//                        // reserved for future use
//                        union switch (LedgerVersion v)
//                        {
//                        case EMPTY_VERSION:
//                          void;
//                        }
//                        ext;
//                      } success;
//        default:
//            void;
//        };
//
type PaymentResult struct {
	Code    PaymentResultCode     `json:"code,omitempty"`
	Success *PaymentResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PaymentResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PaymentResult
func (u PaymentResult) ArmForSwitch(sw int32) (string, bool) {
	switch PaymentResultCode(sw) {
	case PaymentResultCodePaymentSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewPaymentResult creates a new  PaymentResult.
func NewPaymentResult(code PaymentResultCode, value interface{}) (result PaymentResult, err error) {
	result.Code = code
	switch PaymentResultCode(code) {
	case PaymentResultCodePaymentSuccess:
		tv, ok := value.(PaymentResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be PaymentResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u PaymentResult) MustSuccess() PaymentResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PaymentResult) GetSuccess() (result PaymentResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// RecoverOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type RecoverOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RecoverOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RecoverOpExt
func (u RecoverOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewRecoverOpExt creates a new  RecoverOpExt.
func NewRecoverOpExt(v LedgerVersion, value interface{}) (result RecoverOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// RecoverOp is an XDR Struct defines as:
//
//   struct RecoverOp
//        {
//            AccountID account;
//            AccountID newSigner;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type RecoverOp struct {
	Account   AccountId    `json:"account,omitempty"`
	NewSigner AccountId    `json:"newSigner,omitempty"`
	Ext       RecoverOpExt `json:"ext,omitempty"`
}

// RecoverResultCode is an XDR Enum defines as:
//
//   enum RecoverResultCode
//        {
//            // codes considered as "success" for the operation
//            RECOVER_SUCCESS = 0,
//
//            // codes considered as "failure" for the operation
//            RECOVER_MALFORMED = -1,
//            RECOVER_IS_NOT_PLATFORM = -2,
//            RECOVER_ACCOUNT_NOT_FOUND = -3,
//            RECOVER_SIGNER_ALREADY_EXISTS = -4,
//            RECOVER_ACCOUNT_TYPE_NOT_ALLOWED = -5
//        };
//
type RecoverResultCode int32

const (
	RecoverResultCodeRecoverSuccess               RecoverResultCode = 0
	RecoverResultCodeRecoverMalformed             RecoverResultCode = -1
	RecoverResultCodeRecoverIsNotPlatform         RecoverResultCode = -2
	RecoverResultCodeRecoverAccountNotFound       RecoverResultCode = -3
	RecoverResultCodeRecoverSignerAlreadyExists   RecoverResultCode = -4
	RecoverResultCodeRecoverAccountTypeNotAllowed RecoverResultCode = -5
)

var RecoverResultCodeAll = []RecoverResultCode{
	RecoverResultCodeRecoverSuccess,
	RecoverResultCodeRecoverMalformed,
	RecoverResultCodeRecoverIsNotPlatform,
	RecoverResultCodeRecoverAccountNotFound,
	RecoverResultCodeRecoverSignerAlreadyExists,
	RecoverResultCodeRecoverAccountTypeNotAllowed,
}

var recoverResultCodeMap = map[int32]string{
	0:  "RecoverResultCodeRecoverSuccess",
	-1: "RecoverResultCodeRecoverMalformed",
	-2: "RecoverResultCodeRecoverIsNotPlatform",
	-3: "RecoverResultCodeRecoverAccountNotFound",
	-4: "RecoverResultCodeRecoverSignerAlreadyExists",
	-5: "RecoverResultCodeRecoverAccountTypeNotAllowed",
}

var recoverResultCodeShortMap = map[int32]string{
	0:  "recover_success",
	-1: "recover_malformed",
	-2: "recover_is_not_platform",
	-3: "recover_account_not_found",
	-4: "recover_signer_already_exists",
	-5: "recover_account_type_not_allowed",
}

var recoverResultCodeRevMap = map[string]int32{
	"RecoverResultCodeRecoverSuccess":               0,
	"RecoverResultCodeRecoverMalformed":             -1,
	"RecoverResultCodeRecoverIsNotPlatform":         -2,
	"RecoverResultCodeRecoverAccountNotFound":       -3,
	"RecoverResultCodeRecoverSignerAlreadyExists":   -4,
	"RecoverResultCodeRecoverAccountTypeNotAllowed": -5,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for RecoverResultCode
func (e RecoverResultCode) ValidEnum(v int32) bool {
	_, ok := recoverResultCodeMap[v]
	return ok
}
func (e RecoverResultCode) isFlag() bool {
	for i := len(RecoverResultCodeAll) - 1; i >= 0; i-- {
		expected := RecoverResultCode(2) << uint64(len(RecoverResultCodeAll)-1) >> uint64(len(RecoverResultCodeAll)-i)
		if expected != RecoverResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e RecoverResultCode) String() string {
	name, _ := recoverResultCodeMap[int32(e)]
	return name
}

func (e RecoverResultCode) ShortString() string {
	name, _ := recoverResultCodeShortMap[int32(e)]
	return name
}

func (e RecoverResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range RecoverResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *RecoverResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = RecoverResultCode(t.Value)
	return nil
}

// RecoverSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type RecoverSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RecoverSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RecoverSuccessExt
func (u RecoverSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewRecoverSuccessExt creates a new  RecoverSuccessExt.
func NewRecoverSuccessExt(v LedgerVersion, value interface{}) (result RecoverSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// RecoverSuccess is an XDR Struct defines as:
//
//   struct RecoverSuccess
//        {
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type RecoverSuccess struct {
	Ext RecoverSuccessExt `json:"ext,omitempty"`
}

// RecoverResult is an XDR Union defines as:
//
//   union RecoverResult switch (RecoverResultCode code)
//        {
//        case RECOVER_SUCCESS:
//            RecoverSuccess success;
//        default:
//            void;
//        };
//
type RecoverResult struct {
	Code    RecoverResultCode `json:"code,omitempty"`
	Success *RecoverSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RecoverResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RecoverResult
func (u RecoverResult) ArmForSwitch(sw int32) (string, bool) {
	switch RecoverResultCode(sw) {
	case RecoverResultCodeRecoverSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewRecoverResult creates a new  RecoverResult.
func NewRecoverResult(code RecoverResultCode, value interface{}) (result RecoverResult, err error) {
	result.Code = code
	switch RecoverResultCode(code) {
	case RecoverResultCodeRecoverSuccess:
		tv, ok := value.(RecoverSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be RecoverSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u RecoverResult) MustSuccess() RecoverSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RecoverResult) GetSuccess() (result RecoverSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// RequestBuyNowOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type RequestBuyNowOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RequestBuyNowOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RequestBuyNowOpExt
func (u RequestBuyNowOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewRequestBuyNowOpExt creates a new  RequestBuyNowOpExt.
func NewRequestBuyNowOpExt(v LedgerVersion, value interface{}) (result RequestBuyNowOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// RequestBuyNowOp is an XDR Struct defines as:
//
//   struct RequestBuyNowOp
//        {
//            uint64 lotID;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type RequestBuyNowOp struct {
	LotId Uint64             `json:"lotID,omitempty"`
	Ext   RequestBuyNowOpExt `json:"ext,omitempty"`
}

// RequestBuyNowResultCode is an XDR Enum defines as:
//
//   enum RequestBuyNowResultCode
//        {
//            REQUEST_BUY_NOW_SUCCESS = 0,
//            REQUEST_BUY_NOW_LOT_NOT_FOUND = -1,
//            REQUEST_BUY_NOW_UNSUPPORTED = -2,       // lot doesn't support buy now
//            REQUEST_BUY_NOW_UNREGISTERED = -3,      // register to lot first, before submitting buy now
//            REQUEST_BUY_NOW_TOO_LATE = -4,          // auction is not active
//            REQUEST_BUY_NOW_ALREADY_REQUESTED = -5, // source already requested buy now
//            REQUEST_BUY_NOW_ALREADY_TAKEN = -6,     // buy now has already been requested by another participant
//            REQUEST_BUY_NOW_TOO_LOW_LIMIT = -7      // bid limit is less than buy now price
//        };
//
type RequestBuyNowResultCode int32

const (
	RequestBuyNowResultCodeRequestBuyNowSuccess          RequestBuyNowResultCode = 0
	RequestBuyNowResultCodeRequestBuyNowLotNotFound      RequestBuyNowResultCode = -1
	RequestBuyNowResultCodeRequestBuyNowUnsupported      RequestBuyNowResultCode = -2
	RequestBuyNowResultCodeRequestBuyNowUnregistered     RequestBuyNowResultCode = -3
	RequestBuyNowResultCodeRequestBuyNowTooLate          RequestBuyNowResultCode = -4
	RequestBuyNowResultCodeRequestBuyNowAlreadyRequested RequestBuyNowResultCode = -5
	RequestBuyNowResultCodeRequestBuyNowAlreadyTaken     RequestBuyNowResultCode = -6
	RequestBuyNowResultCodeRequestBuyNowTooLowLimit      RequestBuyNowResultCode = -7
)

var RequestBuyNowResultCodeAll = []RequestBuyNowResultCode{
	RequestBuyNowResultCodeRequestBuyNowSuccess,
	RequestBuyNowResultCodeRequestBuyNowLotNotFound,
	RequestBuyNowResultCodeRequestBuyNowUnsupported,
	RequestBuyNowResultCodeRequestBuyNowUnregistered,
	RequestBuyNowResultCodeRequestBuyNowTooLate,
	RequestBuyNowResultCodeRequestBuyNowAlreadyRequested,
	RequestBuyNowResultCodeRequestBuyNowAlreadyTaken,
	RequestBuyNowResultCodeRequestBuyNowTooLowLimit,
}

var requestBuyNowResultCodeMap = map[int32]string{
	0:  "RequestBuyNowResultCodeRequestBuyNowSuccess",
	-1: "RequestBuyNowResultCodeRequestBuyNowLotNotFound",
	-2: "RequestBuyNowResultCodeRequestBuyNowUnsupported",
	-3: "RequestBuyNowResultCodeRequestBuyNowUnregistered",
	-4: "RequestBuyNowResultCodeRequestBuyNowTooLate",
	-5: "RequestBuyNowResultCodeRequestBuyNowAlreadyRequested",
	-6: "RequestBuyNowResultCodeRequestBuyNowAlreadyTaken",
	-7: "RequestBuyNowResultCodeRequestBuyNowTooLowLimit",
}

var requestBuyNowResultCodeShortMap = map[int32]string{
	0:  "request_buy_now_success",
	-1: "request_buy_now_lot_not_found",
	-2: "request_buy_now_unsupported",
	-3: "request_buy_now_unregistered",
	-4: "request_buy_now_too_late",
	-5: "request_buy_now_already_requested",
	-6: "request_buy_now_already_taken",
	-7: "request_buy_now_too_low_limit",
}

var requestBuyNowResultCodeRevMap = map[string]int32{
	"RequestBuyNowResultCodeRequestBuyNowSuccess":          0,
	"RequestBuyNowResultCodeRequestBuyNowLotNotFound":      -1,
	"RequestBuyNowResultCodeRequestBuyNowUnsupported":      -2,
	"RequestBuyNowResultCodeRequestBuyNowUnregistered":     -3,
	"RequestBuyNowResultCodeRequestBuyNowTooLate":          -4,
	"RequestBuyNowResultCodeRequestBuyNowAlreadyRequested": -5,
	"RequestBuyNowResultCodeRequestBuyNowAlreadyTaken":     -6,
	"RequestBuyNowResultCodeRequestBuyNowTooLowLimit":      -7,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for RequestBuyNowResultCode
func (e RequestBuyNowResultCode) ValidEnum(v int32) bool {
	_, ok := requestBuyNowResultCodeMap[v]
	return ok
}
func (e RequestBuyNowResultCode) isFlag() bool {
	for i := len(RequestBuyNowResultCodeAll) - 1; i >= 0; i-- {
		expected := RequestBuyNowResultCode(2) << uint64(len(RequestBuyNowResultCodeAll)-1) >> uint64(len(RequestBuyNowResultCodeAll)-i)
		if expected != RequestBuyNowResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e RequestBuyNowResultCode) String() string {
	name, _ := requestBuyNowResultCodeMap[int32(e)]
	return name
}

func (e RequestBuyNowResultCode) ShortString() string {
	name, _ := requestBuyNowResultCodeShortMap[int32(e)]
	return name
}

func (e RequestBuyNowResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range RequestBuyNowResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *RequestBuyNowResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = RequestBuyNowResultCode(t.Value)
	return nil
}

// RequestBuyNowSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                  void;
//            }
//
type RequestBuyNowSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RequestBuyNowSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RequestBuyNowSuccessExt
func (u RequestBuyNowSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewRequestBuyNowSuccessExt creates a new  RequestBuyNowSuccessExt.
func NewRequestBuyNowSuccessExt(v LedgerVersion, value interface{}) (result RequestBuyNowSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// RequestBuyNowSuccess is an XDR Struct defines as:
//
//   struct RequestBuyNowSuccess
//        {
//            uint64 requestID;
//            bool fulfilled;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                  void;
//            }
//            ext;
//        };
//
type RequestBuyNowSuccess struct {
	RequestId Uint64                  `json:"requestID,omitempty"`
	Fulfilled bool                    `json:"fulfilled,omitempty"`
	Ext       RequestBuyNowSuccessExt `json:"ext,omitempty"`
}

// RequestBuyNowResult is an XDR Union defines as:
//
//   union RequestBuyNowResult switch (RequestBuyNowResultCode code)
//        {
//        case REQUEST_BUY_NOW_SUCCESS:
//            RequestBuyNowSuccess success;
//        default:
//            void;
//        };
//
type RequestBuyNowResult struct {
	Code    RequestBuyNowResultCode `json:"code,omitempty"`
	Success *RequestBuyNowSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RequestBuyNowResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RequestBuyNowResult
func (u RequestBuyNowResult) ArmForSwitch(sw int32) (string, bool) {
	switch RequestBuyNowResultCode(sw) {
	case RequestBuyNowResultCodeRequestBuyNowSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewRequestBuyNowResult creates a new  RequestBuyNowResult.
func NewRequestBuyNowResult(code RequestBuyNowResultCode, value interface{}) (result RequestBuyNowResult, err error) {
	result.Code = code
	switch RequestBuyNowResultCode(code) {
	case RequestBuyNowResultCodeRequestBuyNowSuccess:
		tv, ok := value.(RequestBuyNowSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be RequestBuyNowSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u RequestBuyNowResult) MustSuccess() RequestBuyNowSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RequestBuyNowResult) GetSuccess() (result RequestBuyNowSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// RequestParticipationOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type RequestParticipationOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RequestParticipationOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RequestParticipationOpExt
func (u RequestParticipationOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewRequestParticipationOpExt creates a new  RequestParticipationOpExt.
func NewRequestParticipationOpExt(v LedgerVersion, value interface{}) (result RequestParticipationOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// RequestParticipationOp is an XDR Struct defines as:
//
//   struct RequestParticipationOp
//        {
//            uint64 lotID;
//            bool isBuyNow; // if set, buy now will be automatically requested
//
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type RequestParticipationOp struct {
	LotId    Uint64                    `json:"lotID,omitempty"`
	IsBuyNow bool                      `json:"isBuyNow,omitempty"`
	Ext      RequestParticipationOpExt `json:"ext,omitempty"`
}

// RequestParticipationResultCode is an XDR Enum defines as:
//
//   enum RequestParticipationResultCode
//        {
//            REQUEST_PARTICIPATION_SUCCESS = 0,
//            REQUEST_PARTICIPATION_NO_LOT = -1,
//            REQUEST_PARTICIPATION_ALREADY_REQUESTED = -2,
//            REQUEST_PARTICIPATION_ALREADY_PARTICIPANT = -3,
//            REQUEST_PARTICIPATION_TOO_LATE = -4,
//            REQUEST_PARTICIPATION_MALFORMED = -5,
//            REQUEST_PARTICIPATION_INVALID_VERSION = -6,
//            REQUEST_PARTICIPATION_OWNER_CANT_PARTICIPATE = -7,      // owner is not allowed to participate
//            REQUEST_PARTICIPATION_NO_BUY_NOW = -8,                  // attempt to request buy now on lot that doesn't support it
//            REQUEST_PARTICIPATION_BUY_NOW_ALREADY_TAKEN = -9,       // buy now for shelf auction is already taken
//            REQUEST_PARTICIPATION_TOO_EARLY = -10
//        };
//
type RequestParticipationResultCode int32

const (
	RequestParticipationResultCodeRequestParticipationSuccess              RequestParticipationResultCode = 0
	RequestParticipationResultCodeRequestParticipationNoLot                RequestParticipationResultCode = -1
	RequestParticipationResultCodeRequestParticipationAlreadyRequested     RequestParticipationResultCode = -2
	RequestParticipationResultCodeRequestParticipationAlreadyParticipant   RequestParticipationResultCode = -3
	RequestParticipationResultCodeRequestParticipationTooLate              RequestParticipationResultCode = -4
	RequestParticipationResultCodeRequestParticipationMalformed            RequestParticipationResultCode = -5
	RequestParticipationResultCodeRequestParticipationInvalidVersion       RequestParticipationResultCode = -6
	RequestParticipationResultCodeRequestParticipationOwnerCantParticipate RequestParticipationResultCode = -7
	RequestParticipationResultCodeRequestParticipationNoBuyNow             RequestParticipationResultCode = -8
	RequestParticipationResultCodeRequestParticipationBuyNowAlreadyTaken   RequestParticipationResultCode = -9
	RequestParticipationResultCodeRequestParticipationTooEarly             RequestParticipationResultCode = -10
)

var RequestParticipationResultCodeAll = []RequestParticipationResultCode{
	RequestParticipationResultCodeRequestParticipationSuccess,
	RequestParticipationResultCodeRequestParticipationNoLot,
	RequestParticipationResultCodeRequestParticipationAlreadyRequested,
	RequestParticipationResultCodeRequestParticipationAlreadyParticipant,
	RequestParticipationResultCodeRequestParticipationTooLate,
	RequestParticipationResultCodeRequestParticipationMalformed,
	RequestParticipationResultCodeRequestParticipationInvalidVersion,
	RequestParticipationResultCodeRequestParticipationOwnerCantParticipate,
	RequestParticipationResultCodeRequestParticipationNoBuyNow,
	RequestParticipationResultCodeRequestParticipationBuyNowAlreadyTaken,
	RequestParticipationResultCodeRequestParticipationTooEarly,
}

var requestParticipationResultCodeMap = map[int32]string{
	0:   "RequestParticipationResultCodeRequestParticipationSuccess",
	-1:  "RequestParticipationResultCodeRequestParticipationNoLot",
	-2:  "RequestParticipationResultCodeRequestParticipationAlreadyRequested",
	-3:  "RequestParticipationResultCodeRequestParticipationAlreadyParticipant",
	-4:  "RequestParticipationResultCodeRequestParticipationTooLate",
	-5:  "RequestParticipationResultCodeRequestParticipationMalformed",
	-6:  "RequestParticipationResultCodeRequestParticipationInvalidVersion",
	-7:  "RequestParticipationResultCodeRequestParticipationOwnerCantParticipate",
	-8:  "RequestParticipationResultCodeRequestParticipationNoBuyNow",
	-9:  "RequestParticipationResultCodeRequestParticipationBuyNowAlreadyTaken",
	-10: "RequestParticipationResultCodeRequestParticipationTooEarly",
}

var requestParticipationResultCodeShortMap = map[int32]string{
	0:   "request_participation_success",
	-1:  "request_participation_no_lot",
	-2:  "request_participation_already_requested",
	-3:  "request_participation_already_participant",
	-4:  "request_participation_too_late",
	-5:  "request_participation_malformed",
	-6:  "request_participation_invalid_version",
	-7:  "request_participation_owner_cant_participate",
	-8:  "request_participation_no_buy_now",
	-9:  "request_participation_buy_now_already_taken",
	-10: "request_participation_too_early",
}

var requestParticipationResultCodeRevMap = map[string]int32{
	"RequestParticipationResultCodeRequestParticipationSuccess":              0,
	"RequestParticipationResultCodeRequestParticipationNoLot":                -1,
	"RequestParticipationResultCodeRequestParticipationAlreadyRequested":     -2,
	"RequestParticipationResultCodeRequestParticipationAlreadyParticipant":   -3,
	"RequestParticipationResultCodeRequestParticipationTooLate":              -4,
	"RequestParticipationResultCodeRequestParticipationMalformed":            -5,
	"RequestParticipationResultCodeRequestParticipationInvalidVersion":       -6,
	"RequestParticipationResultCodeRequestParticipationOwnerCantParticipate": -7,
	"RequestParticipationResultCodeRequestParticipationNoBuyNow":             -8,
	"RequestParticipationResultCodeRequestParticipationBuyNowAlreadyTaken":   -9,
	"RequestParticipationResultCodeRequestParticipationTooEarly":             -10,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for RequestParticipationResultCode
func (e RequestParticipationResultCode) ValidEnum(v int32) bool {
	_, ok := requestParticipationResultCodeMap[v]
	return ok
}
func (e RequestParticipationResultCode) isFlag() bool {
	for i := len(RequestParticipationResultCodeAll) - 1; i >= 0; i-- {
		expected := RequestParticipationResultCode(2) << uint64(len(RequestParticipationResultCodeAll)-1) >> uint64(len(RequestParticipationResultCodeAll)-i)
		if expected != RequestParticipationResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e RequestParticipationResultCode) String() string {
	name, _ := requestParticipationResultCodeMap[int32(e)]
	return name
}

func (e RequestParticipationResultCode) ShortString() string {
	name, _ := requestParticipationResultCodeShortMap[int32(e)]
	return name
}

func (e RequestParticipationResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range RequestParticipationResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *RequestParticipationResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = RequestParticipationResultCode(t.Value)
	return nil
}

// RequestParticipationSuccessExt is an XDR NestedUnion defines as:
//
//   union switch(LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type RequestParticipationSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RequestParticipationSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RequestParticipationSuccessExt
func (u RequestParticipationSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewRequestParticipationSuccessExt creates a new  RequestParticipationSuccessExt.
func NewRequestParticipationSuccessExt(v LedgerVersion, value interface{}) (result RequestParticipationSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// RequestParticipationSuccess is an XDR Struct defines as:
//
//   struct RequestParticipationSuccess
//        {
//            uint64 requestID;
//            bool fulfilled;
//
//            // reserved for future use
//            union switch(LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type RequestParticipationSuccess struct {
	RequestId Uint64                         `json:"requestID,omitempty"`
	Fulfilled bool                           `json:"fulfilled,omitempty"`
	Ext       RequestParticipationSuccessExt `json:"ext,omitempty"`
}

// RequestParticipationResult is an XDR Union defines as:
//
//   union RequestParticipationResult switch (RequestParticipationResultCode code)
//        {
//        case REQUEST_PARTICIPATION_SUCCESS:
//            RequestParticipationSuccess success;
//        default:
//            void;
//        };
//
type RequestParticipationResult struct {
	Code    RequestParticipationResultCode `json:"code,omitempty"`
	Success *RequestParticipationSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RequestParticipationResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RequestParticipationResult
func (u RequestParticipationResult) ArmForSwitch(sw int32) (string, bool) {
	switch RequestParticipationResultCode(sw) {
	case RequestParticipationResultCodeRequestParticipationSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewRequestParticipationResult creates a new  RequestParticipationResult.
func NewRequestParticipationResult(code RequestParticipationResultCode, value interface{}) (result RequestParticipationResult, err error) {
	result.Code = code
	switch RequestParticipationResultCode(code) {
	case RequestParticipationResultCodeRequestParticipationSuccess:
		tv, ok := value.(RequestParticipationSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be RequestParticipationSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u RequestParticipationResult) MustSuccess() RequestParticipationSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RequestParticipationResult) GetSuccess() (result RequestParticipationSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ReviewRequestAction is an XDR Enum defines as:
//
//   enum ReviewRequestAction
//    {
//        REVIEW_REQUEST_ACTION_APPROVE = 1,
//        REVIEW_REQUEST_ACTION_REJECT = 2,
//        REVIEW_REQUEST_ACTION_PERMANENT_REJECT = 3
//    };
//
type ReviewRequestAction int32

const (
	ReviewRequestActionReviewRequestActionApprove         ReviewRequestAction = 1
	ReviewRequestActionReviewRequestActionReject          ReviewRequestAction = 2
	ReviewRequestActionReviewRequestActionPermanentReject ReviewRequestAction = 3
)

var ReviewRequestActionAll = []ReviewRequestAction{
	ReviewRequestActionReviewRequestActionApprove,
	ReviewRequestActionReviewRequestActionReject,
	ReviewRequestActionReviewRequestActionPermanentReject,
}

var reviewRequestActionMap = map[int32]string{
	1: "ReviewRequestActionReviewRequestActionApprove",
	2: "ReviewRequestActionReviewRequestActionReject",
	3: "ReviewRequestActionReviewRequestActionPermanentReject",
}

var reviewRequestActionShortMap = map[int32]string{
	1: "review_request_action_approve",
	2: "review_request_action_reject",
	3: "review_request_action_permanent_reject",
}

var reviewRequestActionRevMap = map[string]int32{
	"ReviewRequestActionReviewRequestActionApprove":         1,
	"ReviewRequestActionReviewRequestActionReject":          2,
	"ReviewRequestActionReviewRequestActionPermanentReject": 3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ReviewRequestAction
func (e ReviewRequestAction) ValidEnum(v int32) bool {
	_, ok := reviewRequestActionMap[v]
	return ok
}
func (e ReviewRequestAction) isFlag() bool {
	for i := len(ReviewRequestActionAll) - 1; i >= 0; i-- {
		expected := ReviewRequestAction(2) << uint64(len(ReviewRequestActionAll)-1) >> uint64(len(ReviewRequestActionAll)-i)
		if expected != ReviewRequestActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ReviewRequestAction) String() string {
	name, _ := reviewRequestActionMap[int32(e)]
	return name
}

func (e ReviewRequestAction) ShortString() string {
	name, _ := reviewRequestActionShortMap[int32(e)]
	return name
}

func (e ReviewRequestAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ReviewRequestActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ReviewRequestAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ReviewRequestAction(t.Value)
	return nil
}

// ReviewRequestOpRequestDetails is an XDR NestedUnion defines as:
//
//   union switch(ReviewableRequestType requestType)
//        {
//         case REQUEST_TYPE_CREATE_LOT:
//            void;
//         case REQUEST_TYPE_CREATE_BID:
//            void;
//         case REQUEST_TYPE_REQUEST_PARTICIPATION:
//            void;
//        }
//
type ReviewRequestOpRequestDetails struct {
	RequestType ReviewableRequestType `json:"requestType,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewRequestOpRequestDetails) SwitchFieldName() string {
	return "RequestType"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewRequestOpRequestDetails
func (u ReviewRequestOpRequestDetails) ArmForSwitch(sw int32) (string, bool) {
	switch ReviewableRequestType(sw) {
	case ReviewableRequestTypeRequestTypeCreateLot:
		return "", true
	case ReviewableRequestTypeRequestTypeCreateBid:
		return "", true
	case ReviewableRequestTypeRequestTypeRequestParticipation:
		return "", true
	}
	return "-", false
}

// NewReviewRequestOpRequestDetails creates a new  ReviewRequestOpRequestDetails.
func NewReviewRequestOpRequestDetails(requestType ReviewableRequestType, value interface{}) (result ReviewRequestOpRequestDetails, err error) {
	result.RequestType = requestType
	switch ReviewableRequestType(requestType) {
	case ReviewableRequestTypeRequestTypeCreateLot:
		// void
	case ReviewableRequestTypeRequestTypeCreateBid:
		// void
	case ReviewableRequestTypeRequestTypeRequestParticipation:
		// void
	}
	return
}

// ReviewRequestOpExt is an XDR NestedUnion defines as:
//
//   union switch(LedgerVersion v)
//        {
//         case EMPTY_VERSION:
//            void;
//        }
//
type ReviewRequestOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewRequestOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewRequestOpExt
func (u ReviewRequestOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewReviewRequestOpExt creates a new  ReviewRequestOpExt.
func NewReviewRequestOpExt(v LedgerVersion, value interface{}) (result ReviewRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ReviewRequestOp is an XDR Struct defines as:
//
//   struct ReviewRequestOp
//    {
//        uint64 requestID;
//        Hash hash; // hash of the body of a request
//        ReviewRequestAction action;
//        string256 rejectReason;
//
//        union switch(ReviewableRequestType requestType)
//        {
//         case REQUEST_TYPE_CREATE_LOT:
//            void;
//         case REQUEST_TYPE_CREATE_BID:
//            void;
//         case REQUEST_TYPE_REQUEST_PARTICIPATION:
//            void;
//        } requestDetails;
//
//        // reserved for future use
//        union switch(LedgerVersion v)
//        {
//         case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type ReviewRequestOp struct {
	RequestId      Uint64                        `json:"requestID,omitempty"`
	Hash           Hash                          `json:"hash,omitempty"`
	Action         ReviewRequestAction           `json:"action,omitempty"`
	RejectReason   String256                     `json:"rejectReason,omitempty"`
	RequestDetails ReviewRequestOpRequestDetails `json:"requestDetails,omitempty"`
	Ext            ReviewRequestOpExt            `json:"ext,omitempty"`
}

// ReviewRequestResultCode is an XDR Enum defines as:
//
//   enum ReviewRequestResultCode
//    {
//        REVIEW_REQUEST_SUCCESS = 0,
//
//        REVIEW_REQUEST_REQUEST_NOT_FOUND = -1,
//        REVIEW_REQUEST_INVALID_REJECT_REASON = -2,
//        REVIEW_REQUEST_TYPE_MISMATCHED = -3,
//        REVIEW_REQUEST_HASH_MISMATCHED = -4,
//        REVIEW_REQUEST_REJECT_NOT_ALLOWED = -5,
//        REVIEW_REQUEST_PERMANENT_REJECT_NOT_ALLOWED = -6,
//        REVIEW_REQUEST_INVALID_VERSION = -7,
//        REVIEW_REQUEST_NO_REQUESTOR = -8,
//        REVIEW_REQUEST_REQUEST_PARTICIPATION_TOO_LATE = -9,
//        REVIEW_REQUEST_REQUEST_PARTICIPATION_BUY_NOW_ALREADY_TAKEN = -10
//    };
//
type ReviewRequestResultCode int32

const (
	ReviewRequestResultCodeReviewRequestSuccess                                ReviewRequestResultCode = 0
	ReviewRequestResultCodeReviewRequestRequestNotFound                        ReviewRequestResultCode = -1
	ReviewRequestResultCodeReviewRequestInvalidRejectReason                    ReviewRequestResultCode = -2
	ReviewRequestResultCodeReviewRequestTypeMismatched                         ReviewRequestResultCode = -3
	ReviewRequestResultCodeReviewRequestHashMismatched                         ReviewRequestResultCode = -4
	ReviewRequestResultCodeReviewRequestRejectNotAllowed                       ReviewRequestResultCode = -5
	ReviewRequestResultCodeReviewRequestPermanentRejectNotAllowed              ReviewRequestResultCode = -6
	ReviewRequestResultCodeReviewRequestInvalidVersion                         ReviewRequestResultCode = -7
	ReviewRequestResultCodeReviewRequestNoRequestor                            ReviewRequestResultCode = -8
	ReviewRequestResultCodeReviewRequestRequestParticipationTooLate            ReviewRequestResultCode = -9
	ReviewRequestResultCodeReviewRequestRequestParticipationBuyNowAlreadyTaken ReviewRequestResultCode = -10
)

var ReviewRequestResultCodeAll = []ReviewRequestResultCode{
	ReviewRequestResultCodeReviewRequestSuccess,
	ReviewRequestResultCodeReviewRequestRequestNotFound,
	ReviewRequestResultCodeReviewRequestInvalidRejectReason,
	ReviewRequestResultCodeReviewRequestTypeMismatched,
	ReviewRequestResultCodeReviewRequestHashMismatched,
	ReviewRequestResultCodeReviewRequestRejectNotAllowed,
	ReviewRequestResultCodeReviewRequestPermanentRejectNotAllowed,
	ReviewRequestResultCodeReviewRequestInvalidVersion,
	ReviewRequestResultCodeReviewRequestNoRequestor,
	ReviewRequestResultCodeReviewRequestRequestParticipationTooLate,
	ReviewRequestResultCodeReviewRequestRequestParticipationBuyNowAlreadyTaken,
}

var reviewRequestResultCodeMap = map[int32]string{
	0:   "ReviewRequestResultCodeReviewRequestSuccess",
	-1:  "ReviewRequestResultCodeReviewRequestRequestNotFound",
	-2:  "ReviewRequestResultCodeReviewRequestInvalidRejectReason",
	-3:  "ReviewRequestResultCodeReviewRequestTypeMismatched",
	-4:  "ReviewRequestResultCodeReviewRequestHashMismatched",
	-5:  "ReviewRequestResultCodeReviewRequestRejectNotAllowed",
	-6:  "ReviewRequestResultCodeReviewRequestPermanentRejectNotAllowed",
	-7:  "ReviewRequestResultCodeReviewRequestInvalidVersion",
	-8:  "ReviewRequestResultCodeReviewRequestNoRequestor",
	-9:  "ReviewRequestResultCodeReviewRequestRequestParticipationTooLate",
	-10: "ReviewRequestResultCodeReviewRequestRequestParticipationBuyNowAlreadyTaken",
}

var reviewRequestResultCodeShortMap = map[int32]string{
	0:   "review_request_success",
	-1:  "review_request_request_not_found",
	-2:  "review_request_invalid_reject_reason",
	-3:  "review_request_type_mismatched",
	-4:  "review_request_hash_mismatched",
	-5:  "review_request_reject_not_allowed",
	-6:  "review_request_permanent_reject_not_allowed",
	-7:  "review_request_invalid_version",
	-8:  "review_request_no_requestor",
	-9:  "review_request_request_participation_too_late",
	-10: "review_request_request_participation_buy_now_already_taken",
}

var reviewRequestResultCodeRevMap = map[string]int32{
	"ReviewRequestResultCodeReviewRequestSuccess":                                0,
	"ReviewRequestResultCodeReviewRequestRequestNotFound":                        -1,
	"ReviewRequestResultCodeReviewRequestInvalidRejectReason":                    -2,
	"ReviewRequestResultCodeReviewRequestTypeMismatched":                         -3,
	"ReviewRequestResultCodeReviewRequestHashMismatched":                         -4,
	"ReviewRequestResultCodeReviewRequestRejectNotAllowed":                       -5,
	"ReviewRequestResultCodeReviewRequestPermanentRejectNotAllowed":              -6,
	"ReviewRequestResultCodeReviewRequestInvalidVersion":                         -7,
	"ReviewRequestResultCodeReviewRequestNoRequestor":                            -8,
	"ReviewRequestResultCodeReviewRequestRequestParticipationTooLate":            -9,
	"ReviewRequestResultCodeReviewRequestRequestParticipationBuyNowAlreadyTaken": -10,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ReviewRequestResultCode
func (e ReviewRequestResultCode) ValidEnum(v int32) bool {
	_, ok := reviewRequestResultCodeMap[v]
	return ok
}
func (e ReviewRequestResultCode) isFlag() bool {
	for i := len(ReviewRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := ReviewRequestResultCode(2) << uint64(len(ReviewRequestResultCodeAll)-1) >> uint64(len(ReviewRequestResultCodeAll)-i)
		if expected != ReviewRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ReviewRequestResultCode) String() string {
	name, _ := reviewRequestResultCodeMap[int32(e)]
	return name
}

func (e ReviewRequestResultCode) ShortString() string {
	name, _ := reviewRequestResultCodeShortMap[int32(e)]
	return name
}

func (e ReviewRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ReviewRequestResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ReviewRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ReviewRequestResultCode(t.Value)
	return nil
}

// ReviewRequestResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch(LedgerVersion v)
//            {
//             case EMPTY_VERSION:
//                void;
//            }
//
type ReviewRequestResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewRequestResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewRequestResultSuccessExt
func (u ReviewRequestResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewReviewRequestResultSuccessExt creates a new  ReviewRequestResultSuccessExt.
func NewReviewRequestResultSuccessExt(v LedgerVersion, value interface{}) (result ReviewRequestResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ReviewRequestResultSuccess is an XDR NestedStruct defines as:
//
//   struct
//        {
//            //reserved for future use
//            union switch(LedgerVersion v)
//            {
//             case EMPTY_VERSION:
//                void;
//            } ext;
//        }
//
type ReviewRequestResultSuccess struct {
	Ext ReviewRequestResultSuccessExt `json:"ext,omitempty"`
}

// ReviewRequestResult is an XDR Union defines as:
//
//   union ReviewRequestResult switch(ReviewRequestResultCode code)
//    {
//     case REVIEW_REQUEST_SUCCESS:
//        struct
//        {
//            //reserved for future use
//            union switch(LedgerVersion v)
//            {
//             case EMPTY_VERSION:
//                void;
//            } ext;
//        } success;
//     default:
//        void;
//    };
//
type ReviewRequestResult struct {
	Code    ReviewRequestResultCode     `json:"code,omitempty"`
	Success *ReviewRequestResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewRequestResult
func (u ReviewRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch ReviewRequestResultCode(sw) {
	case ReviewRequestResultCodeReviewRequestSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewReviewRequestResult creates a new  ReviewRequestResult.
func NewReviewRequestResult(code ReviewRequestResultCode, value interface{}) (result ReviewRequestResult, err error) {
	result.Code = code
	switch ReviewRequestResultCode(code) {
	case ReviewRequestResultCodeReviewRequestSuccess:
		tv, ok := value.(ReviewRequestResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReviewRequestResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ReviewRequestResult) MustSuccess() ReviewRequestResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewRequestResult) GetSuccess() (result ReviewRequestResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// SendMessageOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type SendMessageOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SendMessageOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SendMessageOpExt
func (u SendMessageOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSendMessageOpExt creates a new  SendMessageOpExt.
func NewSendMessageOpExt(v LedgerVersion, value interface{}) (result SendMessageOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SendMessageOp is an XDR Struct defines as:
//
//   struct SendMessageOp
//        {
//            AccountID receiverID;
//            uint64 lotID;
//
//            string text<1000>;
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        };
//
type SendMessageOp struct {
	ReceiverId AccountId        `json:"receiverID,omitempty"`
	LotId      Uint64           `json:"lotID,omitempty"`
	Text       string           `json:"text,omitempty" xdrmaxsize:"1000"`
	Ext        SendMessageOpExt `json:"ext,omitempty"`
}

// SendMessageResultCode is an XDR Enum defines as:
//
//   enum SendMessageResultCode
//        {
//            SEND_MESSAGE_SUCCESS = 0,
//            SEND_MESSAGE_NO_LOT = -1,           // lot with lotID not found
//            SEND_MESSAGE_NO_RECEIVER = -2,      // receiver not found or not the participant/organizer of the lot
//            SEND_MESSAGE_MALFORMED = -3         // empty message
//        };
//
type SendMessageResultCode int32

const (
	SendMessageResultCodeSendMessageSuccess    SendMessageResultCode = 0
	SendMessageResultCodeSendMessageNoLot      SendMessageResultCode = -1
	SendMessageResultCodeSendMessageNoReceiver SendMessageResultCode = -2
	SendMessageResultCodeSendMessageMalformed  SendMessageResultCode = -3
)

var SendMessageResultCodeAll = []SendMessageResultCode{
	SendMessageResultCodeSendMessageSuccess,
	SendMessageResultCodeSendMessageNoLot,
	SendMessageResultCodeSendMessageNoReceiver,
	SendMessageResultCodeSendMessageMalformed,
}

var sendMessageResultCodeMap = map[int32]string{
	0:  "SendMessageResultCodeSendMessageSuccess",
	-1: "SendMessageResultCodeSendMessageNoLot",
	-2: "SendMessageResultCodeSendMessageNoReceiver",
	-3: "SendMessageResultCodeSendMessageMalformed",
}

var sendMessageResultCodeShortMap = map[int32]string{
	0:  "send_message_success",
	-1: "send_message_no_lot",
	-2: "send_message_no_receiver",
	-3: "send_message_malformed",
}

var sendMessageResultCodeRevMap = map[string]int32{
	"SendMessageResultCodeSendMessageSuccess":    0,
	"SendMessageResultCodeSendMessageNoLot":      -1,
	"SendMessageResultCodeSendMessageNoReceiver": -2,
	"SendMessageResultCodeSendMessageMalformed":  -3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for SendMessageResultCode
func (e SendMessageResultCode) ValidEnum(v int32) bool {
	_, ok := sendMessageResultCodeMap[v]
	return ok
}
func (e SendMessageResultCode) isFlag() bool {
	for i := len(SendMessageResultCodeAll) - 1; i >= 0; i-- {
		expected := SendMessageResultCode(2) << uint64(len(SendMessageResultCodeAll)-1) >> uint64(len(SendMessageResultCodeAll)-i)
		if expected != SendMessageResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e SendMessageResultCode) String() string {
	name, _ := sendMessageResultCodeMap[int32(e)]
	return name
}

func (e SendMessageResultCode) ShortString() string {
	name, _ := sendMessageResultCodeShortMap[int32(e)]
	return name
}

func (e SendMessageResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range SendMessageResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *SendMessageResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = SendMessageResultCode(t.Value)
	return nil
}

// SendMessageResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                        {
//                        case EMPTY_VERSION:
//                          void;
//                        }
//
type SendMessageResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SendMessageResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SendMessageResultSuccessExt
func (u SendMessageResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSendMessageResultSuccessExt creates a new  SendMessageResultSuccessExt.
func NewSendMessageResultSuccessExt(v LedgerVersion, value interface{}) (result SendMessageResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SendMessageResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//                        uint64 messageID;
//
//                        // reserved for future use
//                        union switch (LedgerVersion v)
//                        {
//                        case EMPTY_VERSION:
//                          void;
//                        }
//                        ext;
//                      }
//
type SendMessageResultSuccess struct {
	MessageId Uint64                      `json:"messageID,omitempty"`
	Ext       SendMessageResultSuccessExt `json:"ext,omitempty"`
}

// SendMessageResult is an XDR Union defines as:
//
//   union SendMessageResult switch (SendMessageResultCode code)
//        {
//        case SEND_MESSAGE_SUCCESS:
//            struct {
//                        uint64 messageID;
//
//                        // reserved for future use
//                        union switch (LedgerVersion v)
//                        {
//                        case EMPTY_VERSION:
//                          void;
//                        }
//                        ext;
//                      } success;
//        default:
//            void;
//        };
//
type SendMessageResult struct {
	Code    SendMessageResultCode     `json:"code,omitempty"`
	Success *SendMessageResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SendMessageResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SendMessageResult
func (u SendMessageResult) ArmForSwitch(sw int32) (string, bool) {
	switch SendMessageResultCode(sw) {
	case SendMessageResultCodeSendMessageSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewSendMessageResult creates a new  SendMessageResult.
func NewSendMessageResult(code SendMessageResultCode, value interface{}) (result SendMessageResult, err error) {
	result.Code = code
	switch SendMessageResultCode(code) {
	case SendMessageResultCodeSendMessageSuccess:
		tv, ok := value.(SendMessageResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be SendMessageResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u SendMessageResult) MustSuccess() SendMessageResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SendMessageResult) GetSuccess() (result SendMessageResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// SetMaxBidOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type SetMaxBidOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SetMaxBidOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SetMaxBidOpExt
func (u SetMaxBidOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSetMaxBidOpExt creates a new  SetMaxBidOpExt.
func NewSetMaxBidOpExt(v LedgerVersion, value interface{}) (result SetMaxBidOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SetMaxBidOp is an XDR Struct defines as:
//
//   struct SetMaxBidOp
//        {
//            uint64 lotID;
//            uint64 maxBid;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type SetMaxBidOp struct {
	LotId  Uint64         `json:"lotID,omitempty"`
	MaxBid Uint64         `json:"maxBid,omitempty"`
	Ext    SetMaxBidOpExt `json:"ext,omitempty"`
}

// SetMaxBidResultCode is an XDR Enum defines as:
//
//   enum SetMaxBidResultCode
//        {
//            SET_MAX_BID_SUCCESS = 0,
//
//            SET_MAX_BID_INVALID_VERSION = -1,           // current ledger version does not support this op
//            SET_MAX_BID_NO_PARTICIPANT = -2,            // participant not found
//            SET_MAX_BID_NO_AUCTION = -3,                // type of dest lot does not allow max bid
//            SET_MAX_BID_TOO_LATE = -4,                  // dest lot is closed or finished
//            SET_MAX_BID_TOO_SMALL = -5,                 // max bid is too low
//            SET_MAX_BID_LIMIT_EXCEEDED = -6,            // max bid exceeds bidding limit
//            SET_MAX_BID_NOT_ALLOWED = -7                // participant is not in pending state
//        };
//
type SetMaxBidResultCode int32

const (
	SetMaxBidResultCodeSetMaxBidSuccess        SetMaxBidResultCode = 0
	SetMaxBidResultCodeSetMaxBidInvalidVersion SetMaxBidResultCode = -1
	SetMaxBidResultCodeSetMaxBidNoParticipant  SetMaxBidResultCode = -2
	SetMaxBidResultCodeSetMaxBidNoAuction      SetMaxBidResultCode = -3
	SetMaxBidResultCodeSetMaxBidTooLate        SetMaxBidResultCode = -4
	SetMaxBidResultCodeSetMaxBidTooSmall       SetMaxBidResultCode = -5
	SetMaxBidResultCodeSetMaxBidLimitExceeded  SetMaxBidResultCode = -6
	SetMaxBidResultCodeSetMaxBidNotAllowed     SetMaxBidResultCode = -7
)

var SetMaxBidResultCodeAll = []SetMaxBidResultCode{
	SetMaxBidResultCodeSetMaxBidSuccess,
	SetMaxBidResultCodeSetMaxBidInvalidVersion,
	SetMaxBidResultCodeSetMaxBidNoParticipant,
	SetMaxBidResultCodeSetMaxBidNoAuction,
	SetMaxBidResultCodeSetMaxBidTooLate,
	SetMaxBidResultCodeSetMaxBidTooSmall,
	SetMaxBidResultCodeSetMaxBidLimitExceeded,
	SetMaxBidResultCodeSetMaxBidNotAllowed,
}

var setMaxBidResultCodeMap = map[int32]string{
	0:  "SetMaxBidResultCodeSetMaxBidSuccess",
	-1: "SetMaxBidResultCodeSetMaxBidInvalidVersion",
	-2: "SetMaxBidResultCodeSetMaxBidNoParticipant",
	-3: "SetMaxBidResultCodeSetMaxBidNoAuction",
	-4: "SetMaxBidResultCodeSetMaxBidTooLate",
	-5: "SetMaxBidResultCodeSetMaxBidTooSmall",
	-6: "SetMaxBidResultCodeSetMaxBidLimitExceeded",
	-7: "SetMaxBidResultCodeSetMaxBidNotAllowed",
}

var setMaxBidResultCodeShortMap = map[int32]string{
	0:  "set_max_bid_success",
	-1: "set_max_bid_invalid_version",
	-2: "set_max_bid_no_participant",
	-3: "set_max_bid_no_auction",
	-4: "set_max_bid_too_late",
	-5: "set_max_bid_too_small",
	-6: "set_max_bid_limit_exceeded",
	-7: "set_max_bid_not_allowed",
}

var setMaxBidResultCodeRevMap = map[string]int32{
	"SetMaxBidResultCodeSetMaxBidSuccess":        0,
	"SetMaxBidResultCodeSetMaxBidInvalidVersion": -1,
	"SetMaxBidResultCodeSetMaxBidNoParticipant":  -2,
	"SetMaxBidResultCodeSetMaxBidNoAuction":      -3,
	"SetMaxBidResultCodeSetMaxBidTooLate":        -4,
	"SetMaxBidResultCodeSetMaxBidTooSmall":       -5,
	"SetMaxBidResultCodeSetMaxBidLimitExceeded":  -6,
	"SetMaxBidResultCodeSetMaxBidNotAllowed":     -7,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for SetMaxBidResultCode
func (e SetMaxBidResultCode) ValidEnum(v int32) bool {
	_, ok := setMaxBidResultCodeMap[v]
	return ok
}
func (e SetMaxBidResultCode) isFlag() bool {
	for i := len(SetMaxBidResultCodeAll) - 1; i >= 0; i-- {
		expected := SetMaxBidResultCode(2) << uint64(len(SetMaxBidResultCodeAll)-1) >> uint64(len(SetMaxBidResultCodeAll)-i)
		if expected != SetMaxBidResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e SetMaxBidResultCode) String() string {
	name, _ := setMaxBidResultCodeMap[int32(e)]
	return name
}

func (e SetMaxBidResultCode) ShortString() string {
	name, _ := setMaxBidResultCodeShortMap[int32(e)]
	return name
}

func (e SetMaxBidResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range SetMaxBidResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *SetMaxBidResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = SetMaxBidResultCode(t.Value)
	return nil
}

// SetMaxBidSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type SetMaxBidSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SetMaxBidSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SetMaxBidSuccessExt
func (u SetMaxBidSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSetMaxBidSuccessExt creates a new  SetMaxBidSuccessExt.
func NewSetMaxBidSuccessExt(v LedgerVersion, value interface{}) (result SetMaxBidSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SetMaxBidSuccess is an XDR Struct defines as:
//
//   struct SetMaxBidSuccess
//        {
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type SetMaxBidSuccess struct {
	Ext SetMaxBidSuccessExt `json:"ext,omitempty"`
}

// SetMaxBidResult is an XDR Union defines as:
//
//   union SetMaxBidResult switch (SetMaxBidResultCode code)
//        {
//        case SET_MAX_BID_SUCCESS:
//           SetMaxBidSuccess success;
//        default:
//            void;
//        };
//
type SetMaxBidResult struct {
	Code    SetMaxBidResultCode `json:"code,omitempty"`
	Success *SetMaxBidSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SetMaxBidResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SetMaxBidResult
func (u SetMaxBidResult) ArmForSwitch(sw int32) (string, bool) {
	switch SetMaxBidResultCode(sw) {
	case SetMaxBidResultCodeSetMaxBidSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewSetMaxBidResult creates a new  SetMaxBidResult.
func NewSetMaxBidResult(code SetMaxBidResultCode, value interface{}) (result SetMaxBidResult, err error) {
	result.Code = code
	switch SetMaxBidResultCode(code) {
	case SetMaxBidResultCodeSetMaxBidSuccess:
		tv, ok := value.(SetMaxBidSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetMaxBidSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u SetMaxBidResult) MustSuccess() SetMaxBidSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SetMaxBidResult) GetSuccess() (result SetMaxBidSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// SetOptionsOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type SetOptionsOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SetOptionsOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SetOptionsOpExt
func (u SetOptionsOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSetOptionsOpExt creates a new  SetOptionsOpExt.
func NewSetOptionsOpExt(v LedgerVersion, value interface{}) (result SetOptionsOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SetOptionsOp is an XDR Struct defines as:
//
//   struct SetOptionsOp
//        {
//            AccountID* inflationDest; // sets the inflation destination
//
//            uint32* clearFlags; // which flags to clear
//            uint32* setFlags;   // which flags to set
//
//            // account threshold manipulation
//            uint32* masterWeight; // weight of the master account
//            uint32* lowThreshold;
//            uint32* medThreshold;
//            uint32* highThreshold;
//
//            string32* homeDomain; // sets the home domain
//
//            // Add, update or remove a signer for the account
//            // signer is deleted if the weight is 0
//            Signer* signer;
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        };
//
type SetOptionsOp struct {
	InflationDest *AccountId      `json:"inflationDest,omitempty"`
	ClearFlags    *Uint32         `json:"clearFlags,omitempty"`
	SetFlags      *Uint32         `json:"setFlags,omitempty"`
	MasterWeight  *Uint32         `json:"masterWeight,omitempty"`
	LowThreshold  *Uint32         `json:"lowThreshold,omitempty"`
	MedThreshold  *Uint32         `json:"medThreshold,omitempty"`
	HighThreshold *Uint32         `json:"highThreshold,omitempty"`
	HomeDomain    *String32       `json:"homeDomain,omitempty"`
	Signer        *Signer         `json:"signer,omitempty"`
	Ext           SetOptionsOpExt `json:"ext,omitempty"`
}

// SetOptionsResultCode is an XDR Enum defines as:
//
//   enum SetOptionsResultCode
//        {
//            // codes considered as "success" for the operation
//            SET_OPTIONS_SUCCESS = 0,
//            // codes considered as "failure" for the operation
//            SET_OPTIONS_LOW_RESERVE = -1,      // not enough funds to add a signer
//            SET_OPTIONS_TOO_MANY_SIGNERS = -2, // max number of signers already reached
//            SET_OPTIONS_BAD_FLAGS = -3,        // invalid combination of clear/set flags
//            SET_OPTIONS_INVALID_INFLATION = -4,      // inflation account does not exist
//            SET_OPTIONS_CANT_CHANGE = -5,            // can no longer change this option
//            SET_OPTIONS_UNKNOWN_FLAG = -6,           // can't set an unknown flag
//            SET_OPTIONS_THRESHOLD_OUT_OF_RANGE = -7, // bad value for weight/threshold
//            SET_OPTIONS_BAD_SIGNER = -8,             // signer cannot be masterkey
//            SET_OPTIONS_INVALID_HOME_DOMAIN = -9     // malformed home domain
//        };
//
type SetOptionsResultCode int32

const (
	SetOptionsResultCodeSetOptionsSuccess             SetOptionsResultCode = 0
	SetOptionsResultCodeSetOptionsLowReserve          SetOptionsResultCode = -1
	SetOptionsResultCodeSetOptionsTooManySigners      SetOptionsResultCode = -2
	SetOptionsResultCodeSetOptionsBadFlags            SetOptionsResultCode = -3
	SetOptionsResultCodeSetOptionsInvalidInflation    SetOptionsResultCode = -4
	SetOptionsResultCodeSetOptionsCantChange          SetOptionsResultCode = -5
	SetOptionsResultCodeSetOptionsUnknownFlag         SetOptionsResultCode = -6
	SetOptionsResultCodeSetOptionsThresholdOutOfRange SetOptionsResultCode = -7
	SetOptionsResultCodeSetOptionsBadSigner           SetOptionsResultCode = -8
	SetOptionsResultCodeSetOptionsInvalidHomeDomain   SetOptionsResultCode = -9
)

var SetOptionsResultCodeAll = []SetOptionsResultCode{
	SetOptionsResultCodeSetOptionsSuccess,
	SetOptionsResultCodeSetOptionsLowReserve,
	SetOptionsResultCodeSetOptionsTooManySigners,
	SetOptionsResultCodeSetOptionsBadFlags,
	SetOptionsResultCodeSetOptionsInvalidInflation,
	SetOptionsResultCodeSetOptionsCantChange,
	SetOptionsResultCodeSetOptionsUnknownFlag,
	SetOptionsResultCodeSetOptionsThresholdOutOfRange,
	SetOptionsResultCodeSetOptionsBadSigner,
	SetOptionsResultCodeSetOptionsInvalidHomeDomain,
}

var setOptionsResultCodeMap = map[int32]string{
	0:  "SetOptionsResultCodeSetOptionsSuccess",
	-1: "SetOptionsResultCodeSetOptionsLowReserve",
	-2: "SetOptionsResultCodeSetOptionsTooManySigners",
	-3: "SetOptionsResultCodeSetOptionsBadFlags",
	-4: "SetOptionsResultCodeSetOptionsInvalidInflation",
	-5: "SetOptionsResultCodeSetOptionsCantChange",
	-6: "SetOptionsResultCodeSetOptionsUnknownFlag",
	-7: "SetOptionsResultCodeSetOptionsThresholdOutOfRange",
	-8: "SetOptionsResultCodeSetOptionsBadSigner",
	-9: "SetOptionsResultCodeSetOptionsInvalidHomeDomain",
}

var setOptionsResultCodeShortMap = map[int32]string{
	0:  "set_options_success",
	-1: "set_options_low_reserve",
	-2: "set_options_too_many_signers",
	-3: "set_options_bad_flags",
	-4: "set_options_invalid_inflation",
	-5: "set_options_cant_change",
	-6: "set_options_unknown_flag",
	-7: "set_options_threshold_out_of_range",
	-8: "set_options_bad_signer",
	-9: "set_options_invalid_home_domain",
}

var setOptionsResultCodeRevMap = map[string]int32{
	"SetOptionsResultCodeSetOptionsSuccess":             0,
	"SetOptionsResultCodeSetOptionsLowReserve":          -1,
	"SetOptionsResultCodeSetOptionsTooManySigners":      -2,
	"SetOptionsResultCodeSetOptionsBadFlags":            -3,
	"SetOptionsResultCodeSetOptionsInvalidInflation":    -4,
	"SetOptionsResultCodeSetOptionsCantChange":          -5,
	"SetOptionsResultCodeSetOptionsUnknownFlag":         -6,
	"SetOptionsResultCodeSetOptionsThresholdOutOfRange": -7,
	"SetOptionsResultCodeSetOptionsBadSigner":           -8,
	"SetOptionsResultCodeSetOptionsInvalidHomeDomain":   -9,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for SetOptionsResultCode
func (e SetOptionsResultCode) ValidEnum(v int32) bool {
	_, ok := setOptionsResultCodeMap[v]
	return ok
}
func (e SetOptionsResultCode) isFlag() bool {
	for i := len(SetOptionsResultCodeAll) - 1; i >= 0; i-- {
		expected := SetOptionsResultCode(2) << uint64(len(SetOptionsResultCodeAll)-1) >> uint64(len(SetOptionsResultCodeAll)-i)
		if expected != SetOptionsResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e SetOptionsResultCode) String() string {
	name, _ := setOptionsResultCodeMap[int32(e)]
	return name
}

func (e SetOptionsResultCode) ShortString() string {
	name, _ := setOptionsResultCodeShortMap[int32(e)]
	return name
}

func (e SetOptionsResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range SetOptionsResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *SetOptionsResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = SetOptionsResultCode(t.Value)
	return nil
}

// SetOptionsResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                        {
//                        case EMPTY_VERSION:
//                          void;
//                        }
//
type SetOptionsResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SetOptionsResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SetOptionsResultSuccessExt
func (u SetOptionsResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSetOptionsResultSuccessExt creates a new  SetOptionsResultSuccessExt.
func NewSetOptionsResultSuccessExt(v LedgerVersion, value interface{}) (result SetOptionsResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SetOptionsResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//                        // reserved for future use
//                        union switch (LedgerVersion v)
//                        {
//                        case EMPTY_VERSION:
//                          void;
//                        }
//                        ext;
//                      }
//
type SetOptionsResultSuccess struct {
	Ext SetOptionsResultSuccessExt `json:"ext,omitempty"`
}

// SetOptionsResult is an XDR Union defines as:
//
//   union SetOptionsResult switch (SetOptionsResultCode code)
//        {
//        case SET_OPTIONS_SUCCESS:
//            struct {
//                        // reserved for future use
//                        union switch (LedgerVersion v)
//                        {
//                        case EMPTY_VERSION:
//                          void;
//                        }
//                        ext;
//                      } success;
//        default:
//            void;
//        };
//
type SetOptionsResult struct {
	Code    SetOptionsResultCode     `json:"code,omitempty"`
	Success *SetOptionsResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SetOptionsResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SetOptionsResult
func (u SetOptionsResult) ArmForSwitch(sw int32) (string, bool) {
	switch SetOptionsResultCode(sw) {
	case SetOptionsResultCodeSetOptionsSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewSetOptionsResult creates a new  SetOptionsResult.
func NewSetOptionsResult(code SetOptionsResultCode, value interface{}) (result SetOptionsResult, err error) {
	result.Code = code
	switch SetOptionsResultCode(code) {
	case SetOptionsResultCodeSetOptionsSuccess:
		tv, ok := value.(SetOptionsResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetOptionsResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u SetOptionsResult) MustSuccess() SetOptionsResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SetOptionsResult) GetSuccess() (result SetOptionsResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// LotDetailsExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LotDetailsExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LotDetailsExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LotDetailsExt
func (u LotDetailsExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLotDetailsExt creates a new  LotDetailsExt.
func NewLotDetailsExt(v LedgerVersion, value interface{}) (result LotDetailsExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LotDetails is an XDR Struct defines as:
//
//   struct LotDetails
//        {
//            string details<>;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type LotDetails struct {
	Details string        `json:"details,omitempty"`
	Ext     LotDetailsExt `json:"ext,omitempty"`
}

// UpdateLotOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            case LEDGER_VERSION_LOT_STATE_NEGOTIATIONS:
//                LotState* state;
//            }
//
type UpdateLotOpExt struct {
	V     LedgerVersion `json:"v,omitempty"`
	State **LotState    `json:"state,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateLotOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateLotOpExt
func (u UpdateLotOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionLedgerVersionLotStateNegotiations:
		return "State", true
	}
	return "-", false
}

// NewUpdateLotOpExt creates a new  UpdateLotOpExt.
func NewUpdateLotOpExt(v LedgerVersion, value interface{}) (result UpdateLotOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionLedgerVersionLotStateNegotiations:
		tv, ok := value.(*LotState)
		if !ok {
			err = fmt.Errorf("invalid value, must be *LotState")
			return
		}
		result.State = &tv
	}
	return
}

// MustState retrieves the State value from the union,
// panicing if the value is not set.
func (u UpdateLotOpExt) MustState() *LotState {
	val, ok := u.GetState()

	if !ok {
		panic("arm State is not set")
	}

	return val
}

// GetState retrieves the State value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u UpdateLotOpExt) GetState() (result *LotState, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "State" {
		result = *u.State
		ok = true
	}

	return
}

// UpdateLotOp is an XDR Struct defines as:
//
//   struct UpdateLotOp
//        {
//            uint64 lotID;
//
//            LotDetails* details;     // JSON string with full lot description
//
//            int64* startTime;        // auction start time in Unix epoch format
//            uint64* duration;        // duration in seconds (optional for certain type of lot)
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            case LEDGER_VERSION_LOT_STATE_NEGOTIATIONS:
//                LotState* state;
//            }
//            ext;
//        };
//
type UpdateLotOp struct {
	LotId     Uint64         `json:"lotID,omitempty"`
	Details   *LotDetails    `json:"details,omitempty"`
	StartTime *Int64         `json:"startTime,omitempty"`
	Duration  *Uint64        `json:"duration,omitempty"`
	Ext       UpdateLotOpExt `json:"ext,omitempty"`
}

// UpdateLotSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type UpdateLotSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateLotSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateLotSuccessExt
func (u UpdateLotSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewUpdateLotSuccessExt creates a new  UpdateLotSuccessExt.
func NewUpdateLotSuccessExt(v LedgerVersion, value interface{}) (result UpdateLotSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// UpdateLotSuccess is an XDR Struct defines as:
//
//   struct UpdateLotSuccess {
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type UpdateLotSuccess struct {
	Ext UpdateLotSuccessExt `json:"ext,omitempty"`
}

// UpdateLotResultCode is an XDR Enum defines as:
//
//   enum UpdateLotResultCode
//        {
//            UPDATE_LOT_SUCCESS = 0, // lot was created
//
//            UPDATE_LOT_NOT_FOUND = -1,
//            UPDATE_LOT_NOT_ALLOWED = -2,
//            UPDATE_LOT_MALFORMED_START_TIME = -3,
//            UPDATE_LOT_MALFORMED_DURATION = -4,
//            UPDATE_LOT_INVALID_JSON = -5
//        };
//
type UpdateLotResultCode int32

const (
	UpdateLotResultCodeUpdateLotSuccess            UpdateLotResultCode = 0
	UpdateLotResultCodeUpdateLotNotFound           UpdateLotResultCode = -1
	UpdateLotResultCodeUpdateLotNotAllowed         UpdateLotResultCode = -2
	UpdateLotResultCodeUpdateLotMalformedStartTime UpdateLotResultCode = -3
	UpdateLotResultCodeUpdateLotMalformedDuration  UpdateLotResultCode = -4
	UpdateLotResultCodeUpdateLotInvalidJson        UpdateLotResultCode = -5
)

var UpdateLotResultCodeAll = []UpdateLotResultCode{
	UpdateLotResultCodeUpdateLotSuccess,
	UpdateLotResultCodeUpdateLotNotFound,
	UpdateLotResultCodeUpdateLotNotAllowed,
	UpdateLotResultCodeUpdateLotMalformedStartTime,
	UpdateLotResultCodeUpdateLotMalformedDuration,
	UpdateLotResultCodeUpdateLotInvalidJson,
}

var updateLotResultCodeMap = map[int32]string{
	0:  "UpdateLotResultCodeUpdateLotSuccess",
	-1: "UpdateLotResultCodeUpdateLotNotFound",
	-2: "UpdateLotResultCodeUpdateLotNotAllowed",
	-3: "UpdateLotResultCodeUpdateLotMalformedStartTime",
	-4: "UpdateLotResultCodeUpdateLotMalformedDuration",
	-5: "UpdateLotResultCodeUpdateLotInvalidJson",
}

var updateLotResultCodeShortMap = map[int32]string{
	0:  "update_lot_success",
	-1: "update_lot_not_found",
	-2: "update_lot_not_allowed",
	-3: "update_lot_malformed_start_time",
	-4: "update_lot_malformed_duration",
	-5: "update_lot_invalid_json",
}

var updateLotResultCodeRevMap = map[string]int32{
	"UpdateLotResultCodeUpdateLotSuccess":            0,
	"UpdateLotResultCodeUpdateLotNotFound":           -1,
	"UpdateLotResultCodeUpdateLotNotAllowed":         -2,
	"UpdateLotResultCodeUpdateLotMalformedStartTime": -3,
	"UpdateLotResultCodeUpdateLotMalformedDuration":  -4,
	"UpdateLotResultCodeUpdateLotInvalidJson":        -5,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for UpdateLotResultCode
func (e UpdateLotResultCode) ValidEnum(v int32) bool {
	_, ok := updateLotResultCodeMap[v]
	return ok
}
func (e UpdateLotResultCode) isFlag() bool {
	for i := len(UpdateLotResultCodeAll) - 1; i >= 0; i-- {
		expected := UpdateLotResultCode(2) << uint64(len(UpdateLotResultCodeAll)-1) >> uint64(len(UpdateLotResultCodeAll)-i)
		if expected != UpdateLotResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e UpdateLotResultCode) String() string {
	name, _ := updateLotResultCodeMap[int32(e)]
	return name
}

func (e UpdateLotResultCode) ShortString() string {
	name, _ := updateLotResultCodeShortMap[int32(e)]
	return name
}

func (e UpdateLotResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range UpdateLotResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *UpdateLotResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = UpdateLotResultCode(t.Value)
	return nil
}

// UpdateLotResult is an XDR Union defines as:
//
//   union UpdateLotResult switch (UpdateLotResultCode code)
//        {
//        case UPDATE_LOT_SUCCESS:
//            UpdateLotSuccess success;
//        default:
//            void;
//        };
//
type UpdateLotResult struct {
	Code    UpdateLotResultCode `json:"code,omitempty"`
	Success *UpdateLotSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateLotResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateLotResult
func (u UpdateLotResult) ArmForSwitch(sw int32) (string, bool) {
	switch UpdateLotResultCode(sw) {
	case UpdateLotResultCodeUpdateLotSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewUpdateLotResult creates a new  UpdateLotResult.
func NewUpdateLotResult(code UpdateLotResultCode, value interface{}) (result UpdateLotResult, err error) {
	result.Code = code
	switch UpdateLotResultCode(code) {
	case UpdateLotResultCodeUpdateLotSuccess:
		tv, ok := value.(UpdateLotSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateLotSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u UpdateLotResult) MustSuccess() UpdateLotSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u UpdateLotResult) GetSuccess() (result UpdateLotSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ErrorCode is an XDR Enum defines as:
//
//   enum ErrorCode
//    {
//        ERR_MISC = 0, // Unspecific error
//        ERR_DATA = 1, // Malformed data
//        ERR_CONF = 2, // Misconfiguration error
//        ERR_AUTH = 3, // Authentication failure
//        ERR_LOAD = 4  // System overloaded
//    };
//
type ErrorCode int32

const (
	ErrorCodeErrMisc ErrorCode = 0
	ErrorCodeErrData ErrorCode = 1
	ErrorCodeErrConf ErrorCode = 2
	ErrorCodeErrAuth ErrorCode = 3
	ErrorCodeErrLoad ErrorCode = 4
)

var ErrorCodeAll = []ErrorCode{
	ErrorCodeErrMisc,
	ErrorCodeErrData,
	ErrorCodeErrConf,
	ErrorCodeErrAuth,
	ErrorCodeErrLoad,
}

var errorCodeMap = map[int32]string{
	0: "ErrorCodeErrMisc",
	1: "ErrorCodeErrData",
	2: "ErrorCodeErrConf",
	3: "ErrorCodeErrAuth",
	4: "ErrorCodeErrLoad",
}

var errorCodeShortMap = map[int32]string{
	0: "err_misc",
	1: "err_data",
	2: "err_conf",
	3: "err_auth",
	4: "err_load",
}

var errorCodeRevMap = map[string]int32{
	"ErrorCodeErrMisc": 0,
	"ErrorCodeErrData": 1,
	"ErrorCodeErrConf": 2,
	"ErrorCodeErrAuth": 3,
	"ErrorCodeErrLoad": 4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ErrorCode
func (e ErrorCode) ValidEnum(v int32) bool {
	_, ok := errorCodeMap[v]
	return ok
}
func (e ErrorCode) isFlag() bool {
	for i := len(ErrorCodeAll) - 1; i >= 0; i-- {
		expected := ErrorCode(2) << uint64(len(ErrorCodeAll)-1) >> uint64(len(ErrorCodeAll)-i)
		if expected != ErrorCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ErrorCode) String() string {
	name, _ := errorCodeMap[int32(e)]
	return name
}

func (e ErrorCode) ShortString() string {
	name, _ := errorCodeShortMap[int32(e)]
	return name
}

func (e ErrorCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ErrorCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ErrorCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ErrorCode(t.Value)
	return nil
}

// ErrorExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type ErrorExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ErrorExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ErrorExt
func (u ErrorExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewErrorExt creates a new  ErrorExt.
func NewErrorExt(v LedgerVersion, value interface{}) (result ErrorExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// Error is an XDR Struct defines as:
//
//   struct Error
//    {
//        ErrorCode code;
//        string msg<100>;
//        union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//    };
//
type Error struct {
	Code ErrorCode `json:"code,omitempty"`
	Msg  string    `json:"msg,omitempty" xdrmaxsize:"100"`
	Ext  ErrorExt  `json:"ext,omitempty"`
}

// AuthCertExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type AuthCertExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AuthCertExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AuthCertExt
func (u AuthCertExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAuthCertExt creates a new  AuthCertExt.
func NewAuthCertExt(v LedgerVersion, value interface{}) (result AuthCertExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AuthCert is an XDR Struct defines as:
//
//   struct AuthCert
//    {
//        Curve25519Public pubkey;
//        uint64 expiration;
//        Signature sig;
//        union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//    };
//
type AuthCert struct {
	Pubkey     Curve25519Public `json:"pubkey,omitempty"`
	Expiration Uint64           `json:"expiration,omitempty"`
	Sig        Signature        `json:"sig,omitempty"`
	Ext        AuthCertExt      `json:"ext,omitempty"`
}

// HelloExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type HelloExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u HelloExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of HelloExt
func (u HelloExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewHelloExt creates a new  HelloExt.
func NewHelloExt(v LedgerVersion, value interface{}) (result HelloExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// Hello is an XDR Struct defines as:
//
//   struct Hello
//    {
//        uint32 ledgerVersion;
//        uint32 overlayVersion;
//        uint32 overlayMinVersion;
//        Hash networkID;
//        string versionStr<100>;
//        int listeningPort;
//        NodeID peerID;
//        AuthCert cert;
//        uint256 nonce;
//        union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//    };
//
type Hello struct {
	LedgerVersion     Uint32   `json:"ledgerVersion,omitempty"`
	OverlayVersion    Uint32   `json:"overlayVersion,omitempty"`
	OverlayMinVersion Uint32   `json:"overlayMinVersion,omitempty"`
	NetworkId         Hash     `json:"networkID,omitempty"`
	VersionStr        string   `json:"versionStr,omitempty" xdrmaxsize:"100"`
	ListeningPort     int32    `json:"listeningPort,omitempty"`
	PeerId            NodeId   `json:"peerID,omitempty"`
	Cert              AuthCert `json:"cert,omitempty"`
	Nonce             Uint256  `json:"nonce,omitempty"`
	Ext               HelloExt `json:"ext,omitempty"`
}

// AuthExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type AuthExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AuthExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AuthExt
func (u AuthExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAuthExt creates a new  AuthExt.
func NewAuthExt(v LedgerVersion, value interface{}) (result AuthExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// Auth is an XDR Struct defines as:
//
//   struct Auth
//    {
//        // Empty message, just to confirm
//        // establishment of MAC keys.
//        int unused;
//        union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//    };
//
type Auth struct {
	Unused int32   `json:"unused,omitempty"`
	Ext    AuthExt `json:"ext,omitempty"`
}

// IpAddrType is an XDR Enum defines as:
//
//   enum IPAddrType
//    {
//        IPv4 = 0,
//        IPv6 = 1
//    };
//
type IpAddrType int32

const (
	IpAddrTypeIPv4 IpAddrType = 0
	IpAddrTypeIPv6 IpAddrType = 1
)

var IpAddrTypeAll = []IpAddrType{
	IpAddrTypeIPv4,
	IpAddrTypeIPv6,
}

var ipAddrTypeMap = map[int32]string{
	0: "IpAddrTypeIPv4",
	1: "IpAddrTypeIPv6",
}

var ipAddrTypeShortMap = map[int32]string{
	0: "i_pv4",
	1: "i_pv6",
}

var ipAddrTypeRevMap = map[string]int32{
	"IpAddrTypeIPv4": 0,
	"IpAddrTypeIPv6": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for IpAddrType
func (e IpAddrType) ValidEnum(v int32) bool {
	_, ok := ipAddrTypeMap[v]
	return ok
}
func (e IpAddrType) isFlag() bool {
	for i := len(IpAddrTypeAll) - 1; i >= 0; i-- {
		expected := IpAddrType(2) << uint64(len(IpAddrTypeAll)-1) >> uint64(len(IpAddrTypeAll)-i)
		if expected != IpAddrTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e IpAddrType) String() string {
	name, _ := ipAddrTypeMap[int32(e)]
	return name
}

func (e IpAddrType) ShortString() string {
	name, _ := ipAddrTypeShortMap[int32(e)]
	return name
}

func (e IpAddrType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range IpAddrTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *IpAddrType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = IpAddrType(t.Value)
	return nil
}

// PeerAddressIp is an XDR NestedUnion defines as:
//
//   union switch (IPAddrType type)
//        {
//        case IPv4:
//            opaque ipv4[4];
//        case IPv6:
//            opaque ipv6[16];
//        }
//
type PeerAddressIp struct {
	Type IpAddrType `json:"type,omitempty"`
	Ipv4 *[4]byte   `json:"ipv4,omitempty"`
	Ipv6 *[16]byte  `json:"ipv6,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PeerAddressIp) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PeerAddressIp
func (u PeerAddressIp) ArmForSwitch(sw int32) (string, bool) {
	switch IpAddrType(sw) {
	case IpAddrTypeIPv4:
		return "Ipv4", true
	case IpAddrTypeIPv6:
		return "Ipv6", true
	}
	return "-", false
}

// NewPeerAddressIp creates a new  PeerAddressIp.
func NewPeerAddressIp(aType IpAddrType, value interface{}) (result PeerAddressIp, err error) {
	result.Type = aType
	switch IpAddrType(aType) {
	case IpAddrTypeIPv4:
		tv, ok := value.([4]byte)
		if !ok {
			err = fmt.Errorf("invalid value, must be [4]byte")
			return
		}
		result.Ipv4 = &tv
	case IpAddrTypeIPv6:
		tv, ok := value.([16]byte)
		if !ok {
			err = fmt.Errorf("invalid value, must be [16]byte")
			return
		}
		result.Ipv6 = &tv
	}
	return
}

// MustIpv4 retrieves the Ipv4 value from the union,
// panicing if the value is not set.
func (u PeerAddressIp) MustIpv4() [4]byte {
	val, ok := u.GetIpv4()

	if !ok {
		panic("arm Ipv4 is not set")
	}

	return val
}

// GetIpv4 retrieves the Ipv4 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PeerAddressIp) GetIpv4() (result [4]byte, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ipv4" {
		result = *u.Ipv4
		ok = true
	}

	return
}

// MustIpv6 retrieves the Ipv6 value from the union,
// panicing if the value is not set.
func (u PeerAddressIp) MustIpv6() [16]byte {
	val, ok := u.GetIpv6()

	if !ok {
		panic("arm Ipv6 is not set")
	}

	return val
}

// GetIpv6 retrieves the Ipv6 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PeerAddressIp) GetIpv6() (result [16]byte, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ipv6" {
		result = *u.Ipv6
		ok = true
	}

	return
}

// PeerAddressExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type PeerAddressExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PeerAddressExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PeerAddressExt
func (u PeerAddressExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPeerAddressExt creates a new  PeerAddressExt.
func NewPeerAddressExt(v LedgerVersion, value interface{}) (result PeerAddressExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PeerAddress is an XDR Struct defines as:
//
//   struct PeerAddress
//    {
//        union switch (IPAddrType type)
//        {
//        case IPv4:
//            opaque ipv4[4];
//        case IPv6:
//            opaque ipv6[16];
//        }
//        ip;
//        uint32 port;
//        uint32 numFailures;
//        union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//    };
//
type PeerAddress struct {
	Ip          PeerAddressIp  `json:"ip,omitempty"`
	Port        Uint32         `json:"port,omitempty"`
	NumFailures Uint32         `json:"numFailures,omitempty"`
	Ext         PeerAddressExt `json:"ext,omitempty"`
}

// MessageType is an XDR Enum defines as:
//
//   enum MessageType
//    {
//        ERROR_MSG = 0,
//        AUTH = 2,
//        DONT_HAVE = 3,
//
//        GET_PEERS = 4, // gets a list of peers this guy knows about
//        PEERS = 5,
//
//        GET_TX_SET = 6, // gets a particular txset by hash
//        TX_SET = 7,
//
//        TRANSACTION = 8, // pass on a tx you have heard about
//
//        // SCP
//        GET_SCP_QUORUMSET = 9,
//        SCP_QUORUMSET = 10,
//        SCP_MESSAGE = 11,
//        GET_SCP_STATE = 12,
//
//        // new messages
//        HELLO = 13
//    };
//
type MessageType int32

const (
	MessageTypeErrorMsg        MessageType = 0
	MessageTypeAuth            MessageType = 2
	MessageTypeDontHave        MessageType = 3
	MessageTypeGetPeers        MessageType = 4
	MessageTypePeers           MessageType = 5
	MessageTypeGetTxSet        MessageType = 6
	MessageTypeTxSet           MessageType = 7
	MessageTypeTransaction     MessageType = 8
	MessageTypeGetScpQuorumset MessageType = 9
	MessageTypeScpQuorumset    MessageType = 10
	MessageTypeScpMessage      MessageType = 11
	MessageTypeGetScpState     MessageType = 12
	MessageTypeHello           MessageType = 13
)

var MessageTypeAll = []MessageType{
	MessageTypeErrorMsg,
	MessageTypeAuth,
	MessageTypeDontHave,
	MessageTypeGetPeers,
	MessageTypePeers,
	MessageTypeGetTxSet,
	MessageTypeTxSet,
	MessageTypeTransaction,
	MessageTypeGetScpQuorumset,
	MessageTypeScpQuorumset,
	MessageTypeScpMessage,
	MessageTypeGetScpState,
	MessageTypeHello,
}

var messageTypeMap = map[int32]string{
	0:  "MessageTypeErrorMsg",
	2:  "MessageTypeAuth",
	3:  "MessageTypeDontHave",
	4:  "MessageTypeGetPeers",
	5:  "MessageTypePeers",
	6:  "MessageTypeGetTxSet",
	7:  "MessageTypeTxSet",
	8:  "MessageTypeTransaction",
	9:  "MessageTypeGetScpQuorumset",
	10: "MessageTypeScpQuorumset",
	11: "MessageTypeScpMessage",
	12: "MessageTypeGetScpState",
	13: "MessageTypeHello",
}

var messageTypeShortMap = map[int32]string{
	0:  "error_msg",
	2:  "auth",
	3:  "dont_have",
	4:  "get_peers",
	5:  "peers",
	6:  "get_tx_set",
	7:  "tx_set",
	8:  "transaction",
	9:  "get_scp_quorumset",
	10: "scp_quorumset",
	11: "scp_message",
	12: "get_scp_state",
	13: "hello",
}

var messageTypeRevMap = map[string]int32{
	"MessageTypeErrorMsg":        0,
	"MessageTypeAuth":            2,
	"MessageTypeDontHave":        3,
	"MessageTypeGetPeers":        4,
	"MessageTypePeers":           5,
	"MessageTypeGetTxSet":        6,
	"MessageTypeTxSet":           7,
	"MessageTypeTransaction":     8,
	"MessageTypeGetScpQuorumset": 9,
	"MessageTypeScpQuorumset":    10,
	"MessageTypeScpMessage":      11,
	"MessageTypeGetScpState":     12,
	"MessageTypeHello":           13,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for MessageType
func (e MessageType) ValidEnum(v int32) bool {
	_, ok := messageTypeMap[v]
	return ok
}
func (e MessageType) isFlag() bool {
	for i := len(MessageTypeAll) - 1; i >= 0; i-- {
		expected := MessageType(2) << uint64(len(MessageTypeAll)-1) >> uint64(len(MessageTypeAll)-i)
		if expected != MessageTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e MessageType) String() string {
	name, _ := messageTypeMap[int32(e)]
	return name
}

func (e MessageType) ShortString() string {
	name, _ := messageTypeShortMap[int32(e)]
	return name
}

func (e MessageType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range MessageTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *MessageType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = MessageType(t.Value)
	return nil
}

// DontHave is an XDR Struct defines as:
//
//   struct DontHave
//    {
//        MessageType type;
//        uint256 reqHash;
//    };
//
type DontHave struct {
	Type    MessageType `json:"type,omitempty"`
	ReqHash Uint256     `json:"reqHash,omitempty"`
}

// StellarMessage is an XDR Union defines as:
//
//   union StellarMessage switch (MessageType type)
//    {
//    case ERROR_MSG:
//        Error error;
//    case HELLO:
//        Hello hello;
//    case AUTH:
//        Auth auth;
//    case DONT_HAVE:
//        DontHave dontHave;
//    case GET_PEERS:
//        void;
//    case PEERS:
//        PeerAddress peers<>;
//
//    case GET_TX_SET:
//        uint256 txSetHash;
//    case TX_SET:
//        TransactionSet txSet;
//
//    case TRANSACTION:
//        TransactionEnvelope transaction;
//
//    // SCP
//    case GET_SCP_QUORUMSET:
//        uint256 qSetHash;
//    case SCP_QUORUMSET:
//        SCPQuorumSet qSet;
//    case SCP_MESSAGE:
//        SCPEnvelope envelope;
//    case GET_SCP_STATE:
//        uint32 getSCPLedgerSeq; // ledger seq requested ; if 0, requests the latest
//    };
//
type StellarMessage struct {
	Type            MessageType          `json:"type,omitempty"`
	Error           *Error               `json:"error,omitempty"`
	Hello           *Hello               `json:"hello,omitempty"`
	Auth            *Auth                `json:"auth,omitempty"`
	DontHave        *DontHave            `json:"dontHave,omitempty"`
	Peers           *[]PeerAddress       `json:"peers,omitempty"`
	TxSetHash       *Uint256             `json:"txSetHash,omitempty"`
	TxSet           *TransactionSet      `json:"txSet,omitempty"`
	Transaction     *TransactionEnvelope `json:"transaction,omitempty"`
	QSetHash        *Uint256             `json:"qSetHash,omitempty"`
	QSet            *ScpQuorumSet        `json:"qSet,omitempty"`
	Envelope        *ScpEnvelope         `json:"envelope,omitempty"`
	GetScpLedgerSeq *Uint32              `json:"getSCPLedgerSeq,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u StellarMessage) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of StellarMessage
func (u StellarMessage) ArmForSwitch(sw int32) (string, bool) {
	switch MessageType(sw) {
	case MessageTypeErrorMsg:
		return "Error", true
	case MessageTypeHello:
		return "Hello", true
	case MessageTypeAuth:
		return "Auth", true
	case MessageTypeDontHave:
		return "DontHave", true
	case MessageTypeGetPeers:
		return "", true
	case MessageTypePeers:
		return "Peers", true
	case MessageTypeGetTxSet:
		return "TxSetHash", true
	case MessageTypeTxSet:
		return "TxSet", true
	case MessageTypeTransaction:
		return "Transaction", true
	case MessageTypeGetScpQuorumset:
		return "QSetHash", true
	case MessageTypeScpQuorumset:
		return "QSet", true
	case MessageTypeScpMessage:
		return "Envelope", true
	case MessageTypeGetScpState:
		return "GetScpLedgerSeq", true
	}
	return "-", false
}

// NewStellarMessage creates a new  StellarMessage.
func NewStellarMessage(aType MessageType, value interface{}) (result StellarMessage, err error) {
	result.Type = aType
	switch MessageType(aType) {
	case MessageTypeErrorMsg:
		tv, ok := value.(Error)
		if !ok {
			err = fmt.Errorf("invalid value, must be Error")
			return
		}
		result.Error = &tv
	case MessageTypeHello:
		tv, ok := value.(Hello)
		if !ok {
			err = fmt.Errorf("invalid value, must be Hello")
			return
		}
		result.Hello = &tv
	case MessageTypeAuth:
		tv, ok := value.(Auth)
		if !ok {
			err = fmt.Errorf("invalid value, must be Auth")
			return
		}
		result.Auth = &tv
	case MessageTypeDontHave:
		tv, ok := value.(DontHave)
		if !ok {
			err = fmt.Errorf("invalid value, must be DontHave")
			return
		}
		result.DontHave = &tv
	case MessageTypeGetPeers:
		// void
	case MessageTypePeers:
		tv, ok := value.([]PeerAddress)
		if !ok {
			err = fmt.Errorf("invalid value, must be []PeerAddress")
			return
		}
		result.Peers = &tv
	case MessageTypeGetTxSet:
		tv, ok := value.(Uint256)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint256")
			return
		}
		result.TxSetHash = &tv
	case MessageTypeTxSet:
		tv, ok := value.(TransactionSet)
		if !ok {
			err = fmt.Errorf("invalid value, must be TransactionSet")
			return
		}
		result.TxSet = &tv
	case MessageTypeTransaction:
		tv, ok := value.(TransactionEnvelope)
		if !ok {
			err = fmt.Errorf("invalid value, must be TransactionEnvelope")
			return
		}
		result.Transaction = &tv
	case MessageTypeGetScpQuorumset:
		tv, ok := value.(Uint256)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint256")
			return
		}
		result.QSetHash = &tv
	case MessageTypeScpQuorumset:
		tv, ok := value.(ScpQuorumSet)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpQuorumSet")
			return
		}
		result.QSet = &tv
	case MessageTypeScpMessage:
		tv, ok := value.(ScpEnvelope)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpEnvelope")
			return
		}
		result.Envelope = &tv
	case MessageTypeGetScpState:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.GetScpLedgerSeq = &tv
	}
	return
}

// MustError retrieves the Error value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustError() Error {
	val, ok := u.GetError()

	if !ok {
		panic("arm Error is not set")
	}

	return val
}

// GetError retrieves the Error value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetError() (result Error, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Error" {
		result = *u.Error
		ok = true
	}

	return
}

// MustHello retrieves the Hello value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustHello() Hello {
	val, ok := u.GetHello()

	if !ok {
		panic("arm Hello is not set")
	}

	return val
}

// GetHello retrieves the Hello value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetHello() (result Hello, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Hello" {
		result = *u.Hello
		ok = true
	}

	return
}

// MustAuth retrieves the Auth value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustAuth() Auth {
	val, ok := u.GetAuth()

	if !ok {
		panic("arm Auth is not set")
	}

	return val
}

// GetAuth retrieves the Auth value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetAuth() (result Auth, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Auth" {
		result = *u.Auth
		ok = true
	}

	return
}

// MustDontHave retrieves the DontHave value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustDontHave() DontHave {
	val, ok := u.GetDontHave()

	if !ok {
		panic("arm DontHave is not set")
	}

	return val
}

// GetDontHave retrieves the DontHave value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetDontHave() (result DontHave, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "DontHave" {
		result = *u.DontHave
		ok = true
	}

	return
}

// MustPeers retrieves the Peers value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustPeers() []PeerAddress {
	val, ok := u.GetPeers()

	if !ok {
		panic("arm Peers is not set")
	}

	return val
}

// GetPeers retrieves the Peers value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetPeers() (result []PeerAddress, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Peers" {
		result = *u.Peers
		ok = true
	}

	return
}

// MustTxSetHash retrieves the TxSetHash value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustTxSetHash() Uint256 {
	val, ok := u.GetTxSetHash()

	if !ok {
		panic("arm TxSetHash is not set")
	}

	return val
}

// GetTxSetHash retrieves the TxSetHash value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetTxSetHash() (result Uint256, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "TxSetHash" {
		result = *u.TxSetHash
		ok = true
	}

	return
}

// MustTxSet retrieves the TxSet value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustTxSet() TransactionSet {
	val, ok := u.GetTxSet()

	if !ok {
		panic("arm TxSet is not set")
	}

	return val
}

// GetTxSet retrieves the TxSet value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetTxSet() (result TransactionSet, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "TxSet" {
		result = *u.TxSet
		ok = true
	}

	return
}

// MustTransaction retrieves the Transaction value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustTransaction() TransactionEnvelope {
	val, ok := u.GetTransaction()

	if !ok {
		panic("arm Transaction is not set")
	}

	return val
}

// GetTransaction retrieves the Transaction value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetTransaction() (result TransactionEnvelope, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Transaction" {
		result = *u.Transaction
		ok = true
	}

	return
}

// MustQSetHash retrieves the QSetHash value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustQSetHash() Uint256 {
	val, ok := u.GetQSetHash()

	if !ok {
		panic("arm QSetHash is not set")
	}

	return val
}

// GetQSetHash retrieves the QSetHash value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetQSetHash() (result Uint256, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "QSetHash" {
		result = *u.QSetHash
		ok = true
	}

	return
}

// MustQSet retrieves the QSet value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustQSet() ScpQuorumSet {
	val, ok := u.GetQSet()

	if !ok {
		panic("arm QSet is not set")
	}

	return val
}

// GetQSet retrieves the QSet value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetQSet() (result ScpQuorumSet, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "QSet" {
		result = *u.QSet
		ok = true
	}

	return
}

// MustEnvelope retrieves the Envelope value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustEnvelope() ScpEnvelope {
	val, ok := u.GetEnvelope()

	if !ok {
		panic("arm Envelope is not set")
	}

	return val
}

// GetEnvelope retrieves the Envelope value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetEnvelope() (result ScpEnvelope, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Envelope" {
		result = *u.Envelope
		ok = true
	}

	return
}

// MustGetScpLedgerSeq retrieves the GetScpLedgerSeq value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustGetScpLedgerSeq() Uint32 {
	val, ok := u.GetGetScpLedgerSeq()

	if !ok {
		panic("arm GetScpLedgerSeq is not set")
	}

	return val
}

// GetGetScpLedgerSeq retrieves the GetScpLedgerSeq value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetGetScpLedgerSeq() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "GetScpLedgerSeq" {
		result = *u.GetScpLedgerSeq
		ok = true
	}

	return
}

// AuthenticatedMessageV0 is an XDR NestedStruct defines as:
//
//   struct
//    {
//       uint64 sequence;
//       StellarMessage message;
//       HmacSha256Mac mac;
//        }
//
type AuthenticatedMessageV0 struct {
	Sequence Uint64         `json:"sequence,omitempty"`
	Message  StellarMessage `json:"message,omitempty"`
	Mac      HmacSha256Mac  `json:"mac,omitempty"`
}

// AuthenticatedMessage is an XDR Union defines as:
//
//   union AuthenticatedMessage switch (LedgerVersion v)
//    {
//    case EMPTY_VERSION:
//        struct
//    {
//       uint64 sequence;
//       StellarMessage message;
//       HmacSha256Mac mac;
//        } v0;
//    };
//
type AuthenticatedMessage struct {
	V  LedgerVersion           `json:"v,omitempty"`
	V0 *AuthenticatedMessageV0 `json:"v0,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AuthenticatedMessage) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AuthenticatedMessage
func (u AuthenticatedMessage) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "V0", true
	}
	return "-", false
}

// NewAuthenticatedMessage creates a new  AuthenticatedMessage.
func NewAuthenticatedMessage(v LedgerVersion, value interface{}) (result AuthenticatedMessage, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		tv, ok := value.(AuthenticatedMessageV0)
		if !ok {
			err = fmt.Errorf("invalid value, must be AuthenticatedMessageV0")
			return
		}
		result.V0 = &tv
	}
	return
}

// MustV0 retrieves the V0 value from the union,
// panicing if the value is not set.
func (u AuthenticatedMessage) MustV0() AuthenticatedMessageV0 {
	val, ok := u.GetV0()

	if !ok {
		panic("arm V0 is not set")
	}

	return val
}

// GetV0 retrieves the V0 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AuthenticatedMessage) GetV0() (result AuthenticatedMessageV0, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "V0" {
		result = *u.V0
		ok = true
	}

	return
}

// BidRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type BidRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u BidRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of BidRequestExt
func (u BidRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewBidRequestExt creates a new  BidRequestExt.
func NewBidRequestExt(v LedgerVersion, value interface{}) (result BidRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// BidRequest is an XDR Struct defines as:
//
//   struct BidRequest
//    {
//        uint64 lotID;
//        uint64 amount;      // amount to pay
//        bool isBuyNow;      // buy now request
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type BidRequest struct {
	LotId    Uint64        `json:"lotID,omitempty"`
	Amount   Uint64        `json:"amount,omitempty"`
	IsBuyNow bool          `json:"isBuyNow,omitempty"`
	Ext      BidRequestExt `json:"ext,omitempty"`
}

// LotCreationRequestExt is an XDR NestedUnion defines as:
//
//   union switch(LedgerVersion v)
//        {
//         case EMPTY_VERSION:
//            void;
//        }
//
type LotCreationRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LotCreationRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LotCreationRequestExt
func (u LotCreationRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLotCreationRequestExt creates a new  LotCreationRequestExt.
func NewLotCreationRequestExt(v LedgerVersion, value interface{}) (result LotCreationRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LotCreationRequest is an XDR Struct defines as:
//
//   struct LotCreationRequest
//    {
//        LotEntry lot;
//
//        // reserved for future use
//        union switch(LedgerVersion v)
//        {
//         case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type LotCreationRequest struct {
	Lot LotEntry              `json:"lot,omitempty"`
	Ext LotCreationRequestExt `json:"ext,omitempty"`
}

// ParticipationRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ParticipationRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ParticipationRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ParticipationRequestExt
func (u ParticipationRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewParticipationRequestExt creates a new  ParticipationRequestExt.
func NewParticipationRequestExt(v LedgerVersion, value interface{}) (result ParticipationRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ParticipationRequest is an XDR Struct defines as:
//
//   struct ParticipationRequest
//    {
//        uint64 lotID;
//        bool isBuyNow;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ParticipationRequest struct {
	LotId    Uint64                  `json:"lotID,omitempty"`
	IsBuyNow bool                    `json:"isBuyNow,omitempty"`
	Ext      ParticipationRequestExt `json:"ext,omitempty"`
}

// OperationType is an XDR Enum defines as:
//
//   enum OperationType
//        {
//            CREATE_ACCOUNT = 0,
//            PAYMENT = 1,
//            PATH_PAYMENT = 2,
//            MANAGE_OFFER = 3,
//            CREATE_PASSIVE_OFFER = 4,
//            SET_OPTIONS = 5,
//            CHANGE_TRUST = 6,
//            ALLOW_TRUST = 7,
//            CREATE_LOT = 8,
//            REQUEST_PARTICIPATION = 9,
//            SEND_MESSAGE = 10,
//            CREATE_BID = 11,
//            RECOVER = 12,
//            MANAGE_MIGRATION = 13,
//            REVIEW_REQUEST = 14,
//            CANCEL_REQUEST = 15,
//            MANAGE_PARTICIPANT = 16,
//            REQUEST_BUY_NOW = 17,
//            CLOSE_AUCTION = 18,
//            UPDATE_LOT = 19,
//            MANAGE_PERMISSIONS = 20,
//            MANAGE_BID_LIMIT = 21,
//            SET_MAX_BID = 22
//        };
//
type OperationType int32

const (
	OperationTypeCreateAccount        OperationType = 0
	OperationTypePayment              OperationType = 1
	OperationTypePathPayment          OperationType = 2
	OperationTypeManageOffer          OperationType = 3
	OperationTypeCreatePassiveOffer   OperationType = 4
	OperationTypeSetOptions           OperationType = 5
	OperationTypeChangeTrust          OperationType = 6
	OperationTypeAllowTrust           OperationType = 7
	OperationTypeCreateLot            OperationType = 8
	OperationTypeRequestParticipation OperationType = 9
	OperationTypeSendMessage          OperationType = 10
	OperationTypeCreateBid            OperationType = 11
	OperationTypeRecover              OperationType = 12
	OperationTypeManageMigration      OperationType = 13
	OperationTypeReviewRequest        OperationType = 14
	OperationTypeCancelRequest        OperationType = 15
	OperationTypeManageParticipant    OperationType = 16
	OperationTypeRequestBuyNow        OperationType = 17
	OperationTypeCloseAuction         OperationType = 18
	OperationTypeUpdateLot            OperationType = 19
	OperationTypeManagePermissions    OperationType = 20
	OperationTypeManageBidLimit       OperationType = 21
	OperationTypeSetMaxBid            OperationType = 22
)

var OperationTypeAll = []OperationType{
	OperationTypeCreateAccount,
	OperationTypePayment,
	OperationTypePathPayment,
	OperationTypeManageOffer,
	OperationTypeCreatePassiveOffer,
	OperationTypeSetOptions,
	OperationTypeChangeTrust,
	OperationTypeAllowTrust,
	OperationTypeCreateLot,
	OperationTypeRequestParticipation,
	OperationTypeSendMessage,
	OperationTypeCreateBid,
	OperationTypeRecover,
	OperationTypeManageMigration,
	OperationTypeReviewRequest,
	OperationTypeCancelRequest,
	OperationTypeManageParticipant,
	OperationTypeRequestBuyNow,
	OperationTypeCloseAuction,
	OperationTypeUpdateLot,
	OperationTypeManagePermissions,
	OperationTypeManageBidLimit,
	OperationTypeSetMaxBid,
}

var operationTypeMap = map[int32]string{
	0:  "OperationTypeCreateAccount",
	1:  "OperationTypePayment",
	2:  "OperationTypePathPayment",
	3:  "OperationTypeManageOffer",
	4:  "OperationTypeCreatePassiveOffer",
	5:  "OperationTypeSetOptions",
	6:  "OperationTypeChangeTrust",
	7:  "OperationTypeAllowTrust",
	8:  "OperationTypeCreateLot",
	9:  "OperationTypeRequestParticipation",
	10: "OperationTypeSendMessage",
	11: "OperationTypeCreateBid",
	12: "OperationTypeRecover",
	13: "OperationTypeManageMigration",
	14: "OperationTypeReviewRequest",
	15: "OperationTypeCancelRequest",
	16: "OperationTypeManageParticipant",
	17: "OperationTypeRequestBuyNow",
	18: "OperationTypeCloseAuction",
	19: "OperationTypeUpdateLot",
	20: "OperationTypeManagePermissions",
	21: "OperationTypeManageBidLimit",
	22: "OperationTypeSetMaxBid",
}

var operationTypeShortMap = map[int32]string{
	0:  "create_account",
	1:  "payment",
	2:  "path_payment",
	3:  "manage_offer",
	4:  "create_passive_offer",
	5:  "set_options",
	6:  "change_trust",
	7:  "allow_trust",
	8:  "create_lot",
	9:  "request_participation",
	10: "send_message",
	11: "create_bid",
	12: "recover",
	13: "manage_migration",
	14: "review_request",
	15: "cancel_request",
	16: "manage_participant",
	17: "request_buy_now",
	18: "close_auction",
	19: "update_lot",
	20: "manage_permissions",
	21: "manage_bid_limit",
	22: "set_max_bid",
}

var operationTypeRevMap = map[string]int32{
	"OperationTypeCreateAccount":        0,
	"OperationTypePayment":              1,
	"OperationTypePathPayment":          2,
	"OperationTypeManageOffer":          3,
	"OperationTypeCreatePassiveOffer":   4,
	"OperationTypeSetOptions":           5,
	"OperationTypeChangeTrust":          6,
	"OperationTypeAllowTrust":           7,
	"OperationTypeCreateLot":            8,
	"OperationTypeRequestParticipation": 9,
	"OperationTypeSendMessage":          10,
	"OperationTypeCreateBid":            11,
	"OperationTypeRecover":              12,
	"OperationTypeManageMigration":      13,
	"OperationTypeReviewRequest":        14,
	"OperationTypeCancelRequest":        15,
	"OperationTypeManageParticipant":    16,
	"OperationTypeRequestBuyNow":        17,
	"OperationTypeCloseAuction":         18,
	"OperationTypeUpdateLot":            19,
	"OperationTypeManagePermissions":    20,
	"OperationTypeManageBidLimit":       21,
	"OperationTypeSetMaxBid":            22,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for OperationType
func (e OperationType) ValidEnum(v int32) bool {
	_, ok := operationTypeMap[v]
	return ok
}
func (e OperationType) isFlag() bool {
	for i := len(OperationTypeAll) - 1; i >= 0; i-- {
		expected := OperationType(2) << uint64(len(OperationTypeAll)-1) >> uint64(len(OperationTypeAll)-i)
		if expected != OperationTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e OperationType) String() string {
	name, _ := operationTypeMap[int32(e)]
	return name
}

func (e OperationType) ShortString() string {
	name, _ := operationTypeShortMap[int32(e)]
	return name
}

func (e OperationType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range OperationTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *OperationType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = OperationType(t.Value)
	return nil
}

// OperationBody is an XDR NestedUnion defines as:
//
//   union switch (OperationType type)
//            {
//            case CREATE_ACCOUNT:
//                CreateAccountOp createAccountOp;
//            case PAYMENT:
//                PaymentOp paymentOp;
//            case PATH_PAYMENT:
//                PathPaymentOp pathPaymentOp;
//            case MANAGE_OFFER:
//                ManageOfferOp manageOfferOp;
//            case CREATE_PASSIVE_OFFER:
//                CreatePassiveOfferOp createPassiveOfferOp;
//            case SET_OPTIONS:
//                SetOptionsOp setOptionsOp;
//            case MANAGE_PERMISSIONS:
//                ManagePermissionsOp managePermissionsOp;
//            case CHANGE_TRUST:
//                ChangeTrustOp changeTrustOp;
//            case ALLOW_TRUST:
//                AllowTrustOp allowTrustOp;
//            case CREATE_LOT:
//                CreateLotOp createLotOp;
//            case UPDATE_LOT:
//                UpdateLotOp updateLotOp;
//            case REQUEST_PARTICIPATION:
//                RequestParticipationOp requestParticipationOp;
//            case SEND_MESSAGE:
//                SendMessageOp sendMessageOp;
//            case CREATE_BID:
//                CreateBidOp createBidOp;
//            case RECOVER:
//                RecoverOp recoverOp;
//            case MANAGE_MIGRATION:
//                ManageMigrationOp manageMigrationOp;
//            case REVIEW_REQUEST:
//                ReviewRequestOp reviewRequestOp;
//            case CANCEL_REQUEST:
//                CancelRequestOp cancelRequestOp;
//            case MANAGE_PARTICIPANT:
//                ManageParticipantOp manageParticipantOp;
//            case MANAGE_BID_LIMIT:
//                ManageBidLimitOp manageBidLimitOp;
//            case CLOSE_AUCTION:
//                CloseAuctionOp closeAuctionOp;
//            case REQUEST_BUY_NOW:
//                RequestBuyNowOp requestBuyNowOp;
//            case SET_MAX_BID:
//                SetMaxBidOp setMaxBidOp;
//            }
//
type OperationBody struct {
	Type                   OperationType           `json:"type,omitempty"`
	CreateAccountOp        *CreateAccountOp        `json:"createAccountOp,omitempty"`
	PaymentOp              *PaymentOp              `json:"paymentOp,omitempty"`
	PathPaymentOp          *PathPaymentOp          `json:"pathPaymentOp,omitempty"`
	ManageOfferOp          *ManageOfferOp          `json:"manageOfferOp,omitempty"`
	CreatePassiveOfferOp   *CreatePassiveOfferOp   `json:"createPassiveOfferOp,omitempty"`
	SetOptionsOp           *SetOptionsOp           `json:"setOptionsOp,omitempty"`
	ManagePermissionsOp    *ManagePermissionsOp    `json:"managePermissionsOp,omitempty"`
	ChangeTrustOp          *ChangeTrustOp          `json:"changeTrustOp,omitempty"`
	AllowTrustOp           *AllowTrustOp           `json:"allowTrustOp,omitempty"`
	CreateLotOp            *CreateLotOp            `json:"createLotOp,omitempty"`
	UpdateLotOp            *UpdateLotOp            `json:"updateLotOp,omitempty"`
	RequestParticipationOp *RequestParticipationOp `json:"requestParticipationOp,omitempty"`
	SendMessageOp          *SendMessageOp          `json:"sendMessageOp,omitempty"`
	CreateBidOp            *CreateBidOp            `json:"createBidOp,omitempty"`
	RecoverOp              *RecoverOp              `json:"recoverOp,omitempty"`
	ManageMigrationOp      *ManageMigrationOp      `json:"manageMigrationOp,omitempty"`
	ReviewRequestOp        *ReviewRequestOp        `json:"reviewRequestOp,omitempty"`
	CancelRequestOp        *CancelRequestOp        `json:"cancelRequestOp,omitempty"`
	ManageParticipantOp    *ManageParticipantOp    `json:"manageParticipantOp,omitempty"`
	ManageBidLimitOp       *ManageBidLimitOp       `json:"manageBidLimitOp,omitempty"`
	CloseAuctionOp         *CloseAuctionOp         `json:"closeAuctionOp,omitempty"`
	RequestBuyNowOp        *RequestBuyNowOp        `json:"requestBuyNowOp,omitempty"`
	SetMaxBidOp            *SetMaxBidOp            `json:"setMaxBidOp,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u OperationBody) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of OperationBody
func (u OperationBody) ArmForSwitch(sw int32) (string, bool) {
	switch OperationType(sw) {
	case OperationTypeCreateAccount:
		return "CreateAccountOp", true
	case OperationTypePayment:
		return "PaymentOp", true
	case OperationTypePathPayment:
		return "PathPaymentOp", true
	case OperationTypeManageOffer:
		return "ManageOfferOp", true
	case OperationTypeCreatePassiveOffer:
		return "CreatePassiveOfferOp", true
	case OperationTypeSetOptions:
		return "SetOptionsOp", true
	case OperationTypeManagePermissions:
		return "ManagePermissionsOp", true
	case OperationTypeChangeTrust:
		return "ChangeTrustOp", true
	case OperationTypeAllowTrust:
		return "AllowTrustOp", true
	case OperationTypeCreateLot:
		return "CreateLotOp", true
	case OperationTypeUpdateLot:
		return "UpdateLotOp", true
	case OperationTypeRequestParticipation:
		return "RequestParticipationOp", true
	case OperationTypeSendMessage:
		return "SendMessageOp", true
	case OperationTypeCreateBid:
		return "CreateBidOp", true
	case OperationTypeRecover:
		return "RecoverOp", true
	case OperationTypeManageMigration:
		return "ManageMigrationOp", true
	case OperationTypeReviewRequest:
		return "ReviewRequestOp", true
	case OperationTypeCancelRequest:
		return "CancelRequestOp", true
	case OperationTypeManageParticipant:
		return "ManageParticipantOp", true
	case OperationTypeManageBidLimit:
		return "ManageBidLimitOp", true
	case OperationTypeCloseAuction:
		return "CloseAuctionOp", true
	case OperationTypeRequestBuyNow:
		return "RequestBuyNowOp", true
	case OperationTypeSetMaxBid:
		return "SetMaxBidOp", true
	}
	return "-", false
}

// NewOperationBody creates a new  OperationBody.
func NewOperationBody(aType OperationType, value interface{}) (result OperationBody, err error) {
	result.Type = aType
	switch OperationType(aType) {
	case OperationTypeCreateAccount:
		tv, ok := value.(CreateAccountOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAccountOp")
			return
		}
		result.CreateAccountOp = &tv
	case OperationTypePayment:
		tv, ok := value.(PaymentOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be PaymentOp")
			return
		}
		result.PaymentOp = &tv
	case OperationTypePathPayment:
		tv, ok := value.(PathPaymentOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be PathPaymentOp")
			return
		}
		result.PathPaymentOp = &tv
	case OperationTypeManageOffer:
		tv, ok := value.(ManageOfferOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageOfferOp")
			return
		}
		result.ManageOfferOp = &tv
	case OperationTypeCreatePassiveOffer:
		tv, ok := value.(CreatePassiveOfferOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreatePassiveOfferOp")
			return
		}
		result.CreatePassiveOfferOp = &tv
	case OperationTypeSetOptions:
		tv, ok := value.(SetOptionsOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetOptionsOp")
			return
		}
		result.SetOptionsOp = &tv
	case OperationTypeManagePermissions:
		tv, ok := value.(ManagePermissionsOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManagePermissionsOp")
			return
		}
		result.ManagePermissionsOp = &tv
	case OperationTypeChangeTrust:
		tv, ok := value.(ChangeTrustOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ChangeTrustOp")
			return
		}
		result.ChangeTrustOp = &tv
	case OperationTypeAllowTrust:
		tv, ok := value.(AllowTrustOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be AllowTrustOp")
			return
		}
		result.AllowTrustOp = &tv
	case OperationTypeCreateLot:
		tv, ok := value.(CreateLotOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateLotOp")
			return
		}
		result.CreateLotOp = &tv
	case OperationTypeUpdateLot:
		tv, ok := value.(UpdateLotOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateLotOp")
			return
		}
		result.UpdateLotOp = &tv
	case OperationTypeRequestParticipation:
		tv, ok := value.(RequestParticipationOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be RequestParticipationOp")
			return
		}
		result.RequestParticipationOp = &tv
	case OperationTypeSendMessage:
		tv, ok := value.(SendMessageOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be SendMessageOp")
			return
		}
		result.SendMessageOp = &tv
	case OperationTypeCreateBid:
		tv, ok := value.(CreateBidOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateBidOp")
			return
		}
		result.CreateBidOp = &tv
	case OperationTypeRecover:
		tv, ok := value.(RecoverOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be RecoverOp")
			return
		}
		result.RecoverOp = &tv
	case OperationTypeManageMigration:
		tv, ok := value.(ManageMigrationOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageMigrationOp")
			return
		}
		result.ManageMigrationOp = &tv
	case OperationTypeReviewRequest:
		tv, ok := value.(ReviewRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReviewRequestOp")
			return
		}
		result.ReviewRequestOp = &tv
	case OperationTypeCancelRequest:
		tv, ok := value.(CancelRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CancelRequestOp")
			return
		}
		result.CancelRequestOp = &tv
	case OperationTypeManageParticipant:
		tv, ok := value.(ManageParticipantOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageParticipantOp")
			return
		}
		result.ManageParticipantOp = &tv
	case OperationTypeManageBidLimit:
		tv, ok := value.(ManageBidLimitOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageBidLimitOp")
			return
		}
		result.ManageBidLimitOp = &tv
	case OperationTypeCloseAuction:
		tv, ok := value.(CloseAuctionOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CloseAuctionOp")
			return
		}
		result.CloseAuctionOp = &tv
	case OperationTypeRequestBuyNow:
		tv, ok := value.(RequestBuyNowOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be RequestBuyNowOp")
			return
		}
		result.RequestBuyNowOp = &tv
	case OperationTypeSetMaxBid:
		tv, ok := value.(SetMaxBidOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetMaxBidOp")
			return
		}
		result.SetMaxBidOp = &tv
	}
	return
}

// MustCreateAccountOp retrieves the CreateAccountOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateAccountOp() CreateAccountOp {
	val, ok := u.GetCreateAccountOp()

	if !ok {
		panic("arm CreateAccountOp is not set")
	}

	return val
}

// GetCreateAccountOp retrieves the CreateAccountOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateAccountOp() (result CreateAccountOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateAccountOp" {
		result = *u.CreateAccountOp
		ok = true
	}

	return
}

// MustPaymentOp retrieves the PaymentOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustPaymentOp() PaymentOp {
	val, ok := u.GetPaymentOp()

	if !ok {
		panic("arm PaymentOp is not set")
	}

	return val
}

// GetPaymentOp retrieves the PaymentOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetPaymentOp() (result PaymentOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PaymentOp" {
		result = *u.PaymentOp
		ok = true
	}

	return
}

// MustPathPaymentOp retrieves the PathPaymentOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustPathPaymentOp() PathPaymentOp {
	val, ok := u.GetPathPaymentOp()

	if !ok {
		panic("arm PathPaymentOp is not set")
	}

	return val
}

// GetPathPaymentOp retrieves the PathPaymentOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetPathPaymentOp() (result PathPaymentOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PathPaymentOp" {
		result = *u.PathPaymentOp
		ok = true
	}

	return
}

// MustManageOfferOp retrieves the ManageOfferOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageOfferOp() ManageOfferOp {
	val, ok := u.GetManageOfferOp()

	if !ok {
		panic("arm ManageOfferOp is not set")
	}

	return val
}

// GetManageOfferOp retrieves the ManageOfferOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageOfferOp() (result ManageOfferOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageOfferOp" {
		result = *u.ManageOfferOp
		ok = true
	}

	return
}

// MustCreatePassiveOfferOp retrieves the CreatePassiveOfferOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreatePassiveOfferOp() CreatePassiveOfferOp {
	val, ok := u.GetCreatePassiveOfferOp()

	if !ok {
		panic("arm CreatePassiveOfferOp is not set")
	}

	return val
}

// GetCreatePassiveOfferOp retrieves the CreatePassiveOfferOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreatePassiveOfferOp() (result CreatePassiveOfferOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreatePassiveOfferOp" {
		result = *u.CreatePassiveOfferOp
		ok = true
	}

	return
}

// MustSetOptionsOp retrieves the SetOptionsOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustSetOptionsOp() SetOptionsOp {
	val, ok := u.GetSetOptionsOp()

	if !ok {
		panic("arm SetOptionsOp is not set")
	}

	return val
}

// GetSetOptionsOp retrieves the SetOptionsOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetSetOptionsOp() (result SetOptionsOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SetOptionsOp" {
		result = *u.SetOptionsOp
		ok = true
	}

	return
}

// MustManagePermissionsOp retrieves the ManagePermissionsOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManagePermissionsOp() ManagePermissionsOp {
	val, ok := u.GetManagePermissionsOp()

	if !ok {
		panic("arm ManagePermissionsOp is not set")
	}

	return val
}

// GetManagePermissionsOp retrieves the ManagePermissionsOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManagePermissionsOp() (result ManagePermissionsOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManagePermissionsOp" {
		result = *u.ManagePermissionsOp
		ok = true
	}

	return
}

// MustChangeTrustOp retrieves the ChangeTrustOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustChangeTrustOp() ChangeTrustOp {
	val, ok := u.GetChangeTrustOp()

	if !ok {
		panic("arm ChangeTrustOp is not set")
	}

	return val
}

// GetChangeTrustOp retrieves the ChangeTrustOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetChangeTrustOp() (result ChangeTrustOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ChangeTrustOp" {
		result = *u.ChangeTrustOp
		ok = true
	}

	return
}

// MustAllowTrustOp retrieves the AllowTrustOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustAllowTrustOp() AllowTrustOp {
	val, ok := u.GetAllowTrustOp()

	if !ok {
		panic("arm AllowTrustOp is not set")
	}

	return val
}

// GetAllowTrustOp retrieves the AllowTrustOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetAllowTrustOp() (result AllowTrustOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AllowTrustOp" {
		result = *u.AllowTrustOp
		ok = true
	}

	return
}

// MustCreateLotOp retrieves the CreateLotOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateLotOp() CreateLotOp {
	val, ok := u.GetCreateLotOp()

	if !ok {
		panic("arm CreateLotOp is not set")
	}

	return val
}

// GetCreateLotOp retrieves the CreateLotOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateLotOp() (result CreateLotOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateLotOp" {
		result = *u.CreateLotOp
		ok = true
	}

	return
}

// MustUpdateLotOp retrieves the UpdateLotOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustUpdateLotOp() UpdateLotOp {
	val, ok := u.GetUpdateLotOp()

	if !ok {
		panic("arm UpdateLotOp is not set")
	}

	return val
}

// GetUpdateLotOp retrieves the UpdateLotOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetUpdateLotOp() (result UpdateLotOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateLotOp" {
		result = *u.UpdateLotOp
		ok = true
	}

	return
}

// MustRequestParticipationOp retrieves the RequestParticipationOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustRequestParticipationOp() RequestParticipationOp {
	val, ok := u.GetRequestParticipationOp()

	if !ok {
		panic("arm RequestParticipationOp is not set")
	}

	return val
}

// GetRequestParticipationOp retrieves the RequestParticipationOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetRequestParticipationOp() (result RequestParticipationOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RequestParticipationOp" {
		result = *u.RequestParticipationOp
		ok = true
	}

	return
}

// MustSendMessageOp retrieves the SendMessageOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustSendMessageOp() SendMessageOp {
	val, ok := u.GetSendMessageOp()

	if !ok {
		panic("arm SendMessageOp is not set")
	}

	return val
}

// GetSendMessageOp retrieves the SendMessageOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetSendMessageOp() (result SendMessageOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SendMessageOp" {
		result = *u.SendMessageOp
		ok = true
	}

	return
}

// MustCreateBidOp retrieves the CreateBidOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateBidOp() CreateBidOp {
	val, ok := u.GetCreateBidOp()

	if !ok {
		panic("arm CreateBidOp is not set")
	}

	return val
}

// GetCreateBidOp retrieves the CreateBidOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateBidOp() (result CreateBidOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateBidOp" {
		result = *u.CreateBidOp
		ok = true
	}

	return
}

// MustRecoverOp retrieves the RecoverOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustRecoverOp() RecoverOp {
	val, ok := u.GetRecoverOp()

	if !ok {
		panic("arm RecoverOp is not set")
	}

	return val
}

// GetRecoverOp retrieves the RecoverOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetRecoverOp() (result RecoverOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RecoverOp" {
		result = *u.RecoverOp
		ok = true
	}

	return
}

// MustManageMigrationOp retrieves the ManageMigrationOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageMigrationOp() ManageMigrationOp {
	val, ok := u.GetManageMigrationOp()

	if !ok {
		panic("arm ManageMigrationOp is not set")
	}

	return val
}

// GetManageMigrationOp retrieves the ManageMigrationOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageMigrationOp() (result ManageMigrationOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageMigrationOp" {
		result = *u.ManageMigrationOp
		ok = true
	}

	return
}

// MustReviewRequestOp retrieves the ReviewRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustReviewRequestOp() ReviewRequestOp {
	val, ok := u.GetReviewRequestOp()

	if !ok {
		panic("arm ReviewRequestOp is not set")
	}

	return val
}

// GetReviewRequestOp retrieves the ReviewRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetReviewRequestOp() (result ReviewRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ReviewRequestOp" {
		result = *u.ReviewRequestOp
		ok = true
	}

	return
}

// MustCancelRequestOp retrieves the CancelRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCancelRequestOp() CancelRequestOp {
	val, ok := u.GetCancelRequestOp()

	if !ok {
		panic("arm CancelRequestOp is not set")
	}

	return val
}

// GetCancelRequestOp retrieves the CancelRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCancelRequestOp() (result CancelRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CancelRequestOp" {
		result = *u.CancelRequestOp
		ok = true
	}

	return
}

// MustManageParticipantOp retrieves the ManageParticipantOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageParticipantOp() ManageParticipantOp {
	val, ok := u.GetManageParticipantOp()

	if !ok {
		panic("arm ManageParticipantOp is not set")
	}

	return val
}

// GetManageParticipantOp retrieves the ManageParticipantOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageParticipantOp() (result ManageParticipantOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageParticipantOp" {
		result = *u.ManageParticipantOp
		ok = true
	}

	return
}

// MustManageBidLimitOp retrieves the ManageBidLimitOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageBidLimitOp() ManageBidLimitOp {
	val, ok := u.GetManageBidLimitOp()

	if !ok {
		panic("arm ManageBidLimitOp is not set")
	}

	return val
}

// GetManageBidLimitOp retrieves the ManageBidLimitOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageBidLimitOp() (result ManageBidLimitOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageBidLimitOp" {
		result = *u.ManageBidLimitOp
		ok = true
	}

	return
}

// MustCloseAuctionOp retrieves the CloseAuctionOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCloseAuctionOp() CloseAuctionOp {
	val, ok := u.GetCloseAuctionOp()

	if !ok {
		panic("arm CloseAuctionOp is not set")
	}

	return val
}

// GetCloseAuctionOp retrieves the CloseAuctionOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCloseAuctionOp() (result CloseAuctionOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CloseAuctionOp" {
		result = *u.CloseAuctionOp
		ok = true
	}

	return
}

// MustRequestBuyNowOp retrieves the RequestBuyNowOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustRequestBuyNowOp() RequestBuyNowOp {
	val, ok := u.GetRequestBuyNowOp()

	if !ok {
		panic("arm RequestBuyNowOp is not set")
	}

	return val
}

// GetRequestBuyNowOp retrieves the RequestBuyNowOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetRequestBuyNowOp() (result RequestBuyNowOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RequestBuyNowOp" {
		result = *u.RequestBuyNowOp
		ok = true
	}

	return
}

// MustSetMaxBidOp retrieves the SetMaxBidOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustSetMaxBidOp() SetMaxBidOp {
	val, ok := u.GetSetMaxBidOp()

	if !ok {
		panic("arm SetMaxBidOp is not set")
	}

	return val
}

// GetSetMaxBidOp retrieves the SetMaxBidOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetSetMaxBidOp() (result SetMaxBidOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SetMaxBidOp" {
		result = *u.SetMaxBidOp
		ok = true
	}

	return
}

// OperationExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type OperationExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u OperationExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of OperationExt
func (u OperationExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewOperationExt creates a new  OperationExt.
func NewOperationExt(v LedgerVersion, value interface{}) (result OperationExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// Operation is an XDR Struct defines as:
//
//   struct Operation
//        {
//            // sourceAccount is the account used to run the operation
//            // if not set, the runtime defaults to "sourceAccount" specified at
//            // the transaction level
//            AccountID* sourceAccount;
//
//            union switch (OperationType type)
//            {
//            case CREATE_ACCOUNT:
//                CreateAccountOp createAccountOp;
//            case PAYMENT:
//                PaymentOp paymentOp;
//            case PATH_PAYMENT:
//                PathPaymentOp pathPaymentOp;
//            case MANAGE_OFFER:
//                ManageOfferOp manageOfferOp;
//            case CREATE_PASSIVE_OFFER:
//                CreatePassiveOfferOp createPassiveOfferOp;
//            case SET_OPTIONS:
//                SetOptionsOp setOptionsOp;
//            case MANAGE_PERMISSIONS:
//                ManagePermissionsOp managePermissionsOp;
//            case CHANGE_TRUST:
//                ChangeTrustOp changeTrustOp;
//            case ALLOW_TRUST:
//                AllowTrustOp allowTrustOp;
//            case CREATE_LOT:
//                CreateLotOp createLotOp;
//            case UPDATE_LOT:
//                UpdateLotOp updateLotOp;
//            case REQUEST_PARTICIPATION:
//                RequestParticipationOp requestParticipationOp;
//            case SEND_MESSAGE:
//                SendMessageOp sendMessageOp;
//            case CREATE_BID:
//                CreateBidOp createBidOp;
//            case RECOVER:
//                RecoverOp recoverOp;
//            case MANAGE_MIGRATION:
//                ManageMigrationOp manageMigrationOp;
//            case REVIEW_REQUEST:
//                ReviewRequestOp reviewRequestOp;
//            case CANCEL_REQUEST:
//                CancelRequestOp cancelRequestOp;
//            case MANAGE_PARTICIPANT:
//                ManageParticipantOp manageParticipantOp;
//            case MANAGE_BID_LIMIT:
//                ManageBidLimitOp manageBidLimitOp;
//            case CLOSE_AUCTION:
//                CloseAuctionOp closeAuctionOp;
//            case REQUEST_BUY_NOW:
//                RequestBuyNowOp requestBuyNowOp;
//            case SET_MAX_BID:
//                SetMaxBidOp setMaxBidOp;
//            }
//            body;
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type Operation struct {
	SourceAccount *AccountId    `json:"sourceAccount,omitempty"`
	Body          OperationBody `json:"body,omitempty"`
	Ext           OperationExt  `json:"ext,omitempty"`
}

// MemoType is an XDR Enum defines as:
//
//   enum MemoType
//        {
//            MEMO_NONE = 0,
//            MEMO_TEXT = 1,
//            MEMO_ID = 2,
//            MEMO_HASH = 3,
//            MEMO_RETURN = 4,
//            MEMO_ADDITIONAL_SIGNATURE = 5
//        };
//
type MemoType int32

const (
	MemoTypeMemoNone                MemoType = 0
	MemoTypeMemoText                MemoType = 1
	MemoTypeMemoId                  MemoType = 2
	MemoTypeMemoHash                MemoType = 3
	MemoTypeMemoReturn              MemoType = 4
	MemoTypeMemoAdditionalSignature MemoType = 5
)

var MemoTypeAll = []MemoType{
	MemoTypeMemoNone,
	MemoTypeMemoText,
	MemoTypeMemoId,
	MemoTypeMemoHash,
	MemoTypeMemoReturn,
	MemoTypeMemoAdditionalSignature,
}

var memoTypeMap = map[int32]string{
	0: "MemoTypeMemoNone",
	1: "MemoTypeMemoText",
	2: "MemoTypeMemoId",
	3: "MemoTypeMemoHash",
	4: "MemoTypeMemoReturn",
	5: "MemoTypeMemoAdditionalSignature",
}

var memoTypeShortMap = map[int32]string{
	0: "memo_none",
	1: "memo_text",
	2: "memo_id",
	3: "memo_hash",
	4: "memo_return",
	5: "memo_additional_signature",
}

var memoTypeRevMap = map[string]int32{
	"MemoTypeMemoNone":                0,
	"MemoTypeMemoText":                1,
	"MemoTypeMemoId":                  2,
	"MemoTypeMemoHash":                3,
	"MemoTypeMemoReturn":              4,
	"MemoTypeMemoAdditionalSignature": 5,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for MemoType
func (e MemoType) ValidEnum(v int32) bool {
	_, ok := memoTypeMap[v]
	return ok
}
func (e MemoType) isFlag() bool {
	for i := len(MemoTypeAll) - 1; i >= 0; i-- {
		expected := MemoType(2) << uint64(len(MemoTypeAll)-1) >> uint64(len(MemoTypeAll)-i)
		if expected != MemoTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e MemoType) String() string {
	name, _ := memoTypeMap[int32(e)]
	return name
}

func (e MemoType) ShortString() string {
	name, _ := memoTypeShortMap[int32(e)]
	return name
}

func (e MemoType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range MemoTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *MemoType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = MemoType(t.Value)
	return nil
}

// Memo is an XDR Union defines as:
//
//   union Memo switch (MemoType type)
//        {
//        case MEMO_NONE:
//            void;
//        case MEMO_TEXT:
//            string text<28>;
//        case MEMO_ID:
//            uint64 id;
//        case MEMO_HASH:
//            Hash hash; // the hash of what to pull from the content server
//        case MEMO_RETURN:
//            Hash retHash; // the hash of the tx you are rejecting
//        case MEMO_ADDITIONAL_SIGNATURE:
//            string additionalSign<>;
//        };
//
type Memo struct {
	Type           MemoType `json:"type,omitempty"`
	Text           *string  `json:"text,omitempty" xdrmaxsize:"28"`
	Id             *Uint64  `json:"id,omitempty"`
	Hash           *Hash    `json:"hash,omitempty"`
	RetHash        *Hash    `json:"retHash,omitempty"`
	AdditionalSign *string  `json:"additionalSign,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u Memo) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of Memo
func (u Memo) ArmForSwitch(sw int32) (string, bool) {
	switch MemoType(sw) {
	case MemoTypeMemoNone:
		return "", true
	case MemoTypeMemoText:
		return "Text", true
	case MemoTypeMemoId:
		return "Id", true
	case MemoTypeMemoHash:
		return "Hash", true
	case MemoTypeMemoReturn:
		return "RetHash", true
	case MemoTypeMemoAdditionalSignature:
		return "AdditionalSign", true
	}
	return "-", false
}

// NewMemo creates a new  Memo.
func NewMemo(aType MemoType, value interface{}) (result Memo, err error) {
	result.Type = aType
	switch MemoType(aType) {
	case MemoTypeMemoNone:
		// void
	case MemoTypeMemoText:
		tv, ok := value.(string)
		if !ok {
			err = fmt.Errorf("invalid value, must be string")
			return
		}
		result.Text = &tv
	case MemoTypeMemoId:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.Id = &tv
	case MemoTypeMemoHash:
		tv, ok := value.(Hash)
		if !ok {
			err = fmt.Errorf("invalid value, must be Hash")
			return
		}
		result.Hash = &tv
	case MemoTypeMemoReturn:
		tv, ok := value.(Hash)
		if !ok {
			err = fmt.Errorf("invalid value, must be Hash")
			return
		}
		result.RetHash = &tv
	case MemoTypeMemoAdditionalSignature:
		tv, ok := value.(string)
		if !ok {
			err = fmt.Errorf("invalid value, must be string")
			return
		}
		result.AdditionalSign = &tv
	}
	return
}

// MustText retrieves the Text value from the union,
// panicing if the value is not set.
func (u Memo) MustText() string {
	val, ok := u.GetText()

	if !ok {
		panic("arm Text is not set")
	}

	return val
}

// GetText retrieves the Text value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Memo) GetText() (result string, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Text" {
		result = *u.Text
		ok = true
	}

	return
}

// MustId retrieves the Id value from the union,
// panicing if the value is not set.
func (u Memo) MustId() Uint64 {
	val, ok := u.GetId()

	if !ok {
		panic("arm Id is not set")
	}

	return val
}

// GetId retrieves the Id value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Memo) GetId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Id" {
		result = *u.Id
		ok = true
	}

	return
}

// MustHash retrieves the Hash value from the union,
// panicing if the value is not set.
func (u Memo) MustHash() Hash {
	val, ok := u.GetHash()

	if !ok {
		panic("arm Hash is not set")
	}

	return val
}

// GetHash retrieves the Hash value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Memo) GetHash() (result Hash, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Hash" {
		result = *u.Hash
		ok = true
	}

	return
}

// MustRetHash retrieves the RetHash value from the union,
// panicing if the value is not set.
func (u Memo) MustRetHash() Hash {
	val, ok := u.GetRetHash()

	if !ok {
		panic("arm RetHash is not set")
	}

	return val
}

// GetRetHash retrieves the RetHash value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Memo) GetRetHash() (result Hash, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RetHash" {
		result = *u.RetHash
		ok = true
	}

	return
}

// MustAdditionalSign retrieves the AdditionalSign value from the union,
// panicing if the value is not set.
func (u Memo) MustAdditionalSign() string {
	val, ok := u.GetAdditionalSign()

	if !ok {
		panic("arm AdditionalSign is not set")
	}

	return val
}

// GetAdditionalSign retrieves the AdditionalSign value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Memo) GetAdditionalSign() (result string, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AdditionalSign" {
		result = *u.AdditionalSign
		ok = true
	}

	return
}

// TimeBoundsExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type TimeBoundsExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TimeBoundsExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TimeBoundsExt
func (u TimeBoundsExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTimeBoundsExt creates a new  TimeBoundsExt.
func NewTimeBoundsExt(v LedgerVersion, value interface{}) (result TimeBoundsExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TimeBounds is an XDR Struct defines as:
//
//   struct TimeBounds
//        {
//            uint64 minTime;
//            uint64 maxTime;
//            union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        };
//
type TimeBounds struct {
	MinTime Uint64        `json:"minTime,omitempty"`
	MaxTime Uint64        `json:"maxTime,omitempty"`
	Ext     TimeBoundsExt `json:"ext,omitempty"`
}

// TransactionExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type TransactionExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionExt
func (u TransactionExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTransactionExt creates a new  TransactionExt.
func NewTransactionExt(v LedgerVersion, value interface{}) (result TransactionExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// Transaction is an XDR Struct defines as:
//
//   struct Transaction
//        {
//            // account used to run the transaction
//            AccountID sourceAccount;
//
//            // the fee the sourceAccount will pay
//            uint32 fee;
//
//            // transaction's salt
//            Salt salt;
//
//            // validity range (inclusive) for the last ledger close time
//            TimeBounds* timeBounds;
//
//            Memo memo;
//
//            Operation operations<100>;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type Transaction struct {
	SourceAccount AccountId      `json:"sourceAccount,omitempty"`
	Fee           Uint32         `json:"fee,omitempty"`
	Salt          Salt           `json:"salt,omitempty"`
	TimeBounds    *TimeBounds    `json:"timeBounds,omitempty"`
	Memo          Memo           `json:"memo,omitempty"`
	Operations    []Operation    `json:"operations,omitempty" xdrmaxsize:"100"`
	Ext           TransactionExt `json:"ext,omitempty"`
}

// TransactionEnvelopeExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type TransactionEnvelopeExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionEnvelopeExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionEnvelopeExt
func (u TransactionEnvelopeExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTransactionEnvelopeExt creates a new  TransactionEnvelopeExt.
func NewTransactionEnvelopeExt(v LedgerVersion, value interface{}) (result TransactionEnvelopeExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TransactionEnvelope is an XDR Struct defines as:
//
//   struct TransactionEnvelope
//        {
//            Transaction tx;
//            DecoratedSignature signatures<20>;
//            Memo additionalSignatures<10>;
//            // reserved for future use
//                union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//        };
//
type TransactionEnvelope struct {
	Tx                   Transaction            `json:"tx,omitempty"`
	Signatures           []DecoratedSignature   `json:"signatures,omitempty" xdrmaxsize:"20"`
	AdditionalSignatures []Memo                 `json:"additionalSignatures,omitempty" xdrmaxsize:"10"`
	Ext                  TransactionEnvelopeExt `json:"ext,omitempty"`
}

// OperationResultCode is an XDR Enum defines as:
//
//   enum OperationResultCode
//        {
//            opINNER = 0, // inner object result is valid
//
//            opBAD_AUTH = -1,  // too few valid signatures / wrong network
//            opNO_ACCOUNT = -2, // source account was not found
//            opNOT_ALLOWED = -3 // account type is not allowed
//        };
//
type OperationResultCode int32

const (
	OperationResultCodeOpInner      OperationResultCode = 0
	OperationResultCodeOpBadAuth    OperationResultCode = -1
	OperationResultCodeOpNoAccount  OperationResultCode = -2
	OperationResultCodeOpNotAllowed OperationResultCode = -3
)

var OperationResultCodeAll = []OperationResultCode{
	OperationResultCodeOpInner,
	OperationResultCodeOpBadAuth,
	OperationResultCodeOpNoAccount,
	OperationResultCodeOpNotAllowed,
}

var operationResultCodeMap = map[int32]string{
	0:  "OperationResultCodeOpInner",
	-1: "OperationResultCodeOpBadAuth",
	-2: "OperationResultCodeOpNoAccount",
	-3: "OperationResultCodeOpNotAllowed",
}

var operationResultCodeShortMap = map[int32]string{
	0:  "op_inner",
	-1: "op_bad_auth",
	-2: "op_no_account",
	-3: "op_not_allowed",
}

var operationResultCodeRevMap = map[string]int32{
	"OperationResultCodeOpInner":      0,
	"OperationResultCodeOpBadAuth":    -1,
	"OperationResultCodeOpNoAccount":  -2,
	"OperationResultCodeOpNotAllowed": -3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for OperationResultCode
func (e OperationResultCode) ValidEnum(v int32) bool {
	_, ok := operationResultCodeMap[v]
	return ok
}
func (e OperationResultCode) isFlag() bool {
	for i := len(OperationResultCodeAll) - 1; i >= 0; i-- {
		expected := OperationResultCode(2) << uint64(len(OperationResultCodeAll)-1) >> uint64(len(OperationResultCodeAll)-i)
		if expected != OperationResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e OperationResultCode) String() string {
	name, _ := operationResultCodeMap[int32(e)]
	return name
}

func (e OperationResultCode) ShortString() string {
	name, _ := operationResultCodeShortMap[int32(e)]
	return name
}

func (e OperationResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range OperationResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *OperationResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = OperationResultCode(t.Value)
	return nil
}

// OperationResultTr is an XDR NestedUnion defines as:
//
//   union switch (OperationType type)
//            {
//            case CREATE_ACCOUNT:
//                CreateAccountResult createAccountResult;
//            case PAYMENT:
//                PaymentResult paymentResult;
//            case PATH_PAYMENT:
//                PathPaymentResult pathPaymentResult;
//            case MANAGE_OFFER:
//                ManageOfferResult manageOfferResult;
//            case CREATE_PASSIVE_OFFER:
//                ManageOfferResult createPassiveOfferResult;
//            case MANAGE_PERMISSIONS:
//                ManagePermissionsResult managePermissionsResult;
//            case SET_OPTIONS:
//                SetOptionsResult setOptionsResult;
//            case CHANGE_TRUST:
//                ChangeTrustResult changeTrustResult;
//            case ALLOW_TRUST:
//                AllowTrustResult allowTrustResult;
//            case CREATE_LOT:
//                CreateLotResult createLotResult;
//            case UPDATE_LOT:
//                UpdateLotResult updateLotResult;
//            case REQUEST_PARTICIPATION:
//                RequestParticipationResult requestParticipationResult;
//            case SEND_MESSAGE:
//                SendMessageResult sendMessageResult;
//            case CREATE_BID:
//                CreateBidResult createBidResult;
//            case RECOVER:
//                RecoverResult recoverResult;
//            case MANAGE_MIGRATION:
//                ManageMigrationResult manageMigrationResult;
//            case REVIEW_REQUEST:
//                ReviewRequestResult reviewRequestResult;
//            case CANCEL_REQUEST:
//                CancelRequestResult cancelRequestResult;
//            case MANAGE_PARTICIPANT:
//                ManageParticipantResult manageParticipantResult;
//            case MANAGE_BID_LIMIT:
//                ManageBidLimitResult manageBidLimitResult;
//            case CLOSE_AUCTION:
//                CloseAuctionResult closeAuctionResult;
//            case REQUEST_BUY_NOW:
//                RequestBuyNowResult requestBuyNowResult;
//            case SET_MAX_BID:
//                SetMaxBidResult setMaxBidResult;
//            }
//
type OperationResultTr struct {
	Type                       OperationType               `json:"type,omitempty"`
	CreateAccountResult        *CreateAccountResult        `json:"createAccountResult,omitempty"`
	PaymentResult              *PaymentResult              `json:"paymentResult,omitempty"`
	PathPaymentResult          *PathPaymentResult          `json:"pathPaymentResult,omitempty"`
	ManageOfferResult          *ManageOfferResult          `json:"manageOfferResult,omitempty"`
	CreatePassiveOfferResult   *ManageOfferResult          `json:"createPassiveOfferResult,omitempty"`
	ManagePermissionsResult    *ManagePermissionsResult    `json:"managePermissionsResult,omitempty"`
	SetOptionsResult           *SetOptionsResult           `json:"setOptionsResult,omitempty"`
	ChangeTrustResult          *ChangeTrustResult          `json:"changeTrustResult,omitempty"`
	AllowTrustResult           *AllowTrustResult           `json:"allowTrustResult,omitempty"`
	CreateLotResult            *CreateLotResult            `json:"createLotResult,omitempty"`
	UpdateLotResult            *UpdateLotResult            `json:"updateLotResult,omitempty"`
	RequestParticipationResult *RequestParticipationResult `json:"requestParticipationResult,omitempty"`
	SendMessageResult          *SendMessageResult          `json:"sendMessageResult,omitempty"`
	CreateBidResult            *CreateBidResult            `json:"createBidResult,omitempty"`
	RecoverResult              *RecoverResult              `json:"recoverResult,omitempty"`
	ManageMigrationResult      *ManageMigrationResult      `json:"manageMigrationResult,omitempty"`
	ReviewRequestResult        *ReviewRequestResult        `json:"reviewRequestResult,omitempty"`
	CancelRequestResult        *CancelRequestResult        `json:"cancelRequestResult,omitempty"`
	ManageParticipantResult    *ManageParticipantResult    `json:"manageParticipantResult,omitempty"`
	ManageBidLimitResult       *ManageBidLimitResult       `json:"manageBidLimitResult,omitempty"`
	CloseAuctionResult         *CloseAuctionResult         `json:"closeAuctionResult,omitempty"`
	RequestBuyNowResult        *RequestBuyNowResult        `json:"requestBuyNowResult,omitempty"`
	SetMaxBidResult            *SetMaxBidResult            `json:"setMaxBidResult,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u OperationResultTr) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of OperationResultTr
func (u OperationResultTr) ArmForSwitch(sw int32) (string, bool) {
	switch OperationType(sw) {
	case OperationTypeCreateAccount:
		return "CreateAccountResult", true
	case OperationTypePayment:
		return "PaymentResult", true
	case OperationTypePathPayment:
		return "PathPaymentResult", true
	case OperationTypeManageOffer:
		return "ManageOfferResult", true
	case OperationTypeCreatePassiveOffer:
		return "CreatePassiveOfferResult", true
	case OperationTypeManagePermissions:
		return "ManagePermissionsResult", true
	case OperationTypeSetOptions:
		return "SetOptionsResult", true
	case OperationTypeChangeTrust:
		return "ChangeTrustResult", true
	case OperationTypeAllowTrust:
		return "AllowTrustResult", true
	case OperationTypeCreateLot:
		return "CreateLotResult", true
	case OperationTypeUpdateLot:
		return "UpdateLotResult", true
	case OperationTypeRequestParticipation:
		return "RequestParticipationResult", true
	case OperationTypeSendMessage:
		return "SendMessageResult", true
	case OperationTypeCreateBid:
		return "CreateBidResult", true
	case OperationTypeRecover:
		return "RecoverResult", true
	case OperationTypeManageMigration:
		return "ManageMigrationResult", true
	case OperationTypeReviewRequest:
		return "ReviewRequestResult", true
	case OperationTypeCancelRequest:
		return "CancelRequestResult", true
	case OperationTypeManageParticipant:
		return "ManageParticipantResult", true
	case OperationTypeManageBidLimit:
		return "ManageBidLimitResult", true
	case OperationTypeCloseAuction:
		return "CloseAuctionResult", true
	case OperationTypeRequestBuyNow:
		return "RequestBuyNowResult", true
	case OperationTypeSetMaxBid:
		return "SetMaxBidResult", true
	}
	return "-", false
}

// NewOperationResultTr creates a new  OperationResultTr.
func NewOperationResultTr(aType OperationType, value interface{}) (result OperationResultTr, err error) {
	result.Type = aType
	switch OperationType(aType) {
	case OperationTypeCreateAccount:
		tv, ok := value.(CreateAccountResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAccountResult")
			return
		}
		result.CreateAccountResult = &tv
	case OperationTypePayment:
		tv, ok := value.(PaymentResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be PaymentResult")
			return
		}
		result.PaymentResult = &tv
	case OperationTypePathPayment:
		tv, ok := value.(PathPaymentResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be PathPaymentResult")
			return
		}
		result.PathPaymentResult = &tv
	case OperationTypeManageOffer:
		tv, ok := value.(ManageOfferResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageOfferResult")
			return
		}
		result.ManageOfferResult = &tv
	case OperationTypeCreatePassiveOffer:
		tv, ok := value.(ManageOfferResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageOfferResult")
			return
		}
		result.CreatePassiveOfferResult = &tv
	case OperationTypeManagePermissions:
		tv, ok := value.(ManagePermissionsResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManagePermissionsResult")
			return
		}
		result.ManagePermissionsResult = &tv
	case OperationTypeSetOptions:
		tv, ok := value.(SetOptionsResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetOptionsResult")
			return
		}
		result.SetOptionsResult = &tv
	case OperationTypeChangeTrust:
		tv, ok := value.(ChangeTrustResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ChangeTrustResult")
			return
		}
		result.ChangeTrustResult = &tv
	case OperationTypeAllowTrust:
		tv, ok := value.(AllowTrustResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be AllowTrustResult")
			return
		}
		result.AllowTrustResult = &tv
	case OperationTypeCreateLot:
		tv, ok := value.(CreateLotResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateLotResult")
			return
		}
		result.CreateLotResult = &tv
	case OperationTypeUpdateLot:
		tv, ok := value.(UpdateLotResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateLotResult")
			return
		}
		result.UpdateLotResult = &tv
	case OperationTypeRequestParticipation:
		tv, ok := value.(RequestParticipationResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be RequestParticipationResult")
			return
		}
		result.RequestParticipationResult = &tv
	case OperationTypeSendMessage:
		tv, ok := value.(SendMessageResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be SendMessageResult")
			return
		}
		result.SendMessageResult = &tv
	case OperationTypeCreateBid:
		tv, ok := value.(CreateBidResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateBidResult")
			return
		}
		result.CreateBidResult = &tv
	case OperationTypeRecover:
		tv, ok := value.(RecoverResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be RecoverResult")
			return
		}
		result.RecoverResult = &tv
	case OperationTypeManageMigration:
		tv, ok := value.(ManageMigrationResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageMigrationResult")
			return
		}
		result.ManageMigrationResult = &tv
	case OperationTypeReviewRequest:
		tv, ok := value.(ReviewRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReviewRequestResult")
			return
		}
		result.ReviewRequestResult = &tv
	case OperationTypeCancelRequest:
		tv, ok := value.(CancelRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CancelRequestResult")
			return
		}
		result.CancelRequestResult = &tv
	case OperationTypeManageParticipant:
		tv, ok := value.(ManageParticipantResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageParticipantResult")
			return
		}
		result.ManageParticipantResult = &tv
	case OperationTypeManageBidLimit:
		tv, ok := value.(ManageBidLimitResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageBidLimitResult")
			return
		}
		result.ManageBidLimitResult = &tv
	case OperationTypeCloseAuction:
		tv, ok := value.(CloseAuctionResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CloseAuctionResult")
			return
		}
		result.CloseAuctionResult = &tv
	case OperationTypeRequestBuyNow:
		tv, ok := value.(RequestBuyNowResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be RequestBuyNowResult")
			return
		}
		result.RequestBuyNowResult = &tv
	case OperationTypeSetMaxBid:
		tv, ok := value.(SetMaxBidResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetMaxBidResult")
			return
		}
		result.SetMaxBidResult = &tv
	}
	return
}

// MustCreateAccountResult retrieves the CreateAccountResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateAccountResult() CreateAccountResult {
	val, ok := u.GetCreateAccountResult()

	if !ok {
		panic("arm CreateAccountResult is not set")
	}

	return val
}

// GetCreateAccountResult retrieves the CreateAccountResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateAccountResult() (result CreateAccountResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateAccountResult" {
		result = *u.CreateAccountResult
		ok = true
	}

	return
}

// MustPaymentResult retrieves the PaymentResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustPaymentResult() PaymentResult {
	val, ok := u.GetPaymentResult()

	if !ok {
		panic("arm PaymentResult is not set")
	}

	return val
}

// GetPaymentResult retrieves the PaymentResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetPaymentResult() (result PaymentResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PaymentResult" {
		result = *u.PaymentResult
		ok = true
	}

	return
}

// MustPathPaymentResult retrieves the PathPaymentResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustPathPaymentResult() PathPaymentResult {
	val, ok := u.GetPathPaymentResult()

	if !ok {
		panic("arm PathPaymentResult is not set")
	}

	return val
}

// GetPathPaymentResult retrieves the PathPaymentResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetPathPaymentResult() (result PathPaymentResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PathPaymentResult" {
		result = *u.PathPaymentResult
		ok = true
	}

	return
}

// MustManageOfferResult retrieves the ManageOfferResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageOfferResult() ManageOfferResult {
	val, ok := u.GetManageOfferResult()

	if !ok {
		panic("arm ManageOfferResult is not set")
	}

	return val
}

// GetManageOfferResult retrieves the ManageOfferResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageOfferResult() (result ManageOfferResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageOfferResult" {
		result = *u.ManageOfferResult
		ok = true
	}

	return
}

// MustCreatePassiveOfferResult retrieves the CreatePassiveOfferResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreatePassiveOfferResult() ManageOfferResult {
	val, ok := u.GetCreatePassiveOfferResult()

	if !ok {
		panic("arm CreatePassiveOfferResult is not set")
	}

	return val
}

// GetCreatePassiveOfferResult retrieves the CreatePassiveOfferResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreatePassiveOfferResult() (result ManageOfferResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreatePassiveOfferResult" {
		result = *u.CreatePassiveOfferResult
		ok = true
	}

	return
}

// MustManagePermissionsResult retrieves the ManagePermissionsResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManagePermissionsResult() ManagePermissionsResult {
	val, ok := u.GetManagePermissionsResult()

	if !ok {
		panic("arm ManagePermissionsResult is not set")
	}

	return val
}

// GetManagePermissionsResult retrieves the ManagePermissionsResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManagePermissionsResult() (result ManagePermissionsResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManagePermissionsResult" {
		result = *u.ManagePermissionsResult
		ok = true
	}

	return
}

// MustSetOptionsResult retrieves the SetOptionsResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustSetOptionsResult() SetOptionsResult {
	val, ok := u.GetSetOptionsResult()

	if !ok {
		panic("arm SetOptionsResult is not set")
	}

	return val
}

// GetSetOptionsResult retrieves the SetOptionsResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetSetOptionsResult() (result SetOptionsResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SetOptionsResult" {
		result = *u.SetOptionsResult
		ok = true
	}

	return
}

// MustChangeTrustResult retrieves the ChangeTrustResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustChangeTrustResult() ChangeTrustResult {
	val, ok := u.GetChangeTrustResult()

	if !ok {
		panic("arm ChangeTrustResult is not set")
	}

	return val
}

// GetChangeTrustResult retrieves the ChangeTrustResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetChangeTrustResult() (result ChangeTrustResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ChangeTrustResult" {
		result = *u.ChangeTrustResult
		ok = true
	}

	return
}

// MustAllowTrustResult retrieves the AllowTrustResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustAllowTrustResult() AllowTrustResult {
	val, ok := u.GetAllowTrustResult()

	if !ok {
		panic("arm AllowTrustResult is not set")
	}

	return val
}

// GetAllowTrustResult retrieves the AllowTrustResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetAllowTrustResult() (result AllowTrustResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AllowTrustResult" {
		result = *u.AllowTrustResult
		ok = true
	}

	return
}

// MustCreateLotResult retrieves the CreateLotResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateLotResult() CreateLotResult {
	val, ok := u.GetCreateLotResult()

	if !ok {
		panic("arm CreateLotResult is not set")
	}

	return val
}

// GetCreateLotResult retrieves the CreateLotResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateLotResult() (result CreateLotResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateLotResult" {
		result = *u.CreateLotResult
		ok = true
	}

	return
}

// MustUpdateLotResult retrieves the UpdateLotResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustUpdateLotResult() UpdateLotResult {
	val, ok := u.GetUpdateLotResult()

	if !ok {
		panic("arm UpdateLotResult is not set")
	}

	return val
}

// GetUpdateLotResult retrieves the UpdateLotResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetUpdateLotResult() (result UpdateLotResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateLotResult" {
		result = *u.UpdateLotResult
		ok = true
	}

	return
}

// MustRequestParticipationResult retrieves the RequestParticipationResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustRequestParticipationResult() RequestParticipationResult {
	val, ok := u.GetRequestParticipationResult()

	if !ok {
		panic("arm RequestParticipationResult is not set")
	}

	return val
}

// GetRequestParticipationResult retrieves the RequestParticipationResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetRequestParticipationResult() (result RequestParticipationResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RequestParticipationResult" {
		result = *u.RequestParticipationResult
		ok = true
	}

	return
}

// MustSendMessageResult retrieves the SendMessageResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustSendMessageResult() SendMessageResult {
	val, ok := u.GetSendMessageResult()

	if !ok {
		panic("arm SendMessageResult is not set")
	}

	return val
}

// GetSendMessageResult retrieves the SendMessageResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetSendMessageResult() (result SendMessageResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SendMessageResult" {
		result = *u.SendMessageResult
		ok = true
	}

	return
}

// MustCreateBidResult retrieves the CreateBidResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateBidResult() CreateBidResult {
	val, ok := u.GetCreateBidResult()

	if !ok {
		panic("arm CreateBidResult is not set")
	}

	return val
}

// GetCreateBidResult retrieves the CreateBidResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateBidResult() (result CreateBidResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateBidResult" {
		result = *u.CreateBidResult
		ok = true
	}

	return
}

// MustRecoverResult retrieves the RecoverResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustRecoverResult() RecoverResult {
	val, ok := u.GetRecoverResult()

	if !ok {
		panic("arm RecoverResult is not set")
	}

	return val
}

// GetRecoverResult retrieves the RecoverResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetRecoverResult() (result RecoverResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RecoverResult" {
		result = *u.RecoverResult
		ok = true
	}

	return
}

// MustManageMigrationResult retrieves the ManageMigrationResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageMigrationResult() ManageMigrationResult {
	val, ok := u.GetManageMigrationResult()

	if !ok {
		panic("arm ManageMigrationResult is not set")
	}

	return val
}

// GetManageMigrationResult retrieves the ManageMigrationResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageMigrationResult() (result ManageMigrationResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageMigrationResult" {
		result = *u.ManageMigrationResult
		ok = true
	}

	return
}

// MustReviewRequestResult retrieves the ReviewRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustReviewRequestResult() ReviewRequestResult {
	val, ok := u.GetReviewRequestResult()

	if !ok {
		panic("arm ReviewRequestResult is not set")
	}

	return val
}

// GetReviewRequestResult retrieves the ReviewRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetReviewRequestResult() (result ReviewRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ReviewRequestResult" {
		result = *u.ReviewRequestResult
		ok = true
	}

	return
}

// MustCancelRequestResult retrieves the CancelRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCancelRequestResult() CancelRequestResult {
	val, ok := u.GetCancelRequestResult()

	if !ok {
		panic("arm CancelRequestResult is not set")
	}

	return val
}

// GetCancelRequestResult retrieves the CancelRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCancelRequestResult() (result CancelRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CancelRequestResult" {
		result = *u.CancelRequestResult
		ok = true
	}

	return
}

// MustManageParticipantResult retrieves the ManageParticipantResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageParticipantResult() ManageParticipantResult {
	val, ok := u.GetManageParticipantResult()

	if !ok {
		panic("arm ManageParticipantResult is not set")
	}

	return val
}

// GetManageParticipantResult retrieves the ManageParticipantResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageParticipantResult() (result ManageParticipantResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageParticipantResult" {
		result = *u.ManageParticipantResult
		ok = true
	}

	return
}

// MustManageBidLimitResult retrieves the ManageBidLimitResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageBidLimitResult() ManageBidLimitResult {
	val, ok := u.GetManageBidLimitResult()

	if !ok {
		panic("arm ManageBidLimitResult is not set")
	}

	return val
}

// GetManageBidLimitResult retrieves the ManageBidLimitResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageBidLimitResult() (result ManageBidLimitResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageBidLimitResult" {
		result = *u.ManageBidLimitResult
		ok = true
	}

	return
}

// MustCloseAuctionResult retrieves the CloseAuctionResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCloseAuctionResult() CloseAuctionResult {
	val, ok := u.GetCloseAuctionResult()

	if !ok {
		panic("arm CloseAuctionResult is not set")
	}

	return val
}

// GetCloseAuctionResult retrieves the CloseAuctionResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCloseAuctionResult() (result CloseAuctionResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CloseAuctionResult" {
		result = *u.CloseAuctionResult
		ok = true
	}

	return
}

// MustRequestBuyNowResult retrieves the RequestBuyNowResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustRequestBuyNowResult() RequestBuyNowResult {
	val, ok := u.GetRequestBuyNowResult()

	if !ok {
		panic("arm RequestBuyNowResult is not set")
	}

	return val
}

// GetRequestBuyNowResult retrieves the RequestBuyNowResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetRequestBuyNowResult() (result RequestBuyNowResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RequestBuyNowResult" {
		result = *u.RequestBuyNowResult
		ok = true
	}

	return
}

// MustSetMaxBidResult retrieves the SetMaxBidResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustSetMaxBidResult() SetMaxBidResult {
	val, ok := u.GetSetMaxBidResult()

	if !ok {
		panic("arm SetMaxBidResult is not set")
	}

	return val
}

// GetSetMaxBidResult retrieves the SetMaxBidResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetSetMaxBidResult() (result SetMaxBidResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SetMaxBidResult" {
		result = *u.SetMaxBidResult
		ok = true
	}

	return
}

// OperationResult is an XDR Union defines as:
//
//   union OperationResult switch (OperationResultCode code)
//        {
//        case opINNER:
//            union switch (OperationType type)
//            {
//            case CREATE_ACCOUNT:
//                CreateAccountResult createAccountResult;
//            case PAYMENT:
//                PaymentResult paymentResult;
//            case PATH_PAYMENT:
//                PathPaymentResult pathPaymentResult;
//            case MANAGE_OFFER:
//                ManageOfferResult manageOfferResult;
//            case CREATE_PASSIVE_OFFER:
//                ManageOfferResult createPassiveOfferResult;
//            case MANAGE_PERMISSIONS:
//                ManagePermissionsResult managePermissionsResult;
//            case SET_OPTIONS:
//                SetOptionsResult setOptionsResult;
//            case CHANGE_TRUST:
//                ChangeTrustResult changeTrustResult;
//            case ALLOW_TRUST:
//                AllowTrustResult allowTrustResult;
//            case CREATE_LOT:
//                CreateLotResult createLotResult;
//            case UPDATE_LOT:
//                UpdateLotResult updateLotResult;
//            case REQUEST_PARTICIPATION:
//                RequestParticipationResult requestParticipationResult;
//            case SEND_MESSAGE:
//                SendMessageResult sendMessageResult;
//            case CREATE_BID:
//                CreateBidResult createBidResult;
//            case RECOVER:
//                RecoverResult recoverResult;
//            case MANAGE_MIGRATION:
//                ManageMigrationResult manageMigrationResult;
//            case REVIEW_REQUEST:
//                ReviewRequestResult reviewRequestResult;
//            case CANCEL_REQUEST:
//                CancelRequestResult cancelRequestResult;
//            case MANAGE_PARTICIPANT:
//                ManageParticipantResult manageParticipantResult;
//            case MANAGE_BID_LIMIT:
//                ManageBidLimitResult manageBidLimitResult;
//            case CLOSE_AUCTION:
//                CloseAuctionResult closeAuctionResult;
//            case REQUEST_BUY_NOW:
//                RequestBuyNowResult requestBuyNowResult;
//            case SET_MAX_BID:
//                SetMaxBidResult setMaxBidResult;
//            }
//            tr;
//        default:
//            void;
//        };
//
type OperationResult struct {
	Code OperationResultCode `json:"code,omitempty"`
	Tr   *OperationResultTr  `json:"tr,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u OperationResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of OperationResult
func (u OperationResult) ArmForSwitch(sw int32) (string, bool) {
	switch OperationResultCode(sw) {
	case OperationResultCodeOpInner:
		return "Tr", true
	default:
		return "", true
	}
}

// NewOperationResult creates a new  OperationResult.
func NewOperationResult(code OperationResultCode, value interface{}) (result OperationResult, err error) {
	result.Code = code
	switch OperationResultCode(code) {
	case OperationResultCodeOpInner:
		tv, ok := value.(OperationResultTr)
		if !ok {
			err = fmt.Errorf("invalid value, must be OperationResultTr")
			return
		}
		result.Tr = &tv
	default:
		// void
	}
	return
}

// MustTr retrieves the Tr value from the union,
// panicing if the value is not set.
func (u OperationResult) MustTr() OperationResultTr {
	val, ok := u.GetTr()

	if !ok {
		panic("arm Tr is not set")
	}

	return val
}

// GetTr retrieves the Tr value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResult) GetTr() (result OperationResultTr, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Tr" {
		result = *u.Tr
		ok = true
	}

	return
}

// TransactionResultCode is an XDR Enum defines as:
//
//   enum TransactionResultCode
//        {
//            txSUCCESS = 0, // all operations succeeded
//
//            txFAILED = -1, // one of the operations failed (none were applied)
//
//            txTOO_EARLY = -2,         // ledger closeTime before minTime
//            txTOO_LATE = -3,          // ledger closeTime after maxTime
//            txMISSING_OPERATION = -4, // no operation was specified
//
//            txBAD_AUTH = -6,             // too few valid signatures / wrong network
//            txINSUFFICIENT_BALANCE = -7, // fee would bring account below reserve
//            txNO_ACCOUNT = -8,           // source account not found
//            txINSUFFICIENT_FEE = -9,     // fee is too small
//            txBAD_AUTH_EXTRA = -10,      // unused signatures attached to transaction
//            txINTERNAL_ERROR = -11       // an unknown error occured
//        };
//
type TransactionResultCode int32

const (
	TransactionResultCodeTxSuccess             TransactionResultCode = 0
	TransactionResultCodeTxFailed              TransactionResultCode = -1
	TransactionResultCodeTxTooEarly            TransactionResultCode = -2
	TransactionResultCodeTxTooLate             TransactionResultCode = -3
	TransactionResultCodeTxMissingOperation    TransactionResultCode = -4
	TransactionResultCodeTxBadAuth             TransactionResultCode = -6
	TransactionResultCodeTxInsufficientBalance TransactionResultCode = -7
	TransactionResultCodeTxNoAccount           TransactionResultCode = -8
	TransactionResultCodeTxInsufficientFee     TransactionResultCode = -9
	TransactionResultCodeTxBadAuthExtra        TransactionResultCode = -10
	TransactionResultCodeTxInternalError       TransactionResultCode = -11
)

var TransactionResultCodeAll = []TransactionResultCode{
	TransactionResultCodeTxSuccess,
	TransactionResultCodeTxFailed,
	TransactionResultCodeTxTooEarly,
	TransactionResultCodeTxTooLate,
	TransactionResultCodeTxMissingOperation,
	TransactionResultCodeTxBadAuth,
	TransactionResultCodeTxInsufficientBalance,
	TransactionResultCodeTxNoAccount,
	TransactionResultCodeTxInsufficientFee,
	TransactionResultCodeTxBadAuthExtra,
	TransactionResultCodeTxInternalError,
}

var transactionResultCodeMap = map[int32]string{
	0:   "TransactionResultCodeTxSuccess",
	-1:  "TransactionResultCodeTxFailed",
	-2:  "TransactionResultCodeTxTooEarly",
	-3:  "TransactionResultCodeTxTooLate",
	-4:  "TransactionResultCodeTxMissingOperation",
	-6:  "TransactionResultCodeTxBadAuth",
	-7:  "TransactionResultCodeTxInsufficientBalance",
	-8:  "TransactionResultCodeTxNoAccount",
	-9:  "TransactionResultCodeTxInsufficientFee",
	-10: "TransactionResultCodeTxBadAuthExtra",
	-11: "TransactionResultCodeTxInternalError",
}

var transactionResultCodeShortMap = map[int32]string{
	0:   "tx_success",
	-1:  "tx_failed",
	-2:  "tx_too_early",
	-3:  "tx_too_late",
	-4:  "tx_missing_operation",
	-6:  "tx_bad_auth",
	-7:  "tx_insufficient_balance",
	-8:  "tx_no_account",
	-9:  "tx_insufficient_fee",
	-10: "tx_bad_auth_extra",
	-11: "tx_internal_error",
}

var transactionResultCodeRevMap = map[string]int32{
	"TransactionResultCodeTxSuccess":             0,
	"TransactionResultCodeTxFailed":              -1,
	"TransactionResultCodeTxTooEarly":            -2,
	"TransactionResultCodeTxTooLate":             -3,
	"TransactionResultCodeTxMissingOperation":    -4,
	"TransactionResultCodeTxBadAuth":             -6,
	"TransactionResultCodeTxInsufficientBalance": -7,
	"TransactionResultCodeTxNoAccount":           -8,
	"TransactionResultCodeTxInsufficientFee":     -9,
	"TransactionResultCodeTxBadAuthExtra":        -10,
	"TransactionResultCodeTxInternalError":       -11,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for TransactionResultCode
func (e TransactionResultCode) ValidEnum(v int32) bool {
	_, ok := transactionResultCodeMap[v]
	return ok
}
func (e TransactionResultCode) isFlag() bool {
	for i := len(TransactionResultCodeAll) - 1; i >= 0; i-- {
		expected := TransactionResultCode(2) << uint64(len(TransactionResultCodeAll)-1) >> uint64(len(TransactionResultCodeAll)-i)
		if expected != TransactionResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e TransactionResultCode) String() string {
	name, _ := transactionResultCodeMap[int32(e)]
	return name
}

func (e TransactionResultCode) ShortString() string {
	name, _ := transactionResultCodeShortMap[int32(e)]
	return name
}

func (e TransactionResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range TransactionResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *TransactionResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = TransactionResultCode(t.Value)
	return nil
}

// TransactionResultResult is an XDR NestedUnion defines as:
//
//   union switch (TransactionResultCode code)
//            {
//            case txSUCCESS:
//            case txFAILED:
//                OperationResult results<>;
//            default:
//                void;
//            }
//
type TransactionResultResult struct {
	Code    TransactionResultCode `json:"code,omitempty"`
	Results *[]OperationResult    `json:"results,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionResultResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionResultResult
func (u TransactionResultResult) ArmForSwitch(sw int32) (string, bool) {
	switch TransactionResultCode(sw) {
	case TransactionResultCodeTxSuccess:
		return "Results", true
	case TransactionResultCodeTxFailed:
		return "Results", true
	default:
		return "", true
	}
}

// NewTransactionResultResult creates a new  TransactionResultResult.
func NewTransactionResultResult(code TransactionResultCode, value interface{}) (result TransactionResultResult, err error) {
	result.Code = code
	switch TransactionResultCode(code) {
	case TransactionResultCodeTxSuccess:
		tv, ok := value.([]OperationResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be []OperationResult")
			return
		}
		result.Results = &tv
	case TransactionResultCodeTxFailed:
		tv, ok := value.([]OperationResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be []OperationResult")
			return
		}
		result.Results = &tv
	default:
		// void
	}
	return
}

// MustResults retrieves the Results value from the union,
// panicing if the value is not set.
func (u TransactionResultResult) MustResults() []OperationResult {
	val, ok := u.GetResults()

	if !ok {
		panic("arm Results is not set")
	}

	return val
}

// GetResults retrieves the Results value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionResultResult) GetResults() (result []OperationResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Results" {
		result = *u.Results
		ok = true
	}

	return
}

// TransactionResultExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type TransactionResultExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionResultExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionResultExt
func (u TransactionResultExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTransactionResultExt creates a new  TransactionResultExt.
func NewTransactionResultExt(v LedgerVersion, value interface{}) (result TransactionResultExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TransactionResult is an XDR Struct defines as:
//
//   struct TransactionResult
//        {
//            int64 feeCharged; // actual fee charged for the transaction
//
//            union switch (TransactionResultCode code)
//            {
//            case txSUCCESS:
//            case txFAILED:
//                OperationResult results<>;
//            default:
//                void;
//            }
//            result;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type TransactionResult struct {
	FeeCharged Int64                   `json:"feeCharged,omitempty"`
	Result     TransactionResultResult `json:"result,omitempty"`
	Ext        TransactionResultExt    `json:"ext,omitempty"`
}

// Hash is an XDR Typedef defines as:
//
//   typedef opaque Hash[32];
//
type Hash [32]byte

// Uint256 is an XDR Typedef defines as:
//
//   typedef opaque uint256[32];
//
type Uint256 [32]byte

// Uint32 is an XDR Typedef defines as:
//
//   typedef unsigned int uint32;
//
type Uint32 uint32

// Int32 is an XDR Typedef defines as:
//
//   typedef int int32;
//
type Int32 int32

// Uint64 is an XDR Typedef defines as:
//
//   typedef unsigned hyper uint64;
//
type Uint64 uint64

// Int64 is an XDR Typedef defines as:
//
//   typedef hyper int64;
//
type Int64 int64

// CryptoKeyType is an XDR Enum defines as:
//
//   enum CryptoKeyType
//    {
//        KEY_TYPE_ED25519 = 0
//    };
//
type CryptoKeyType int32

const (
	CryptoKeyTypeKeyTypeEd25519 CryptoKeyType = 0
)

var CryptoKeyTypeAll = []CryptoKeyType{
	CryptoKeyTypeKeyTypeEd25519,
}

var cryptoKeyTypeMap = map[int32]string{
	0: "CryptoKeyTypeKeyTypeEd25519",
}

var cryptoKeyTypeShortMap = map[int32]string{
	0: "key_type_ed25519",
}

var cryptoKeyTypeRevMap = map[string]int32{
	"CryptoKeyTypeKeyTypeEd25519": 0,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CryptoKeyType
func (e CryptoKeyType) ValidEnum(v int32) bool {
	_, ok := cryptoKeyTypeMap[v]
	return ok
}
func (e CryptoKeyType) isFlag() bool {
	for i := len(CryptoKeyTypeAll) - 1; i >= 0; i-- {
		expected := CryptoKeyType(2) << uint64(len(CryptoKeyTypeAll)-1) >> uint64(len(CryptoKeyTypeAll)-i)
		if expected != CryptoKeyTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CryptoKeyType) String() string {
	name, _ := cryptoKeyTypeMap[int32(e)]
	return name
}

func (e CryptoKeyType) ShortString() string {
	name, _ := cryptoKeyTypeShortMap[int32(e)]
	return name
}

func (e CryptoKeyType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range CryptoKeyTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CryptoKeyType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CryptoKeyType(t.Value)
	return nil
}

// PublicKey is an XDR Union defines as:
//
//   union PublicKey switch (CryptoKeyType type)
//    {
//    case KEY_TYPE_ED25519:
//        uint256 ed25519;
//    };
//
type PublicKey struct {
	Type    CryptoKeyType `json:"type,omitempty"`
	Ed25519 *Uint256      `json:"ed25519,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PublicKey) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PublicKey
func (u PublicKey) ArmForSwitch(sw int32) (string, bool) {
	switch CryptoKeyType(sw) {
	case CryptoKeyTypeKeyTypeEd25519:
		return "Ed25519", true
	}
	return "-", false
}

// NewPublicKey creates a new  PublicKey.
func NewPublicKey(aType CryptoKeyType, value interface{}) (result PublicKey, err error) {
	result.Type = aType
	switch CryptoKeyType(aType) {
	case CryptoKeyTypeKeyTypeEd25519:
		tv, ok := value.(Uint256)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint256")
			return
		}
		result.Ed25519 = &tv
	}
	return
}

// MustEd25519 retrieves the Ed25519 value from the union,
// panicing if the value is not set.
func (u PublicKey) MustEd25519() Uint256 {
	val, ok := u.GetEd25519()

	if !ok {
		panic("arm Ed25519 is not set")
	}

	return val
}

// GetEd25519 retrieves the Ed25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PublicKey) GetEd25519() (result Uint256, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ed25519" {
		result = *u.Ed25519
		ok = true
	}

	return
}

// Signature is an XDR Typedef defines as:
//
//   typedef opaque Signature<64>;
//
type Signature []byte

// SignatureHint is an XDR Typedef defines as:
//
//   typedef opaque SignatureHint[4];
//
type SignatureHint [4]byte

// NodeId is an XDR Typedef defines as:
//
//   typedef PublicKey NodeID;
//
type NodeId PublicKey

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u NodeId) SwitchFieldName() string {
	return PublicKey(u).SwitchFieldName()
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PublicKey
func (u NodeId) ArmForSwitch(sw int32) (string, bool) {
	return PublicKey(u).ArmForSwitch(sw)
}

// NewNodeId creates a new  NodeId.
func NewNodeId(aType CryptoKeyType, value interface{}) (result NodeId, err error) {
	u, err := NewPublicKey(aType, value)
	result = NodeId(u)
	return
}

// MustEd25519 retrieves the Ed25519 value from the union,
// panicing if the value is not set.
func (u NodeId) MustEd25519() Uint256 {
	return PublicKey(u).MustEd25519()
}

// GetEd25519 retrieves the Ed25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u NodeId) GetEd25519() (result Uint256, ok bool) {
	return PublicKey(u).GetEd25519()
}

// Curve25519Secret is an XDR Struct defines as:
//
//   struct Curve25519Secret
//    {
//            opaque key[32];
//    };
//
type Curve25519Secret struct {
	Key [32]byte `json:"key,omitempty"`
}

// Curve25519Public is an XDR Struct defines as:
//
//   struct Curve25519Public
//    {
//            opaque key[32];
//    };
//
type Curve25519Public struct {
	Key [32]byte `json:"key,omitempty"`
}

// HmacSha256Key is an XDR Struct defines as:
//
//   struct HmacSha256Key
//    {
//            opaque key[32];
//    };
//
type HmacSha256Key struct {
	Key [32]byte `json:"key,omitempty"`
}

// HmacSha256Mac is an XDR Struct defines as:
//
//   struct HmacSha256Mac
//    {
//            opaque mac[32];
//    };
//
type HmacSha256Mac struct {
	Mac [32]byte `json:"mac,omitempty"`
}

// AccountId is an XDR Typedef defines as:
//
//   typedef PublicKey AccountID;
//
type AccountId PublicKey

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountId) SwitchFieldName() string {
	return PublicKey(u).SwitchFieldName()
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PublicKey
func (u AccountId) ArmForSwitch(sw int32) (string, bool) {
	return PublicKey(u).ArmForSwitch(sw)
}

// NewAccountId creates a new  AccountId.
func NewAccountId(aType CryptoKeyType, value interface{}) (result AccountId, err error) {
	u, err := NewPublicKey(aType, value)
	result = AccountId(u)
	return
}

// MustEd25519 retrieves the Ed25519 value from the union,
// panicing if the value is not set.
func (u AccountId) MustEd25519() Uint256 {
	return PublicKey(u).MustEd25519()
}

// GetEd25519 retrieves the Ed25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AccountId) GetEd25519() (result Uint256, ok bool) {
	return PublicKey(u).GetEd25519()
}

// Thresholds is an XDR Typedef defines as:
//
//   typedef opaque Thresholds[4];
//
type Thresholds [4]byte

// String32 is an XDR Typedef defines as:
//
//   typedef string string32<32>;
//
type String32 string

// XDRMaxSize implements the Sized interface for String32
func (e String32) XDRMaxSize() int {
	return 32
}

// String64 is an XDR Typedef defines as:
//
//   typedef string string64<64>;
//
type String64 string

// XDRMaxSize implements the Sized interface for String64
func (e String64) XDRMaxSize() int {
	return 64
}

// String256 is an XDR Typedef defines as:
//
//   typedef string string256<256>;
//
type String256 string

// XDRMaxSize implements the Sized interface for String256
func (e String256) XDRMaxSize() int {
	return 256
}

// Salt is an XDR Typedef defines as:
//
//   typedef uint64 Salt;
//
type Salt Uint64

// SequenceNumber is an XDR Typedef defines as:
//
//   typedef uint64 SequenceNumber;
//
type SequenceNumber Uint64

// DecoratedSignature is an XDR Struct defines as:
//
//   struct DecoratedSignature
//        {
//            SignatureHint hint;  // last 4 bytes of the public key, used as a hint
//            Signature signature; // actual signature
//        };
//
type DecoratedSignature struct {
	Hint      SignatureHint `json:"hint,omitempty"`
	Signature Signature     `json:"signature,omitempty"`
}

// LedgerVersion is an XDR Enum defines as:
//
//   enum LedgerVersion {
//            EMPTY_VERSION = 0,
//            LEDGER_VERSION_SIGN_BY_PLATFORM = 1,
//            LEDGER_VERSION_SIMPLIFY_PARTICIPATION = 2,
//            LEDGER_VERSION_ADMIN_PERMISSIONS = 3,
//            LEDGER_VERSION_MULTIPLE_BUY_NOW_ON_SHELF = 4,
//            LEDGER_VERSION_REJECT_SHELF_PARTS_BEFORE_END_TIME = 5,
//            LEDGER_VERSION_B2B_B2C_PRICES = 6,
//            LEDGER_VERSION_LIVE_BID = 7,
//            LEDGER_VERSION_IGNORE_MAX_STEP = 8,
//            LEDGER_VERSION_PLATFORM_PARTICIPANT = 9,
//            LEDGER_VERSION_LOT_STATE_NEGOTIATIONS = 10,
//            LEDGER_VERSION_ORGANIZER_ID = 11
//        };
//
type LedgerVersion int32

const (
	LedgerVersionEmptyVersion                               LedgerVersion = 0
	LedgerVersionLedgerVersionSignByPlatform                LedgerVersion = 1
	LedgerVersionLedgerVersionSimplifyParticipation         LedgerVersion = 2
	LedgerVersionLedgerVersionAdminPermissions              LedgerVersion = 3
	LedgerVersionLedgerVersionMultipleBuyNowOnShelf         LedgerVersion = 4
	LedgerVersionLedgerVersionRejectShelfPartsBeforeEndTime LedgerVersion = 5
	LedgerVersionLedgerVersionB2BB2CPrices                  LedgerVersion = 6
	LedgerVersionLedgerVersionLiveBid                       LedgerVersion = 7
	LedgerVersionLedgerVersionIgnoreMaxStep                 LedgerVersion = 8
	LedgerVersionLedgerVersionPlatformParticipant           LedgerVersion = 9
	LedgerVersionLedgerVersionLotStateNegotiations          LedgerVersion = 10
	LedgerVersionLedgerVersionOrganizerId                   LedgerVersion = 11
)

var LedgerVersionAll = []LedgerVersion{
	LedgerVersionEmptyVersion,
	LedgerVersionLedgerVersionSignByPlatform,
	LedgerVersionLedgerVersionSimplifyParticipation,
	LedgerVersionLedgerVersionAdminPermissions,
	LedgerVersionLedgerVersionMultipleBuyNowOnShelf,
	LedgerVersionLedgerVersionRejectShelfPartsBeforeEndTime,
	LedgerVersionLedgerVersionB2BB2CPrices,
	LedgerVersionLedgerVersionLiveBid,
	LedgerVersionLedgerVersionIgnoreMaxStep,
	LedgerVersionLedgerVersionPlatformParticipant,
	LedgerVersionLedgerVersionLotStateNegotiations,
	LedgerVersionLedgerVersionOrganizerId,
}

var ledgerVersionMap = map[int32]string{
	0:  "LedgerVersionEmptyVersion",
	1:  "LedgerVersionLedgerVersionSignByPlatform",
	2:  "LedgerVersionLedgerVersionSimplifyParticipation",
	3:  "LedgerVersionLedgerVersionAdminPermissions",
	4:  "LedgerVersionLedgerVersionMultipleBuyNowOnShelf",
	5:  "LedgerVersionLedgerVersionRejectShelfPartsBeforeEndTime",
	6:  "LedgerVersionLedgerVersionB2BB2CPrices",
	7:  "LedgerVersionLedgerVersionLiveBid",
	8:  "LedgerVersionLedgerVersionIgnoreMaxStep",
	9:  "LedgerVersionLedgerVersionPlatformParticipant",
	10: "LedgerVersionLedgerVersionLotStateNegotiations",
	11: "LedgerVersionLedgerVersionOrganizerId",
}

var ledgerVersionShortMap = map[int32]string{
	0:  "empty_version",
	1:  "ledger_version_sign_by_platform",
	2:  "ledger_version_simplify_participation",
	3:  "ledger_version_admin_permissions",
	4:  "ledger_version_multiple_buy_now_on_shelf",
	5:  "ledger_version_reject_shelf_parts_before_end_time",
	6:  "ledger_version_b2_bb2_c_prices",
	7:  "ledger_version_live_bid",
	8:  "ledger_version_ignore_max_step",
	9:  "ledger_version_platform_participant",
	10: "ledger_version_lot_state_negotiations",
	11: "ledger_version_organizer_id",
}

var ledgerVersionRevMap = map[string]int32{
	"LedgerVersionEmptyVersion":                               0,
	"LedgerVersionLedgerVersionSignByPlatform":                1,
	"LedgerVersionLedgerVersionSimplifyParticipation":         2,
	"LedgerVersionLedgerVersionAdminPermissions":              3,
	"LedgerVersionLedgerVersionMultipleBuyNowOnShelf":         4,
	"LedgerVersionLedgerVersionRejectShelfPartsBeforeEndTime": 5,
	"LedgerVersionLedgerVersionB2BB2CPrices":                  6,
	"LedgerVersionLedgerVersionLiveBid":                       7,
	"LedgerVersionLedgerVersionIgnoreMaxStep":                 8,
	"LedgerVersionLedgerVersionPlatformParticipant":           9,
	"LedgerVersionLedgerVersionLotStateNegotiations":          10,
	"LedgerVersionLedgerVersionOrganizerId":                   11,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for LedgerVersion
func (e LedgerVersion) ValidEnum(v int32) bool {
	_, ok := ledgerVersionMap[v]
	return ok
}
func (e LedgerVersion) isFlag() bool {
	for i := len(LedgerVersionAll) - 1; i >= 0; i-- {
		expected := LedgerVersion(2) << uint64(len(LedgerVersionAll)-1) >> uint64(len(LedgerVersionAll)-i)
		if expected != LedgerVersionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e LedgerVersion) String() string {
	name, _ := ledgerVersionMap[int32(e)]
	return name
}

func (e LedgerVersion) ShortString() string {
	name, _ := ledgerVersionShortMap[int32(e)]
	return name
}

func (e LedgerVersion) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range LedgerVersionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *LedgerVersion) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = LedgerVersion(t.Value)
	return nil
}

var fmtTest = fmt.Sprint("this is a dummy usage of fmt")
