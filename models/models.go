package models

import "time"

type Todo struct {
	Id       int       `json:"id"`
	Task     string    `json:"task"`
	DueDate  time.Time `json:"due_date"`
	Status   int       `json:"status"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
