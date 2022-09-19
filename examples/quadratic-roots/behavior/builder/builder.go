package builder

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"github.com/discomco/go-cart/sdk/behavior"
)

func BehaviorBuilder(newBehavior behavior.GenBehaviorFtor[schema.QuadraticDoc]) behavior.BehaviorBuilder {
	return func() behavior.IBehavior {
		behavior := newBehavior()
		// TODO: inject apply() and try() functions here
		return behavior
	}
}
