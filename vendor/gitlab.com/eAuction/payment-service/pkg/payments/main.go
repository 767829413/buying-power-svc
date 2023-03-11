package payments

import (
	"net/url"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

// var _ PaymentService = &MockPaymentService{}
var _ PaymentServicer = &MockPaymentServicer{}

//PaymentServicer - helper interface to get payment service from Config
//go:generate mockery --case underscore --name PaymentServicer --inpackage
type PaymentServicer interface {
	PaymentService() PaymentService
}

//Config - represents Config for live bidder
type Config struct {
	URL *url.URL `fig:"url,required"`
}

type paymentServiceer struct {
	getter kv.Getter
	once   comfig.Once
	log    *logan.Entry
}

func NewPaymentServicer(getter kv.Getter, log *logan.Entry) PaymentServicer {
	return &paymentServiceer{
		getter: getter,
		log:    log.WithField("service", "payment-service"),
	}
}

//EnsureValid - panics if config is not valid
func (cfg Config) EnsureValid() {
	if cfg.URL == nil {
		panic(errors.New("expected payment service URL"))
	}
}

//PaymentService - returns usable instance of paymentService
func (c *paymentServiceer) PaymentService() PaymentService {
	return c.once.Do(func() interface{} {
		var cfg struct {
			URL *url.URL `fig:"url,required"`
		}

		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "payment_service")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out payment_service cfg"))
		}

		return NewPaymentService(*cfg.URL, c.log)
	}).(PaymentService)
}
