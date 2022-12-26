package mwgo

import (
	"context"
	"log"
	"time"

	"github.com/rrgaya/mongodb-wrapper-go/tools"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const OperationTimeOut = 5

func NewCollection(uriConnection, dbName, collectionName string) (*mongo.Collection, error) {
	Ctx, cancel := context.WithTimeout(context.Background(), OperationTimeOut*time.Second)
	defer cancel()
	client, err := mongo.Connect(Ctx, options.Client().ApplyURI(tools.URLConnection(uriConnection)))
	if err != nil {
		return nil, err
	}
	collection := client.Database(dbName).Collection(collectionName)
	log.Print("Connection Established")
	return collection, nil
}
