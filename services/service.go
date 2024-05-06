package services

import (
	"github.com/jmoiron/sqlx"
)

type TodoAppSev struct {
	db *sqlx.DB
}

func NewTodoAppSev(db *sqlx.DB) *TodoAppSev {
	return &TodoAppSev{db: db}
}
