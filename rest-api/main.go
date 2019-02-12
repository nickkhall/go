package main

import (
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"

	"google.golang.org/api/option"
)

func main() {
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	client, err := app.Firestore(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/todos/{id}", GetTodo).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))

	defer client.Close()
}
