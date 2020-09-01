package controller

import (
	"github.com/gofiber/fiber"

	careTakerCommands "github.com/markonesgava/take-care/care-taker/domain/commands"
	careTakerQueries "github.com/markonesgava/take-care/care-taker/domain/queries"

	"github.com/markonesgava/take-care/domain/commands"
	"github.com/markonesgava/take-care/domain/queries"
)

type careTakerController struct {
	commandBusier commands.CommandBusier
	queryBusier   queries.QueryBusier
}

func (ctrl *careTakerController) HelloTaker(c *fiber.Ctx) {
	query := careTakerQueries.NewGetTakerHello(c.Params("name"))
	result, _ := ctrl.queryBusier.Send(query)

	c.Send(result)
}

func (ctrl *careTakerController) CreateTaker(c *fiber.Ctx) {
	command := careTakerCommands.NewCreateTaker(c.Params("name"))
	ctrl.commandBusier.Send(command)
	c.Send("CRIADO!!!")
}
