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
	InitializedLinkName = "link.Initialized.CalculateRoots"
)

type IInitializedLink interface {
	comps.IBehaviorLink
}

type initLink struct {
	*comps.BehaviorLink
}

func newInitLink(newCH comps.CmdHandlerFtor) *initLink {
	l := &initLink{}
	b := comps.NewBehaviorLink(
		InitializedLinkName,
		behavior.EvtTopic,
		l.linkFunc, newCH)
	l.BehaviorLink = b
	return l
}

func InitializedLink(newCH comps.CmdHandlerFtor) IInitializedLink {
	return newInitLink(newCH)
}

func (l *initLink) linkFunc(ctx context.Context, evt sdk_behavior.IEvt) error {
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
