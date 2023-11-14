package routes

import "github.com/a-frank/web-dev/todos"

type RouterEnv struct {
	Store        todos.TodoStore
	TemplatePath string
}
