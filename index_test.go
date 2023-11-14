package main

import (
	"github.com/a-frank/web-dev/routes"
	"github.com/a-frank/web-dev/test"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetIndex(t *testing.T) {
	testRecorder := httptest.NewRecorder()
	_, testEngine := gin.CreateTestContext(testRecorder)

	todoStore := &test.MemStore{}
	_, _ = todoStore.Add("New Todo")

	testEngine.GET("/", func(context *gin.Context) {
		routes.GetIndex(context, todoStore)
	})

	request := httptest.NewRequest("GET", "/", nil)
	testEngine.ServeHTTP(testRecorder, request)

	if testRecorder.Code != 200 {
		t.Errorf("Request not successful. Got code %d", testRecorder.Code)
	}

	body := string(testRecorder.Body.Bytes())

	if !strings.Contains(body, "<tr id=\"todo_1\"") {
		t.Error("No row found for todo_1")
	}

	if !strings.Contains(body, "<td>New Todo<b/td>") {
		t.Error("Didn't find cell with todo text")
	}
}
