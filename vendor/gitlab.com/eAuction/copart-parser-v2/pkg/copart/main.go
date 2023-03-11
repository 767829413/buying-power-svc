package copart

import (
	"context"
	"net/url"
	"sync"
	"time"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/eAuction/copart-parser-v2/internal/services/copart/cookies"
)

var _ Connector = &MockConnector{}
var _ LiveBidderer = &MockLiveBidderer{}
var _ LiveBidder = &MockLiveBidder{}

//LiveBidder - helper interface to access copart live bidding api
//go:generate mockery -case underscore -name LiveBidder -inpkg
type LiveBidder interface {
	Send(ctx context.Context, request ParserCommand) error
	// Subscribe - creates new channel for copart messages. Use context cancel to unsubscribe
	Subscribe(ctx context.Context, id string) <-chan ParserEvent
	BuyerNumber() string
	Close()
	Done() <-chan struct{}
}

//LiveBidderer - helper interface to get LiveBidder from Config
//go:generate mockery -case underscore -name LiveBidderer -inpkg
type LiveBidderer interface {
	LiveBidder(ctx context.Context) (LiveBidder, error)
	Connector(ctx context.Context) (Connector, error)
	GetConfig() Config
}

//Config - represents Config for live bidder
type Config struct {
	IsAnonymous      bool     `fig:"is_anonymous"`
	Login            *string  `fig:"login"`
	Password         *string  `fig:"password"`
	Parser           *url.URL `fig:"parser,required"`
	DisableKeepAlive bool     `fig:"disable_keep_alive"`
}

type liveBidderer struct {
	getter      kv.Getter
	appLevelCtx context.Context
	log         *logan.Entry
	cfg         comfig.Once
	getterName  string

	lock      sync.Mutex
	bidder    LiveBidder
	connector Connector
}

func NewLiveBidderer(ctx context.Context, getter kv.Getter, getterName string, log *logan.Entry) LiveBidderer {
	return &liveBidderer{
		getter:      getter,
		log:         log.WithField("service", getterName),
		appLevelCtx: ctx,
		getterName:  getterName,
	}
}

func (c *liveBidderer) GetConfig() Config {
	return c.cfg.Do(func() interface{} {
		var cfg Config
		vals := kv.MustGetStringMap(c.getter, c.getterName)
		if err := figure.Out(&cfg).From(vals).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out copart-live-bidder Config"))
		}

		if !cfg.IsAnonymous && (cfg.Login == nil || cfg.Password == nil) {
			panic(errors.New("expected login and password to be specified in non anonymous mode"))
		}

		return cfg
	}).(Config)
}

//Connector - returns usable instance of connector
func (c *liveBidderer) Connector(ctx context.Context) (Connector, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	log := c.log.WithField("service", "copart-live-bidderer-connector")
	var err error
	var result Connector
	running.UntilSuccess(ctx, log, "copart-live-bidderer", func(ctx context.Context) (bool, error) {
		result, err = c.tryGetConnector(ctx, log)
		return err == nil, err
	}, time.Second, time.Minute)
	return result, errors.Wrap(err, "failed to get connector")
}

func (c *liveBidderer) tryGetConnector(ctx context.Context, log *logan.Entry) (Connector, error) {
	if c.connector != nil {
		select {
		case <-c.connector.Done():
			c.connector.Close()
			c.connector = nil
		default:
			return c.connector, nil
		}
	}

	cfg := c.GetConfig()
	var err error
	log.Debug("creating cookies provider")
	cookiesProvider := cookies.NewProvider(*cfg.Parser)
	log.Debug("getting cookies")
	jar, err := cookiesProvider.GetNewCookies(cfg.Login, cfg.Password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get cookies")
	}

	log.Debug("creating connector")
	c.connector, err = NewConnector(c.appLevelCtx, log, jar, cfg.IsAnonymous, cfg.DisableKeepAlive)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create connector")
	}

	return c.connector, nil
}

func (c *liveBidderer) LiveBidder(ctx context.Context) (LiveBidder, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	var err error
	var result LiveBidder
	log := c.log.WithField("service", "copart-live-bidderer")
	running.UntilSuccess(ctx, log, "copart-live-bidderer", func(ctx context.Context) (bool, error) {
		result, err = c.tryGetBidder(ctx, log)
		return err == nil, err
	}, 30*time.Second, time.Minute)

	return result, errors.Wrap(err, "failed to get bidder")
}

func (c *liveBidderer) tryGetBidder(ctx context.Context, log *logan.Entry) (_ LiveBidder, err error) {
	if c.bidder != nil {
		select {
		case <-c.bidder.Done():
			c.bidder.Close()
			c.bidder = nil
		default:
			return c.bidder, nil
		}
	}

	defer func() {
		if err == nil {
			return
		}

		if c.bidder != nil {
			c.bidder.Close()
		}
	}()

	cfg := c.GetConfig()
	c.bidder, err = newLiveBidder(c.appLevelCtx, log, cfg.Parser.String(), cfg.Login, cfg.Password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create live biddeer")
	}

	log.Debug("created live bidder")
	return c.bidder, nil
}
