package doc_must

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	app_doc "github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"github.com/discomco/go-cart/sdk/contract"
	sdk_schema "github.com/discomco/go-cart/sdk/schema"
	"github.com/discomco/go-status"
)

func BeInitialized(doc sdk_schema.ISchema, fbk contract.IFbk) {
	s := doc.(*schema.QuadraticDoc)
	if status.NotHasStatus(s.Status, app_doc.Initialized) {
		fbk.SetError("Calculation is not initialized")
	}
}
