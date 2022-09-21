package ftor

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	sdk_behavior "github.com/discomco/go-cart/sdk/behavior"
	sdk_schema "github.com/discomco/go-cart/sdk/schema"
)

const (
	behaviorName = "quadratic-roots"
)

// IBehavior is the injection discriminator for the quadratic-roots behavior.
type IBehavior interface {
	sdk_behavior.IBehavior
}

// BehaviorFtor returns a function that creates an empty behavior, using the DocFtor to instantiate a new State for the Behavior.
func BehaviorFtor(newDoc sdk_schema.DocFtor[schema.QuadraticDoc]) sdk_behavior.GenBehaviorFtor[schema.QuadraticDoc] {
	return func() sdk_behavior.IBehavior {
		return sdk_behavior.NewBehavior(behaviorName, newDoc())
	}
}
