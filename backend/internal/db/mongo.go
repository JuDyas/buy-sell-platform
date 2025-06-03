package db

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"time"
)

type Mongo struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewMongo(uri, dbName string) *Mongo {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		log.Printf("failed to connect to mongo: %v", err)
		return nil
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Printf("failed to ping mongo: %v", err)
		return nil
	}

	db := client.Database(dbName)
	return &Mongo{
		Client: client,
		DB:     db,
	}
}
