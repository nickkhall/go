package todo

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Todo Struct : Data structure for a Todo
type Todo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

// PopulateTodos : Populates todo slice
func PopulateTodos() []Todo {
	todos := []Todo{
		{
			ID:        "1234",
			Name:      "Clean house",
			Completed: false,
		},
		{
			ID:        "1233",
			Name:      "learn golang",
			Completed: false,
		},
	}

	return todos
}

// GetTodo : Gets a single todo
func GetTodo(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["id"]
	todos := PopulateTodos()

	for _, todo := range todos {
		if todo.ID == todoID {
			json.NewEncoder(w).Encode(todo)
		}
	}
}
