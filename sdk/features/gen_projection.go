package features

import (
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/model"
)

type GenProjectionFtor[TEvt domain.IEvt, TState model.IReadModel] func() IGenProjection[TEvt, TState]

type IGenProjection[TEvt domain.IEvt, TState model.IReadModel] interface {
	IProjection
}
