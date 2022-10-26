package serverfiles

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

type JsonData struct {
	id    int
	date  string
	input string
	time  string
}

func SaveSubmittedData(name string, data interface{}) {
	fmt.Println(name)
	if _, err := rh.JSONSet(name+"_json", ".", data); err != nil {
		log.Printf("error occured saving submitted data: %v", err)
	}
}

func RetrieveSubmittedData(name string) interface{} {
	var result interface{}
	fmt.Println(name)
	res, err := redis.Bytes(rh.JSONGet(name+"_json", "."))
	if err != nil {
		log.Printf("Error retrieving data: %v", err)
	}
	err = json.Unmarshal(res, &result)
	if err != nil {
		log.Printf("%v", err)
	}

	return result
}

func AcceptSubmittedData(name string) {
	_, err := client.Del(ctx, name+"_"+"json").Result()
	if err != nil {
		log.Fatalf("something went wrong deleting key: %v", err)
	}
}
