package repository

import (
	"go.uber.org/dig"
)

func ProvideRepositories(container *dig.Container) error {
	return container.Provide(newTakerMongoDBRepository)
}
