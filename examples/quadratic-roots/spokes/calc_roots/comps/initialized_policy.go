package comps

import (
	"context"
	calc_roots_behavior "github.com/discomco/go-cart/examples/quadratic-roots/spokes/calc_roots/behavior"
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/calc_roots/contract"
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/behavior"
	sdk_behavior "github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/pkg/errors"
)

const (
	InitializedPolicyName = "policy(Initialized ~> CalculateRoots)"
)

type IInitializedPolicy interface {
	comps.IPolicy
}

type initializedPolicy struct {
	*comps.Policy
}

func newInitializedPolicy(newCH comps.CmdHandlerFtor) *initializedPolicy {
	l := &initializedPolicy{}
	b := comps.NewPolicy(
		InitializedPolicyName,
		behavior.EvtTopic,
		l.linkFunc, newCH)
	l.Policy = b
	return l
}

func InitializedPolicy(newCH comps.CmdHandlerFtor) IInitializedPolicy {
	return newInitializedPolicy(newCH)
}

func (l *initializedPolicy) linkFunc(ctx context.Context, evt sdk_behavior.IEvt) error {
	calcID, err := evt.GetBehaviorID()
	if err != nil {
		return err
	}
	calcPl := contract.NewHopePayload()
	calcCmd, err := calc_roots_behavior.NewCmd(calcID, *calcPl)
	if err != nil {
		return err
	}
	ch := l.NewCH()
	fbk := ch.Handle(ctx, calcCmd)
	if !fbk.IsSuccess() {
		err := errors.Wrapf(err, "failed(%+v)", fbk.GetFlattenedErrors())
		l.GetLogger().Errorf("failed(%+v)", fbk.GetFlattenedErrors())
		return err
	}
	return nil

}
