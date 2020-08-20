package mongodb

import (
	"context"
	"log"

	"github.com/markonesgava/take-care/config"
	"go.uber.org/dig"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newConnection(configuration *config.Configuration) *mongo.Client {
	clientOptions := options.Client().ApplyURI(configuration.MongoDB.ConnectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func ProvideServices(container *dig.Container) error {
	return container.Provide(newConnection)
}
