package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rikuya98/goTodoApp/config"
	"github.com/rikuya98/goTodoApp/services"
)

func main() {
	db := config.InitDB()
	appSev := services.NewTodoAppSev(db)

	r := mux.NewRouter()
	r.HandleFunc("/todo", appSev.GetTodoHandler).Methods(http.MethodGet)
	r.HandleFunc("/todo", appSev.PostTodoHandler).Methods(http.MethodPost)
	r.HandleFunc("/todo/{id}", appSev.PutTodoHandler).Methods(http.MethodPut)
	r.HandleFunc("/todo/{id}", appSev.DeleteTodoHandler).Methods(http.MethodDelete)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
