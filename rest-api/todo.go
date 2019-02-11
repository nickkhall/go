package main

import (
	"encoding/json"
	"net/http"
)

// Todo Struct : Data structure for a Todo
type Todo struct {
	ID        string `json:"id",omitempty`
	Name      string `json:"name,omitempty`
	Completed bool   `json:"completed"`
}

var todos []Todo

// GetAllTodos : Returns all current Todos
func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos = append(todos, Todo{
		ID:        "1234",
		Name:      "take out the trash",
		Completed: false,
	})
	todos = append(todos, Todo{
		ID:        "1235",
		Name:      "sweep the cat food",
		Completed: true,
	})

	json.NewEncoder(w).Encode(todos)
}
