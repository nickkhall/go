package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Todo Struct : Data structure for a Todo
type Todo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

// GetTodo : Gets a single todo
func GetTodo(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["id"]

	dsnap, err := client.Collection("react-redux-todos").Doc("todos").Get(todoId)
	if err != nil {
		log.Fatalln(err)
	}
}
