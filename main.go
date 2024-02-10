package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rohanthewiz/element"
	"log"
)

func main() {
	var app *fiber.App
	app = fiber.New()

	app.Get("/",
		func(c *fiber.Ctx) error {
			c.Set("Content-Type", "text/html")
			return c.SendString(showForm())
		})

	/*	app.Get("/:form_name",
		func(c *fiber.Ctx) error {
			//println(c.Params("name"))
			return c.SendString("Hello via form, " + c.Params("name"))
		})
	*/
	err := app.Listen(":3000")
	log.Println(err)
}

func showForm() string {
	b := element.NewBuilder()

	e := b.Ele
	t := b.Text

	e("html").R(
		e("body").R(
			e("p", "style", "color:maroon;font-weight:bold;font-size:2rem").R(
				t("this is the body")),
			e("p").R(
				t("this is another line")),
		),
	)
	return b.String()
}
