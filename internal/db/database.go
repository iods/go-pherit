package db

import (
	"context"

	"github.com/iods/go-pherit/internal/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx = context.TODO()
	mongoURI = "mongodb://localhost:27017"
)

func connect() (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	errors.HandleError(err)
	err = client.Ping(ctx, readpref.Primary())
	errors.HandleError(err)
	return client, nil
}

func Session() (*mongo.Collection, error) {
	client, err := connect()
	errors.HandleError(err)
	session := client.Database("testdb").Collection("records")
	return session, nil
}