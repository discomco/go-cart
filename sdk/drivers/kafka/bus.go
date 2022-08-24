package kafka

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/discomco/go-cart/features"
)

type IKafkaBus interface {
	features.IBus
}

type bus struct {
	*kafka.Consumer
	*kafka.Producer
}

func (b *bus) Publish(topic string, data []byte) {
	//TODO implement me
	panic("implement me")
}

func (b *bus) Listen(topic string, facts chan []byte) {
	//TODO implement me
	panic("implement me")
}

func (b *bus) Request(topic string, data []byte, timeout time.Duration) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func newKafkaBus() *bus {
	return &bus{}
}
