package main

import (
	"github.com/mongodb/mongo-go-driver/mongo"
)

func CreateClient() {
	client, err := mongo.NewClient("mongodb://localhost:3000")
}
