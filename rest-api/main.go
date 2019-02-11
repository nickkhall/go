package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/todos", GetAllTodos).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}
