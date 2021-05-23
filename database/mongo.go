package database

import (
	"context"
	"fmt"
	"time"

	painterconfig "github.com/just-arun/painter-config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetMongoInstance() *mongo.Client {
	// Replace the uri string with your MongoDB deployment's connection string.
	uri := painterconfig.AppConfig.Database.Mongo.MongoURI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("MongoDB Connected established...")
	return client
}
