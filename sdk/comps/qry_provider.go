package comps

import (
	"context"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-cart/sdk/schema"
	"golang.org/x/sync/errgroup"
)

type QryWorker[TReadModel schema.ISchema] func(ctx context.Context, store behavior.IReadModelStore[TReadModel], qry contract.IReq) contract.IRsp

type QryProvider[TReadModel schema.ISchema] struct {
	*Component
	newStore     behavior.StoreFtor[TReadModel]
	newQryWorker QryWorker[TReadModel]
}

func (p *QryProvider[TReadModel]) IAmQryProvider() {}

func (p *QryProvider[TReadModel]) RunQuery(ctx context.Context, qry contract.IReq) contract.IRsp {
	var result contract.IRsp
	result, err := contract.NewRsp(qry.GetId(), nil)
	if err != nil {
		return nil
	}
	resp := make(chan contract.IRsp)
	gr, ctx := errgroup.WithContext(ctx)
	gr.Go(p.runWorker(ctx, qry, resp))
	if err != nil {
		return nil
	}
	go func(r chan contract.IRsp) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case result = <-r:
		}
		return nil
	}(resp)
	err = gr.Wait()
	if err != nil {
		result.SetError(err.Error())
	}
	return result
}

func (p *QryProvider[TReadModel]) runWorker(ctx context.Context, qry contract.IReq, resp chan contract.IRsp) func() error {
	return func() error {
		store := p.newStore()
		rsp := p.newQryWorker(ctx, store, qry)
		resp <- rsp
		return nil
	}
}

func NewQryProvider[TReadModel schema.ISchema](
	name schema.Name,
	storeFtor behavior.StoreFtor[TReadModel],
	qryWorker QryWorker[TReadModel],
) *QryProvider[TReadModel] {
	b := NewComponent(name)
	p := &QryProvider[TReadModel]{
		newStore:     storeFtor,
		newQryWorker: qryWorker,
	}
	p.Component = b
	return p
}
