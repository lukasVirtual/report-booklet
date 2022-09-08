package server

import "github.com/gofiber/websocket/v2"

func SendToInstructor(userID int, c *websocket.Conn, message []byte) {
	for client := range clients {
		if client.ID == uint64(userID) {
			c.WriteMessage(1, message)
		}
	}
}
