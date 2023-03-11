package ws

import (
	"time"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	writeWait  = 10 * time.Second
	readWait   = 20 * time.Second
	pingPeriod = readWait / 2
)

var errConnectionClosed = errors.New("clients client has been closed")
