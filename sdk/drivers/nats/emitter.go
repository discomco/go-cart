package nats

import (
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/mediator"
	"github.com/nats-io/nats.go"
)

type IEmitter interface {
	comps.IEmitter
}

type Emitter struct {
	*comps.EventReactor
	natsBus  INATSBus
	mediator mediator.IMediator
	Topic    behavior.EventType
}

func NewEmitter(
	topic behavior.EventType,
	emitFact comps.OnEvtFunc,
) (*Emitter, error) {
	eh := comps.NewEventReactor(topic, emitFact)
	var b INATSBus
	var err error
	dig := ioc.SingleIoC()
	err = dig.Invoke(func(newBus comps.GenBusFtor[*nats.Conn, *nats.Msg]) {
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
