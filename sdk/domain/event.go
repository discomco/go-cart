package domain

import (
	"encoding/json"
	"fmt"
	"github.com/discomco/go-cart/sdk/core"
	uuid "github.com/satori/go.uuid"
	"time"
)

type EventType string

const (
	AllTopics EventType = "*"
)

type CommandType string

type Event struct {
	EventID       string
	EventType     EventType
	Data          []byte
	Timestamp     time.Time
	AggregateType AggregateType
	AggregateID   string
	Version       int64
	Metadata      []byte
}

func NewEvt(aggregate IAggregate, eventType EventType) IEvt {
	return newEvent(aggregate, eventType)
}

// newEvent new -base Event constructor with configured EventID, IAggregate properties and Timestamp.
func newEvent(aggregate IAggregate, eventType EventType) *Event {
	if aggregate == nil {
		panic(ErrAggregateCannotBeNil)
	}
	id := ""
	if aggregate.GetID() != nil {
		id = aggregate.GetID().Id()
	}
	eID, _ := uuid.NewV4()
	return &Event{
		EventID:       eID.String(),
		AggregateType: aggregate.GetAggregateType(),
		AggregateID:   id,
		Version:       aggregate.GetVersion(),
		EventType:     eventType,
		Timestamp:     time.Now().UTC(),
	}
}

func (e *Event) GetAggregateID() (core.IIdentity, error) {
	return core.IdentityFromPrefixedId(e.AggregateID)
}

func (e *Event) GetEventType() EventType {
	return e.EventType
}

func (e *Event) GetEventID() uuid.UUID {
	return uuid.FromStringOrNil(e.EventID)
}

func (e *Event) GetStreamID() string {
	return e.EventID
}

func (e *Event) EventNumber() uint64 {
	return uint64(e.Version)
}

func (e *Event) CreatedDate() time.Time {
	return e.Timestamp
}

func (e *Event) GetEventId() string {
	return e.EventID
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

// GetJsonData json unmarshal data attached to the Event.
func (e *Event) GetJsonData(data interface{}) error {
	return json.Unmarshal(e.GetData(), data)
}

// SetJsonData serialize to json and set data attached to the Event.
func (e *Event) SetJsonData(data interface{}) error {
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

// GetAggregateType is the AggregateType that the Event can be applied to.
func (e *Event) GetAggregateType() AggregateType {
	return e.AggregateType
}

// SetAggregateType set the AggregateType that the Event can be applied to.
func (e *Event) SetAggregateType(aggregateType AggregateType) {
	e.AggregateType = aggregateType
}

// GeTAID is the Id of the IAggregate that the Event belongs to
func (e *Event) GetAggregateId() string {
	return e.AggregateID
}

// GetVersion is the version of the IAggregate after the Event has been applied.
func (e *Event) GetVersion() int64 {
	return e.Version
}

// SetVersion set the version of the IAggregate.
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

func (e *Event) SetAggregateId(id string) {
	e.AggregateID = id
}

func (e *Event) String() string {
	return fmt.Sprintf("(Event): aggregateID: {%s}, Version: {%d}, EventType: {%s}, AggregateType: {%s}, Metadata: {%s}, TimeStamp: {%s}",
		e.AggregateID,
		e.Version,
		e.EventType,
		e.AggregateType,
		string(e.Metadata),
		e.Timestamp.UTC().String(),
	)
}
