package behavior

import (
	"github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-cart/sdk/schema"
	"time"
)

type Evt2DocFtor[TEvt IEvt, TDoc schema.ISchema] func() Evt2DocFunc[TEvt, TDoc]

type Evt2DocFunc[TEvt IEvt, TDoc schema.ISchema] func(evt TEvt, model *TDoc) error
type Evt2CmdFunc func(evt IEvt) (ICmd, error)
type Evt2FactFunc func(evt IEvt) (contract.IFact, error)
type GenEvt2FactFunc[TFact contract.IFact] func(evt IEvt) (TFact, error)

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
