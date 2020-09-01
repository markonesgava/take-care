package controller

import (
	"github.com/gofiber/fiber"
	"github.com/markonesgava/take-care/domain/commands"
	"github.com/markonesgava/take-care/domain/queries"
)

type CareTakerController interface {
	HelloTaker(c *fiber.Ctx)
	CreateTaker(c *fiber.Ctx)
}

func NewController(commandBusier commands.CommandBusier, queryBusier queries.QueryBusier) CareTakerController {
	return &careTakerController{
		commandBusier: commandBusier,
		queryBusier:   queryBusier,
	}
}
