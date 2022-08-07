package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

// Read messages from websocket connection and send them to all clients through pool entity
func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		payload := Payload{}

		err = json.Unmarshal(p, &payload)
		if err != nil {
			log.Println(err)
		}

		message := Message{Type: messageType, Payload: payload}
		c.Pool.Broadcast <- message
	}
}
