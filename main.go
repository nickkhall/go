package main

import (
    "fmt"
    "encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/gorilla/mux"
)

type Todo struct {
  Name      string    `json:"name"`
  Completed bool      `json:"completed"`
  Due       time.Time `json:"due"`
}

type Todos []Todo

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/todos", TodoIndex)
    router.HandleFunc("/todos/{todoId}", TodoShow)

    fmt.Printf("Server running on Port 3000...")


    log.Fatal(http.ListenAndServe(":3000", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Welcome to the homepage!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
  todos := Todos{
    Todo{Name: "Learn GoLang"},
    Todo{Name: "Punch Synnux"},
    Todo{Name: "Tell Norpyx HE FUCKKUNN SUUUUUUUUUUUUCgsZ"},
  }

  json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  todoId := vars["todoId"]
  fmt.Fprintln(w, "Here is the todo:", todoId)
}
