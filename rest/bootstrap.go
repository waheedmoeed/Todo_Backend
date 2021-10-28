package rest

import (
	"fmt"

	"github.com/abdulwaheed/todobackend/service"
)

//StartServer initate server
func StartServer(container *service.Container) *HttpServer {
	//Inject services instance from ServiceContainer
	todoController := NewTodoController(container.TodoService)

	httpServer := NewHttpServer()

	//Inject controller instance to server
	httpServer.todoController = todoController

	go httpServer.Start()
	fmt.Println("rest server ok")
	return httpServer
}
