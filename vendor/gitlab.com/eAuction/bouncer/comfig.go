package bouncer

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Bouncerer interface {
	Bouncer() Bouncer
}

func NewBouncerer(getter kv.Getter) Bouncerer {
	return &bouncerer{
		getter: getter,
	}
}

type bouncerer struct {
	getter kv.Getter
	comfig.Once
}

func (b *bouncerer) Bouncer() Bouncer {
	return b.Do(func() interface{} {
		var config struct {
			SkipCheck bool `fig:"skip_checks"`
		}
		err := figure.
			Out(&config).
			From(kv.MustGetStringMap(b.getter, "bouncer")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out bouncer"))
		}
		return New(Opts{
			SkipChecks: config.SkipCheck,
		})
	}).(Bouncer)
}
