package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rikuya98/goTodoApp/api"
)

func main() {
	r := api.NewRouter()
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
