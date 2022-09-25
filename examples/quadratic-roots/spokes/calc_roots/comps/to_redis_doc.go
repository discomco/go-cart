package comps

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/calc_roots/behavior"
	sdk_behavior "github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	sdk_schema "github.com/discomco/go-cart/sdk/schema"
)

const (
	ProjectionName = "toRedisDoc.RootsCalculated"
)

type IToRedisDoc interface {
	comps.IGenProjection[behavior.IEvt, schema.QuadraticDoc]
}

func ToRedisDoc(
	newStore sdk_behavior.StoreFtor[schema.QuadraticDoc],
	evt2Doc sdk_behavior.Evt2DocFunc[behavior.IEvt, schema.QuadraticDoc],
	newDoc sdk_schema.DocFtor[schema.QuadraticDoc],
) IToRedisDoc {
	return comps.NewProjection[behavior.IEvt, schema.QuadraticDoc](
		ProjectionName,
		behavior.EvtTopic,
		newStore,
		evt2Doc,
		newDoc,
		nil,
	)
}
