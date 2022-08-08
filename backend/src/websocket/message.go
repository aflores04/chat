package websocket

import (
	"github.com/aflores04/chat/src/chat/domain"
)

type WebsocketMessage struct {
	Type    int            `json:"type"`
	Payload domain.Message `json:"payload"`
}
