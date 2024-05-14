package repositories

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/rikuya98/goTodoApp/repositories/testdata"
	"regexp"
	"testing"
)

func TestGetTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	var todo = testdata.TodoData
	rows := sqlmock.NewRows([]string{"id", "task", "due_date", "status", "created_at", "updated_at"}).
		AddRow(todo[0].Id, todo[0].Task, todo[0].DueDate, todo[0].Status, todo[0].CreatedAt, todo[0].UpdatedAt).
		AddRow(todo[1].Id, todo[1].Task, todo[1].DueDate, todo[1].Status, todo[1].CreatedAt, todo[1].UpdatedAt)

	mock.ExpectQuery("SELECT id,task,due_date, status,created_at,updated_at FROM todos").WillReturnRows(rows)

	_, err = GetTodos(sqlxDB)
	if err != nil {
		t.Errorf("error was not expected while getting todos: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	var expectedTodo = testdata.TodoData[0]

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO todos (task,due_date,status,created_at) VALUES (?, ?, 0, now()`)).
		WithArgs(expectedTodo.Task, expectedTodo.DueDate).WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = InsertTodo(sqlxDB, expectedTodo)
	if err != nil {
		t.Errorf("error was not expected while inserting todo: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
