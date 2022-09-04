package builder

import (
	"github.com/discomco/go-cart/robby/execute-game/schema"
	initialize_game "github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/behavior"
	"github.com/discomco/go-cart/sdk/behavior"
)

func BehaviorBuilder(newAgg behavior.GenBehaviorFtor[schema.GameDoc]) behavior.BehaviorBuilder {
	return func() behavior.IBehavior {
		agg := newAgg()
		return agg.Inject(agg,
			initialize_game.ApplyEvt,
			initialize_game.TryCmd,
		)
	}
}
