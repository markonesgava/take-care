package domain

import (
	events "github.com/markonesgava/take-care/domain/events"
)

type Aggregator interface {
	Entitier
	PendingEvents() []events.Eventer
}
