package ws

import (
	"context"
	"encoding/json"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/lorem"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/responses"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/ws/copart"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/eAuction/bouncer/jwt"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
)

type client struct {
	conn         *websocket.Conn
	log          *logan.Entry
	cfg          config.Config
	isSimulation bool
	claims       jwt.Claims
	connectionID string


	writeQueue    chan []byte
	requestsQueue chan resources.EnvelopeRequestResponse
	ctx           context.Context
	cancel context.CancelFunc
	liveBidder    comfig.Once

}

//RunClient - starts client related routines. Blocks until client is valid.
func RunClient(ctx context.Context, conn *websocket.Conn, claims jwt.Claims, cfg config.Config, isSimulation bool) {
	defer func() {
		_ = conn.Close()
	}()
	ctx, cancel := context.WithCancel(ctx)
	connectionID := lorem.ULID()
	log := cfg.Log().WithFields(logan.F{
		"connection_id": connectionID,
		"account_id":    claims.AccountID,
	})

	c := &client{
		conn:         conn,
		log:          log,
		cfg:          cfg,
		isSimulation: isSimulation,
		claims:       claims,
		connectionID:connectionID,

		requestsQueue: make(chan resources.EnvelopeRequestResponse, 16),
		writeQueue:    make(chan []byte, 16),
		ctx:           ctx,
		cancel:        cancel,
	}

	go c.runRead()
	go c.runWriter()
	go c.runRequestHandler()
	participantMsgs := cfg.ParticipantsRouter().Subscribe(ctx, connectionID, claims.AccountID)
	go c.runListener("participants_listener", participantMsgs)
	<-ctx.Done()
	log.Debug("client stopped")
}

//close - stops all runners and releases resources
func (c *client) close() {
	c.cancel()
}

func (c *client) getLiveBidder() *copart.LiveBidder {
	return c.liveBidder.Do(func() interface{} {
		log := c.log.WithField("runner", "live-bidder-provider")
		log.Debug("getting live bidder")
		ctx, cancel := context.WithTimeout(c.ctx, time.Minute*2)
		defer cancel()
		log.Debug("getting platform")
		platform := c.cfg.Platformer().MustPlatform(c.claims.PlatformID)
		var err error
		var liveBidder *copart.LiveBidder
		if c.isSimulation {
			log.Debug("getting simulation")
			liveBidder, err = c.cfg.Simulation(ctx, platform.Data.Attributes.CountryCode)
		} else {
			log.Debug("getting real live bidder")
			liveBidder, err = c.cfg.Real(ctx, platform.Data.Attributes.CountryCode)
		}

		if err != nil {
			panic(errors.Wrap(err, "failed to get live bidder", logan.F{
				"is_simulation": c.isSimulation,
			}))
		}

		log.Debug("got live bidder")
		if !liveBidder.IsPlatformSupported(c.claims.PlatformID) {
			resp := responses.PopulateError("", http.StatusMethodNotAllowed, errors.New("live bidding is not supported for the platform"))
			err = c.scheduleWrite(*resp)
			if err != nil {
				panic(errors.Wrap(err, "failed to send msg that live bidding is not supported"))
			}
		}

		log.Debug("subscribing")
		copartMsgs := liveBidder.Subscribe(c.ctx, c.connectionID)
		log.Debug("starting listener")
		go c.runListener("copart_listener", copartMsgs)
		log.Debug("got live bidder")
		return liveBidder
	}).(*copart.LiveBidder)
}

func (c *client) runRead() {
	c.conn.SetPongHandler(func(string) error { _ = c.conn.SetReadDeadline(time.Now().Add(readWait)); return nil })
	defer c.close()
	defer close(c.requestsQueue)
	err := running.RunSafely(c.ctx, "client-reader", func(ctx context.Context) error {
		for {
			err := c.conn.SetReadDeadline(time.Now().Add(readWait))
			if err != nil {
				return errors.Wrap(err, "failed to set read dead line")
			}

			c.log.Debug("reading new request")
			_, rawRequest, err := c.conn.ReadMessage()
			if err != nil {
				return errors.Wrap(err, "failed to get next reader")
			}

			c.log.WithField("request", string(rawRequest)).Debug("received new request")
			var request resources.EnvelopeRequestResponse
			err = json.Unmarshal(rawRequest, &request)
			if err != nil {
				c.log.WithError(err).Info("failed to unmarshal request")
				return nil
			}

			immediateResponse, err := c.addRequestIntoQueue(request)
			if err != nil {
				return errors.Wrap(err, "failed to handle request")
			}

			if immediateResponse != nil {
				err = c.scheduleWrite(*immediateResponse)
				if err != nil {
					return errors.Wrap(err, "failed to schedule write")
				}
				c.log.Debug("added response to queue")
				continue
			}

			c.log.Debug("no sync response for the request")
		}
	})
	c.log.WithError(err).Debug("client-reader died")
}

func (c *client) addRequestIntoQueue(request resources.EnvelopeRequestResponse) (*resources.EnvelopeResponse, error) {
	// shortcut for the ping messages to ensure that client does not closes connection
	if request.Data.Relationships.Ping != nil {
		return responses.PopulateOk(request.Data.ID), nil
	}

	c.log.Debug("adding request to queue")
	select {
	case <-c.ctx.Done():
		return nil, errConnectionClosed
	case c.requestsQueue <- request:
		if request.Data.Relationships.Subscribe != nil {
			obj := request.Data.Relationships.Subscribe.Data
			if obj == nil {
				return responses.PopulateError(request.Data.ID, http.StatusBadRequest, errors.New("expected non nil Subscribe data")), nil
			}
			return responses.PopulateOk(request.Data.ID), nil
		}
		return nil, nil
	default:
		c.log.Debug("requests queue is full")
		return responses.PopulateError(request.Data.ID, http.StatusTooManyRequests, errors.New("requests queue is full - try again later")), nil
	}
}

func (c *client) runRequestHandler() {
	defer c.close()
	err := running.RunSafely(c.ctx, "requests-handler", func(ctx context.Context) error {
		handler := newRequestsHandler(c.log, c.connectionID, c.claims.AccountID, c.isSimulation, c.cfg.Storage().Clone(), func() CopartWriter {
			return c.getLiveBidder()
		})
		for {
			c.log.Debug("ready to handle next request")
			select {
			case <-c.ctx.Done():
				return errConnectionClosed
			case request, ok := <-c.requestsQueue:
				if !ok {
					return errors.New("request queue has been closed")
				}

				if request.Data.ID == "" {
					request.Data.ID = "locally_assigned:" + lorem.ULID()
				}

				requestAsJSON, _ := json.Marshal(request)
				c.log.WithField("request", string(requestAsJSON)).Debug("handling scheduled request")

				response := handler.HandleRequest(c.ctx, request)
				if response != nil {
					c.log.WithField("request_id", request.Data.ID).Debug("got response for request - scheduling it's write")
					err := c.scheduleWrite(*response)
					if err != nil {
						return errors.Wrap(err, "failed to schedule write")
					}
				}
			}
		}
	})
	c.log.WithError(err).Debug("request handler died")
}

func (c *client) runWriter() {
	defer c.close()
	err := running.RunSafely(c.ctx, "client-writer", func(ctx context.Context) error {
		ticker := time.NewTicker(pingPeriod)
		defer func() {
			ticker.Stop()
		}()
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case msg, ok := <-c.writeQueue:
				if !ok {
					return errors.New("scheduleWrite queue has been closed")
				}

				c.log.WithField("msg", string(msg)).Debug("writing response")
				err := c.conn.SetWriteDeadline(time.Now().Add(writeWait))
				if err != nil {
					return errors.Wrap(err, "failed to set scheduleWrite deadline")
				}

				err = c.conn.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					return errors.Wrap(err, "failed to scheduleWrite message")
				}
				c.log.Debug("wrote response")
			case <-ticker.C:
				err := c.conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(writeWait))
				if err != nil {
					return errors.Wrap(err, "failed to scheduleWrite ping")
				}
			}
		}
	})
	c.log.WithError(err).Debug("client-writer died")
}

func (c *client) runListener(name string, msgs <-chan []byte) {
	// TODO: send corresponding message instead of closing connection
	defer c.close()

	err := running.RunSafely(c.ctx, name, func(ctx context.Context) error {
		for {
			select {
			case <-c.ctx.Done():
				return ctx.Err()
			case msg, ok := <-msgs:
				if !ok {
					return errors.New("msgs has been closed")
				}

				select {
				case <-c.ctx.Done():
					return errConnectionClosed
				case c.writeQueue <- msg:
				}

			}
		}
	})
	c.log.WithError(err).Debug(name)
}

func (c *client) scheduleWrite(msg resources.EnvelopeResponse) error {
	asJSON, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "failed to marshal msg")
	}
	log := c.log.WithField("response_id", msg.Data.ID)
	log.Debug("scheduling write")
	select {
	case <-c.ctx.Done():
		return errConnectionClosed
	case c.writeQueue <- asJSON:
		log.Debug("scheduled write")
		return nil
	}
}
