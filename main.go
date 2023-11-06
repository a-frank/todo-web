package main

import (
	"fmt"
	"github.com/a-frank/web-dev/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := gin.Default()
	ginServer.GET("/", routes.GetIndex)
	ginServer.POST("/todo", routes.AddTodo)
	err := ginServer.Run()
	fmt.Printf("Error with server %s", err.Error())
}
