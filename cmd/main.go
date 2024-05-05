package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rikuya98/goTodoApp/config"
)

func getTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get Todo"))
}
func postTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post Todo"))
}
func putTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("put Todo"))
}
func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete Todo"))
}

func main() {
	config.LoadDBConfig()
	r := mux.NewRouter()
	r.HandleFunc("/todo", getTodoHandler).Methods(http.MethodGet)
	r.HandleFunc("/todo", postTodoHandler).Methods(http.MethodPost)
	r.HandleFunc("/todo", putTodoHandler).Methods(http.MethodPut)
	r.HandleFunc("/todo", deleteTodoHandler).Methods(http.MethodDelete)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
