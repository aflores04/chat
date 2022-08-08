package websocket

import (
	"github.com/aflores04/chat/backend/src/chat_messages/domain"
)

type WebsocketMessage struct {
	Type    int            `json:"type"`
	Payload domain.Message `json:"payload"`
}
