package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rikuya98/goTodoApp/models"
	"github.com/rikuya98/goTodoApp/services"
)

type TodoController struct {
	service services.TodoAppSev
}

func NewTodoController(service *services.TodoAppSev) *TodoController {
	return &TodoController{service: *service}
}

func (s *TodoController) GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := s.service.GetTodoServices()
	if err != nil {
		log.Println("Failed to get todos", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (s *TodoController) PostTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	newTodo, err := s.service.PostTodoServices(todo)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTodo)
}
func (s *TodoController) PutTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		println("Failed to decode json", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	todo.Id, _ = strconv.Atoi(mux.Vars(r)["id"])
	updatedTodo, err := s.service.PutTodoServices(todo)
	if err != nil {
		println("Failed to update todo", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTodo)
}

func (s *TodoController) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = s.service.DeleteTodoServices(id)
	if err != nil {
		println("Failed to delete todo", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
