package comps

import (
	"context"
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/change_game_settings/behavior"
	contract2 "github.com/discomco/go-cart/examples/robby/execute-game/spokes/change_game_settings/contract"
	behavior2 "github.com/discomco/go-cart/examples/robby/execute-game/spokes/initialize_game/behavior"
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/initialize_game/contract"
	sdk_behavior "github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/pkg/errors"
	"sync"
)

const LinkName = "change_game_settings.InitializedPolicy"

type IInitializedLink interface {
	comps.IPolicy
}

type link struct {
	*comps.Policy
}

var wMutex = &sync.Mutex{} //

func (l *link) onEvtFunc(ctx context.Context, evt sdk_behavior.IEvt) error {
	wMutex.Lock()
	defer wMutex.Unlock()
	docID, err := evt.GetBehaviorID()
	if err != nil {
		return errors.Wrapf(err, "failed to get aggregate Id from event %v", evt)
	}
	var pl contract.Payload
	err = evt.GetPayload(&pl)
	settings := schema.NewSettings(pl.MapSize, pl.NbrOfPlayers)
	settingsPl := contract2.NewPayload(settings)
	cmd, err := behavior.NewCmd(docID, *settingsPl)
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
	l.Policy = comps.NewPolicy(
		LinkName, behavior2.EvtTopic, l.onEvtFunc, newCH)
	return l
}

func InitializedLink(newCH comps.CmdHandlerFtor) IInitializedLink {
	return newLink(newCH)
}
