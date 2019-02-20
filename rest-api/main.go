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

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "todos"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
	"password=%s dbname=%s",
	host, port, user, password, dbname)

	var err error

	database.DBCon, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	defer database.DBCon.Close()

	err = database.DBCon.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfull connected! Running server on port 3000")

	router := mux.NewRouter()
	router.HandleFunc("/todos", todo.GetTodos).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
