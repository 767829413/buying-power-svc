package kafka

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Kafkaer interface {
	Kafka() *Config
}

type kafkaer struct {
	getter kv.Getter
	once   comfig.Once
}

func NewKafkaer(getter kv.Getter) Kafkaer {
	return &kafkaer{
		getter: getter,
	}
}

func (k *kafkaer) Kafka() *Config {
	return k.once.Do(func() interface{} {
		config := Config{
			// sane default since we are reinventing batch write.
			// higher value will make batch write to take N*BatchTimeout seconds
			BatchSize: 1,
		}
		err := figure.
			Out(&config).
			From(kv.MustGetStringMap(k.getter, "kafka")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out kafka config"))
		}
		return &config
	}).(*Config)
}
