package copart

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"time"
)

type reader struct {
	conn *websocket.Conn
	log  *logan.Entry

	notifier *notifier
	ctx      context.Context
	cancel   context.CancelFunc
}

func newReader(ctx context.Context, conn *websocket.Conn, log *logan.Entry) *reader {
	ctx, cancel := context.WithCancel(ctx)
	log = log.WithField("runner", "reader")
	r := &reader{
		conn:     conn,
		log:      log,
		notifier: newNotifier(ctx, log),
		ctx:      ctx,
		cancel:   cancel,
	}

	go r.run()
	return r
}

//Done - signals when reader is not usable
func (r *reader) Done() <-chan struct{} {
	return r.ctx.Done()
}

func (r *reader) close() {
	r.cancel()
	r.notifier.CloseAll()
	_ = r.conn.Close()
}

//Subscribe - returns channel with all messages received from copart. Panics, if same id used twice
func (r *reader) Subscribe(ctx context.Context, id string) <-chan ParserEvent {
	return r.notifier.Subscribe(ctx, id)
}

//Run - runs reader. Must only be called once per instance
func (r *reader) run() {
	err := running.RunSafely(r.ctx, "reader", func(ctx context.Context) error {
		defer r.close()
		for {
			err := r.readMsg(ctx)
			if err != nil {
				return errors.Wrap(err, "failed to read msg")
			}
		}
	})
	if err != nil {
		r.log.WithError(err).Error("reader failed")
	}
}

func (r *reader) readMsg(ctx context.Context) error {
	if ctx.Err() != nil {
		r.log.Info("received signal to stop")
		return ctx.Err()
	}

	err := r.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		return errors.Wrap(err, "failed to set read deadline")
	}
	msgType, rawMsg, err := r.conn.ReadMessage()
	if err != nil {
		return errors.Wrap(err, "failed to read msg")
	}

	r.log.WithField("msg", string(rawMsg)).Debug("read new msg")
	if msgType != websocket.TextMessage {
		return errors.From(errors.New("received unexpected message type"), logan.F{
			"msg_type": msgType,
			"msg":      string(rawMsg),
		})
	}

	var event ParserEvent
	err = json.Unmarshal(rawMsg, &event)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal event", logan.F{
			"raw": string(rawMsg),
		})
	}

	err = r.notifier.NotifyAll(ctx, event)
	if err != nil {
		return errors.Wrap(err, "failed to notify all")
	}

	return nil
}
