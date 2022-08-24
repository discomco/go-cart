package mediator

import (
	"github.com/asaskevich/EventBus"
	"log"
	"sync"
)

var (
	cMutex = &sync.Mutex{}
	pMutex = &sync.Mutex{}
	sMutex = &sync.Mutex{}
)

type mediator struct {
	bus         EventBus.Bus
	knownTopics map[string]interface{}
}

var singleMediator IMediator

func newMediator() *mediator {
	return &mediator{
		bus:         EventBus.New(),
		knownTopics: make(map[string]interface{}, 0),
	}
}

func SingletonMediator() IMediator {
	if singleMediator == nil {
		cMutex.Lock()
		defer cMutex.Unlock()
		singleMediator = newMediator()
	}
	return singleMediator
}

func TransientDECBus() IMediator {
	return newMediator()
}

func (b *mediator) KnownTopics() map[string]interface{} {
	return b.knownTopics
}

func (b *mediator) setSubscription(topic string, fn interface{}) {
	sMutex.Lock()
	defer sMutex.Unlock()
	if b.knownTopics[topic] != nil {
		return
	}
	b.knownTopics[topic] = fn
}

func (b *mediator) Register(topic string, fn interface{}) error {
	b.setSubscription(topic, fn)
	return b.bus.Subscribe(topic, fn)
}

func (b *mediator) unsetSubscription(topic string) {
	if b.knownTopics[topic] == nil {
		return
	}
	delete(b.knownTopics, topic)
}

func (b *mediator) RegisterOnce(topic string, fn interface{}) error {
	return b.bus.SubscribeOnce(topic, fn)
}

func (b *mediator) HasCallback(topic string) bool {
	return b.bus.HasCallback(topic)
}

func (b *mediator) Unregister(topic string, fn interface{}) error {
	b.unsetSubscription(topic)
	return b.bus.Unsubscribe(topic, fn)
}

func (b *mediator) Broadcast(topic string, msg ...interface{}) {
	pMutex.Lock()
	defer pMutex.Unlock()
	b.bus.Publish(topic, msg...)
}

func (b *mediator) RegisterAsync(topic string, fn interface{}, transactional bool) error {
	b.setSubscription(topic, fn)
	err := b.bus.SubscribeAsync(topic, fn, transactional)
	if err != nil {
		log.Fatal(err)
		return err
	}
	//	b.bus.WaitAsync()
	return nil
}

func (b *mediator) RegisterOnceAsync(topic string, fn interface{}) error {
	return b.bus.SubscribeOnceAsync(topic, fn)
}

func (b *mediator) WaitAsync() {
	b.bus.WaitAsync()
}
