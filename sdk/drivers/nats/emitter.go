package nats

import (
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/mediator"
	"github.com/discomco/go-cart/sdk/reactors"
	"github.com/nats-io/nats.go"
)

type IEmitter interface {
	reactors.IEmitter
}

type Emitter struct {
	*reactors.EventReactor
	natsBus  INATSBus
	mediator mediator.IMediator
	Topic    behavior.EventType
}

func NewEmitter(
	topic behavior.EventType,
	emitFact reactors.OnEvtFunc,
) (*Emitter, error) {
	eh := reactors.NewEventReactor(topic, emitFact)
	var b INATSBus
	var err error
	dig := ioc.SingleIoC()
	err = dig.Invoke(func(newBus reactors.GenBusFtor[*nats.Conn, *nats.Msg]) {
		b, err = newBus()
	})
	if err != nil {
		return nil, err
	}
	e := &Emitter{
		EventReactor: eh,
		natsBus:      b,
		Topic:        topic,
	}
	return e, nil
}

func (e *Emitter) IAmEmitter() {}
