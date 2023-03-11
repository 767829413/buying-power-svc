package data

import (
	"context"
	"sync"
	"time"

	"github.com/golang/groupcache/lru"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

//LotsCache - helper interface to access lots from storage with cache
type LotsCache interface {
	Get(key string, isSimulation bool) (*Lot, error)
}

//NewLotsCacheByExternalIDPrefix - creates new lots cache. Note: need to close context or call close to release resources
func NewLotsCacheByExternalIDPrefix(ctx context.Context, storage Storage, log *logan.Entry) LotsCache {
	ctx, cancel := context.WithCancel(ctx)
	return &lotsCacheByExternalIDPrefix{
		ctx:    ctx,
		cancel: cancel,
		log:    log.WithField("service", "lot-cache-by-external-id-prefix"),
		q:      storage.Lots(),
		lock:   sync.Mutex{},
		cache:  lru.New(256),
	}
}

type lotsCacheByExternalIDPrefix struct {
	ctx    context.Context
	cancel context.CancelFunc

	log   *logan.Entry
	q     Lots
	lock  sync.Mutex
	cache *lru.Cache
}

//Get - returns lot by external ID prefix
func (l *lotsCacheByExternalIDPrefix) Get(prefix string, isSimulation bool) (*Lot, error) {
	if l.ctx.Err() != nil {
		return nil, errors.Wrap(l.ctx.Err(), "cache is not usable as context has been cancled")
	}
	l.lock.Lock()
	defer l.lock.Unlock()
	rawLot, ok := l.cache.Get(prefix)
	if ok {
		if rawLot == nil {
			return nil, nil
		}
		return rawLot.(*Lot), nil
	}

	lot, err := l.q.ByExternalIDPrefix(prefix, isSimulation)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load lot")
	}

	l.cache.Add(prefix, lot)
	return lot, nil
}

func (l *lotsCacheByExternalIDPrefix) runCleaner() {
	// ensures that we do not end up with old version of the lot
	running.WithBackOff(l.ctx, l.log, "lots-cache-cleaner", func(ctx context.Context) error {
		l.cache.Clear()
		return nil
	}, time.Hour*12, time.Second, time.Hour)
}
