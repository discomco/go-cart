package domain

import "github.com/discomco/go-cart/sdk/domain"

func AggBuilder(newAgg AggFtor) domain.AggBuilder {
	return func() domain.IAggregate {
		agg := newAgg()
		return agg.Inject(agg)
	}
}
