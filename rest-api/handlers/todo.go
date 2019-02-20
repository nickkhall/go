package todo

import (
	"encoding/json"
	"net/http"
	"log"
  "io/ioutil"

  "github.com/gorilla/mux"
	database "github.com/nickkhall/go/rest-api/database"
)

type Todo struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type Error struct {
  Status  int64 `json:"status"`
  Message string `json:"message"`
}

// ReturnError : Returns the Error struct with a custom message
func ReturnError(error *Error, errStatus int64, errMsg string) Error {
  *Error.Status = errStatus
  *Error.Message = errMsg
  return Error
}

// GetTodos : Gets all todos
func GetTodos(w http.ResponseWriter,  r *http.Request) {
	todos := []Todo{}
	rows, err := database.DBCon.Query("SELECT * FROM tododb;")

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
  INSERT INTO tododb (id, name, completed)
  VALUES ($1, $2, $3)
  `

  _, dbErr := database.DBCon.Exec(sqlStatement, int(todo.ID), string(todo.Name), bool(todo.Completed))
  if err != nil {
    log.Fatal(dbErr)
  }
}

// GetTodo : Gets a single Todo
func GetTodo(w http.ResponseWriter, r *http.Request) {
  todoId := mux.Vars(r)["id"]

  sqlStatement := `
  SELECT EXISTS (SELECT * FROM tododb WHERE id=$1);
  `

  rows, dbErr := database.DBCon.Query(sqlStatement, todoId)
  if dbErr != nil {
    log.Fatal(dbErr)
  }

  todo := Todo{}

  defer rows.Close()

  for rows.Next() {
    var id        int64
    var name      string
    var completed bool

    err := rows.Scan(&id, &name, &completed)
    if err != nil {
      getErr := ReturnError(Error, 404, "Todo does not exist.")
      json.NewEncoder(w).Encode(Error)
    }

    todo = Todo{id, name, completed}
  }

  json.NewEncoder(w).Encode(todo)
}

