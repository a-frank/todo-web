package main

import (
	"fmt"
	"github.com/a-frank/web-dev/routes"
	"github.com/a-frank/web-dev/todos"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

var portenv, _ = strconv.Atoi(os.Getenv("DB_PORT"))
var (
	host     = os.Getenv("DB_HOST")
	port     = portenv
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

func main() {
	todoStore, err := todos.NewTodoStore(host, port, dbname, user, password, false)
	if err != nil {
		panic(err)
	}
	defer todoStore.Close()

	ginServer := gin.Default()

	ginServer.Static("/css", "./templates/css")
	ginServer.Static("/images", "./templates/images")

	ginServer.GET("/", func(context *gin.Context) {
		routes.GetIndex(context, todoStore)
	})
	ginServer.POST("/todo", func(context *gin.Context) {
		routes.AddTodo(context, todoStore)
	})
	ginServer.POST("/todo/:id/toggle-done", func(context *gin.Context) {
		routes.ToggleDone(context, todoStore)
	})
	ginServer.DELETE("/todo/:id", func(context *gin.Context) {
		routes.DeleteTodo(context, todoStore)
	})

	err = ginServer.Run()
	fmt.Printf("Error with server %s", err.Error())
}
