package ftor

import (
	"github.com/discomco/go-cart/robby/execute-game/schema"
	sdk_domain "github.com/discomco/go-cart/sdk/behavior"
	sdk_schema "github.com/discomco/go-cart/sdk/schema"
)

const BehaviorName = "robby.execute-game"

type IBehavior interface {
	sdk_domain.IBehavior
}

func BehaviorFtor(newRoot sdk_schema.DocFtor[schema.GameDoc]) sdk_domain.GenBehaviorFtor[schema.GameDoc] {
	return func() sdk_domain.IBehavior {
		return sdk_domain.NewBehavior(BehaviorName, newRoot())
	}
}
