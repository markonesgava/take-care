package controller

import (
	"github.com/gofiber/fiber"

	careTakerCommands "github.com/markonesgava/take-care/care-taker/domain/commands"
	careTakerQueries "github.com/markonesgava/take-care/care-taker/domain/queries"

	"github.com/markonesgava/take-care/domain/commands"
	"github.com/markonesgava/take-care/domain/queries"
)

type careTakerController struct {
	commandBuser commands.CommandBuser
	queryBusier  queries.QueryBusier
}

func (ctrl *careTakerController) HelloTaker(c *fiber.Ctx) {
	query := careTakerQueries.NewGetTakerHello(c.Params("name"))
	result, _ := ctrl.queryBusier.Send(query)

	c.Send(result)
}

func (ctrl *careTakerController) CreateTaker(c *fiber.Ctx) {
	command := careTakerCommands.NewCreateTaker(c.Params("name"))
	taker, _ := ctrl.commandBuser.Send(command)
	c.JSON(taker)
}
