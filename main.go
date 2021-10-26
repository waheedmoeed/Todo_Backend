package main

import (
	"fmt"

	"github.com/abdulwaheed/todobackend/rest"
	"github.com/abdulwaheed/todobackend/service"
)

func main() {
	/*
	* Initiate Service Layer Container
	 */

	serviceContainer := service.NewServiceContainer()

	/*
	* Initiate Rest Server
	 */
	rest.StartServer(serviceContainer)

	fmt.Println("========== Rest Server Started ============")
	select {}
}
