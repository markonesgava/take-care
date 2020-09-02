package commands

import (
	"github.com/markonesgava/take-care/domain/commands"
	"go.uber.org/dig"
)

func ProvideCommands(container *dig.Container) error {
	err := container.Provide(newCreateTakeHandler)

	if err != nil {
		return err
	}

	return container.Invoke(func(
		commandBus commands.CommandBuser,
		createTakerHandler *createTakerHandler,
	) {
		commandBus.Attach(createTakerHandler)
	})
}
