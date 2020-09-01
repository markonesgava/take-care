package entities

import (
	"github.com/markonesgava/take-care/domain"
)

type Taker struct {
	domain.Aggregate `bson:",inline"`
	Name             string `json:"name,omitempty"`
}

func NewTaker(name string) *Taker {
	return &Taker{
		Name:      name,
		Aggregate: domain.NewAggregate(),
	}
}
