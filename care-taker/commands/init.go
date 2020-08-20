package commands

import (
	"go.uber.org/dig"
)

func ProvideCommands(container *dig.Container) error {
	err := container.Provide(NewCreateTakerCommandHandler)

	if err != nil {
		return err
	}

	err = container.Provide(NewSendTakerHelloCommandHandler)

	return err
}
