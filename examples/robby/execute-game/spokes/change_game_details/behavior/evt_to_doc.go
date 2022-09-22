package behavior

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/change_game_details/contract"
	"github.com/discomco/go-cart/sdk/behavior"
	sdk_schema "github.com/discomco/go-cart/sdk/schema"
	"github.com/pkg/errors"
)

func EvtToDoc() behavior.Evt2DocFunc[IEvt, schema.GameDoc] {
	return func(evt IEvt, schema *schema.GameDoc) error {
		return evt2Doc(evt, schema)
	}
}

func evt2Doc(evt IEvt, doc *schema.GameDoc) error {
	aggID, err := evt.GetBehaviorID()
	if err != nil {
		return errors.Wrapf(err, "(change_game_details.evt2Doc) failed to get aggregate Id from evt")
	}
	doc.ID = aggID.(*sdk_schema.Identity)
	var pl contract.Payload
	err = evt.GetPayload(&pl)
	if err != nil {
		return errors.Wrapf(err, "(change_game_details.evt2Doc) failed to extract payload from Event")
	}
	doc.Details = pl.Details
	return err
}
