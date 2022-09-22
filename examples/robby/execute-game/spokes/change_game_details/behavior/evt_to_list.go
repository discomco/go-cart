package behavior

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/change_game_details/contract"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/pkg/errors"
)

func EvtToList() behavior.Evt2DocFunc[IEvt, schema.GameList] {
	return func(evt IEvt, list *schema.GameList) error {
		return evtToList(evt, list)
	}
}

func evtToList(evt IEvt, list *schema.GameList) error {
	gameId := evt.GetBehaviorId()
	gameIt := list.GetItem(gameId)
	var pl contract.Payload
	err := evt.GetPayload(&pl)
	if err != nil {
		return errors.Wrapf(err, "(change_game_details.evtToList) failed to get payload from event %v", evt)
	}
	gameIt.Name = pl.Details.Name
	list.Items[gameId] = gameIt
	return nil
}
