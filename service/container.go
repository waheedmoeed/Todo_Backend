package service

import "github.com/abdulwaheed/todobackend/repository"

type Container struct {
	Repository repository.Repository

	//services
	TodoService TodoService
}

func NewContainer() *Container {
	repo := repository.NewRepository()

	todoService := NewTodoService()

	container := &Container{
		Repository:  repo,
		TodoService: todoService,
	}

	return container
}
