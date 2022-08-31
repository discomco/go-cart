package kafka

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/utils/convert"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/features"
	"github.com/pkg/errors"
)

type IEmitter interface {
	features.IFactEmitter
}

type emitter struct {
	*features.EventHandler
	evt2Fact domain.Evt2FactFunc
	producer *kafka.Producer
}

func (e *emitter) IAmEmitter() {}

func (e *emitter) emit(ctx context.Context, evt domain.IEvt) error {
	fact, err := e.evt2Fact(evt)
	if err != nil {
		return errors.Wrapf(err, "(%+v) could not convert event to fact", e.GetName())
	}
	data, err := convert.Any2Data(fact)
	if err != nil {
		return errors.Wrapf(err, "Failed to convert Fact %+v", err)
	}
	topic := string(e.GetEventType())
	e.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: data,
	}, nil)
	e.GetLogger().Infof("[%+v] emitted fact [%+v, %+v]", e.GetName(), e.GetEventType(), fact.GetId())
	return nil
}

func newEmitter(name features.Name,
	eventType domain.EventType,
	evt2Fact domain.Evt2FactFunc) (*emitter, error) {
	e := &emitter{
		evt2Fact: evt2Fact,
	}
	base := features.NewEventHandler(eventType, e.emit)
	var err error
	var p *kafka.Producer
	dig := ioc.SingleIoC()
	err = dig.Invoke(func(newProducer ProducerFtor) {
		p, err = newProducer()
		e.producer = p
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create producer")
	}
	e.EventHandler = base
	e.Name = name
	return e, nil
}

func NewEmitter(name features.Name,
	eventType domain.EventType,
	evt2Fact domain.Evt2FactFunc) (IEmitter, error) {
	return newEmitter(name, eventType, evt2Fact)
}
