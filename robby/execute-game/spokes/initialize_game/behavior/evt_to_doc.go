package behavior

import (
	"github.com/discomco/go-cart/robby/execute-game/schema"
	doc2 "github.com/discomco/go-cart/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/core"
	"github.com/discomco/go-cart/sdk/core/utils/status"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/pkg/errors"
)

func EvtToDoc() domain.Evt2ModelFunc[IEvt, schema.GameDoc] {
	return func(evt IEvt, schema *schema.GameDoc) error {
		return evt2Doc(evt, schema)
	}
}

func evt2Doc(evt IEvt, doc *schema.GameDoc) error {
	aggID, err := evt.GetAggregateID()
	if err != nil {
		return errors.Wrapf(err, "(initialize_game.evt2Doc) failed to get aggregate ID from evt")
	}
	doc.ID = aggID.(*core.Identity)
	var pl contract.Payload
	err = evt.GetJsonData(&pl)
	if err != nil {
		return errors.Wrapf(err, "(initialize_game.evt2Doc) failed to extract payload from Event")
	}
	doc.Details = pl.Details
	status.SetFlag(&doc.Status, doc2.Initialized)
	return err
}
