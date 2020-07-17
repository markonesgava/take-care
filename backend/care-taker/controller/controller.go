package controller

import (
	"fmt"

	"github.com/gofiber/fiber"
)

type careTakerController struct {
}

type CareTakerController interface {
	HelloTaker(c *fiber.Ctx)
}

func (careTakerController) HelloTaker(c *fiber.Ctx) {
	msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
	c.Send(msg) // => Hello john 👋!
}

func NewController() CareTakerController {
	return &careTakerController{}
}
