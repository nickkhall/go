package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	database "github.com/nickkhall/go/rest-api/database"
	todo "github.com/nickkhall/go/rest-api/handlers"
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
	router.HandleFunc("/todos/{id}", todo.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", todo.DeleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
