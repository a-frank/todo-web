package main

import (
	"flag"
	"fmt"
	"github.com/a-frank/web-dev/routes"
	"github.com/a-frank/web-dev/todos"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang.org/x/term"
	"os"
	"strconv"
	"syscall"
)

func main() {
	todoStore, err := setupTodoStore()
	if err != nil {
		panic(err)
	}
	defer todoStore.Close()

	env := &routes.RouterEnv{
		Store:        todoStore,
		TemplatePath: "./templates/",
	}

	ginServer := gin.Default()

	ginServer.Static("/css", "./templates/css")
	ginServer.Static("/images", "./templates/images")

	ginServer.GET("/", func(context *gin.Context) {
		routes.GetIndex(context, env)
	})
	ginServer.POST("/todo", func(context *gin.Context) {
		routes.AddTodo(context, env)
	})
	ginServer.POST("/todo/:id/toggle-done", func(context *gin.Context) {
		routes.ToggleDone(context, env)
	})
	ginServer.DELETE("/todo/:id", func(context *gin.Context) {
		routes.DeleteTodo(context, env)
	})

	err = ginServer.Run()
	fmt.Printf("Error with server %s", err.Error())
}

func setupTodoStore() (todos.TodoStore, error) {
	var host string
	var port int
	var dbname string
	var user string
	var password = os.Getenv("DB_PASSWORD")

	flag.StringVar(&host, "h", "", "Host of the DB")
	flag.IntVar(&port, "p", 0, "Port to connect to the DB")
	flag.StringVar(&dbname, "db", "", "DB name to connect to")
	flag.StringVar(&user, "u", "", "User for the selected DB")
	flag.Parse()

	if host == "" {
		host = os.Getenv("DB_HOST")
	}
	if port == 0 {
		port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	}
	if dbname == "" {
		dbname = os.Getenv("DB_NAME")
	}
	if user == "" {
		user = os.Getenv("DB_USER")
	}
	if password == "" {
		fmt.Print("Enter the password for your DB user: ")
		pwdBytes, err := term.ReadPassword(syscall.Stdin)
		if err != nil {
			panic(fmt.Sprintf("Password input failed: %s", err.Error()))
		}
		password = string(pwdBytes)
	}

	return todos.NewTodoStore(host, port, dbname, user, password, false)
}
