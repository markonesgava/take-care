package controller

import (
	"github.com/gofiber/fiber"
	"github.com/markonesgava/take-care/care-taker/commands"
)

type careTakerController struct {
	helloCommand  *commands.SendTakerHello
	createCommand *commands.CreateTaker
}

func (ctrl *careTakerController) HelloTaker(c *fiber.Ctx) {
	// msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))

	_, msg := ctrl.helloCommand.Handle(c.Params("name"))
	c.Send(msg) // => Hello john 👋!
}

func (ctrl *careTakerController) CreateTaker(c *fiber.Ctx) {
	// msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))

	ctrl.createCommand.Handle(c.Params("name"))
	c.Send("CRIADO!!!") // => Hello john 👋!
}
