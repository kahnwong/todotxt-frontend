package api

import (
	"fmt"

	todo "github.com/1set/todotxt"
)

var todotxtSanitizedPath = "/tmp/todo.txt"

type Todo struct {
	ID      int    `json:"id"`
	Context string `json:"context"`
	Project string `json:"project"`
	Todo    string `json:"todo"`
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
			t.ID,
			context,
			project,
			t.Todo,
		})
	}

	return todos
}

func getTodos(group string) []Todo {
	var todos []Todo

	sanitizeTodo() // strip leading `https://` which results in the todo body returning null
	if tasklist, err := todo.LoadFromPath(todotxtSanitizedPath); err != nil {
		fmt.Printf("Error reading todo.txt: %s", err)
	} else {
		switch group {
		case "today":
			todayTasks := tasklist.Filter(todo.FilterNotCompleted).Filter(todo.FilterDueToday, todo.FilterOverdue)
			_ = todayTasks.Sort(todo.SortPriorityAsc, todo.SortProjectAsc)
			todos = parseTodos(todayTasks)
		case "tinkering":
			tinkeringTasks := tasklist.Filter(todo.FilterNotCompleted).Filter(todo.FilterByContext("tinkering"))
			_ = tinkeringTasks.Sort(todo.SortPriorityAsc, todo.SortProjectAsc)
			todos = parseTodos(tinkeringTasks)
		}
	}
	return todos
}
