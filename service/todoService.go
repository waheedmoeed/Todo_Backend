package service

type TodoService interface {
	AddTask(task string) error
	GetTask(task string) ([]string, error)
}

type todoService struct {
}

func NewTodoService() TodoService {
	return &todoService{}
}

func (c *todoService) AddTask(task string) error {
	return nil
}

func (c *todoService) GetTask(task string) ([]string, error) {
	return nil, nil
}
