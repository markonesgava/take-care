package main

import (
	"github.com/gofiber/fiber"
	"github.com/markonesgava/take-care/authorization"
	caretaker "github.com/markonesgava/take-care/care-taker"
	"github.com/markonesgava/take-care/config"
	"github.com/markonesgava/take-care/mongodb"
	"go.uber.org/dig"
)

func buildContainer() (*dig.Container, error) {
	container := dig.New()

	err := container.Provide(fiber.New)
	if err != nil {
		return nil, err
	}

	return container, nil
}

func main() {
	container, err := buildContainer()

	if err != nil {
		panic(err)
	}

	err = config.ProvideServices(container)
	if err != nil {
		panic(err)
	}

	err = mongodb.ProvideServices(container)

	if err != nil {
		panic(err)
	}

	err = caretaker.ProvideServices(container)

	if err != nil {
		panic(err)
	}

	err = authorization.ProvideServices(container)

	if err != nil {
		panic(err)
	}

	container.Invoke(func(app *fiber.App, configuration *config.Configuration) {
		app.Listen(configuration.Port)
	})
}
