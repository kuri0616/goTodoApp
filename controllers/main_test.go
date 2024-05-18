package controllers_test

import (
	"github.com/rikuya98/goTodoApp/controllers"
	"github.com/rikuya98/goTodoApp/controllers/testdata"
	"testing"
)

var con *controllers.TodoController

func TestMain(m *testing.M) {
	sev := testdata.NewMockTodoSev()
	con = controllers.NewTodoController(sev)
	m.Run()
}
