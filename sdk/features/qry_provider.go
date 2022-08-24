package features

import (
	"context"

	"github.com/discomco/go-cart/domain"
	"github.com/discomco/go-cart/dtos"
	"github.com/discomco/go-cart/model"
	"golang.org/x/sync/errgroup"
)

type QryWorker[TReadModel model.IReadModel] func(ctx context.Context, store domain.IReadModelStore[TReadModel], qry dtos.IReq) dtos.IRsp

type QryProvider[TReadModel model.IReadModel] struct {
	*AppComponent
	newStore     domain.StoreFtor[TReadModel]
	newQryWorker QryWorker[TReadModel]
}

func (p *QryProvider[TReadModel]) IAmQryProvider() {}

func (p *QryProvider[TReadModel]) RunQuery(ctx context.Context, qry dtos.IReq) dtos.IRsp {
	var result dtos.IRsp
	result, err := dtos.NewRsp(qry.GetId(), nil)
	if err != nil {
		return nil
	}
	resp := make(chan dtos.IRsp)
	gr, ctx := errgroup.WithContext(ctx)
	gr.Go(p.runWorker(ctx, qry, resp))
	if err != nil {
		return nil
	}
	go func(r chan dtos.IRsp) error {
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

func (p *QryProvider[TReadModel]) runWorker(ctx context.Context, qry dtos.IReq, resp chan dtos.IRsp) func() error {
	return func() error {
		store := p.newStore()
		rsp := p.newQryWorker(ctx, store, qry)
		resp <- rsp
		return nil
	}
}

func NewQryProvider[TReadModel model.IReadModel](
	name Name,
	storeFtor domain.StoreFtor[TReadModel],
	qryWorker QryWorker[TReadModel],
) *QryProvider[TReadModel] {
	b := NewAppComponent(name)
	p := &QryProvider[TReadModel]{
		newStore:     storeFtor,
		newQryWorker: qryWorker,
	}
	p.AppComponent = b
	return p
}
