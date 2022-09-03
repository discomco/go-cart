package ftor

import (
	"github.com/discomco/go-cart/robby/execute-game/-shared/model"

	sdk_domain "github.com/discomco/go-cart/sdk/domain"
)

const AggregateType = "robby.execute-game"

type IAggregate interface {
	sdk_domain.IAggregate
}

func AggFtor(newRoot model.DocFtor) sdk_domain.GenAggFtor[model.Root] {
	return func() sdk_domain.IAggregate {
		return sdk_domain.NewAggregate(AggregateType, newRoot())
	}
}
