package rest

import (
	"testing"

	mock_service "github.com/abdulwaheed/todobackend/mocks/service"
	"github.com/golang/mock/gomock"
)

type testHttpServer struct {
	ctrl            *gomock.Controller
	mockTodoService *mock_service.MockTodoService
}

func NewTestHttpServer(t *testing.T) *testHttpServer {
	mockCtrl := gomock.NewController(t)

	mockTodoService := mock_service.NewMockTodoService(mockCtrl)

	return &testHttpServer{
		ctrl:            mockCtrl,
		mockTodoService: mockTodoService,
	}
}

func (test *testHttpServer) GetTodoRestController() *HttpServer {
	todoController := NewTodoController(test.mockTodoService)
	return &HttpServer{
		addr:           "",
		todoController: todoController,
	}
}
