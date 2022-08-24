package domain

import (
	"context"
)

type IGetRoot interface {
	GetRoot() IAggregate
}

type ILoad interface {
	Load(events []Event) error
}

type Whener interface {
	When(ctx context.Context, evt IEvt) error
}

type GenWhener[TEvt IEvt] interface {
	When(ctx context.Context, evt TEvt) error
}

type IGenWhen[TEvt IEvt] interface {
	GenWhener[TEvt]
}

func ImplementsIAggregate(aggregate IAggregate) bool {
	return true
}
