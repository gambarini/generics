package main

import (
	"generics/events"
	"generics/model"
	"generics/proc"
	"log"
	"math/rand"
	"time"
)

func main() {

	process := proc.NewProcessManager("main")

	consumerStream := events.NewEventStream[model.CustomerEventData]()
	orderStream := events.NewEventStream[model.OrderEventData]()

	consumerStream.AddConsumer(process.Ctx, func(event events.Event[model.CustomerEventData]) {
		log.Printf("from consumerStream: %+v", event)
	})

	orderStream.AddConsumer(process.Ctx, func(event events.Event[model.OrderEventData]) {
		log.Printf("from orderStream: %+v", event)
	})

	go func() {

		for {

			if process.Ctx.Err() != nil {
				log.Printf("stopping producers...")
				break
			}

			time.Sleep(time.Second / 2)

			e1 := events.NewEvent(model.CustomerEventData{
				CustomerID: rand.Int(),
				CreatedAt:  time.Now(),
			}, model.CustomerCreatedEventType)

			e2 := events.NewEvent(model.OrderEventData{
				OrderNumber: rand.Intn(5000),
				CreatedAt:   time.Now(),
			}, model.OrderCreatedEventType)

			consumerStream.Produce(e1)
			orderStream.Produce(e2)
		}

	}()

	process.Start()

}
