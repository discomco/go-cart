package ftor

import (
	"github.com/discomco/go-cart/robby/execute-game/-shared/schema"

	sdk_domain "github.com/discomco/go-cart/sdk/domain"
)

const BehaviorName = "robby.execute-game"

type IBehavior interface {
	sdk_domain.IAggregate
}

func BehaviorFtor(newRoot schema.DocFtor) sdk_domain.GenAggFtor[schema.Root] {
	return func() sdk_domain.IAggregate {
		return sdk_domain.NewAggregate(BehaviorName, newRoot())
	}
}
