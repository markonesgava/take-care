package caretaker

import (
	"github.com/markonesgava/take-care/care-taker/commands"
	"github.com/markonesgava/take-care/care-taker/controller"
	"github.com/markonesgava/take-care/care-taker/routes"
	"go.uber.org/dig"
)

func provideControllers(container *dig.Container) error {
	return container.Provide(controller.NewController)
}

func provideRoutes(container *dig.Container) error {
	return container.Provide(routes.NewRouter)
}

func provideCommands(container *dig.Container) error {
	return commands.ProvideCommands(container)
}

func ProvideServices(container *dig.Container) error {
	err := provideRoutes(container)

	if err != nil {
		return err
	}

	err = provideControllers(container)
	if err != nil {
		return err
	}

	err = provideCommands(container)
	if err != nil {
		return err
	}

	return container.Invoke(func(_ routes.CareTakerRoutes) {

	})
}
