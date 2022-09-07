package convert

import (
	"github.com/EventStore/EventStore-Client-Go/v2/esdb"
	"github.com/discomco/go-cart/sdk/behavior"
)

func Recorded2Evt(recorded *esdb.RecordedEvent) behavior.IEvt {
	return &behavior.Event{
		EventId:   recorded.EventID.String(),
		EventType: behavior.EventType(recorded.EventType),
		//		BehaviorType: domain.BehaviorType(recorded.ContentType),
		Data:        recorded.Data,
		Timestamp:   recorded.CreatedDate,
		AggregateID: recorded.StreamID,
		Version:     int64(recorded.EventNumber),
		Metadata:    recorded.UserMetadata,
	}
}

/*func Resolved2Evt(recorded *esdb.ResolvedEvent) domain.IEvt {
	return &domain.Event{
		EventId:     recorded.Event.EventId.String(),
		EventType:   domain.EventType(recorded.EventType),
		Data:        recorded.Data,
		Timestamp:   recorded.CreatedDate,
		Id: recorded.StreamID,
		Version:     int64(recorded.EventNumber),
		Metadata:    recorded.UserMetadata,
	}
}
*/

func EventData2Evt(eventData esdb.EventData) behavior.IEvt {
	return &behavior.Event{
		EventId:   eventData.EventID.String(),
		EventType: behavior.EventType(eventData.EventType),
		Data:      eventData.Data,
		Metadata:  eventData.Metadata,
	}
}

func Evt2EventData(evt behavior.IEvt) esdb.EventData {
	return esdb.EventData{
		EventType:   string(evt.GetEventType()),
		ContentType: esdb.JsonContentType,
		Data:        evt.GetData(),
		Metadata:    evt.GetMetadata(),
	}
}
