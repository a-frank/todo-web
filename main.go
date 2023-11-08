package main

import (
	"database/sql"
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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	todos.DbConnection = db
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	ginServer := gin.Default()

	ginServer.Static("/css", "./templates/css")
	ginServer.Static("/images", "./templates/images")

	ginServer.GET("/", routes.GetIndex)
	ginServer.POST("/todo", routes.AddTodo)
	ginServer.POST("/todo/:id/toggle-done", routes.ToggleDone)
	ginServer.DELETE("/todo/:id", routes.DeleteTodo)

	err = ginServer.Run()
	fmt.Printf("Error with server %s", err.Error())
}
