package behavior

import (
	"github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-cart/sdk/schema"
	"time"
)

type Evt2SchemaFtor[TEvt IEvt, TSchema schema.ISchema] func() FEvt2Schema[TEvt, TSchema]

type FEvt2Schema[TEvt IEvt, TSchema schema.ISchema] func(evt TEvt, model *TSchema) error
type FEvt2Cmd func(evt IEvt) (ICmd, error)
type FEvt2Fact func(evt IEvt) (contract.IFact, error)
type FGenEvt2Fact[TFact contract.IFact] func(evt IEvt) (TFact, error)

type IEvt interface {
	IGetEvtType
	IGetBehaviorType
	ISetBehaviorType
	GetBehaviorID() (schema.IIdentity, error)
	GetStreamId() string
	EventNumber() uint64
	CreatedDate() time.Time
	GetEventId() string
	GetTimeStamp() time.Time
	GetData() []byte
	SetData(data []byte) *Event
	GetPayload(data interface{}) error
	SetPayload(data interface{}) error
	GetEventTypeString() string
	GetBehaviorId() string
	GetVersion() int64
	SetVersion(version int64)
	GetMetadata() []byte
	SetMetadata(metaData interface{}) error
	GetJsonMetadata(metaData interface{}) error
	GetString() string
	String() string
	SetBehaviorId(id string)
}

type ISetBehaviorType interface {
	SetBehaviorType(behaviorType BehaviorType)
}

type IGetBehaviorType interface {
	GetBehaviorType() BehaviorType
}
