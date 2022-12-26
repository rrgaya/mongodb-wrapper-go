package mwgo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestNewCollection(t *testing.T) {
	collectionTest, err := NewCollection("mongodb://admin:password123@localhost:27017", "test", "test")
	collectionTest.InsertOne(context.Background(), bson.D{})
	assert.Nil(t, err)
}
