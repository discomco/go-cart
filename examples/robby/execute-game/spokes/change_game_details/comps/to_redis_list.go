package comps

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/change_game_details/behavior"
	behavior2 "github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	schema2 "github.com/discomco/go-cart/sdk/schema"
)

const (
	ToRedisListProjectionName = "toRedisList.GameDetailsChanged"
)

type IToRedisList interface {
	comps.IGenProjection[behavior.IEvt, schema.GameList]
}

func ToRedisList(
	newStore behavior2.StoreFtor[schema.GameList],
	newMapper behavior2.Evt2DocFunc[behavior.IEvt, schema.GameList],
	newList schema2.DocFtor[schema.GameList],
) IToRedisList {
	return comps.NewProjection[behavior.IEvt, schema.GameList](
		ToRedisListProjectionName,
		behavior.EVT_TOPIC,
		newStore,
		newMapper,
		newList,
		schema.GameListKey)
}
