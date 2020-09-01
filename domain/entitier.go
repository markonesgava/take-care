package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Entitier is the basic interface for an entity
type Entitier interface {
	Id() primitive.ObjectID
	Enabled() bool
	CreatedAt() time.Time
	UpdatedAt() time.Time
	DeletedAt() time.Time
}
