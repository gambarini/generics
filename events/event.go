package events

import (
	"encoding/json"
	"math/rand"
)

type (
	EventID int

	EventData interface {
		GetAggregateID() int
		GetTopic() EventTopic
	}

	EventTopic string

	EventType string

	Event[ED EventData] struct {
		EventID   EventID   `json:"event_id"`
		EventData ED        `json:"event_data"`
		EventType EventType `json:"event_type"`
	}

	EventJSON[ED EventData] json.RawMessage
)

func NewEvent[ED EventData](eventData ED, eventType EventType) (event Event[ED]) {

	event.EventID = EventID(rand.Int())
	event.EventData = eventData
	event.EventType = eventType

	return event
}

func (e *Event[ED]) ToJSON() (eventJSON EventJSON[ED]) {

	eventJSON, _ = json.Marshal(e)

	return eventJSON
}

func (eJSON EventJSON[ED]) EventFromJSON() (event Event[ED]) {
	_ = json.Unmarshal(eJSON, &event)

	return event
}
