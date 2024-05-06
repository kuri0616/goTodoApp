package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/rikuya98/goTodoApp/models"
)

func InsertTodo(db *sqlx.DB, todo models.Todo) (models.Todo, error) {
	var newTodo models.Todo
	result, err := db.Exec("INSERT INTO todos (task,due_date,status,created_at) VALUES (?, ?, 0, now())", todo.Task, todo.DueDate)
	if err != nil {
		return models.Todo{}, err
	}
	id, _ := result.LastInsertId()
	newTodo = models.Todo{
		Id:      int(id),
		Task:    todo.Task,
		DueDate: todo.DueDate,
		Status:  0,
	}
	return newTodo, nil
}
