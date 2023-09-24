package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDb struct {
	Client *mongo.Database
}

func NewMongoDb() (*MongoDb, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	cleanup := func() {
		defer cancel()
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}
	return &MongoDb{Client: client.Database("portfolio")}, cleanup
}
