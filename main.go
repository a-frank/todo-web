package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	ginServer := gin.Default()
	
	ginServer.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	
	err := ginServer.Run()
	fmt.Printf("Error with server %s", err.Error())
}