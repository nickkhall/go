package todo

import (
	"encoding/json"
	"net/http"
	"log"
	"io/ioutil"
	"context"

	"github.com/gorilla/mux"
	database "github.com/nickkhall/go/rest-api/database"
	errors "github.com/nickkhall/go/rest-api/errors"
)

type Todo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

// GetTodos : Gets all todos
func GetTodos(w http.ResponseWriter,  r *http.Request) {
	todos := []Todo{}
	rows, err := database.DBCon.Query("SELECT * FROM todos;")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id        string
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

// GetTodo : Gets a single Todo
func GetTodo(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["id"]

	var id        string
	var name      string
	var completed bool
  	var todo Todo

	err := database.DBCon.QueryRowContext(context.Background(), "SELECT * FROM todos WHERE id = $1", todoId).Scan(&id, &name, &completed)

	switch {
		case err != nil:
			e := errors.CustomError{404, "Todo does not exist"}
			json.NewEncoder(w).Encode(e)
			return
		default:
			todo = Todo{id, name, completed}
	}

	json.NewEncoder(w).Encode(todo)
}

// CreateTodo : Creates a Todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var todo Todo

	err = json.Unmarshal(reqBody, &todo)
	if err != nil {
		log.Fatal(err)
	}

	sqlStatement := `
	INSERT INTO todos (id, name, completed)
	VALUES ($1, $2, $3)
	`

	_, dbErr := database.DBCon.Exec(sqlStatement, string(todo.ID), string(todo.Name), bool(todo.Completed))
	if err != nil {
		log.Fatal(dbErr)
	}

  json.NewEncoder(w).Encode(todo)
}


// UpdateTodo : Updates an existing Todo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e := errors.CustomError{400, "Bad Request"}
		json.NewEncoder(w).Encode(e)
	}

	var todo Todo

	err = json.Unmarshal(reqBody, &todo)
	if err != nil {
		log.Fatal(err)
	}

	sqlStatement := `
	UPDATE todos SET name = $2, completed = $3 WHERE id = $1;
	`

	_, dbErr := database.DBCon.Exec(sqlStatement, string(todo.ID), string(todo.Name), bool(todo.Completed))
	if dbErr != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(todo)
}

// DeleteTodo : Deletes a Todo
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["id"]

	sqlStatement := `
	DELETE FROM todos WHERE id = $1;
	`

	_, dbErr := database.DBCon.Exec(sqlStatement, todoId)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
}
