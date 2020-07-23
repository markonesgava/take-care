package controller

import (
	"fmt"

	"github.com/gofiber/fiber"
)

type careTakerController struct {
}

func (careTakerController) HelloTaker(c *fiber.Ctx) {
	msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
	c.Send(msg) // => Hello john 👋!
}
