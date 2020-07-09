// package caretaker

// import (
// 	"fmt"

// 	"github.com/gofiber/fiber"
// )

// func InitRoute(app *fiber.App) {
// 	app.Get("/taker/:name", func(c *fiber.Ctx) {
// 		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
// 		c.Send(msg) // => Hello john 👋!
// 	})
// }
