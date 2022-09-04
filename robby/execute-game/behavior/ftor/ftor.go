package ftor

import (
	"github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/sdk/model"

	sdk_domain "github.com/discomco/go-cart/sdk/domain"
)

const BehaviorName = "robby.execute-game"

type IBehavior interface {
	sdk_domain.IAggregate
}

func BehaviorFtor(newRoot model.DocFtor[schema.GameDoc]) sdk_domain.GenAggFtor[schema.GameDoc] {
	return func() sdk_domain.IAggregate {
		return sdk_domain.NewAggregate(BehaviorName, newRoot())
	}
}
