package events

import (
	"time"

	"github.com/google/uuid"
)

type Eventer interface {
	Id() uuid.UUID
	Name() string
	Timestamp() time.Time
}
