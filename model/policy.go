package model

import (
	"generics/events"
	"time"
)

const (
	OrderTopic            = events.EventTopic("order")
	OrderCreatedEventType = events.EventType("order.created")
)

type (
	OrderEventData struct {
		OrderNumber int       `json:"order_number"`
		CreatedAt   time.Time `json:"created_at"`
	}
)

func (c OrderEventData) GetAggregateID() int {
	return c.OrderNumber
}

func (c OrderEventData) GetTopic() events.EventTopic {
	return OrderTopic
}
