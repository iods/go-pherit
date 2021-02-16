package main

import (
	"context"
	"fmt"
	"log"

	"github.com/iods/go-pherit/internal/database"
)

type Record struct {
	Id int
	Name string
	Description string
}

func main() {
	singleRecord()
	multipleRecords()
}

func singleRecord() {
	client, err := database.GetMongoDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("testdb").Collection("records")

	one := Record{1, "Rye", "This is the first record in the DB."}

	result, err := collection.InsertOne(context.TODO(), one)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("You are connected motha fucka!, you inserted %s", result.InsertedID)
}

func multipleRecords() {

	client, err := database.GetMongoDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("dbtest").Collection("records")

	two := Record{2, "Two", "This is the second record added to the collection."}
	three := Record{3, "Three", "This is the third record added to the collection."}

	collections := []interface{}{two, three}

	result, err := collection.InsertMany(context.TODO(), collections)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("You inserted", result.InsertedIDs)
}