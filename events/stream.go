package events

import (
	"context"
)

type (
	EventStream[ED EventData] struct {
		Events chan EventJSON[ED]
	}

	ConsumerHandler[ED EventData] func(event Event[ED])
)

func NewEventStream[ED EventData]() (es *EventStream[ED]) {

	es = &EventStream[ED]{
		Events: make(chan EventJSON[ED]),
	}
	return es
}

func (es EventStream[ED]) AddConsumer(ctx context.Context, handler ConsumerHandler[ED]) {

	go func() {
		for {
			select {
			case eJSON := <-es.Events:
				handler(eJSON.EventFromJSON())
			case <-ctx.Done():
				break
			}
		}

	}()
}

func (es EventStream[ED]) Produce(event Event[ED]) {

	es.Events <- event.ToJSON()

}
