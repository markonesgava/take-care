package commands

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type CreateTaker struct {
	connection *mongo.Client
}

type Taker struct {
	Name string
}

func NewCreateTakerCommandHandler(connection *mongo.Client) *CreateTaker {
	return &CreateTaker{
		connection: connection,
	}
}

func (c *CreateTaker) Handle(takerName string) error {
	if len(takerName) == 0 {
		return fmt.Errorf("TakerName is required")
	}

	taker := Taker{
		Name: takerName,
	}
	collection := c.connection.Database("care-taker").Collection("takers")

	ctx, _ := context.WithTimeout(context.TODO(), time.Second)
	result, err := collection.InsertOne(ctx, taker)

	fmt.Printf("result >>> %v", result)

	return err
}
