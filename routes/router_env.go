package routes

import "github.com/a-frank/todo-web/todos"

type RouterEnv struct {
	Store        todos.TodoStore
	TemplatePath string
}
