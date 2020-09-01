package queries

import (
	"github.com/markonesgava/take-care/domain/queries"
	"go.uber.org/dig"
)

func ProvideQueries(container *dig.Container) error {
	err := container.Provide(newGetTakerHelloHandler)

	if err != nil {
		return err
	}

	return container.Invoke(func(
		commandBus queries.QueryBusier,
		getTakerHelloHandler *getTakerHelloHandler,
	) {
		commandBus.Attach(getTakerHelloHandler)
	})
}
