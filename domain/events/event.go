package events

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Eventer
	id        uuid.UUID
	name      string
	timestamp time.Time
}

func (event *Event) Name() string {
	return event.name
}

func (event *Event) Id() uuid.UUID {
	return event.id
}

func (event *Event) Timestamp() time.Time {
	return event.timestamp
}

func NewEvent(name string) Event {
	id, _ := uuid.NewUUID()
	return Event{
		id:        id,
		name:      name,
		timestamp: time.Now(),
	}
}
