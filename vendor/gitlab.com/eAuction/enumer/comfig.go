package enumer

import (
	"net/url"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"gitlab.com/eAuction/enumer/pkg/state/remote"
)

type Enumerer interface {
	InjectRemoteStateGetter()
}

// NewEnumerer returns enumer instantiator.
func NewEnumerer(getter kv.Getter) Enumerer {
	return &enumerer{getter: getter}
}

type enumerer struct {
	getter kv.Getter
	comfig.Once
}

// InjectRemoteStateGetter constructs enumer instance from config.
func (e *enumerer) InjectRemoteStateGetter() {
	e.Do(func() interface{} {
		var cfg struct {
			URL          string `fig:"url,required"`
			FailOnUpdate bool   `fig:"fail_on_update"`
		}
		err := figure.Out(&cfg).From(kv.MustGetStringMap(e.getter, "enumer")).Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out enumer config"))
		}

		parsedURL, err := url.Parse(cfg.URL)
		if err != nil {
			panic(errors.Wrap(err, "failed to parse enumer url"))
		}

		remoteGetter, err := remote.NewGetter(*parsedURL, cfg.FailOnUpdate)
		if err != nil {
			panic(errors.Wrap(err, "failed to instan"))
		}
		InjectStateGetter(remoteGetter)
		return nil
	})
}
