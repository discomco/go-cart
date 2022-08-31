package features

import (
	"context"
	"github.com/discomco/go-cart/sdk/dtos"
)

type IQueryProvider interface {
	IFeaturePlugin
	IAmQueryProvider()
	RunQuery(ctx context.Context, qry dtos.IReq) dtos.IRsp
}

type IQryFeature interface {
	IFeature
}
