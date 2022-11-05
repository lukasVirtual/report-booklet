package serverfiles

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func SaveSubmittedData(name, date string, data []interface{}) {
	coll := client.Database("serverfiles").Collection(name + date + "_json")
	coll.Drop(ctx)
	_, err := coll.InsertMany(ctx, data)
	if err != nil {
		log.Println(err)
	}
}

func RetrieveSubmittedData(name, date string) interface{} {
	var result []ReturnData
	coll := client.Database("serverfiles").Collection(name + date + "_json")

	curr, err := coll.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}

	if err := curr.All(ctx, &result); err != nil {
		log.Fatalf("Something went wrong retrieving data: %v", err)
	}
	return result
}
