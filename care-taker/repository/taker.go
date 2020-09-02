package repository

import (
	"context"
	"time"

	"github.com/markonesgava/take-care/care-taker/domain/entities"
	"github.com/markonesgava/take-care/care-taker/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

type takerMongoDBRepository struct {
	connection *mongo.Client
	collection *mongo.Collection
}

func newTakerMongoDBRepository(connection *mongo.Client) repository.TakerRepository {
	return &takerMongoDBRepository{
		connection: connection,
		collection: connection.Database("care-taker").Collection("takers"),
	}
}

func (repo *takerMongoDBRepository) Add(taker *entities.Taker) error {
	ctx, _ := context.WithTimeout(context.TODO(), time.Second)
	result, err := repo.collection.InsertOne(ctx, taker)

	if err != nil {
		return err
	}

	taker.ID = result.InsertedID.(primitive.ObjectID)

	return nil
}

func (repo *takerMongoDBRepository) Update(id primitive.ObjectID, taker *entities.Taker) error {
	return nil
}

func (repo *takerMongoDBRepository) GetOne(id primitive.ObjectID) (*entities.Taker, error) {
	ctx, _ := context.WithTimeout(context.TODO(), time.Second)

	var taker *entities.Taker
	err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&taker)

	return taker, err
}

func (repo *takerMongoDBRepository) GetAll() ([]*entities.Taker, error) {
	ctx, _ := context.WithTimeout(context.TODO(), time.Second)

	cursor, err := repo.collection.Find(ctx, bson.M{"deletedAt": nil})
	if err != nil {
		return nil, err
	}

	var takers []*entities.Taker
	if err = cursor.All(ctx, &takers); err != nil {
		return nil, err
	}

	return takers, nil

}
