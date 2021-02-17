package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"github.com/iods/go-pherit/internal/common"
	"github.com/iods/go-pherit/internal/database"
)

type Record struct {
	Type        string `bson:"type"`
	Name 		string `bson:"name"`
	Description string `bson:"description"`
}

var (
	ctx = context.TODO()
)

func main() {
	insertRecord(Record{"One", "One", "This is the first record in the db."})
	insertRecords()
	readRecord()
	updateRecord()
	readRecords()
}

func insertRecord(one Record) {
	c, err := database.Session()
	common.HandleError(err)
	result, err := c.InsertOne(ctx, one)
	common.HandleError(err)
	fmt.Println("You have inserted one record:", result.InsertedID)
}

func insertRecords() {
	c, err := database.Session()
	common.HandleError(err)
	two := Record{"Two", "Two", "This is the second record added to the collection."}
	three := Record{"Three", "Three", "This is the third record added to the collection."}
	r := []interface{}{two, three}
	result, err := c.InsertMany(ctx, r)
	common.HandleError(err)
	fmt.Println("You inserted multiple records: ", result.InsertedIDs)
}

func readRecord() {
	filter := bson.M{"type": "One"}
	var record Record
	c, err := database.Session()
	common.HandleError(err)
	result := c.FindOne(ctx, filter)
	if err := result.Decode(&record); err != nil {
		log.Fatal(err)
	}
	log.Println(record)
}

func readRecords() {
	c, err := database.Session()
	common.HandleError(err)
	if result, err := c.Find(ctx, bson.D{}); err == nil {
		for result.Next(ctx) {
			var record Record
			if err := result.Decode(&record); err != nil {
				log.Fatal(err)
			}
			log.Println(record)
		}
	} else {
		log.Fatal(err)
	}
}

func updateRecord() {
	c, err := database.Session()
	common.HandleError(err)
	update := bson.M{}
	update = bson.M{
		"$inc": bson.M{
			"type": 1,
		},
	}
	_, err = c.UpdateOne(ctx, bson.M{"name": "Updated"}, update)
}