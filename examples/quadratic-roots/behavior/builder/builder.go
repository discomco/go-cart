package builder

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	initialize_calc "github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/behavior"
	"github.com/discomco/go-cart/sdk/behavior"
)

// BehaviorBuilder returns a builder function that composes the behavior for the process.
func BehaviorBuilder(newCalculation behavior.GenBehaviorFtor[schema.QuadraticDoc]) behavior.BehaviorBuilder {
	return func() behavior.IBehavior {
		calculation := newCalculation()
		return calculation.Inject(calculation,
			initialize_calc.TryCmd,
			initialize_calc.ApplyEvt,
		)
	}
}
