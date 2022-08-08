package websocket

import (
	"encoding/json"
	"github.com/AlekSi/pointer"
	"github.com/aflores04/chat/backend/src/chat_messages/domain"
	"github.com/gorilla/websocket"
	"log"
	"time"
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

		payload := domain.Message{}
		payload.Timestamp = pointer.ToTime(time.Now())

		err = json.Unmarshal(p, &payload)
		if err != nil {
			log.Println(err)
		}

		message := WebsocketMessage{Type: messageType, Payload: payload}
		c.Pool.Broadcast <- message
	}
}
