package mwgo

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func skipCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}

func TestNewCollection(t *testing.T) {
	skipCI(t)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	collectionTest, err := NewCollection("mongodb://admin:password123@localhost:27017", "test", "test")
	collectionTest.InsertOne(ctx, bson.D{})
	assert.Nil(t, err)
}
