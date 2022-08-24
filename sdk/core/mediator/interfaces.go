package mediator

import "github.com/dustin/go-broadcast"

// IMediator is an injection that represents an in-memory eventing bus which
// allows for internal decoupling of components

type IMediator interface {
	Register(topic string, fn interface{}) error
	RegisterOnce(topic string, fn interface{}) error
	HasCallback(topic string) bool
	Unregister(topic string, fn interface{}) error
	Broadcast(topic string, msg ...interface{})
	RegisterAsync(topic string, fn interface{}, transactional bool) error
	RegisterOnceAsync(topic string, fn interface{}) error
	WaitAsync()
	KnownTopics() map[string]interface{}
}

type IBroadcaster interface {
	broadcast.Broadcaster
}
