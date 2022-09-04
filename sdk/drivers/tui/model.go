package tui

import (
	"fmt"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/schema"
)

const (
	DocChangedFmt  = "%+v.docChanged"
	ListChangedFmt = "%+v.listChanged"
	//ViewStateChanged = "%+v.viewStateChanged"
)

type ModelFtor func() IModel
type GenModelFtor[TDoc schema.IReadModel, TList schema.IReadModel] func() IGenModel[TDoc, TList]

type IModel interface {
	comps.IComponent
	IAmModel()
}

type IGenModel[TDoc schema.IReadModel, TList schema.IReadModel] interface {
	IModel
	GetDocChangedTopic() string
	GetListChangedTopic() string
	GetDoc() *TDoc
	GetList() *TList
	SetDoc(doc *TDoc)
	SetList(list *TList)
}

type innerModel[TDoc schema.IReadModel, TList schema.IReadModel] struct {
	*comps.Component
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

func newModel[TDoc schema.IReadModel, TList schema.IReadModel](
	name schema.Name,
	docFtor schema.DocFtor[TDoc],
	listFtor schema.DocFtor[TList],
) *innerModel[TDoc, TList] {
	m := &innerModel[TDoc, TList]{
		doc:  docFtor(),
		list: listFtor(),
	}
	b := comps.NewComponent(name)
	m.Component = b
	return m
}

func NewModel[TDoc schema.IReadModel, TList schema.IReadModel](
	name schema.Name,
	docFtor schema.DocFtor[TDoc],
	listFtor schema.DocFtor[TList],
) IGenModel[TDoc, TList] {
	return newModel[TDoc, TList](name, docFtor, listFtor)
}
