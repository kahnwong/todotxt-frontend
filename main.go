package main

import (
	"fmt"

	"github.com/kahnwong/todotxt/api"

	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// init
	app := fiber.New()

	// render site
	app.Static("/", "./frontend/dist/spa")
	app.Get("/api/todo/today", api.TodoTodayController)
	app.Get("/api/todo/tinkering", api.TodoTinkeringController)

	// start server
	err := app.Listen(os.Getenv("LISTEN_ADDR"))
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
