package comps

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	calc_roots_behavior "github.com/discomco/go-cart/examples/quadratic-roots/spokes/calc_roots/behavior"
	sdk_behavior "github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	sdk_schema "github.com/discomco/go-cart/sdk/schema"
)

const (
	ToRedisListProjectionName = "toRedisList.RootsCalculated"
)

type IToRedisList interface {
	comps.IGenProjection[calc_roots_behavior.IEvt, schema.QuadraticList]
}

func ToRedisList(
	newListStore sdk_behavior.StoreFtor[schema.QuadraticList],
	evt2List sdk_behavior.Evt2DocFunc[calc_roots_behavior.IEvt, schema.QuadraticList],
	newList sdk_schema.DocFtor[schema.QuadraticList],
) IToRedisList {
	return comps.NewProjection[calc_roots_behavior.IEvt, schema.QuadraticList](
		ToRedisListProjectionName,
		calc_roots_behavior.EvtTopic,
		newListStore,
		evt2List,
		newList,
		schema.DefaultCalcListId,
	)
}
