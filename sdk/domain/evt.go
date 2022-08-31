package domain

import (
	"github.com/discomco/go-cart/sdk/core"
	"github.com/discomco/go-cart/sdk/dtos"
	"github.com/discomco/go-cart/sdk/model"
	"github.com/satori/go.uuid"
	"time"
)

type Evt2ModelFtor[TEvt IEvt, TReadModel model.IReadModel] func() Evt2ModelFunc[TEvt, TReadModel]

type Evt2ModelFunc[TEvt IEvt, TReadModel model.IReadModel] func(evt TEvt, model *TReadModel) error
type Evt2CmdFunc func(evt IEvt) (ICmd, error)
type Evt2FactFunc func(evt IEvt) (dtos.IFact, error)
type GenEvt2FactFunc[TFact dtos.IFact] func(evt IEvt) (TFact, error)

type IEvt interface {
	EvtTypeGetter
	AggregateTypeGetter
	AggregateTypeSetter
	GetAggregateID() (core.IIdentity, error)
	GetEventID() uuid.UUID
	GetStreamID() string
	EventNumber() uint64
	CreatedDate() time.Time
	GetEventId() string
	GetTimeStamp() time.Time
	GetData() []byte
	SetData(data []byte) *Event
	GetJsonData(data interface{}) error
	SetJsonData(data interface{}) error
	GetEventTypeString() string
	GetAggregateId() string
	GetVersion() int64
	SetVersion(aggregateVersion int64)
	GetMetadata() []byte
	SetMetadata(metaData interface{}) error
	GetJsonMetadata(metaData interface{}) error
	GetString() string
	String() string
	SetAggregateId(id string)
}

type AggregateTypeSetter interface {
	SetAggregateType(aggregateType AggregateType)
}

type AggregateTypeGetter interface {
	GetAggregateType() AggregateType
}
