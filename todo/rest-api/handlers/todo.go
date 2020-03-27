package todo

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	database "github.com/nickkhall/go/rest-api/database"
	errors "github.com/nickkhall/go/rest-api/errors"
)

// Todo : Todo Structure
type Todo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	CreatedAt int64  `json:"createdAt"`
}

// GetTodos : Gets all todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{}
	rows, err := database.DBCon.Query("SELECT * FROM todos;")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id        string
			name      string
			completed bool
			createdAt int64
		)

		err = rows.Scan(&id, &name, &completed, &createdAt)
		if err != nil {
			log.Fatal(err)
		}

		todo := Todo{id, name, completed, createdAt}
		todos = append(todos, todo)
	}

	json.NewEncoder(w).Encode(todos)
}

// GetTodo : Gets a single Todo
func GetTodo(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["id"]

	var (
		id        string
		name      string
		completed bool
		createdAt int64
		todo      Todo
	)

	err := database.DBCon.QueryRowContext(context.Background(), "SELECT * FROM todos WHERE id = $1", todoID).Scan(&id, &name, &completed, &createdAt)

	switch {
	case err != nil:
		e := errors.CustomError{
			Status:  404,
			Message: "Todo does not exist",
		}

		json.NewEncoder(w).Encode(e)
		return
	default:
		todo = Todo{id, name, completed, createdAt}
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
	timestamp := time.Now().Unix()
	todoID := uuid.New().String()
	todo.ID = todoID
	todo.CreatedAt = timestamp

	err = json.Unmarshal(reqBody, &todo)
	if err != nil {
		log.Fatal(err)
	}

	sqlStatement := `
	INSERT INTO todos (id, name, completed, createdAt)
	VALUES ($1, $2, $3, $4)
	`

	_, dbErr := database.DBCon.Exec(sqlStatement, string(todo.ID), string(todo.Name), false, timestamp)
	if err != nil {
		log.Fatal(dbErr)
	}

	json.NewEncoder(w).Encode(todo)
}

// UpdateTodo : Updates an existing Todo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e := errors.CustomError{
			Status:  400,
			Message: "Bad Request",
		}

		json.NewEncoder(w).Encode(e)
	}

	var todo Todo
	todoID := mux.Vars(r)["id"]

	err = json.Unmarshal(reqBody, &todo)
	if err != nil {
		log.Fatal(err)
	}

	todo.ID = todoID

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
	todoID := mux.Vars(r)["id"]

	sqlStatement := `
	DELETE FROM todos WHERE id = $1;
	`

	_, dbErr := database.DBCon.Exec(sqlStatement, todoID)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
}
