package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func check(err error, context *gin.Context) error {
	if err != nil {
		_ = context.Error(err)
		context.Writer.WriteHeader(http.StatusInternalServerError)
		log.Default().Printf("Template error: %s", err.Error())
		return err
	}
	return nil
}