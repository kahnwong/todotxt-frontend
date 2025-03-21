package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kahnwong/todotxt-frontend/core"
)

func main() {
	// init
	app := fiber.New()

	// render site
	//app.Static("/assets", "./static/assets")
	app.Get("/api/todo/today", core.TodoTodayController)
	app.Get("api/todo/tinkering", core.TodoTinkeringController)

	// start server
	err := app.Listen(os.Getenv("LISTEN_ADDR"))
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
