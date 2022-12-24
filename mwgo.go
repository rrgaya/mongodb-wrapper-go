package mwgo

import (
	"context"
	"log"
	"time"

	"github.com/rrgaya/mwgo/tools"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const OperationTimeOut = 5

func NewCollection(uriConnection, dbName, collectionName string) *mongo.Collection {
	Ctx, cancel := context.WithTimeout(context.Background(), OperationTimeOut*time.Second)
	defer cancel()
	client, err := mongo.Connect(Ctx, options.Client().ApplyURI(tools.URLConnection(uriConnection)))
	if err != nil {
		log.Panic(err)
	}
	log.Print("Connection Established")
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}
