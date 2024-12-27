package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	todo "github.com/1set/todotxt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

type Todo struct {
	Context string
	Project string
	Todo    string
}

func getTodos() []Todo {
	var todos []Todo

	if tasklist, err := todo.LoadFromPath(os.Getenv("TODO_PATH")); err != nil {
		fmt.Print("Error reading todo.txt")
	} else {
		tasks := tasklist.Filter(todo.FilterNotCompleted).Filter(todo.FilterDueToday, todo.FilterOverdue)
		_ = tasks.Sort(todo.SortPriorityAsc, todo.SortProjectAsc)
		for _, t := range tasks {
			// context
			var context string
			if len(t.Contexts) > 0 {
				context = fmt.Sprintf("@%s", t.Contexts[0])
			}

			// project
			var project string
			if len(t.Projects) > 0 {
				project = fmt.Sprintf("+%s", t.Projects[0])
			}

			// append
			todos = append(todos, Todo{
				context, project,
				t.Todo})
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
		// init template
		tmpl, err := template.ParseFiles("./static/index.gohtml")
		if err != nil {
			fmt.Println("Error parsing template:", err)
		}

		// render template
		var tpl bytes.Buffer
		err = tmpl.Execute(&tpl, getTodos())
		if err != nil {
			fmt.Println("Error rendering html:", err)
		}

		// return response
		c.Response().Header.Add("Content-Type", "text/html; charset=utf-8")
		_, err = c.Write(tpl.Bytes())

		return err
	})

	// start server
	err := app.Listen(os.Getenv("LISTEN_ADDR"))
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
