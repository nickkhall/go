package todo

import (
	"encoding/json"
	"net/http"
	"log"

	database "github.com/nickkhall/go/rest-api/database"
)

type Todo struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

// GetTodos : Gets all todos
func GetTodos(w http.ResponseWriter,  r *http.Request) {
	todos := []Todo{}
	rows, err := database.DBCon.Query("SELECT id, name, completed FROM tododb")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id        int64
		var name      string
		var completed bool

		err = rows.Scan(&id, &name, &completed)
		if err != nil {
			log.Fatal(err)
		}

		todo := Todo{id, name, completed}
		todos = append(todos, todo)

	}

	json.NewEncoder(w).Encode(todos)
}

