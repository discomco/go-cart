package comps

import (
	"context"
	"github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/robby/execute-game/spokes/change_game_details/behavior"
	contract2 "github.com/discomco/go-cart/robby/execute-game/spokes/change_game_details/contract"
	behavior2 "github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/behavior"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/contract"
	sdk_behavior "github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/pkg/errors"
	"sync"
)

const LinkName = "change_game_details.InitializedLink"

type IInitializedLink interface {
	comps.IBehaviorLink
}

type link struct {
	*comps.BehaviorLink
}

var wMutex = &sync.Mutex{} //

func (l *link) onEvtFunc(ctx context.Context, evt sdk_behavior.IEvt) error {
	wMutex.Lock()
	defer wMutex.Unlock()
	docID, err := evt.GetAggregateID()
	if err != nil {
		return errors.Wrapf(err, "failed to get aggregate ID from event %v", evt)
	}
	var pl contract.Payload
	err = evt.GetPayload(&pl)
	details := schema.NewDetails(pl.Details.Name)
	detailsPl := contract2.NewPayload(details)
	cmd, err := behavior.NewCmd(docID, *detailsPl)
	ch := l.NewCH()
	fbk := ch.Handle(ctx, cmd)
	if !fbk.IsSuccess() {
		err := errors.Wrapf(err, "failed(%+v)", fbk.GetErrors()[0])
		l.GetLogger().Errorf("failed(%+v)", fbk.GetErrors()[0])
		return err
	}
	return nil
}

func newLink(newCH comps.CmdHandlerFtor) *link {
	l := &link{}
	l.BehaviorLink = comps.NewBehaviorLink(
		LinkName, behavior2.EVT_TOPIC, l.onEvtFunc, newCH)
	return l
}

func InitializedLink(newCH comps.CmdHandlerFtor) IInitializedLink {
	return newLink(newCH)
}
