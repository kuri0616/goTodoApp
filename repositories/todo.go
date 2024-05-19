package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/rikuya98/goTodoApp/apperrors"
	"github.com/rikuya98/goTodoApp/models"
)

func GetTodos(db *sqlx.DB) ([]models.Todo, error) {
	var todos []models.Todo
	err := db.Select(&todos, "SELECT id,task,due_date, status,created_at,updated_at FROM todos")
	if err != nil {
		return []models.Todo{}, err
	}
	return todos, nil
}

func InsertTodo(db *sqlx.DB, todo models.Todo) (models.Todo, error) {
	var newTodo models.Todo
	result, err := db.Exec("INSERT INTO todos (task,due_date,status,created_at) VALUES (?, ?, 0, now())", todo.Task, todo.DueDate)
	if err != nil {
		return models.Todo{}, err
	}
	id, _ := result.LastInsertId()
	if id == 0 {
		return models.Todo{}, apperrors.ErrNoData
	}
	newTodo = models.Todo{
		Id:      int(id),
		Task:    todo.Task,
		DueDate: todo.DueDate,
		Status:  0,
	}
	return newTodo, nil
}

func UpdateTodo(db *sqlx.DB, todo models.Todo) (models.Todo, error) {
	result, err := db.Exec("UPDATE todos SET task = ?, due_date = ?, status = ? WHERE id = ?", todo.Task, todo.DueDate, todo.Status, todo.Id)

	if err != nil {
		return models.Todo{}, err
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return models.Todo{}, err
	}
	if rowsAffected == 0 {
		err = apperrors.ErrNoData
		return models.Todo{}, err
	}
	return todo, nil
}

func DeleteTodo(db *sqlx.DB, id int) error {
	result, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		err = apperrors.ErrNoData
		return err
	}
	return nil
}
