package config

import (
	"context"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/ws/copart"
	copartParser "gitlab.com/eAuction/copart-parser-v2/pkg/copart"
	"gitlab.com/eAuction/platformer-svc/pkg/platformer"
	"gitlab.com/eAuction/platformer-svc/resources"
	"strings"
	"sync"
)

type CopartLiveBidderer interface {
	Real(ctx context.Context, country string) (*copart.LiveBidder, error)
	Simulation(ctx context.Context, country string) (*copart.LiveBidder, error)
	GetSupportedCountries() []string
}

type liveBidderCfg struct {
	IsEnabled bool     `fig:"is_enabled"`
	Countries []string `fig:"countries"`
}

type copartLiveBidderer struct {
	appLevelCtx context.Context
	lock        sync.Mutex
	log         *logan.Entry
	lots        data.LotsCache
	storage     data.Storage
	getter      kv.Getter
	platforms   platformer.Platformer

	cfg comfig.Once

	real             map[string]*copart.LiveBidder
	realLiveBidderer map[string]copartParser.LiveBidderer
	simulation       map[string]*copart.LiveBidder
}

//NewCopartLiveBidderer -.
func NewCopartLiveBidderer(ctx context.Context, getter kv.Getter, storage data.Storage, log *logan.Entry,
	lots data.LotsCache, platforms platformer.Platformer) CopartLiveBidderer {

	return &copartLiveBidderer{
		appLevelCtx:      ctx,
		log:              log,
		lots:             lots,
		storage:          storage,
		getter:           getter,
		real:             map[string]*copart.LiveBidder{},
		simulation:       map[string]*copart.LiveBidder{},
		realLiveBidderer: map[string]copartParser.LiveBidderer{},
		platforms:        platforms,
	}
}

func (r *copartLiveBidderer) getConfig() liveBidderCfg {
	return r.cfg.Do(func() interface{} {
		raw := kv.MustGetStringMap(r.getter, "copart-live-bidder")

		var cfg liveBidderCfg

		if err := figure.Out(&cfg).With(figure.BaseHooks).From(raw).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out copart-live-bidder config"))
		}

		countries := map[string]struct{}{}
		for _, c := range cfg.Countries {
			c = strings.ToLower(c)
			_, ok := countries[c]
			if ok {
				panic(errors.From(errors.New("countries should not have duplicates"), logan.F{
					"country": c,
				}))
			}

			countries[c] = struct{}{}
		}

		return cfg
	}).(liveBidderCfg)
}

//GetSupportedCountries - returns list of supported countries
func (r *copartLiveBidderer) GetSupportedCountries() []string {
	cfg := r.getConfig()
	if cfg.IsEnabled {
		return cfg.Countries
	}
	return nil
}

type liveBidderCreator func(ctx context.Context, log *logan.Entry, country string, platforms []string) (bidder *copart.LiveBidder, err error)

func (r *copartLiveBidderer) getLiveBidder(ctx context.Context, country string, liveBidders map[string]*copart.LiveBidder,
	createLiveBidder liveBidderCreator) (*copart.LiveBidder, error) {

	log := r.log.WithField("copart-live-bidder-country", country)
	log.Debug("acquiring lock")
	r.lock.Lock()
	defer r.lock.Unlock()
	log.Debug("got lock")
	country = strings.ToLower(country)
	cfg := r.getConfig()
	if !cfg.IsEnabled {
		return nil, errors.New("copart live bidder is disabled")
	}

	liveBidder, ok := liveBidders[country]
	if ok && liveBidder.Ctx().Err() == nil {
		log.Debug("found healthy live bidder")
		return liveBidder, nil
	}

	if !r.isCountrySupported(country) {
		return nil, errors.From(errors.New("live bidding is not supported for country"), logan.F{
			"country": country,
		})
	}
	log.Debug("getting platforms for country")
	platforms := r.getPlatformIDs(func(p resources.Platform) bool {
		return strings.ToLower(p.Attributes.CountryCode) == country
	})

	if len(platforms) == 0 {
		return nil, errors.From(errors.New("no supported platforms for live bidder - not even creating it"), logan.F{
			"country": country,
		})
	}

	var err error
	liveBidder, err = createLiveBidder(ctx, log, country, platforms)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create live bidder")
	}
	log.Debug("got new live bidder")
	liveBidders[country] = liveBidder
	return liveBidder, nil
}

func (r *copartLiveBidderer) createRealLiveBidder(ctx context.Context, log *logan.Entry, country string, platforms []string) (bidder *copart.LiveBidder, err error) {
	liveBidderer, ok := r.realLiveBidderer[country]
	if !ok {
		log.Debug("copart bidderer is not initialized - creating new one")
		liveBidderer = copartParser.NewLiveBidderer(r.appLevelCtx, r.getter, "copart-live-bidder-"+country, log)
		r.realLiveBidderer[country] = liveBidderer
	}

	log.Debug("getting copart live bidder")
	copartLiveBidder, err := liveBidderer.LiveBidder(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get copart live bidder")
	}
	log.Debug("creating buying power live bidder")

	liveBidder, err := copart.NewLiveBidder(r.appLevelCtx, log, copartLiveBidder, platforms, r.storage.Clone(),
		r.lots,false,r.getCopartPlatformID())
	if err != nil {
		return nil, errors.Wrap(err, "failed to init copart bidder")
	}

	return liveBidder, nil
}

func (r *copartLiveBidderer) Real(ctx context.Context, country string) (*copart.LiveBidder, error) {
	return r.getLiveBidder(ctx, country, r.real, r.createRealLiveBidder)
}

func (r *copartLiveBidderer) getCopartPlatformID() string {
	p, err := r.platforms.GetPlatformByCode("COPART")
	if err != nil {
		panic(errors.Wrap(err, "failed to get copart platform"))
	}

	if p == nil {
		panic(errors.New("expected copart platform to exist"))
	}

	return p.Data.ID
}

func (r *copartLiveBidderer) isCountrySupported(country string) bool {
	for _, supported := range r.getConfig().Countries {
		if supported == country {
			return true
		}
	}

	return false
}

func (r *copartLiveBidderer) createSimulationLiveBidder(ctx context.Context, log *logan.Entry, country string, platforms []string) (*copart.LiveBidder, error) {
	/*simulation := copartParser.NewLiveBidderSimulation(r.appLevelCtx, r.storage.Lots(), r.log)

	newHandler := func(shelfSaleID string) copart.Handler {
		return copart.NewHandler(log, r.storage.Clone(), r.lots, strconv.FormatInt(simulation.GetBuyerNumber(), 10), platforms,
			true, r.getCopartPlatformID(), shelfSaleID)
	}

	var err error
	sim := copartParser.NewLiveBidderSimulationV2(ctx, r.storage.Clone().Lots(), r.log)
	liveBidder, err := copart.NewLiveBidder(r.appLevelCtx, r.log, sim, platforms)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create simulation live bidder")
	}

	return liveBidder, nil*/
	return nil, errors.New("simulation is not implemented")
}

func (r *copartLiveBidderer) Simulation(ctx context.Context, country string) (*copart.LiveBidder, error) {
	return r.getLiveBidder(ctx, country, r.simulation, r.createSimulationLiveBidder)
}

func (r *copartLiveBidderer) getPlatformIDs(isOk func(p resources.Platform) bool) []string {
	platforms := r.platforms.MustPlatformList()
	var result []string
	for _, platform := range platforms.Data {
		if isOk(platform) {
			result = append(result, platform.ID)
		}
	}

	return result
}
