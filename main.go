package main

import (
	"fmt"
	"os"

	todo "github.com/1set/todotxt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func getTodos() []string {
	var todos []string

	if tasklist, err := todo.LoadFromPath(os.Getenv("TODO_PATH")); err != nil {
		fmt.Print("Error reading todo.txt")
	} else {
		tasks := tasklist.Filter(todo.FilterNotCompleted).Filter(todo.FilterDueToday, todo.FilterOverdue)
		_ = tasks.Sort(todo.SortPriorityAsc, todo.SortProjectAsc)
		for _, t := range tasks {
			var context string
			if len(t.Contexts) > 0 {
				context = fmt.Sprintf("@%s", t.Contexts[0])
			}

			var project string
			if len(t.Projects) > 0 {
				project = fmt.Sprintf("+%s", t.Projects[0])
			}

			todoStr := fmt.Sprintf("%s %s %s", context, project, t.Todo)
			todos = append(todos, todoStr)
		}
	}

	return todos
}

func main() {
	// init
	app := fiber.New()

	// render site
	app.Static("/assets", "./static/assets")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("./static/index.html", fiber.Map{
			"Items": getTodos()})
	})

	// start server
	err := app.Listen(os.Getenv("LISTEN_ADDR"))
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
