package main

import (
	"github.com/gofiber/fiber"
	caretaker "github.com/markonesgava/take-care/care-taker"
	"go.uber.org/dig"
)

func buildContainer() (*dig.Container, error) {
	container := dig.New()

	err := container.Provide(fiber.New)
	if err != nil {
		return nil, err
	}

	return container, nil
}

func main() {
	container, err := buildContainer()

	if err != nil {
		panic(err)
	}

	err = caretaker.ProvideServices(container)

	if err != nil {
		panic(err)
	}
	// app := fiber.New()

	// // GET /john
	// app.Get("/hello/:name", func(c *fiber.Ctx) {
	// 	msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
	// 	c.Send(msg) // => Hello john 👋!
	// })

	// app.Listen(3900)

	container.Invoke(func(app *fiber.App) {
		app.Listen(3900)
	})
}
