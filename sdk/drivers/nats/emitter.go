package nats

import (
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/mediator"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/features"
	"github.com/nats-io/nats.go"
)

type Emitter struct {
	*features.EventHandler
	natsBus  INATSBus
	mediator mediator.IMediator
	Topic    domain.EventType
}

func NewNatsEmitter(
	topic domain.EventType,
	emitFact features.OnEvtFunc,
) (*Emitter, error) {
	eh := features.NewEventHandler(topic, emitFact)
	var b INATSBus
	var err error
	dig := ioc.SingleIoC()
	err = dig.Invoke(func(newBus features.GenBusFtor[*nats.Conn, *nats.Msg]) {
		b, err = newBus()
	})
	if err != nil {
		return nil, err
	}
	e := &Emitter{
		EventHandler: eh,
		natsBus:      b,
		Topic:        topic,
	}
	return e, nil
}

func (e *Emitter) IAmEmitter() {}
