package copart

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/responses"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/ws/router"
	"gitlab.com/eAuction/go/support/log"
	"sync"
)

type Sales interface {
	AcquireSubscribeLock(listenerID string, sale Sale) bool
	SaleEnded(copartSaleID, reason string)
	GetShelfSaleID(copartSaleID string) string
	IsSubscribed(copartSaleID string) bool
}

// sales - tracks live bidder subscriptions to sale
type sales struct {
	log *logan.Entry
	router router.Router
	salesByCopartID map[string]Sale
	lock sync.Mutex
}

func newSales(log *logan.Entry, router router.Router) *sales {
	return &sales{
		router:          router,
		salesByCopartID: map[string]Sale{},
		log: log,
	}
}

//AcquireSubscribeLock - if returns true - caller must send subscribe request
func (s *sales) AcquireSubscribeLock(listenerID string, sale Sale) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.router.Subscribe(listenerID, sale.CopartSaleID)
	_, ok := s.salesByCopartID[sale.CopartSaleID]
	// subscription has been sent
	if ok {
		return false
	}

	s.salesByCopartID[sale.CopartSaleID] = sale
	return true
}

//SaleEnded - notifies all subscribers that sale ended
func (s *sales) SaleEnded(copartSaleID, reason string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	sale, ok := s.salesByCopartID[copartSaleID]
	if !ok {
		s.log.WithField("copart_sale_id", copartSaleID).Error("trying to handle sale end for sale that is not available")
		return
	}

	delete(s.salesByCopartID, sale.CopartSaleID)
	endSaleAsJson, err := json.Marshal(responses.PopulateSaleEnd(sale.ShelfSaleID, reason))
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal sale end"))
	}
	log.Debug("removing sale handler from map")
	log.Debug("noting listeners that sale has been closed")
	s.router.NotifyResourceListeners(sale.CopartSaleID, [][]byte{endSaleAsJson})
	log.Debug("closing resource")
	s.router.CloseResource(sale.CopartSaleID)
}

//GetShelfSaleID - returns shelf id of the sale
func (s *sales) GetShelfSaleID(copartSaleID string) string {
	s.lock.Lock()
	defer s.lock.Unlock()
	sale := s.salesByCopartID[copartSaleID]
	return sale.ShelfSaleID
}

//IsSubscribed - returns true if subscribed for sale
func (s *sales) IsSubscribed(copartSaleID string) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.salesByCopartID[copartSaleID]
	return ok
}
