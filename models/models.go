package models

import "time"

type Todo struct {
	Id        int       `json:"id" db:"id"`
	Task      string    `json:"task" db:"task"`
	DueDate   time.Time `json:"due_date" db:"due_date"`
	Status    int       `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
