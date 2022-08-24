package features

import (
	"github.com/discomco/go-cart/domain"
	"github.com/discomco/go-cart/model"
)

type GenProjectionFtor[TEvt domain.IEvt, TState model.IReadModel] func() IGenProjection[TEvt, TState]

type IGenProjection[TEvt domain.IEvt, TState model.IReadModel] interface {
	IProjection
}
