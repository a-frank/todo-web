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
	handler, err := todos.NewDbHandler(host, port, dbname, user, password, false)
	if err != nil {
		panic(err)
	}
	defer handler.CloseConnection()

	ginServer := gin.Default()

	ginServer.Static("/css", "./templates/css")
	ginServer.Static("/images", "./templates/images")

	ginServer.GET("/", func(context *gin.Context) {
		routes.GetIndex(context, handler)
	})
	ginServer.POST("/todo", func(context *gin.Context) {
		routes.AddTodo(context, handler)
	})
	ginServer.POST("/todo/:id/toggle-done", func(context *gin.Context) {
		routes.ToggleDone(context, handler)
	})
	ginServer.DELETE("/todo/:id", func(context *gin.Context) {
		routes.DeleteTodo(context, handler)
	})

	err = ginServer.Run()
	fmt.Printf("Error with server %s", err.Error())
}
