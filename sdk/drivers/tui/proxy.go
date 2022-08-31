package tui

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/drivers/tui/app_topics"
	"github.com/discomco/go-cart/sdk/dtos"
	"github.com/discomco/go-cart/sdk/features"
	"github.com/discomco/go-cart/sdk/model"
	"github.com/pkg/errors"
	"time"
)

type IProxy interface {
	features.IMsgHandler
	IAmProxy()
	Inject(requesters ...features.IHopeRequester)
	RefreshList(ctx context.Context, key string) error
	RefreshDoc(ctx context.Context, key string) error
	Request(ctx context.Context, hopeType dtos.HopeType, hope dtos.IHope, timeout time.Duration) dtos.IFbk
}

type Proxy[TDoc model.IReadModel, TList model.IReadModel] struct {
	*features.MsgHandler
	requesters map[dtos.HopeType]features.IHopeRequester
	docStore   domain.IReadModelStore[TDoc]
	listStore  domain.IReadModelStore[TList]
	model      IGenModel[TDoc, TList]
}

//RefreshList calls the listStore and
func (p *Proxy[TDoc, TList]) RefreshList(ctx context.Context, key string) error {
	lst, err := p.listStore.Get(ctx, key)
	if err != nil {
		return errors.Wrapf(err, "(%+v) failed to retrieve list", p.GetName())
	}
	p.model.SetList(lst)
	return err
}

func (p *Proxy[TDoc, TList]) RefreshDoc(ctx context.Context, key string) error {
	doc, err := p.docStore.Get(ctx, key)
	if err != nil {
		return errors.Wrapf(err, "(%+v.RefreshDoc) failed to get document with key %+v", p.GetName(), key)
	}
	p.model.SetDoc(doc)
	return err
}

func (p *Proxy[TDoc, TList]) Request(ctx context.Context, hopeType dtos.HopeType, hope dtos.IHope, timeout time.Duration) dtos.IFbk {
	fbk := dtos.NewFbk(hope.GetId(), -1, "")
	r, ok := p.requesters[hopeType]
	if !ok {
		fbk.SetError(fmt.Sprintf("(%+v.GenRequest) could not find a requester for message %+v", p.GetName(), hope))
		return fbk
	}
	fbk = r.Request(ctx, hope, timeout)
	return fbk
}

func (p *Proxy[TDoc, TList]) Inject(requesters ...features.IHopeRequester) {
	for _, requester := range requesters {
		_, ok := p.requesters[requester.GetHopeType()]
		if !ok {
			p.requesters[requester.GetHopeType()] = requester
		}
	}
}

func (p *Proxy[TDoc, TList]) IAmProxy() {}

func newProxy[TDoc model.IReadModel, TList model.IReadModel](
	name features.Name,
	onAppInitialized features.OnMsgFunc,
	newDocStore domain.StoreFtor[TDoc],
	newListStore domain.StoreFtor[TList],
	newModel GenModelFtor[TDoc, TList],
) *Proxy[TDoc, TList] {
	p := &Proxy[TDoc, TList]{
		docStore:   newDocStore(),
		listStore:  newListStore(),
		requesters: make(map[dtos.HopeType]features.IHopeRequester),
		model:      newModel(),
	}
	b := features.NewMsgHandler(app_topics.AppInitialized, onAppInitialized)
	b.Name = name
	p.MsgHandler = b
	return p
}

func NewProxy[TDoc model.IReadModel, TList model.IReadModel](
	name features.Name,
	onAppInitialized features.OnMsgFunc,
	newDocStore domain.StoreFtor[TDoc],
	newListStore domain.StoreFtor[TList],
	newModel GenModelFtor[TDoc, TList],
) *Proxy[TDoc, TList] {
	return newProxy(name,
		onAppInitialized,
		newDocStore,
		newListStore,
		newModel,
	)
}
