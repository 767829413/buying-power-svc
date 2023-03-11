package copart

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/ws/router"
	"gitlab.com/eAuction/copart-parser-v2/pkg/copart"
)

type eventsListener struct {
	log     *logan.Entry
	events  <-chan copart.ParserEvent
	handler Handler
	ctx     context.Context
	cancel  context.CancelFunc
	router  router.Router
	sales   Sales
}

func newEventsListener(ctx context.Context, log *logan.Entry, events <-chan copart.ParserEvent,
	handler Handler, router  router.Router, sales Sales) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	l := &eventsListener{
		log:     log,
		events:  events,
		handler: handler,
		ctx:     ctx,
		cancel:  cancel,
		router:  router,
		sales:   sales,
	}

	go l.run()
	return ctx
}

func (l *eventsListener) run() {
	defer l.cancel()
	err := running.RunSafely(l.ctx, "events_listener", func(ctx context.Context) error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case msg, ok := <-l.events:
				if !ok {
					l.log.Warn("events has been closed - stopping")
					return nil
				}

				switch msg.Type {
				case copart.ParserEventTypeSaleEnded:
					l.sales.SaleEnded(msg.SaleEnded().SaleID, "received sale ended msg")
					continue
				case copart.ParserEventTypeJoinSaleFailed:
					l.sales.SaleEnded(msg.JoinSaleFailed().SaleID, "join failed")
					continue
				case copart.ParserEventTypeRawWSFrame:
					for _, copartMsg := range msg.RawWSMessage().Messages() {
						if copartMsg.Event == nil {
							continue
						}

						if ctx.Err() != nil {
							return nil
						}

						saleID, results, err := l.handler.Handle(ctx, *copartMsg.Event)
						if err != nil {
							l.log.WithError(err).Error("failed to handle event msg")
							continue
						}

						l.router.NotifyResourceListeners(saleID, results)
					}
				default:
					l.log.WithFields(logan.F{
						"msg":      string(msg.Data),
						"msg_type": msg.Type,
					}).Debug("received not interesting event type")
				}

			}
		}
	})
	if err != nil {
		l.log.WithError(err).Error("events listener died")
		return
	}

	l.log.Debug("events listener stopped")
}
