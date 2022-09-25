package comps

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	calc_roots_behavior "github.com/discomco/go-cart/examples/quadratic-roots/spokes/calc_roots/behavior"
	sdk_behavior "github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	sdk_schema "github.com/discomco/go-cart/sdk/schema"
)

const (
	ToRedisDocProjectionName = "toRedisDoc.RootsCalculated"
)

type IToRedisDoc interface {
	comps.IGenProjection[calc_roots_behavior.IEvt, schema.QuadraticDoc]
}

func ToRedisDoc(
	newStore sdk_behavior.StoreFtor[schema.QuadraticDoc],
	evt2Doc sdk_behavior.Evt2DocFunc[calc_roots_behavior.IEvt, schema.QuadraticDoc],
	newDoc sdk_schema.DocFtor[schema.QuadraticDoc],
) IToRedisDoc {
	return comps.NewProjection[calc_roots_behavior.IEvt, schema.QuadraticDoc](
		ToRedisDocProjectionName,
		calc_roots_behavior.EvtTopic,
		newStore,
		evt2Doc,
		newDoc,
		nil,
	)
}
