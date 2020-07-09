package caretaker

import (
	"fmt"

	"github.com/gofiber/fiber"
	"go.uber.org/dig"
)

type CareTakerController struct {
}

type ICareTakerController interface {
	helloTaker(c *fiber.Ctx)
}

func (CareTakerController) helloTaker(c *fiber.Ctx) {
	msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
	c.Send(msg) // => Hello john 👋!
}

func newController() ICareTakerController {
	return &CareTakerController{}
}

type CareTakerRoutes struct {
}

func newRoute(app *fiber.App, controller ICareTakerController) CareTakerRoutes {
	app.Get("/taker/:name", controller.helloTaker)

	return CareTakerRoutes{}
}

func provideControllers(container *dig.Container) error {
	return container.Provide(newController)
}

func provideRoutes(container *dig.Container) error {
	return container.Provide(newRoute)
}

func ProvideServices(container *dig.Container) error {
	err := provideRoutes(container)

	if err != nil {
		return err
	}

	err = provideControllers(container)
	if err != nil {
		return err
	}

	return container.Invoke(func(_ CareTakerRoutes) {

	})
}
