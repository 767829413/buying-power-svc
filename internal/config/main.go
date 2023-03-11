package config

import (
	"context"

	"gitlab.com/eAuction/buying-power-svc/internal/services/api/ws/participants"

	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/eAuction/bouncer"
	"gitlab.com/eAuction/enumer"
	"gitlab.com/eAuction/events/go/kafka"
	"gitlab.com/eAuction/exrates-svc/pkg/exrates"
	"gitlab.com/eAuction/payment-service/pkg/payments"
	"gitlab.com/eAuction/platformer-svc/pkg/platformer"

	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/data/postgres"
)

//go:generate mockery --case underscore --name Config  --inpackage
type Config interface {
	comfig.Listenerer
	comfig.Logger
	pgdb.Databaser
	bouncer.Bouncerer
	exrates.Converterer
	platformer.Platformerer
	enumer.Enumerer
	CopartLiveBidderer
	payments.PaymentServicer

	Storage() data.Storage

	TBCTransactions() kafka.IngestConfig
	LotsIngester() kafka.IngestConfig
	IdentitiesIngester() kafka.IngestConfig
	LeadsIngester() kafka.IngestConfig
	Deposits() Deposits
	DepositsWriter() kafka.DepositEventWriter
	DepositIDGenerator() data.DepositIDGenerator
	LotsCacheByExternalIDPrefix() data.LotsCache
	AuctionsWriter() kafka.AuctionEventWriter
	ParticipantsRouter() participants.Router
	BuyingPowerRecharger() BuyingPowerRecharger
	AutoRechargeableIdentityTypes() []kafka.Identity_Type
	IsAutoRechargeable(kafka.Identity_Type) bool
}

type config struct {
	ctx context.Context
	comfig.Listenerer
	comfig.Logger
	pgdb.Databaser
	kafka.Kafkaer
	bouncer.Bouncerer
	exrates.Converterer
	platformer.Platformerer
	payments.PaymentServicer
	enumer.Enumerer
	CopartLiveBidderer

	getter               kv.Getter
	tbcTransactions      kafka.Ingester
	lotsIngest           kafka.Ingester
	identitiesIngest     kafka.Ingester
	leadsIngest          kafka.Ingester
	deposits             comfig.Once
	horizon              comfig.Once
	lotsCache            comfig.Once
	participantsRouter   comfig.Once
	buyingPowerRecharger comfig.Once
	copartPlatformID     string
}

func New(ctx context.Context, getter kv.Getter) Config {
	logger := comfig.NewLogger(getter, comfig.LoggerOpts{})
	kafkaer := kafka.NewKafkaer(getter)
	result := &config{
		ctx:             ctx,
		Kafkaer:         kafkaer,
		Listenerer:      comfig.NewListenerer(getter),
		Logger:          logger,
		Databaser:       pgdb.NewDatabaser(getter),
		Platformerer:    platformer.NewPlatformerer(getter),
		getter:          getter,
		Bouncerer:       bouncer.NewBouncerer(getter),
		Enumerer:        enumer.NewEnumerer(getter),
		PaymentServicer: payments.NewPaymentServicer(getter, logger.Log()),
		tbcTransactions: kafka.NewIngester("tbc_transactions_ingest_v2", getter, kafka.IngesterOpts{
			Topic:   kafka.Topic_tbc_transactions_v2,
			Log:     logger.Log(),
			Kafkaer: kafkaer,
		}),
		lotsIngest: kafka.NewIngester("lots_ingest", getter, kafka.IngesterOpts{
			Topic:   kafka.Topic_auctions_v2,
			Log:     logger.Log(),
			Kafkaer: kafkaer,
		}),
		identitiesIngest: kafka.NewIngester("identities", getter, kafka.IngesterOpts{
			Topic:   kafka.Topic_identities,
			Log:     logger.Log(),
			Kafkaer: kafkaer,
		}),
		leadsIngest: kafka.NewIngester("leads", getter, kafka.IngesterOpts{
			Topic:   kafka.Topic_leads,
			Log:     logger.Log(),
			Kafkaer: kafkaer,
		}),
		Converterer: exrates.NewConverterer(getter, exrates.ConvertererOpts{
			Logger: logger,
		}),
	}

	result.CopartLiveBidderer = NewCopartLiveBidderer(ctx, getter, result.Storage().Clone(), result.Log(),
		result.LotsCacheByExternalIDPrefix(), result.Platformer())
	return result
}

func (c *config) Storage() data.Storage {
	return postgres.New(c.DB()).Clone()
}

//TBCTransactions - returns config for tbc transactions ingest
func (c *config) TBCTransactions() kafka.IngestConfig {
	return c.tbcTransactions.Ingest()
}

func (c *config) LotsIngester() kafka.IngestConfig {
	return c.lotsIngest.Ingest()
}

func (c *config) IdentitiesIngester() kafka.IngestConfig {
	return c.identitiesIngest.Ingest()
}

func (c *config) LeadsIngester() kafka.IngestConfig {
	return c.leadsIngest.Ingest()
}

//DepositIDGenerator - returns deposit ID generator
func (c *config) DepositIDGenerator() data.DepositIDGenerator {
	return data.NewDepositIDGenerator()
}

func (c *config) DepositsWriter() kafka.DepositEventWriter {
	return kafka.NewDepositEventWriter(kafka.WriterConfig{
		BatchBytes:   c.Kafka().BatchBytes,
		BatchSize:    c.Kafka().BatchSize,
		BatchTimeout: c.Kafka().BatchTimeout,
		Brokers:      c.Kafka().Brokers,
		Topic:        kafka.Topic_deposits.String(),
	})
}

func (c *config) LotsCacheByExternalIDPrefix() data.LotsCache {
	return c.lotsCache.Do(func() interface{} {
		return data.NewLotsCacheByExternalIDPrefix(context.Background(), c.Storage(), c.Log())
	}).(data.LotsCache)
}

func (c *config) AuctionsWriter() kafka.AuctionEventWriter {
	return kafka.NewAuctionEventWriter(kafka.WriterConfig{
		BatchBytes:   c.Kafka().BatchBytes,
		BatchSize:    c.Kafka().BatchSize,
		BatchTimeout: c.Kafka().BatchTimeout,
		Brokers:      c.Kafka().Brokers,
		Topic:        kafka.Topic_auctions_v2.String(),
	})
}

func (c *config) ParticipantsRouter() participants.Router {
	return c.participantsRouter.Do(func() interface{} {
		return participants.NewRouter(c.ctx, c.Log(), c.Storage(), c.NewListener)
	}).(participants.Router)
}

func (c *config) AutoRechargeableIdentityTypes() []kafka.Identity_Type {
	return []kafka.Identity_Type{kafka.Identity_TYPE_DEALER, kafka.Identity_TYPE_BROKER}
}

func (c *config) IsAutoRechargeable(identityType kafka.Identity_Type) bool {
	for _, t := range c.AutoRechargeableIdentityTypes() {
		if identityType == t {
			return true
		}
	}

	return false
}
