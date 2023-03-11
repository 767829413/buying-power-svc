package router

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
)

type listener struct {
	ctx   context.Context
	id    string
	queue chan []byte
	log   *logan.Entry
}

func newListener(ctx context.Context, log *logan.Entry, id string) *listener {
	return &listener{
		id:    id,
		ctx:   ctx,
		queue: make(chan []byte, 64),
		log:   log.WithField("listener_id", id),
	}
}

//Notify - returns true if sent notify
func (l *listener) Notify(ctx context.Context, msgs [][]byte) bool {
	l.log.Debug("sending batch of messages")
	for _, msg := range msgs {
		select {
		case <-ctx.Done():
			l.log.Debug("failed to notify listener - router context canceled.")
			return true
		case <-l.ctx.Done():
			l.log.Debug("failed to notify listener - listener context canceled")
			return false
		case l.queue <- msg:
		default:
			l.log.Warn("failed to notify listener - listener queue is full")
			return false
		}
	}

	l.log.Debug("done sending batch of messages")
	return true
}

//Close - closes Queue
func (l *listener) Close() {
	close(l.queue)
}

//Queue - returns readonly queue for listener
func (l *listener) Queue() <-chan []byte {
	return l.queue
}
