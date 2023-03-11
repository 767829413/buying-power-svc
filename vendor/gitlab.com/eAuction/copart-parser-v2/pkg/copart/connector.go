package copart

import (
	"context"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"
	"time"

	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/copart-parser-v2/internal/services/copart/cookies"
)

//go:generate mockery -case underscore -name Connector -inpkg
type Connector interface {
	GetBuyerNumber() BuyerNumber
	GetUpcomingSales() ([]Sale, error)
	GetLots() ([]Lot, error)
	IsAnonymous() bool
	Close()
	Done() <-chan struct{}
	GetSales() ([]SaleData, error)
	GetSale(saleID string) (*SaleData, error)
	GetLotDetails(lotID string) (*LotDetails, error)
	GetImagesList(lotID string) (*ImageList, error)
	GetWon(page, pageSize int64) (*WonResponse, error)
	GetLost(page, pageSize int64) (*LostResponse, error)
	GetOpenItems(page, pageSize int64) (*OpenItemsResponse, error)
	GetInvoices(page, pageSize int64) (*InvoicesResponse, error)
	DowloadInvoice(invoiceNumber, paymentType, currencyCode, invoiceType, lotNumber, secretID, bidderNumber string) (io.ReadCloser, error)
	GetDynamicLotDetails(lotID string) (*DynamicLotDetails, error)
	PostMaxBid(request MaxBidRequest) error
	RemoveFromWatchlist(lots []string) error
	GetWatchlist() ([]string, error)
}

//ErrNotFound - resource not found
var ErrNotFound = errors.New("not found")

//connector - provides access to copart REST API
type connector struct {
	userAgent        string
	csrfToken        string
	client           *http.Client
	ctx              context.Context
	cancel           context.CancelFunc
	log              *logan.Entry
	isAnonymous      bool
	disableKeepAlive bool
	proxy            *cookies.Proxy

	drawLock      sync.Mutex
	wonDraw       int64
	lostDraw      int64
	openItemsDraw int64
	invoicesDraw  int64

	buyerNumber comfig.Once

	throttle *time.Timer
}

const requestsDelay = time.Second * 7

//NewConnector - creates new connector
func NewConnector(ctx context.Context, log *logan.Entry, cookieJar *cookies.CookieJar, isAnonymous, disableKeepAlive bool) (Connector, error) {
	log.Debug("creating cookies jar")
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to create cookie jar"))
	}

	if !isAnonymous && cookieJar.CsrfToken == "" {
		return nil, errors.New("expected csfr token to be present")
	}

	ctx, cancel := context.WithCancel(ctx)
	u := cookieJar.Proxy.URL
	u.User = url.UserPassword(*cookieJar.Proxy.Username, *cookieJar.Proxy.Password)
	tr := &http.Transport{
		Proxy: http.ProxyURL(u),
	}

	result := connector{
		csrfToken: cookieJar.CsrfToken,
		userAgent: cookieJar.UserAgent,
		client: &http.Client{
			Timeout:   time.Second * 40,
			Jar:       jar,
			Transport: tr,
		},
		ctx:              ctx,
		cancel:           cancel,
		log:              log,
		throttle:         time.NewTimer(requestsDelay),
		isAnonymous:      isAnonymous,
		disableKeepAlive: disableKeepAlive,
	}

	rawURLs := []string{"https://g2auction.copart.com", "https://www.copart.com"}
	for _, rawURL := range rawURLs {
		var resource *url.URL
		resource, err = url.Parse(rawURL)
		if err != nil {
			panic(errors.Wrap(err, "failed to parse copart url", logan.F{
				"raw_url": rawURL,
			}))
		}

		jar.SetCookies(resource, cookieJar.Cookies)
	}

	go result.runKeepAlive()

	return &result, nil
}

//Close - closes connector
func (c *connector) Close() {
	if c == nil {
		return
	}
	c.cancel()
	c.throttle.Stop()
}

//Done - signals that connector is no longer usable
func (c *connector) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *connector) runKeepAlive() {
	if c.isAnonymous {
		c.log.Info("not running keep alive as in anonymous mode")
		return
	}

	if c.disableKeepAlive {
		c.log.Info("keep alive disabled by config option")
		return
	}
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()
	defer c.Close()
	for {
		select {
		case <-ticker.C:
			_, err := c.GetSales()
			if err != nil {
				c.log.WithError(err).Error("failed to get buyer number from keep alive routine - stopping")
				return
			}
			c.log.Info("refreshed cookies using get sales request")
		case <-c.ctx.Done():
			return
		}
	}
}

//IsAnonymous - returns true if not logged into copart
func (c *connector) IsAnonymous() bool {
	return c.isAnonymous
}

func (c *connector) getNextDraw(draw *int64) int64 {
	c.drawLock.Lock()
	defer c.drawLock.Unlock()
	*draw = (*draw) + 1
	return *draw
}
