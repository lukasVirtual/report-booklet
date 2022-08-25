package mongodb

import (
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertQualis(qualis []interface{}, date string) {
	s := strings.ReplaceAll(date, ".", "")
	var coll = client.Database("qualificationsdb").Collection(s)
	coll.Drop(ctx)

	_, err := coll.InsertMany(ctx, qualis)
	if err != nil {
		log.Println(err)
	}
}

func GetQualis(date string) []QualificationFormReturn {
	s := strings.ReplaceAll(date, ".", "")
	var coll = client.Database("qualificationsdb").Collection(s)

	var qualis []QualificationFormReturn
	cur, err := coll.Find(ctx, bson.M{})
	if err != nil {
		log.Println("could not find: ", err)
	}

	if err := cur.All(ctx, &qualis); err != nil {
		log.Fatalf("Could not get data: %v", err)
	}
	//fmt.Println(qualis)
	return qualis
}
