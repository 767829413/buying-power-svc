package exrates

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/url"
	"time"
)

//Converterer - helper interface to get instance of Converter via config
type Converterer interface {
	ExRatesConverter() *Converter
}

type converterer struct {
	getter kv.Getter
	once   comfig.Once
	opts   ConvertererOpts
}

//ConvertererOpts - defines params to configure Converter
type ConvertererOpts struct {
	Logger comfig.Logger
}

//NewConverterer - returns new instance of Converterer
func NewConverterer(getter kv.Getter, opts ConvertererOpts) Converterer {
	return &converterer{
		getter: getter,
		opts:   opts,
	}
}

//ExRatesConverter - returns configured instance of converter.
func (c *converterer) ExRatesConverter() *Converter {
	return c.once.Do(func() interface{} {
		var config = struct {
			URL          *url.URL      `fig:"url,required"`
			UpdatePeriod time.Duration `fig:"update_period"`
			TTL          time.Duration `fig:"ttl"`
		}{
			UpdatePeriod: time.Minute * 10,
			TTL:          time.Hour,
		}

		err := figure.Out(&config).
			From(kv.MustGetStringMap(c.getter, "exrates_converter")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return NewConverter(c.opts.Logger.Log().WithField("service", "exrates_converter"), *config.URL, config.UpdatePeriod, config.TTL)
	}).(*Converter)
}
