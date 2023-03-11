package lots

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
)

// ErrLotNotFound is a type of error that happens when we can't find lot
// linked to event we're ingesting. It is needed to live somehow with bug about
// unordered kafka so we can skip ingestion of everything linked to lot until we
// got the message about lot creation.
var ErrLotNotFound = errors.New("lot not found")

// ErrParticipantNotFound is a type of error that happens when we can't find
// participant linked to event we're ingesting. It is needed to live somehow with bug about
//// unordered kafka so we can skip ingestion of everything linked to participant until we
//// got the message about participant creation.
var ErrParticipantNotFound = errors.New("participant not found")
