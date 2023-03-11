package router

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"sync"
)

//Router - routes messages among listeners
type Router interface {
	//Notify - notifies specified listener
	Notify(listenerID string, msgs [][]byte)
	//Close - releases all resources. Not usable after close
	Close()
	//AddNewListener - creates new listener. Panics if listener with same ID already exists
	AddNewListener(ctx context.Context, id string) <-chan []byte
	//NotifyResourceListeners - notifies all listeners of the resource
	NotifyResourceListeners(resourceID string, msgs [][]byte)
	//Unsubscribe - unsubscribe listener from specified resource
	Unsubscribe(listenerID, resourceID string)
	//Subscribe - subscribes listener to specified resource
	Subscribe(listenerID, resourceID string)
	//CloseListener - closes specified listener
	CloseListener(id string)
	// CloseResource - unsubscribe all listeners from specified resource
	CloseResource(resourceID string)
}

type router struct {
	ctx                 context.Context
	cancel              context.CancelFunc
	log                 *logan.Entry
	lock                sync.Mutex
	listeners           map[string]*listener
	resourceToListeners map[string]map[string]struct{}
}

//New - creates new instance of the router
func New(ctx context.Context, log *logan.Entry) Router {
	ctx, cancel := context.WithCancel(ctx)
	return &router{
		ctx:                 ctx,
		cancel:              cancel,
		log:                 log.WithField("router", "router"),
		listeners:           map[string]*listener{},
		resourceToListeners: map[string]map[string]struct{}{},
	}
}

//Close - releases all resources. Not usable after close
func (s *router) Close() {
	s.cancel()
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, l := range s.listeners {
		l.Close()
	}

	s.listeners = nil
	s.resourceToListeners = nil
}

//AddNewListener - creates new listener. Panics if listener with same ID already exists
func (s *router) AddNewListener(ctx context.Context, id string) <-chan []byte {
	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.listeners[id]
	if ok {
		panic(errors.From(errors.New("expected listener id to be unique"), logan.F{
			"id": id,
		}))
	}

	l := newListener(ctx, s.log, id)
	if s.ctx.Err() != nil || ctx.Err() != nil {
		l.Close()
		return l.Queue()
	}
	s.listeners[id] = l
	return l.Queue()
}

//NotifyResourceListeners - notifies all listeners of the resource
func (s *router) NotifyResourceListeners(resourceID string, msgs [][]byte) {
	s.lock.Lock()
	defer s.lock.Unlock()
	var removals []string
	listeners, ok := s.resourceToListeners[resourceID]
	if !ok {
		return
	}
	for id := range listeners {
		if s.ctx.Err() != nil {
			return
		}

		var l *listener
		l, ok = s.listeners[id]
		if !ok {
			s.log.WithField("listener_id", id).Warn("resource listener contains listener that is not available in the listeners")
			continue
		}

		notified := l.Notify(s.ctx, msgs)
		if !notified {
			s.log.WithField("listener_id", id).Debug("failed to notify listener - scheduling closing")
			removals = append(removals, id)
		}
	}

	for _, removal := range removals {
		s.closeListener(removal)
	}
}

//CloseListener - closes specified listener
func (s *router) CloseListener(id string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.closeListener(id)
}

func (s *router) closeListener(id string) {
	log := s.log.WithField("listener_id", id)
	log.Debug("closing listener")
	l, ok := s.listeners[id]
	if !ok {
		log.Debug("failed to find listener in listeners - must be already closed")
		return
	}

	delete(s.listeners, id)
	for _, resourceToListeners := range s.resourceToListeners {
		delete(resourceToListeners, id)
	}

	l.Close()
	log.Debug("closed listener")
}

//Unsubscribe - unsubscribe listener from specified resource
func (s *router) Unsubscribe(listenerID, resourceID string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.ctx.Err() != nil {
		return
	}
	if _, ok := s.listeners[listenerID]; !ok {
		return
	}
	resourceListeners, ok := s.resourceToListeners[resourceID]
	if !ok {
		return
	}

	delete(resourceListeners, listenerID)
	if len(resourceListeners) == 0 {
		delete(s.resourceToListeners, resourceID)
	}
}

//CloseResource - unsubscribe all listeners from specified resource
func (s *router) CloseResource(resourceID string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.ctx.Err() != nil {
		return
	}

	delete(s.resourceToListeners, resourceID)
}

//Subscribe - subscribes listener to specified resource
func (s *router) Subscribe(listenerID, resourceID string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.ctx.Err() != nil {
		return
	}
	if _, ok := s.listeners[listenerID]; !ok {
		return
	}

	resourceListeners, ok := s.resourceToListeners[resourceID]
	if !ok {
		resourceListeners = map[string]struct{}{}
		s.resourceToListeners[resourceID] = resourceListeners
	}

	resourceListeners[listenerID] = struct{}{}
}

//Notify - notifies specified listener
func (s *router) Notify(listenerID string, msgs [][]byte) {
	s.lock.Lock()
	defer s.lock.Unlock()
	l, ok := s.listeners[listenerID]
	if !ok {
		return
	}

	notified := l.Notify(s.ctx, msgs)
	if !notified {
		s.log.WithField("listener_id", listenerID).Debug("failed to notify listener - closing")
		s.closeListener(listenerID)
	}
}
