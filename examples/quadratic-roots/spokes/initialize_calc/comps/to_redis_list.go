package comps

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/behavior"
	behavior2 "github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	schema2 "github.com/discomco/go-cart/sdk/schema"
)

const (
	ToRedisListProjectionName = "projection(CalculationInitialized.toRedisList)"
)

type IToRedisList interface {
	comps.IGenProjection[behavior.IEvt, schema.QuadraticList]
}

func ToRedisList(
	newListStore behavior2.StoreFtor[schema.QuadraticList],
	evt2List behavior2.Evt2DocFunc[behavior.IEvt, schema.QuadraticList],
	newList schema2.DocFtor[schema.QuadraticList],
) IToRedisList {
	return comps.NewProjection[behavior.IEvt, schema.QuadraticList](
		ToRedisListProjectionName,
		behavior.EvtTopic,
		newListStore,
		evt2List,
		newList,
		schema.DefaultCalcListId,
	)
}
