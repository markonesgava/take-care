package repository

import (
	"github.com/markonesgava/take-care/care-taker/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TakerRepository interface {
	GetOne(id primitive.ObjectID) (*entities.Taker, error)
	GetAll() ([]*entities.Taker, error)
	Add(entity *entities.Taker) error
	Update(id primitive.ObjectID, entity *entities.Taker) error
}
