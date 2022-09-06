package behavior

import (
	"github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/robby/execute-game/spokes/change_game_settings/contract"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/pkg/errors"
)

func EvtToList() behavior.Evt2ModelFunc[IEvt, schema.GameList] {
	return func(evt IEvt, list *schema.GameList) error {
		return evtToList(evt, list)
	}
}

func evtToList(evt IEvt, list *schema.GameList) error {
	gameId := evt.GetAggregateId()
	gameIt := list.GetItem(gameId)
	var pl contract.Payload
	err := evt.GetPayload(&pl)
	if err != nil {
		return errors.Wrapf(err, "(change_game_settings.evtToList) failed to get payload from event %v", evt)
	}
	gameIt.NumberOfPlayers = pl.Settings.NbrOfPlayers
	list.Items[gameId] = gameIt
	return nil
}
