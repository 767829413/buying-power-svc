package kafka

import (
	"fmt"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"reflect"
	"time"
)

//IngestConfig - configuration options for Ingest
type IngestConfig struct {
	Config
	ingestConfig
	Log   *logan.Entry
	Topic Topic
}

type ingestConfig struct {
	IngestBatchSize int           `fig:"batch_size"`
	MinRetryPeriod  time.Duration `fig:"min_retry_period"`
	MaxRetryPeriod  time.Duration `fig:"max_retry_period"`
	SkipUntilOffset map[int]int64 `fig:"skip_until_offset"`
	// allows to override group id from kafkaer
	GroupID string `fig:"group_id"`
}

//IngesterOpts - helper struct that used to pass into IngestConfig additional values
type IngesterOpts struct {
	Topic   Topic        // required
	Log     *logan.Entry // required
	Kafkaer              //required
}

//Ingester - figures out config for Ingest
type Ingester interface {
	Ingest() IngestConfig
}

type ingester struct {
	getter kv.Getter
	once   comfig.Once
	name   string

	opts IngesterOpts
}

//NewIngester - returns new implementation of Ingester. Panics if one of the required opts are not passed
func NewIngester(name string, getter kv.Getter, opts IngesterOpts) Ingester {
	if opts.Log == nil {
		panic("Log is required opts for Ingester")
	}

	if opts.Kafkaer == nil {
		panic("kafkaer is required opts for Ingester")
	}

	return &ingester{
		getter: getter,
		opts:   opts,
		name:   name,
	}
}

//Ingest - returns config for Ingest
func (i *ingester) Ingest() IngestConfig {
	return i.once.Do(func() interface{} {
		config := ingestConfig{
			IngestBatchSize: 100,
			MinRetryPeriod:  time.Second,
			MaxRetryPeriod:  time.Minute,
			SkipUntilOffset: map[int]int64{},
		}
		err := figure.
			Out(&config).With(figure.BaseHooks, mapHooks).
			From(kv.MustGetStringMap(i.getter, i.name)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out ingester config config", logan.F{
				"name": i.name,
			}))
		}

		kafkaCfg := i.opts.Kafkaer.Kafka()

		groupID := config.GroupID
		if groupID == "" {
			groupID = kafkaCfg.GroupID
		}

		return IngestConfig{
			Config:       *kafkaCfg,
			ingestConfig: config,
			Log: i.opts.Log.WithFields(logan.F{
				"ingest":   i.name,
				"topic":    i.opts.Topic.String(),
				"group_id": groupID,
			}),
			Topic: i.opts.Topic,
		}
	}).(IngestConfig)
}

var mapHooks = figure.Hooks{
	"map[int]int64": func(value interface{}) (reflect.Value, error) {
		switch v := value.(type) {
		case map[string]interface{}:
			result := map[int]int64{}
			for key, value := range v {
				result[cast.ToInt(key)] = cast.ToInt64(value)
			}
			return reflect.ValueOf(result), nil
		default:
			return reflect.Value{}, fmt.Errorf("unsupported conversion from %T to map[int]int64", value)
		}
	},
}
