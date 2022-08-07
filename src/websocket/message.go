package websocket

import "time"

type Payload struct {
	RoomId    string    `json:"room_id"`
	Username  string    `json:"username"`
	Body      string    `json:"body"`
	Timestamp time.Time `json:"timestamp"`
}

type Message struct {
	Type    int     `json:"type"`
	Payload Payload `json:"payload"`
}
