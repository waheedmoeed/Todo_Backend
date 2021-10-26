package rest

import (
	"net/http"

	"github.com/abdulwaheed/todobackend/service"
	"github.com/gin-gonic/gin"
)

type TodoController interface {
	AddTask(ctx *gin.Context)
	GetTasks(ctx *gin.Context)
}

type todoController struct {
	todoService service.TodoService
}

func NewTodoController(todoService service.TodoService) TodoController {
	return &todoController{
		todoService: todoService,
	}
}

func (t *todoController) AddTask(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Ready to GO!")
}

func (t *todoController) GetTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Ready to GO!")
}
