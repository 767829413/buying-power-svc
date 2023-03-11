package copart

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	pongWait  = 90 * time.Second
	writeWait = 10 * time.Second
	// ping must be less than pongWait
	pingPeriod = 60 * time.Second
)

var (

	//ErrSaleNotFound - signals that failed to find sale
	ErrSaleNotFound = errors.New("Sale not found")
	//ErrConnectorClosed - live bidder has been closed
	ErrConnectorClosed = errors.New("liveBidder is closed")
)

type liveBidder struct {
	log    *logan.Entry
	reader *reader
	writer *writer

	ctx         context.Context
	cancel      context.CancelFunc
	buyerNumber string
}

//newLiveBidder - returns live bidder
func newLiveBidder(ctx context.Context, log *logan.Entry, url string, copartLogin, copartPass *string) (LiveBidder, error) {
	ctx, cancel := context.WithCancel(ctx)
	dialer := &websocket.Dialer{
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: true,
	}
	conn, _, err := dialer.DialContext(ctx, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial copart", logan.F{
			"url": url,
		})
	}

	conn.SetPingHandler(nil)
	conn.SetPongHandler(func(string) error {
		return conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	log = log.WithField("service", "live-bidder")

	w := newWriter(ctx, conn, log)
	r := newReader(ctx, conn, log)

	c := &liveBidder{
		log:    log,
		ctx:    ctx,
		reader: r,
		writer: w,
		cancel: cancel,
	}

	go func() {
		defer c.Close()
		select {
		case <-w.Done():
			log.Warn("writer stopped - stopping whole live bidder")
		case <-r.Done():
			log.Warn("reader stopped - stopping whole live bidder")
		}
	}()

	err = c.startSession(ctx, copartLogin, copartPass)
	if err != nil {
		return nil, errors.Wrap(err, "failed to start session")
	}

	return c, nil
}

func (c *liveBidder) startSession(ctx context.Context, login, pass *string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute*5)
	defer cancel()
	events := c.Subscribe(ctx, "wait_for_session_start")
	err := c.Send(ctx, ParserCommandStart{
		CopartUserName: login,
		CopartPassword: pass,
	})
	if err != nil {
		return errors.Wrap(err, "failed to send start session request")
	}
	for {
		select {
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "failed to wait for session start")
		case event, ok := <-events:
			if !ok {
				return errors.New("events channel has been closed")
			}

			switch event.Type {
			case ParserEventTypeSessionStarted:
				c.log.Info("session started - waiting for buyer number")
			case ParserEventTypeSessionFailed:
				return errors.Wrap(errors.New(event.SessionFailed().Error), "session failed")
			case ParserEventTypeBuyerNumberIssued:
				c.buyerNumber = event.BuyerNumberIssued().BuyerNumber
				return nil
			default:
				return errors.From(errors.New("unexpected event type"), logan.F{
					"event_type": event.Type,
				})
			}
		}
	}
}

//Send - sends msg to the copart
func (c *liveBidder) Send(ctx context.Context, request ParserCommand) error {
	cmd := request.Command()
	msg, err := json.Marshal(cmd)
	if err != nil {
		return errors.Wrap(err, "failed to marshal copart request")
	}

	return c.writer.Write(ctx, msg)
}

//Subscribe - returns channel with all messages received from copart. Panics, if same id used twice
func (c *liveBidder) Subscribe(ctx context.Context, id string) <-chan ParserEvent {
	return c.reader.Subscribe(ctx, id)
}

//Done - signals when liveBidder is not usable
func (c *liveBidder) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *liveBidder) Close() {
	c.cancel()
}

func (c *liveBidder) BuyerNumber() string {
	return c.buyerNumber
}
