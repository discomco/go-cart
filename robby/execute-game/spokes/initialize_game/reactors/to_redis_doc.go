package reactors

import (
	"github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/behavior"
	sdk_behavior "github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/reactors"
	sdk_schema "github.com/discomco/go-cart/sdk/schema"
)

const (
	ProjectionName = "toRedisDoc.Initialized"
)

type IToRedisDoc interface {
	reactors.IGenProjection[behavior.IEvt, schema.GameDoc]
}

func ToRedisDoc(
	newStoreFtor sdk_behavior.StoreFtor[schema.GameDoc],
	evt2Doc sdk_behavior.Evt2ModelFunc[behavior.IEvt, schema.GameDoc],
	newDocFtor sdk_schema.DocFtor[schema.GameDoc]) IToRedisDoc {
	return reactors.NewProjection[behavior.IEvt, schema.GameDoc](
		ProjectionName,
		behavior.EVT_TOPIC,
		newStoreFtor,
		evt2Doc,
		newDocFtor,
		nil,
	)
}
