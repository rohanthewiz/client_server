package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	var app *fiber.App
	app = fiber.New()

	app.Get("/",
		func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})

	app.Get("/:name",
		func(c *fiber.Ctx) error {
			//println(c.Params("name"))
			return c.SendString("Hello, " + c.Params("name"))
		})

	err := app.Listen(":3000")
	log.Println(err)
}
