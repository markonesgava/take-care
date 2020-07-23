package routes

import (
	"github.com/gofiber/fiber"
	"github.com/markonesgava/take-care/care-taker/controller"
)

type CareTakerRoutes struct {
}

func NewRouter(app *fiber.App, ctrl controller.CareTakerController) CareTakerRoutes {
	app.Get("/taker/:name", ctrl.HelloTaker)

	return CareTakerRoutes{}
}
