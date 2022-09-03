package builder

import (
	"github.com/discomco/go-cart/robby/execute-game/-shared/schema"
	initialize_game "github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/behavior"
	"github.com/discomco/go-cart/sdk/domain"
)

func BehaviorBuilder(newAgg domain.GenAggFtor[schema.Root]) domain.AggBuilder {
	return func() domain.IAggregate {
		agg := newAgg()
		return agg.Inject(agg,
			initialize_game.ApplyEvt,
			initialize_game.TryCmd,
		)
	}
}
