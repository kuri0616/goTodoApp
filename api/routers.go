package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rikuya98/goTodoApp/config"
	"github.com/rikuya98/goTodoApp/controllers"
	"github.com/rikuya98/goTodoApp/services"
)

func NewRouter() *mux.Router {
	db := config.InitDB()
	appSev := services.NewTodoAppSev(db)
	con := controllers.NewTodoController(appSev)

	r := mux.NewRouter()
	r.HandleFunc("/todo", con.GetTodoHandler).Methods(http.MethodGet)
	r.HandleFunc("/todo", con.PostTodoHandler).Methods(http.MethodPost)
	r.HandleFunc("/todo/{id}", con.PutTodoHandler).Methods(http.MethodPut)
	r.HandleFunc("/todo/{id}", con.DeleteTodoHandler).Methods(http.MethodDelete)
	return r
}
