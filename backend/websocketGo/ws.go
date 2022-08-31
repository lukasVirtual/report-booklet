package websocketGo

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Text string `json: "text"`
}

var upgrade = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrade.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Could not connect to Websocket: %v", err)
	}
	defer conn.Close()

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Could not read json: ", err.Error())
			break
		}

		fmt.Println("Message recieved: ", msg.Text)
	}
}
