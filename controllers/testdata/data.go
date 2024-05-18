package testdata

import (
	"github.com/rikuya98/goTodoApp/models"
	"time"
)

var TodoData = []models.Todo{
	{
		Id:        1,
		Task:      "task1",
		DueDate:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Status:    0,
		CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		Id:        2,
		Task:      "task2",
		DueDate:   time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
		Status:    0,
		CreatedAt: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
	},
}
