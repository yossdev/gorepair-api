package db

import (
	"context"
	"gorepair-rest-api/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB interface {
	DB() *mongo.Client
}

type mongoDB struct {
	client *mongo.Client
}

func NewMongoClient() MongoDB {
	var client *mongo.Client
	// Set client options
	clientOptions := options.Client().ApplyURI(config.Get().MongoDb_Address)

	// Connect to MongoDB
	var e error
	client, e = mongo.Connect(context.TODO(), clientOptions)
	if e != nil {
		log.Fatalln(e)
	}

	// Check the connection
	e = client.Ping(context.TODO(), nil)
	if e != nil {
		log.Fatalln(e)
	}

	return &mongoDB{
		client: client,
	}
}

func (c mongoDB) DB() *mongo.Client {
	return c.client
}