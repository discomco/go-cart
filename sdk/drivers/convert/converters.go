package convert

import (
	"github.com/EventStore/EventStore-Client-Go/v2/esdb"
	"github.com/discomco/go-cart/domain"
)

func Recorded2Evt(recorded *esdb.RecordedEvent) domain.IEvt {
	return &domain.Event{
		EventID:   recorded.EventID.String(),
		EventType: domain.EventType(recorded.EventType),
		//		AggregateType: domain.AggregateType(recorded.ContentType),
		Data:        recorded.Data,
		Timestamp:   recorded.CreatedDate,
		AggregateID: recorded.StreamID,
		Version:     int64(recorded.EventNumber),
		Metadata:    recorded.UserMetadata,
	}
}

/*func Resolved2Evt(recorded *esdb.ResolvedEvent) domain.IEvt {
	return &domain.Event{
		EventID:     recorded.Event.EventID.String(),
		EventType:   domain.EventType(recorded.EventType),
		Data:        recorded.Data,
		Timestamp:   recorded.CreatedDate,
		Id: recorded.StreamID,
		Version:     int64(recorded.EventNumber),
		Metadata:    recorded.UserMetadata,
	}
}
*/

func EventData2Evt(eventData esdb.EventData) domain.IEvt {
	return &domain.Event{
		EventID:   eventData.EventID.String(),
		EventType: domain.EventType(eventData.EventType),
		Data:      eventData.Data,
		Metadata:  eventData.Metadata,
	}
}

func Evt2EventData(evt domain.IEvt) esdb.EventData {
	return esdb.EventData{
		EventType:   string(evt.GetEventType()),
		ContentType: esdb.JsonContentType,
		Data:        evt.GetData(),
		Metadata:    evt.GetMetadata(),
	}
}
