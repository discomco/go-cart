package builder

import (
	"github.com/discomco/go-cart/robby/execute-game/-shared/model"
	"github.com/discomco/go-cart/sdk/domain"
)

func AggBuilder(newAgg domain.GenAggFtor[model.Root]) domain.AggBuilder {
	return func() domain.IAggregate {
		agg := newAgg()
		return agg.Inject(agg)
	}
}
