package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	todo "github.com/nickkhall/go/rest-api/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/todos/{id}", todo.GetTodo).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
