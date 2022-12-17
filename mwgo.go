package mwgo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const OperationTimeOut = 5

func NewCollection(uriConnection, dbName, collectionName string) *mongo.Collection {
	Ctx, _ := context.WithTimeout(context.Background(), OperationTimeOut*time.Second)
	client, err := mongo.Connect(Ctx, options.Client().ApplyURI(uriConnection))
	if err != nil {
		log.Panic(err)
	}
	log.Print("Connection Established")
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}
