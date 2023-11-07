package main

import (
	"fmt"
	"github.com/a-frank/web-dev/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := gin.Default()

	ginServer.Static("/css", "./templates/css")

	ginServer.GET("/", routes.GetIndex)
	ginServer.POST("/todo", routes.AddTodo)
	ginServer.POST("/todo/:id/toggle-done", routes.ToggleDone)

	err := ginServer.Run()
	fmt.Printf("Error with server %s", err.Error())
}
