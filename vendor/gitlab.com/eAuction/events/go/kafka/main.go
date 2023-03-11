package kafka

import (
	"sync"
	"time"

	"gitlab.com/distributed_lab/lorem"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

var (
	NewWriter = kafka.NewWriter
	NewReader = kafka.NewReader
)

type (
	ReaderConfig kafka.ReaderConfig
	WriterConfig kafka.WriterConfig

	Config struct {
		BatchBytes   int           `fig:"batch_bytes"`
		BatchSize    int           `fig:"batch_size"`
		BatchTimeout time.Duration `fig:"batch_timeout"`
		MaxWait      time.Duration `fig:"max_wait"`
		GroupID      string        `fig:"group_id"`
		Brokers      []string      `fig:"brokers,required"`
	}
)

func (cfg Config) reader(topic string) ReaderConfig {
	if cfg.GroupID == "RANDOM" {
		cfg.GroupID = lorem.ULID()
	}
	return ReaderConfig{
		Topic:   topic,
		GroupID: cfg.GroupID,
		Brokers: cfg.Brokers,
		MaxWait: cfg.MaxWait,
	}
}

func (cfg Config) writer(topic string) WriterConfig {
	return WriterConfig{
		BatchBytes:   cfg.BatchBytes,
		BatchSize:    cfg.BatchSize,
		BatchTimeout: cfg.BatchTimeout,
		Topic:        topic,
		Brokers:      cfg.Brokers,
	}
}

// DEPRECATED: consider Kafkaer instead
type ConfigProvider interface {
	Config() Config
}

type configProvider struct {
	getter kv.Getter
	once   sync.Once
	value  Config
	err    error
}

// DEPRECATED: consider NewKafkaer instead
func NewConfigProvider(getter kv.Getter) ConfigProvider {
	return &configProvider{
		getter: getter,
	}
}

func (p *configProvider) Config() Config {
	var cfg Config

	p.once.Do(func() {
		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(p.getter, "kafka")).
			Please()
		if err != nil {
			p.err = errors.Wrap(err, "failed to figure out kafka config")
			return
		}
	})
	if p.err != nil {
		panic(p.err)
	}

	return cfg
}
