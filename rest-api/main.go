package main

import (
	"log"
	"net/http"
	"database/sql"
	"fmt"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	todo "github.com/nickkhall/go/rest-api/handlers"
	database "github.com/nickkhall/go/rest-api/database"
)

func main() {
	var err error

	database.DBCon, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer database.DBCon.Close()

	err = database.DBCon.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected! Running server on port 3000")

	router := mux.NewRouter()
	router.HandleFunc("/todos", todo.GetTodos).Methods("GET")
  	router.HandleFunc("/todos", todo.CreateTodo).Methods("POST")
  	router.HandleFunc("/todos/{id}", todo.GetTodo).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
