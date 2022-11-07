package servercom

import (
	"log"
)

type JsonData struct {
	id    int
	date  string
	input string
	time  string
}

func WriteLog(message, instructor string) {
	_, err := client.LPush(ctx, instructor+".log", message).Result()
	if err != nil {
		log.Printf("Could not save log: %v", err)
	}
}

func RetrieveLatestLogs(instructor string) []string {
	client.LTrim(ctx, instructor+".log", 0, 40)
	res, _ := client.LRange(ctx, instructor+".log", 0, 30).Result()

	return res
}
