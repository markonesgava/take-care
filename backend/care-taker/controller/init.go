package controller

import (
	"github.com/gofiber/fiber"
)

type CareTakerController interface {
	HelloTaker(c *fiber.Ctx)
}

func NewController() CareTakerController {
	return &careTakerController{}
}
