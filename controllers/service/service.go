package service

import "github.com/rikuya98/goTodoApp/models"

type TodoService interface {
	GetTodoServices() ([]models.Todo, error)
	PostTodoServices(todo models.Todo) (models.Todo, error)
	PutTodoServices(todo models.Todo) (models.Todo, error)
	DeleteTodoServices(id int) error
}
