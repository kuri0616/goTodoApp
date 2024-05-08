package services

import (
	"github.com/rikuya98/goTodoApp/apperrors"
	"github.com/rikuya98/goTodoApp/models"
	"github.com/rikuya98/goTodoApp/repositories"
)

func (s *TodoAppSev) GetTodoServices() ([]models.Todo, error) {
	var todos []models.Todo
	todos, err := repositories.GetTodos(s.db)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "Failed to get todos")
		return []models.Todo{}, err
	}
	if len(todos) == 0 {
		err = apperrors.ErrNoData
		apperrors.Nodata.Wrap(err, "No data found")
		return []models.Todo{}, err
	}
	return todos, nil
}
func (s *TodoAppSev) PostTodoServices(todo models.Todo) (models.Todo, error) {
	newTodo, err := repositories.InsertTodo(s.db, todo)
	if err != nil {
		err = apperrors.InserDetaFailed.Wrap(err, "Failed to insert todo")
		return models.Todo{}, err
	}
	return newTodo, nil
}
func (s *TodoAppSev) PutTodoServices(todo models.Todo) (models.Todo, error) {
	updatedTodo, err := repositories.UpdateTodo(s.db, todo)
	if err != nil {
		err = apperrors.UpdateDataFailed.Wrap(err, "Failed to update todo")
		return models.Todo{}, err
	}
	return updatedTodo, nil
}

func (s *TodoAppSev) DeleteTodoServices(id int) error {
	err := repositories.DeleteTodo(s.db, id)
	if err != nil {
		err = apperrors.DeleteDataFailed.Wrap(err, "Failed to delete todo")
		return err
	}
	return nil
}
