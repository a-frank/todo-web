package routes

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func Index(context *gin.Context) {
	t, err := template.ParseFiles("./templates/index.html")
	if check(err, context) != nil {
		return
	}

	err = t.Execute(context.Writer, indexData{
		Name: "Go Gin",
		})
	if check(err, context) != nil {
		return
	}
	context.Status(http.StatusOK)
}

type indexData struct {
	Name string
}
