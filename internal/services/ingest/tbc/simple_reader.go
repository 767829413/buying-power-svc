package tbc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	gKafka "github.com/segmentio/kafka-go"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/eAuction/events/go/kafka"
)

func PrintTBC(ctx context.Context, log *logan.Entry) {
	reader := kafka.NewReader(gKafka.ReaderConfig{
		//Brokers: []string{"localhost:9092"},
		Brokers: []string{"b-2.kafka-prod.yt4k8z.c3.kafka.eu-west-1.amazonaws.com:9092", "b-3.kafka-prod.yt4k8z.c3.kafka.eu-west-1.amazonaws.com:9092", "b-1.kafka-prod.yt4k8z.c3.kafka.eu-west-1.amazonaws.com:9092"},
		Topic:   kafka.Topic_tbc_transactions_v2.String(),
		GroupID: "local:7",
	})

	for {
		msg, err := reader.FetchMessage(ctx)
		if err != nil {
			log.WithError(err).Panic("failed to read message")
		}

		println(fmt.Sprintf("%s %d", msg.Time.String(), msg.Offset))

		var event kafka.TbcTransactionV2
		err = proto.Unmarshal(msg.Value, &event)
		if err != nil {
			log.WithError(err).Panic("failed to unmarhsal tbc tx")
		}

		if event.InvoiceId == "deposit:01F2RB71CA0X25WAMWQMNHDFKQ" || event.InvoiceId == "deposit:01F2R5318B63CZBHDBE7010WZA" {
			asJSON, err := json.Marshal(event)
			if err != nil {
				log.WithError(err).Panic("failed to marshal to json")
			}

			log.Info(string(asJSON))
		}

		err = reader.CommitMessages(ctx, msg)
		if err != nil {
			log.WithError(err).Panic("failed to commit msg")
		}
	}
}
