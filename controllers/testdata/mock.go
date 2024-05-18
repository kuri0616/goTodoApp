package testdata

import "github.com/rikuya98/goTodoApp/models"

type MockTodoSev struct{}

func NewMockTodoSev() *MockTodoSev {
	return &MockTodoSev{}
}

func (m *MockTodoSev) GetTodoServices() ([]models.Todo, error) {
	return TodoData, nil
}

func (m *MockTodoSev) PostTodoServices(todo models.Todo) (models.Todo, error) {
	return todo, nil
}

func (m *MockTodoSev) PutTodoServices(todo models.Todo) (models.Todo, error) {
	return todo, nil
}

func (m *MockTodoSev) DeleteTodoServices(id int) error {
	return nil
}
