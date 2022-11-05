package servercom

import (
	"encoding/json"
	"log"

	"github.com/gomodule/redigo/redis"
)

type JsonData struct {
	id    int
	date  string
	input string
	time  string
}

func SaveSubmittedData(name, date string, data interface{}) {

	if _, err := rh.JSONSet(name+date+"_json", ".", data); err != nil {
		log.Printf("error occured saving submitted data: %v", err)
	}
}

func RetrieveSubmittedData(name, date string) interface{} {
	var result []interface{}

	res, _ := redis.Bytes(rh.JSONGet(name+date+"_json", "."))
	// if err != nil {
	// 	log.Printf("Error retrieving data: %v", err)
	// }

	_ = json.Unmarshal(res, &result)
	// if err != nil {
	// 	log.Printf("%v", err)
	// }

	return result
}

func AcceptSubmittedData(name string) {
	_, err := client.Del(ctx, name+"_"+"json").Result()
	if err != nil {
		log.Fatalf("something went wrong deleting key: %v", err)
	}
}
