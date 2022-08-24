package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func InitMongoDb() {
	fmt.Println("Connecting to MongoDb...")
	var err error
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatalf("error occured when trying to connect to mongodb err = %v", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("error occured when trying to connect to mongodb err = %v", err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("Database error: %v", err)
	}
	defer client.Disconnect(ctx)

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(databases)
	fmt.Println("Successfully connected to mongoDB Server")
}
