package services_test

import (
	"github.com/rikuya98/goTodoApp/config"
	"github.com/rikuya98/goTodoApp/services"
	"testing"
)

var service *services.TodoAppSev

func TestMain(m *testing.M) {
	db := config.InitDB()
	service = services.NewTodoAppSev(db)
	m.Run()
}

func BenchmarkTodoAppSev_GetTodoServices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := service.GetTodoServices()
		if err != nil {
			b.Error(err)
			break
		}
	}
}
