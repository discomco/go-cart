package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/discomco/go-cart/config"
	"github.com/pkg/errors"
)

type ProducerFtor func() (*kafka.Producer, error)

type IProducer interface {
}

func newProducer(cfg config.IKafkaConfig) (*kafka.Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.GetBootstrapServers(),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create Kafka producer: %v", err)
	}
	return p, nil
}

func Producer(appConfig config.IAppConfig) ProducerFtor {
	return func() (*kafka.Producer, error) {
		cfg := appConfig.GetKafkaConfig()
		return newProducer(cfg)
	}
}
