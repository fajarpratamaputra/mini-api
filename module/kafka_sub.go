package module

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"interaction-api/config"
	"strings"
)

type KafkaSub struct {
	subs *kafka.Subscriber
}

var AppKafkaSubs *KafkaSub

func ConfigureSub() *KafkaSub {
	var sbs KafkaSub
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	// equivalent of auto.offset.reset: earliest
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:     strings.Split(config.AppConfig.KafkaConfig.Brokers, ","),
			Unmarshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		fmt.Println("Failed to start kafka publisher")
		//panic(err)
	}
	sbs.subs = subscriber
	AppKafkaSubs = &sbs
	return AppKafkaSubs
}
