package tui

import (
	"fmt"
	"github.com/discomco/go-cart/sdk/features"
	"github.com/discomco/go-cart/sdk/model"
)

const (
	DocChangedFmt  = "%+v.docChanged"
	ListChangedFmt = "%+v.listChanged"
	//ViewStateChanged = "%+v.viewStateChanged"
)

type ModelFtor func() IModel
type GenModelFtor[TDoc model.IReadModel, TList model.IReadModel] func() IGenModel[TDoc, TList]

type IModel interface {
	features.IComponent
	IAmModel()
}

type IGenModel[TDoc model.IReadModel, TList model.IReadModel] interface {
	IModel
	GetDocChangedTopic() string
	GetListChangedTopic() string
	GetDoc() *TDoc
	GetList() *TList
	SetDoc(doc *TDoc)
	SetList(list *TList)
}

type innerModel[TDoc model.IReadModel, TList model.IReadModel] struct {
	*features.AppComponent
	doc  *TDoc
	list *TList
}

func (m *innerModel[TDoc, TList]) IAmModel() {}

func (m *innerModel[TDoc, TList]) GetDoc() *TDoc {
	return m.doc
}

func (m *innerModel[TDoc, TList]) GetList() *TList {
	return m.list
}

func (m *innerModel[TDoc, TList]) GetDocChangedTopic() string {
	return fmt.Sprintf(DocChangedFmt, m.GetName())
}

func (m *innerModel[TDoc, TList]) GetListChangedTopic() string {
	return fmt.Sprintf(ListChangedFmt, m.GetName())
}

func (m *innerModel[TDoc, TList]) SetDoc(doc *TDoc) {
	if doc != nil {
		m.doc = doc
	}
	topic := fmt.Sprintf(DocChangedFmt, m.GetName())
	m.GetMediator().Broadcast(topic, nil, m.doc)
}

func (m *innerModel[TDoc, TList]) SetList(list *TList) {
	if list != nil {
		m.list = list
	}
	m.GetMediator().Broadcast(m.GetListChangedTopic(), nil, m.list)
}

func newModel[TDoc model.IReadModel, TList model.IReadModel](
	name features.Name,
	docFtor features.DocFtor[TDoc],
	listFtor features.DocFtor[TList],
) *innerModel[TDoc, TList] {
	m := &innerModel[TDoc, TList]{
		doc:  docFtor(),
		list: listFtor(),
	}
	b := features.NewAppComponent(name)
	m.AppComponent = b
	return m
}

func NewModel[TDoc model.IReadModel, TList model.IReadModel](
	name features.Name,
	docFtor features.DocFtor[TDoc],
	listFtor features.DocFtor[TList],
) IGenModel[TDoc, TList] {
	return newModel[TDoc, TList](name, docFtor, listFtor)
}
