package module

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"interaction-api/config"
	"interaction-api/domain"
	"log"
	"strings"
)

type KafkaPub struct {
	pubs *kafka.Publisher
}

var AppKafkaPub *KafkaPub

func Configure() *KafkaPub {
	var pbs KafkaPub
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	// equivalent of auto.offset.reset: earliest
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   strings.Split(config.AppConfig.KafkaConfig.Brokers, ","),
			Marshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		fmt.Println("Failed to start kafka publisher")
		//panic(err)
	}
	pbs.pubs = publisher
	AppKafkaPub = &pbs
	return AppKafkaPub
}

func (k *KafkaPub) PublishLike(ctx context.Context, topic string, interaction domain.Interaction) error {
	jsonString, err := json.Marshal(interaction)
	if err != nil {
		return err
	}
	msg := message.NewMessage(watermill.NewUUID(), jsonString)
	if err2 := k.pubs.Publish(topic, msg); err2 != nil {
		return err2
	}
	log.Printf("Success Sending %v to kafka with topic %s ", interaction, topic)
	return nil
}

func (k *KafkaPub) PublishView(ctx context.Context, topic string, interaction domain.InteractionView) error {
	jsonString, err := json.Marshal(interaction)
	if err != nil {
		return err
	}
	msg := message.NewMessage(watermill.NewUUID(), jsonString)
	if err2 := k.pubs.Publish(topic, msg); err2 != nil {
		return err2
	}
	log.Printf("Success Sending %v to kafka with topic %s ", interaction, topic)
	return nil
}

func (k *KafkaPub) PublishFollow(ctx context.Context, topic string, interaction domain.InteractionFollow) error {
	jsonString, err := json.Marshal(interaction)
	if err != nil {
		return err
	}
	msg := message.NewMessage(watermill.NewUUID(), jsonString)
	if err2 := k.pubs.Publish(topic, msg); err2 != nil {
		return err2
	}
	log.Printf("Success Sending %v to kafka with topic %s ", interaction, topic)
	return nil
}
