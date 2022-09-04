package actors

import (
	"github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/behavior"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/features"
	"github.com/discomco/go-cart/sdk/model"
)

const (
	ProjectionName = "toRedisDoc.Initialized"
)

type IToRedisDoc interface {
	features.IGenProjection[behavior.IEvt, schema.GameDoc]
}

func ToRedisDoc(
	newStoreFtor domain.StoreFtor[schema.GameDoc],
	evt2Doc domain.Evt2ModelFunc[behavior.IEvt, schema.GameDoc],
	newDocFtor model.DocFtor[schema.GameDoc]) IToRedisDoc {
	return features.NewProjection[behavior.IEvt, schema.GameDoc](
		ProjectionName,
		behavior.EVT_TOPIC,
		newStoreFtor,
		evt2Doc,
		newDocFtor,
		nil,
	)
}
