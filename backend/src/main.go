package main

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	// GET /john
	app.Get("/hello/:name", func(c *fiber.Ctx) {
		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
		c.Send(msg) // => Hello john 👋!
	})

	app.Listen(3900)
}
