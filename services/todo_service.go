package services

import (
	"encoding/json"
	"net/http"

	"log"

	"github.com/rikuya98/goTodoApp/models"
	"github.com/rikuya98/goTodoApp/repositories"
)

func (s *TodoAppSev) GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	todos, err := repositories.GetTodos(s.db)
	if err != nil {
		log.Println("Failed to get todos", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
func (s *TodoAppSev) PostTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	newTodo, err := repositories.InsertTodo(s.db, todo)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTodo)
}
func (s *TodoAppSev) PutTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("put Todo"))
}
func (s *TodoAppSev) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete Todo"))
}
