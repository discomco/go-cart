package reactors

import (
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/schema"
)

type GenProjectionFtor[TEvt behavior.IEvt, TState schema.IReadModel] func() IGenProjection[TEvt, TState]

type IGenProjection[TEvt behavior.IEvt, TState schema.IReadModel] interface {
	IProjection
}
