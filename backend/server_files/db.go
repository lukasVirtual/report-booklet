package serverfiles

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
var ctx, _ = context.WithTimeout(context.Background(), 24*time.Hour)

type ReturnData struct {
	Id    int    `bson: "id, omitempty"`
	Date  string `bson: "date, omitempty"`
	Input string `bson: "input, omitempty"`
	Time  string `bson: "time, omitempty"`
	Rows  int    `bson: "rows, omitempty"`
}

func InitMongoDb() {
	fmt.Println("Connecting to MongoDb...")
	var err error
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
	// defer client.Disconnect(ctx)

	_, err = client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to mongoDB Server")
}
