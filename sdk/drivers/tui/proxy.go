package tui

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-cart/sdk/drivers/tui/app_topics"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/pkg/errors"
	"time"
)

type IProxy interface {
	comps.IMsgReactor
	IAmProxy()
	Inject(requesters ...comps.IRequester)
	RefreshList(ctx context.Context, key string) error
	RefreshDoc(ctx context.Context, key string) error
	Request(ctx context.Context, hopeType contract.HopeType, hope contract.IHope, timeout time.Duration) contract.IFbk
}

type Proxy[TDoc schema.IReadModel, TList schema.IReadModel] struct {
	*comps.MsgReactor
	requesters map[contract.HopeType]comps.IRequester
	docStore   behavior.IReadModelStore[TDoc]
	listStore  behavior.IReadModelStore[TList]
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

func (p *Proxy[TDoc, TList]) Request(ctx context.Context, hopeType contract.HopeType, hope contract.IHope, timeout time.Duration) contract.IFbk {
	fbk := contract.NewFbk(hope.GetId(), -1, "")
	r, ok := p.requesters[hopeType]
	if !ok {
		fbk.SetError(fmt.Sprintf("(%+v.GenRequest) could not find a requester for message %+v", p.GetName(), hope))
		return fbk
	}
	fbk = r.Request(ctx, hope, timeout)
	return fbk
}

func (p *Proxy[TDoc, TList]) Inject(requesters ...comps.IRequester) {
	for _, requester := range requesters {
		_, ok := p.requesters[requester.GetHopeType()]
		if !ok {
			p.requesters[requester.GetHopeType()] = requester
		}
	}
}

func (p *Proxy[TDoc, TList]) IAmProxy() {}

func newProxy[TDoc schema.IReadModel, TList schema.IReadModel](
	name schema.Name,
	onAppInitialized comps.OnMsgFunc,
	newDocStore behavior.StoreFtor[TDoc],
	newListStore behavior.StoreFtor[TList],
	newModel GenModelFtor[TDoc, TList],
) *Proxy[TDoc, TList] {
	p := &Proxy[TDoc, TList]{
		docStore:   newDocStore(),
		listStore:  newListStore(),
		requesters: make(map[contract.HopeType]comps.IRequester),
		model:      newModel(),
	}
	b := comps.NewMsgReactor(app_topics.AppInitialized, onAppInitialized)
	b.Name = name
	p.MsgReactor = b
	return p
}

func NewProxy[TDoc schema.IReadModel, TList schema.IReadModel](
	name schema.Name,
	onAppInitialized comps.OnMsgFunc,
	newDocStore behavior.StoreFtor[TDoc],
	newListStore behavior.StoreFtor[TList],
	newModel GenModelFtor[TDoc, TList],
) *Proxy[TDoc, TList] {
	return newProxy(name,
		onAppInitialized,
		newDocStore,
		newListStore,
		newModel,
	)
}
