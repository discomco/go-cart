package doc_must

import (
	app_schema "github.com/discomco/go-cart/examples/quadratic-roots/schema"
	app_doc "github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/discomco/go-status"
)

func NotBeInitialized(doc schema.ISchema, fbk contract.IFbk) {
	s := doc.(*app_schema.QuadraticDoc)
	if status.HasStatus(s.Status, app_doc.Initialized) {
		fbk.SetError("Calculation is already initialized")
	}
}
