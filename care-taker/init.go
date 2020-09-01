package caretaker

import (
	"github.com/markonesgava/take-care/care-taker/commands"
	"github.com/markonesgava/take-care/care-taker/controller"
	"github.com/markonesgava/take-care/care-taker/queries"
	"github.com/markonesgava/take-care/care-taker/repository"
	"github.com/markonesgava/take-care/care-taker/routes"
	"go.uber.org/dig"
)

func provideControllers(container *dig.Container) error {
	return container.Provide(controller.NewController)
}

func provideRoutes(container *dig.Container) error {
	return container.Provide(routes.NewRouter)
}

func ProvideServices(container *dig.Container) error {
	err := repository.ProvideRepositories(container)
	if err != nil {
		return err
	}

	err = provideRoutes(container)
	if err != nil {
		return err
	}

	err = provideControllers(container)
	if err != nil {
		return err
	}

	err = commands.ProvideCommands(container)
	if err != nil {
		return err
	}

	err = queries.ProvideQueries(container)
	if err != nil {
		return err
	}

	return container.Invoke(func(_ routes.CareTakerRoutes) {

	})
}
