package routes

import (
	"github.com/a-frank/todo-web/test"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetIndex(t *testing.T) {
	testRecorder := httptest.NewRecorder()
	_, testEngine := gin.CreateTestContext(testRecorder)

	todoStore := test.NewMemStore()
	_, _ = todoStore.Add("New Todo")
	env := &RouterEnv{
		Store:        todoStore,
		TemplatePath: "../templates/",
	}

	testEngine.GET("/", func(context *gin.Context) {
		GetIndex(context, env)
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

	if !strings.Contains(body, "<td>New Todo</td>") {
		t.Error("Didn't find cell with todo text")
	}
}
