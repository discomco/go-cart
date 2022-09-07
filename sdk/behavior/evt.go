package behavior

import (
	"github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-cart/sdk/schema"
	"time"
)

type Evt2ModelFtor[TEvt IEvt, TReadModel schema.IReadModel] func() Evt2ModelFunc[TEvt, TReadModel]

type Evt2ModelFunc[TEvt IEvt, TReadModel schema.IReadModel] func(evt TEvt, model *TReadModel) error
type Evt2CmdFunc func(evt IEvt) (ICmd, error)
type Evt2FactFunc func(evt IEvt) (contract.IFact, error)
type GenEvt2FactFunc[TFact contract.IFact] func(evt IEvt) (TFact, error)

type IEvt interface {
	EvtTypeGetter
	BehaviorTypeGetter
	BehaviorTypeSetter
	GetAggregateID() (schema.IIdentity, error)
	GetStreamID() string
	EventNumber() uint64
	CreatedDate() time.Time
	GetEventId() string
	GetTimeStamp() time.Time
	GetData() []byte
	SetData(data []byte) *Event
	GetPayload(data interface{}) error
	SetPayload(data interface{}) error
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

type BehaviorTypeSetter interface {
	SetBehaviorType(behaviorType BehaviorType)
}

type BehaviorTypeGetter interface {
	GetBehaviorType() BehaviorType
}
