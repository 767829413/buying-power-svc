package copart

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"sync"
)

type client struct {
	queue chan ParserEvent
	ctx   context.Context
}

type notifier struct {
	lock    sync.Mutex
	clients map[string]client
	ctx     context.Context
	log     *logan.Entry
}

func newNotifier(ctx context.Context, log *logan.Entry) *notifier {
	return &notifier{
		clients: map[string]client{},
		ctx:     ctx,
		log:     log,
	}
}

//Subscribe - returns channel with all messages received from copart. Panics, if same id used twice.
// Use cancel of the context to unsubscribe
func (n *notifier) Subscribe(ctx context.Context, id string) <-chan ParserEvent {
	result := make(chan ParserEvent, 64)
	n.lock.Lock()
	defer n.lock.Unlock()
	if n.ctx.Err() != nil {
		close(result)
		return result
	}
	_, ok := n.clients[id]
	if ok {
		panic(errors.From(errors.New("trying to subscribe with ID already user"), logan.F{
			"id": id,
		}))
	}
	n.clients[id] = client{
		queue: result,
		ctx:   ctx,
	}
	return result

}

//NotifyAll - notifies all the clients
func (n *notifier) NotifyAll(ctx context.Context, msg ParserEvent) error {
	n.lock.Lock()
	defer n.lock.Unlock()
	// not feeling safe to modify map during iterating through it
	var toRemove []string
	n.log.Debug("starting to notify all connections")
	for id, c := range n.clients {
		log := n.log.WithField("connection_id", id)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c.ctx.Done():
			log.Debug("connection context closed - adding for removal")
			close(c.queue)
			toRemove = append(toRemove, id)
		case c.queue <- msg:
		default:
			log.Warn("connection's queue is full - adding for removal")
			close(c.queue)
			toRemove = append(toRemove, id)
		}
	}

	for _, id := range toRemove {
		delete(n.clients, id)
	}

	n.log.Debug("done notifying all connections")
	return nil
}

//CloseAll - closes all clients
func (n *notifier) CloseAll() {
	n.lock.Lock()
	defer n.lock.Unlock()
	for _, c := range n.clients {
		close(c.queue)
	}

	n.clients = map[string]client{}
}
