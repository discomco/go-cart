package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/discomco/go-cart/sdk/config"
	"sync"
)

type IConsumer interface {
}

type ConsumerFtor func() (*kafka.Consumer, error)

func newConsumer(cfg config.IKafkaConfig) (*kafka.Consumer, error) {
	return kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.GetBootstrapServers(),
		"group.id":          cfg.GetGroupId(),
		"auto.offset.reset": cfg.GetAutoOffsetReset(),
	})
}

var (
	singletonConsumer *kafka.Consumer
	cMutex            = &sync.Mutex{}
)

func singleConsumer(cfg config.IKafkaConfig) (*kafka.Consumer, error) {
	if singletonConsumer == nil {
		cMutex.Lock()
		cMutex.Unlock()
		s, err := newConsumer(cfg)
		if err != nil {
			return nil, err
		}
		singletonConsumer = s
	}
	return singletonConsumer, nil
}

func SingleConsumer(appConfig config.IAppConfig) ConsumerFtor {
	return func() (*kafka.Consumer, error) {
		return singleConsumer(appConfig.GetKafkaConfig())
	}
}

func TransientConsumer(appConfig config.IAppConfig) ConsumerFtor {
	return func() (*kafka.Consumer, error) {
		return newConsumer(appConfig.GetKafkaConfig())
	}
}
