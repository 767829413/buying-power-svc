package platformer

import (
	"net/url"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/eAuction/platformer-svc/resources"
)

// Platformer is a provider of platforms.
//go:generate mockery -inpkg -name Platformer
type Platformer interface {
	// GetPlatformList returns a list of all available platforms.
	GetPlatformList() (*resources.PlatformListResponse, error)
	// MustPlatformList returns a list of all available platforms. Panics if got error.
	MustPlatformList() *resources.PlatformListResponse
	// GetPlatform returns platform by it's ID. Returns nil, nil,
	// if platform doesn't exist.
	GetPlatform(id string) (*resources.PlatformResponse, error)
	// MustPlatform returns platform by it's ID. Returns nil, nil,
	// if platform doesn't exist. Panics if got error
	MustPlatform(id string) *resources.PlatformResponse
	// GetPlatformIDs returns a list of just platform IDs, without any details.
	GetPlatformIDs() ([]string, error)
	// MustPlatformIDs returns a list of just platform IDs, without any details.. Panics if got error.
	MustPlatformIDs() []string
	//GetPlatformByCode - returns platform by code
	GetPlatformByCode(code string) (*resources.PlatformResponse, error)
}

// Platformerer is a helper interface to get Platformer through config.
type Platformerer interface {
	Platformer() Platformer
}

type platformerer struct {
	getter kv.Getter
	once   comfig.Once
}

// NewPlatformerer returns new instance of Platformerer.
func NewPlatformerer(getter kv.Getter) Platformerer {
	return &platformerer{
		getter: getter,
	}
}

// Platformer figures out new Platformer instance from cfg.
func (c *platformerer) Platformer() Platformer {
	return c.once.Do(func() interface{} {
		var cfg struct {
			URL *url.URL `fig:"url,required"`
		}

		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "platformer")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out platformer cfg"))
		}

		return NewPlatformer(cfg.URL)
	}).(Platformer)
}
