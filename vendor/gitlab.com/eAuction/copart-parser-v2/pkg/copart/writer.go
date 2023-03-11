package copart

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/running"
	"time"
)

type writer struct {
	queue   chan []byte
	conn    *websocket.Conn
	log     *logan.Entry

	ctx    context.Context
	cancel context.CancelFunc
}

func newWriter(ctx context.Context, conn *websocket.Conn, log *logan.Entry) *writer {
	ctx, cancel := context.WithCancel(ctx)
	w := &writer{
		queue:   make(chan []byte, 16),
		conn:    conn,
		log:     log.WithField("runner", "writer"),
		ctx:     ctx,
		cancel:  cancel,
	}

	go w.run()

	return w
}

// Write - writes msg to the queue. Returns ErrConnectorClosed is writer is not usable
func (w *writer) Write(ctx context.Context, msg []byte) error {
	log := w.log.WithField("request_id", RequestID(ctx))
	log.Debug("received new request to add into queue")
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-w.ctx.Done():
		return ErrConnectorClosed
	case w.queue <- msg:
		log.Debug("added new request to queue")
		return nil
	}
}

//Done - signals when writer is not usable
func (w *writer) Done() <-chan struct{} {
	return w.ctx.Done()
}

func (w *writer) close() {
	w.cancel()
	_ = w.conn.Close()
}

//Run - runs writer. Must only be called once per instance
func (w *writer) run() {
	err := running.RunSafely(w.ctx, "writer", func(ctx context.Context) error {
		ticker := time.NewTicker(pingPeriod)
		defer func() {
			ticker.Stop()
			w.close()
		}()

		for {
			select {
			case <-ctx.Done():
				w.log.Info("received signal to stop")
				return w.writeClose()
			case msg, ok := <-w.queue:
				if !ok {
					w.log.Info("write queue has been closed - stopping")
					return w.writeClose()
				}

				err := w.writeMsg(msg)
				if err != nil {
					return errors.Wrap(err, "failed to write msg")
				}

			case <-ticker.C:
				err := w.ping()
				if err != nil {
					return errors.Wrap(err, "failed to write ping msg")
				}
			}
		}
	})
	if err != nil {
		w.log.WithError(err).Error("writer stopped")
	}
}

func (w *writer) writeMsg(msg []byte) error {
	w.log.WithField("msg", string(msg)).Debug("writing msg to copart")
	err := w.conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err != nil {
		return errors.Wrap(err, "failed to set write deadline")
	}

	err = w.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return errors.Wrap(err, "failed to write")
	}

	return nil
}

func (w *writer) ping() error {
	err := w.conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err != nil {
		return errors.Wrap(err, "failed to set write deadline for ping msg")
	}

	err = w.conn.WriteMessage(websocket.PingMessage, nil)
	return errors.Wrap(err, "failed to write ping msg")
}

func (w *writer) writeClose() error {
	w.log.Info("received signal to stop")
	err := w.conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err != nil {
		return errors.Wrap(err, "failed to set write deadline")
	}

	err = w.conn.WriteMessage(websocket.CloseMessage, []byte{})
	if err != nil {
		return errors.Wrap(err, "failed to write close msg")
	}

	err = w.conn.Close()
	if err != nil {
		return errors.Wrap(err, "failed to close connection")
	}

	return nil
}
