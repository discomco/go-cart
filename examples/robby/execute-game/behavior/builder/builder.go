package builder

import (
	"github.com/discomco/go-cart/robby/execute-game/schema"
	change_game_details "github.com/discomco/go-cart/robby/execute-game/spokes/change_game_details/behavior"
	change_game_settings "github.com/discomco/go-cart/robby/execute-game/spokes/change_game_settings/behavior"
	initialize_game "github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/behavior"
	"github.com/discomco/go-cart/sdk/behavior"
)

// BehaviorBuilder is a function that composes the behavior for the Execute Game Context.
func BehaviorBuilder(newAgg behavior.GenBehaviorFtor[schema.GameDoc]) behavior.BehaviorBuilder {
	return func() behavior.IBehavior {
		agg := newAgg()
		return agg.Inject(agg,
			initialize_game.ApplyEvt,
			initialize_game.TryCmd,
			change_game_details.TryCmd,
			change_game_details.ApplyEvt,
			change_game_settings.TryCmd,
			change_game_settings.ApplyEvt,
		)
	}
}
