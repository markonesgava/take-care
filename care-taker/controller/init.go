package controller

import (
	"github.com/gofiber/fiber"
	"github.com/markonesgava/take-care/care-taker/commands"
)

type CareTakerController interface {
	HelloTaker(c *fiber.Ctx)
	CreateTaker(c *fiber.Ctx)
}

func NewController(helloCommand *commands.SendTakerHello, createCommand *commands.CreateTaker) CareTakerController {
	return &careTakerController{
		createCommand: createCommand,
		helloCommand:  helloCommand,
	}
}
