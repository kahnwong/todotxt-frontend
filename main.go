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

func parseTodos(tasks todo.TaskList) []Todo {
	var todos []Todo

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

	return todos
}

func getTodos() ([]Todo, []Todo) {
	var todayTodos []Todo
	var tinkeringTodos []Todo

	if tasklist, err := todo.LoadFromPath(os.Getenv("TODO_PATH")); err != nil {
		fmt.Printf("Error reading todo.txt: %s", err)
	} else {
		// today
		todayTasks := tasklist.Filter(todo.FilterNotCompleted).Filter(todo.FilterDueToday, todo.FilterOverdue)
		_ = todayTasks.Sort(todo.SortPriorityAsc, todo.SortProjectAsc)
		todayTodos = parseTodos(todayTasks)

		// tinkering
		tinkeringTasks := tasklist.Filter(todo.FilterNotCompleted).Filter(todo.FilterByContext("tinkering"))
		_ = tinkeringTasks.Sort(todo.SortPriorityAsc, todo.SortProjectAsc)
		tinkeringTodos = parseTodos(tinkeringTasks)

	}
	return todayTodos, tinkeringTodos
}

func main() {
	// init
	app := fiber.New()

	// render site
	app.Static("/assets", "./static/assets")
	app.Get("/", func(c *fiber.Ctx) error {
		// parse todos
		todayTodos, tinkeringTodos := getTodos()

		// init template
		tmpl, err := template.ParseFiles("./static/index.gohtml")
		if err != nil {
			fmt.Println("Error parsing template:", err)
		}

		// render template
		var tpl bytes.Buffer
		err = tmpl.Execute(&tpl, map[string][]Todo{
			"today":     todayTodos,
			"tinkering": tinkeringTodos,
		})
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
