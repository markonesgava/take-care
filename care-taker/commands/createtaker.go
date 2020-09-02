package commands

import (
	"fmt"

	takerCommands "github.com/markonesgava/take-care/care-taker/domain/commands"
	"github.com/markonesgava/take-care/care-taker/domain/entities"
	"github.com/markonesgava/take-care/care-taker/domain/repository"
	"github.com/markonesgava/take-care/domain/commands"
)

type createTakerHandler struct {
	repository repository.TakerRepository
}

func newCreateTakeHandler(repository repository.TakerRepository) *createTakerHandler {
	return &createTakerHandler{
		repository,
	}
}

func (*createTakerHandler) CommandName() string {
	return takerCommands.CreateTaker{}.CommandName()
}

func (handler *createTakerHandler) Handle(command commands.Commander) (interface{}, error) {
	cmd := command.(takerCommands.CreateTaker)
	if len(cmd.Name()) == 0 {
		return nil, fmt.Errorf("Taker name is required")
	}

	taker := entities.NewTaker(cmd.Name())

	if err := handler.repository.Add(taker); err != nil {
		return nil, err
	}

	return taker, nil
}
