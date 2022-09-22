package behavior

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
	doc2 "github.com/discomco/go-cart/examples/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/behavior"
	sdk_schema "github.com/discomco/go-cart/sdk/schema"
	status "github.com/discomco/go-status"
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
		return errors.Wrapf(err, "(initialize_game.evt2Doc) failed to get aggregate Id from evt")
	}
	doc.ID = aggID.(*sdk_schema.Identity)
	var pl contract.Payload
	err = evt.GetPayload(&pl)
	if err != nil {
		return errors.Wrapf(err, "(initialize_game.evt2Doc) failed to extract payload from Event")
	}
	doc.Details = pl.Details
	status.SetStatus(&doc.Status, doc2.Initialized)
	return err
}
