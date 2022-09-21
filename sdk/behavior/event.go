package behavior

import (
	"encoding/json"
	"fmt"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/google/uuid"
	"time"
)

type EventType string

const (
	AllTopics EventType = "*"
)

type CommandType string

type Event struct {
	EventId       string
	EventType     EventType
	Data          []byte
	Timestamp     time.Time
	AggregateType BehaviorType
	AggregateID   string
	Version       int64
	Metadata      []byte
}

func NewEvt(aggregate IBehavior, eventType EventType) IEvt {
	return newEvent(aggregate, eventType)
}

// newEvent new -base Event constructor with configured EventId, IBehavior properties and Timestamp.
func newEvent(aggregate IBehavior, eventType EventType) *Event {
	if aggregate == nil {
		panic(ErrBehaviorCannotBeNil)
	}
	id := ""
	if aggregate.GetID() != nil {
		id = aggregate.GetID().Id()
	}
	eID, _ := uuid.NewUUID()
	return &Event{
		EventId:       eID.String(),
		AggregateType: aggregate.GetBehaviorType(),
		AggregateID:   id,
		Version:       aggregate.GetVersion(),
		EventType:     eventType,
		Timestamp:     time.Now().UTC(),
	}
}

func (e *Event) GetBehaviorID() (schema.IIdentity, error) {
	return schema.IdentityFromPrefixedId(e.AggregateID)
}

func (e *Event) GetEventType() EventType {
	return e.EventType
}

func (e *Event) GetStreamId() string {
	return e.EventId
}

func (e *Event) EventNumber() uint64 {
	return uint64(e.Version)
}

func (e *Event) CreatedDate() time.Time {
	return e.Timestamp
}

func (e *Event) GetEventId() string {
	return e.EventId
}

// GetTimeStamp get timestamp of the Event.
func (e *Event) GetTimeStamp() time.Time {
	return e.Timestamp
}

// GetData The data attached to the Event serialized to bytes.
func (e *Event) GetData() []byte {
	return e.Data
}

// SetData add the data attached to the Event serialized to bytes.
func (e *Event) SetData(data []byte) *Event {
	e.Data = data
	return e
}

// GetPayload json unmarshal data attached to the Event.
func (e *Event) GetPayload(data interface{}) error {
	return json.Unmarshal(e.GetData(), data)
}

// SetPayload serialize to json and set data attached to the Event.
func (e *Event) SetPayload(data interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	e.SetData(dataBytes)
	return nil
}

// GetEventType returns the EventType of the event.
func (e *Event) GetEventTypeString() string {
	return fmt.Sprintf("%v", e.EventType)
}

// GetAggregateType is the BehaviorType that the Event can be applied to.
func (e *Event) GetBehaviorType() BehaviorType {
	return e.AggregateType
}

// SetAggregateType set the BehaviorType that the Event can be applied to.
func (e *Event) SetBehaviorType(aggregateType BehaviorType) {
	e.AggregateType = aggregateType
}

// GeTAID is the Id of the IBehavior that the Event belongs to
func (e *Event) GetBehaviorId() string {
	return e.AggregateID
}

// GetVersion is the version of the IBehavior after the Event has been applied.
func (e *Event) GetVersion() int64 {
	return e.Version
}

// SetVersion set the version of the IBehavior.
func (e *Event) SetVersion(aggregateVersion int64) {
	e.Version = aggregateVersion
}

// GetMetadata is domain-specific metadata such as request Id, originating user etc.
func (e *Event) GetMetadata() []byte {
	return e.Metadata
}

// SetMetadata add domain-specific metadata serialized as json for the Event.
func (e *Event) SetMetadata(metaData interface{}) error {
	metaDataBytes, err := json.Marshal(metaData)
	if err != nil {
		return err
	}
	e.Metadata = metaDataBytes
	return nil
}

// GetJsonMetadata unmarshal domain-specific metadata serialized as json for the Event.
func (e *Event) GetJsonMetadata(metaData interface{}) error {
	return json.Unmarshal(e.GetMetadata(), metaData)
}

// GetString A string representation of the Event.
func (e *Event) GetString() string {
	return fmt.Sprintf("event: %+v", e)
}

func (e *Event) SetBehaviorId(id string) {
	e.AggregateID = id
}

func (e *Event) String() string {
	return fmt.Sprintf("(Event): behaviorID: {%s}, Version: {%d}, EventType: {%s}, BehaviorType: {%s}, Metadata: {%s}, TimeStamp: {%s}",
		e.AggregateID,
		e.Version,
		e.EventType,
		e.AggregateType,
		string(e.Metadata),
		e.Timestamp.UTC().String(),
	)
}
