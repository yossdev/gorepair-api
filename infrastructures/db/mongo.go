package db

import (
	"context"
	"gorepair-rest-api/config"
	"log"
	"time"

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
	// clientOptions := options.Client().ApplyURI(config.Get().MongoDb_Address) //for local connection
	clientOptions := options.Client().ApplyURI("mongodb+srv://"+config.Get().MongoDb_Username+":"+config.Get().MongoDb_Password+"@cluster0.atngo.mongodb.net/"+config.Get().MongoDb_Name+"?retryWrites=true&w=majority")

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	e := client.Ping(ctx, nil)
	if e != nil {
		log.Fatalln(e)
	}

	// log.Println("Connected to MongoDB!")

	return &mongoDB{
		client: client,
	}
}

func (c mongoDB) DB() *mongo.Client {
	return c.client
}