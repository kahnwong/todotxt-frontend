package main

import (
	"fmt"

	todo "github.com/1set/todotxt"
	"github.com/gofiber/fiber/v2"
)

func getTodos() []string {
	//return []string{"a", "b", "c", "d"}

	var todos []string

	if tasklist, err := todo.LoadFromPath("/opt/syncthing/cloud/Apps/todotxt/todo.txt"); err != nil {
		fmt.Print("Error reading todo.txt")
	} else {
		tasks := tasklist.Filter(todo.FilterNotCompleted).Filter(todo.FilterDueToday, todo.FilterOverdue)
		_ = tasks.Sort(todo.SortPriorityAsc, todo.SortProjectAsc)
		for _, t := range tasks {
			todos = append(todos, t.Todo)
		}
	}

	return todos
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
