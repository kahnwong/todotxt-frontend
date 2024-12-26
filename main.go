package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./spa")

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
