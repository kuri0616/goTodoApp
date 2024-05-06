package services

import "database/sql"

type TodoAppSev struct {
	db *sql.DB
}

func NewTodoAppSev(db *sql.DB) *TodoAppSev {
	return &TodoAppSev{db: db}
}
