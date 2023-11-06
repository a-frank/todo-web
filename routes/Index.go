package routes

import (
	"github.com/a-frank/web-dev/todos"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func GetIndex(context *gin.Context) {
	t, err := template.ParseFiles("./templates/index.html", "./templates/todo.tmpl")
	if check(err, context) != nil {
		return
	}

	allTodos := todos.GetTodos()

	err = t.Execute(context.Writer, indexData{
		Name:  "Go Gin",
		Todos: allTodos,
	})
	if check(err, context) != nil {
		return
	}
	context.Status(http.StatusOK)
}

func AddTodo(context *gin.Context) {
	t, err := template.ParseFiles("./templates/todo.tmpl")
	if check(err, context) != nil {
		return
	}
	err = context.Request.ParseForm()
	if check(err, context) != nil {
		return
	}
	todoText := context.Request.Form.Get("newTodo")
	if todoText == "" {
		context.Status(http.StatusBadRequest)
		return
	}

	newTodo := todos.AddNewTodo(todoText)

	err = t.ExecuteTemplate(context.Writer, "todo", newTodo)
	if check(err, context) != nil {
		return
	}
	context.Status(http.StatusOK)
}

type indexData struct {
	Name  string
	Todos []todos.Todo
}
