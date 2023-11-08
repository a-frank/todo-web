package routes

import (
	"github.com/a-frank/web-dev/todos"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
)

func GetIndex(context *gin.Context) {
	t, err := template.ParseFiles("./templates/index.html", "./templates/todo.tmpl")
	if check(err, context) != nil {
		return
	}

	allTodos, err := todos.GetTodos()
	if check(err, context) != nil {
		return
	}

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

	newTodo, err := todos.AddNewTodo(todoText)
	if check(err, context) != nil {
		return
	}

	err = t.ExecuteTemplate(context.Writer, "todo", newTodo)
	if check(err, context) != nil {
		return
	}
	context.Status(http.StatusOK)
}

func ToggleDone(context *gin.Context) {
	t, err := template.ParseFiles("./templates/todo.tmpl")
	if check(err, context) != nil {
		return
	}

	paramId, _ := strconv.Atoi(context.Param("id"))
	todoId := uint32(paramId)
	updatedTodo, err := todos.ToggleDone(todoId)

	err = t.ExecuteTemplate(context.Writer, "todo", updatedTodo)
	if check(err, context) != nil {
		return
	}
	context.Status(http.StatusOK)
}

func DeleteTodo(context *gin.Context) {
	t, err := template.ParseFiles("./templates/todo.tmpl")
	if check(err, context) != nil {
		return
	}

	paramId, _ := strconv.Atoi(context.Param("id"))
	todoId := uint32(paramId)

	remainingTodos, err := todos.DeleteTodo(todoId)
	if check(err, context) != nil {
		return
	}

	for _, todo := range remainingTodos {
		err = t.ExecuteTemplate(context.Writer, "todo", todo)
		if check(err, context) != nil {
			return
		}
	}

	context.Status(http.StatusOK)
}

type indexData struct {
	Name  string
	Todos []todos.Todo
}
