package mongodb

import (
	"fmt"
	"log"

	"github.com/iods/go-pherit/internal/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Record struct {
	Type        string `bson:"type"`
	Name 		string `bson:"name"`
	Description string `bson:"description"`
}

func mainInit() {
	insertRecord(Record{"One", "One", "This is the first record in the db."})
	insertRecords()
	readRecord()
	updateRecord()
	updateRecords()
	replaceRecord()
	readRecords()
	deleteRecord()
	readRecords()
	deleteRecords()
	readRecords()
}

func insertRecord(one Record) {
	c, err := Session()
	common.HandleError(err)
	result, err := c.InsertOne(ctx, one)
	common.HandleError(err)
	fmt.Println("You have inserted one record:", result.InsertedID)
}

func insertRecords() {
	c, err := Session()
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
	c, err := Session()
	common.HandleError(err)
	result := c.FindOne(ctx, filter)
	if err := result.Decode(&record); err != nil {
		log.Fatal(err)
	}
	log.Println(record)
}

func readRecords() {
	c, err := Session()
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
	c, err := Session()
	common.HandleError(err)
	filter := bson.M{"type": "Two"}
	update := bson.M{}
	update = bson.M{
		"$set": bson.M{
			"name": "Changed",
		},
	}
	if result, err := c.UpdateOne(ctx, filter, update); err == nil {
		log.Println(result)
	} else {
		log.Fatal(err)
	}
}

func updateRecords() {
	c, err := Session()
	common.HandleError(err)
	filter := bson.M{"type": primitive.Regex{Pattern: "Three", Options: ""}}
	update := bson.M{}
	update = bson.M{
		"$set": bson.M{
			"type": "Updated",
		},
	}
	if result, err := c.UpdateMany(ctx, filter, update); err == nil {
		log.Println(result)
	} else {
		log.Fatal(err)
	}
}

func replaceRecord() {
	record := Record{"Replaced", "One", "This is the first record replaced by another."}
	c, err := Session()
	common.HandleError(err)
	if result, err := c.ReplaceOne(ctx, bson.M{"type": "One"}, record); err == nil {
		log.Println(result)
	} else {
		log.Fatal(err)
	}
}

func deleteRecord() {
	c, err := Session()
	common.HandleError(err)
	if result, err := c.DeleteOne(ctx, bson.M{"name": "One"}); err == nil {
		log.Println(result)
	} else {
		log.Fatal(err)
	}
}

func deleteRecords() {
	c, err := Session()
	common.HandleError(err)
	if result, err := c.DeleteMany(ctx, bson.M{"name": primitive.Regex{Pattern: "Three", Options: ""}}); err == nil {
		log.Println(result)
	} else {
		log.Fatal(err)
	}
}
