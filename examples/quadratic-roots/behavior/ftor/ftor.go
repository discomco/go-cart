package ftor

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	sdk_behavior "github.com/discomco/go-cart/sdk/behavior"
	sdk_schema "github.com/discomco/go-cart/sdk/schema"
)

const (
	behaviorName = "quadratic-roots"
)

type IBehavior interface {
	sdk_behavior.IBehavior
}

func BehaviorFtor(newDoc sdk_schema.DocFtor[schema.QuadraticDoc]) sdk_behavior.GenBehaviorFtor[schema.QuadraticDoc] {
	return func() sdk_behavior.IBehavior {
		return sdk_behavior.NewBehavior(behaviorName, newDoc())
	}
}
