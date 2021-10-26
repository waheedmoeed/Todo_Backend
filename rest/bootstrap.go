package rest

import (
	"fmt"

	"github.com/abdulwaheed/todobackend/conf"
	"github.com/abdulwaheed/todobackend/service"
)

//StartServer initate server
func StartServer(container *service.Container) *HttpServer {
	gbeConfig := conf.GetConfig()
	//Inject services instance from ServiceContainer
	todoController := NewTodoController(container.TodoService)

	httpServer := NewHttpServer(
		fmt.Sprintf("%v%v", gbeConfig.Host, gbeConfig.Port),
	)

	//Inject controller instance to server
	httpServer.todoController = todoController

	go httpServer.Start()
	fmt.Println("rest server ok")
	return httpServer
}
