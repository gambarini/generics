package model

import (
	"generics/events"
	"time"
)

const (
	CustomerTopic            = events.EventTopic("customer")
	CustomerCreatedEventType = events.EventType("customer.created")
)

type (
	CustomerEventData struct {
		CustomerID int       `json:"customer_id"`
		CreatedAt  time.Time `json:"created_at"`
	}
)

func (c CustomerEventData) GetAggregateID() int {
	return c.CustomerID
}

func (c CustomerEventData) GetTopic() events.EventTopic {
	return CustomerTopic
}
