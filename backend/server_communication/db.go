package servercom

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
)

var (
	client *redis.Client
	rh     *rejson.Handler
	ctx    = context.Background()
)

func InitRedis() {
	fmt.Println(":: Init Redis Database")

	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	rh = rejson.NewReJSONHandler()
	rh.SetGoRedisClient(client)
	status, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error Connecting to DB")
	}
	fmt.Println(status)
}
