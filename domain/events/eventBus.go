package events

import (
	"sync"
)

//EventBus is a implementation of EventBusier
type EventBus struct {
	EventBusier
	mutex         sync.Mutex
	eventHandlers map[string][]EventHadler
}

//NewEventBus create a new event bus
func NewEventBus() EventBusier {
	return &EventBus{
		eventHandlers: make(map[string][]EventHadler),
	}
}

//Attach includes an handle to event
func (ev *EventBus) Attach(handler EventHadler) error {
	ev.mutex.Lock()
	defer ev.mutex.Unlock()

	for _, eventName := range handler.EventNames() {
		ev.eventHandlers[eventName] = append(ev.eventHandlers[eventName], handler)
	}

	return nil
}

//Publish emit the events
func (ev EventBus) Publish(events ...Eventer) error {
	for _, event := range events {
		if handlers, ok := ev.eventHandlers[event.Name()]; ok {
			for _, handler := range handlers {
				err := handler.Handle(event)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
