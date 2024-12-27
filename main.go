package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func getTodos() []string {
	return []string{"a", "b", "c", "d"}
}

func main() {
	// init
	app := fiber.New()

	// render site
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("./static/index.html", fiber.Map{
			"Items": getTodos()})
	})

	// start server
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
