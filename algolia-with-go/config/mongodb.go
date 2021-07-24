package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConnection struct {
	DBLive *mongo.Client
}

func NewMongoDB() *MongoDBConnection {
	// create a new context
	ctx, cancel := NewMongoDBContextWithTimeout()
	defer cancel()
	// declare client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") //should be use env if not local dev
	// get connect
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Printf("error when get connection, cause: %v", err)
		return nil
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("error when test connection, cause: %v", err)
		return nil
	}
	log.Println("successfully connect to mongodb!")
	return &MongoDBConnection{
		DBLive: client,
	}
}

func NewMongoDBContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
