package server

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

func SendToInstructor(name string, c *websocket.Conn) {
	var (
		mt  int
		msg []byte
		err error
	)
	for client := range clients {
		if client.Name == name {
			for {
				if mt, msg, err = c.ReadMessage(); err != nil {
					log.Println("read: ", err)
					break
				}

				log.Printf("message recieved: %s", msg)

				if err := c.WriteMessage(mt, msg); err != nil {
					log.Println("write:", err)
					break
				}
			}
		} else {
			log.Println("No messages for other users")
		}
	}
}
