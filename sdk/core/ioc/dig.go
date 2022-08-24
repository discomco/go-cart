package ioc

import (
	"go.uber.org/dig"
)

// IDig provides an interface to DI receiver
// this has a dependency on uber-go/dig
type IDig interface {
	Inject(dig IDig, items ...interface{}) IDig
	String() string
	Scope(name string, opts ...dig.ScopeOption) *dig.Scope
	Provide(constructor interface{}, opts ...dig.ProvideOption) error
	Decorate(decorator interface{}, opts ...dig.DecorateOption) error
	Invoke(function interface{}, opts ...dig.InvokeOption) error
}
