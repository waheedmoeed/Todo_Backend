package rest

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestGetTask(t *testing.T) {
	t.Run("Should return greeting", func(t *testing.T) {
		//mocked servives and http server
		mockHttpServer := NewTestHttpServer(t)

		//mocked Http request and recording response of request
		router := gin.Default()
		rr := httptest.NewRecorder()

		router.GET("/api/task", mockHttpServer.GetTodoRestController().todoController.GetTasks)

		request, _ := http.NewRequest(http.MethodGet, "/api/task", nil)
		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)
		//read response and validate response
		assert.Equal(t, 200, rr.Code)

		bodyBytes, err := ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Fail()
		}

		if len(string(bodyBytes)) == 0 {
			t.Fail()
		}
	})
}

func TestAddTask(t *testing.T) {
	t.Run("Should return greeting", func(t *testing.T) {
		//mocked servives and http server
		mockHttpServer := NewTestHttpServer(t)

		//mocked Http request and recording response of request
		router := gin.Default()
		rr := httptest.NewRecorder()

		router.POST("/api/task", mockHttpServer.GetTodoRestController().todoController.GetTasks)

		request, _ := http.NewRequest(http.MethodPost, "/api/task", nil)
		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)
		//read response and validate response
		assert.Equal(t, 200, rr.Code)

		bodyBytes, err := ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Fail()
		}

		if len(string(bodyBytes)) == 0 {
			t.Fail()
		}
	})
}
