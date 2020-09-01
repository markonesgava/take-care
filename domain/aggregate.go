package domain

import (
	"time"

	events "github.com/markonesgava/take-care/domain/events"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Aggregate struct {
	ID            primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Enabled       bool               `bson:"enabled,omitempty"`
	CreatedAt     time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt     time.Time          `bson:"updatedAt,omitempty"`
	DeletedAt     time.Time          `bson:"deletedAt,omitempty"`
	pendingEvents []events.Eventer   `bson:-`
}

func NewAggregate() Aggregate {
	return Aggregate{
		ID:        primitive.NewObjectID(),
		Enabled:   true,
		CreatedAt: time.Now(),
	}
}

// Disable entity
func (ea *Aggregate) Disable() {
	ea.Enabled = false
	ea.UpdatedAt = time.Now()
}

// Enable entity
func (ea *Aggregate) Enable() {
	ea.Enabled = true
	ea.UpdatedAt = time.Now()
}

// Delete entity
func (ea *Aggregate) Delete() {
	ea.DeletedAt = time.Now()
	ea.UpdatedAt = time.Now()
}

//PendingEvents getter
func (ea *Aggregate) PendingEvents() []events.Eventer {
	return ea.pendingEvents
}

//Append a new event
func (ea *Aggregate) Append(event events.Eventer) {
	ea.pendingEvents = append(ea.pendingEvents, event)
}
