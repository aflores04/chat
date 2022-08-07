package websocket

type Payload struct {
	RoomId   int64  `json:"room_id"`
	Username string `json:"username"`
	Body     string `json:"body"`
}

type Message struct {
	Type    int     `json:"type"`
	Payload Payload `json:"payload"`
}
