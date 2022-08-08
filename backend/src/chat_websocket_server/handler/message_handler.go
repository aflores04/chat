package handler

import (
	"context"
	"github.com/aflores04/chat/backend/src/chat_messages/service"
	"github.com/aflores04/chat/backend/src/websocket"
	"log"
)

type MessageHandlerModule struct{}

func (*MessageHandlerModule) ProvideMessageHandler(service service.ChatService) MessageHandler {
	return &messageHandler{chatService: service}
}

type MessageHandler interface {
	Handle(message websocket.WebsocketMessage)
}

type messageHandler struct {
	chatService service.ChatService
}

func (h messageHandler) Handle(wsMessage websocket.WebsocketMessage) {
	ctx := context.Background()

	_, err := h.chatService.StoreMessage(ctx, &wsMessage.Payload)
	if err != nil {
		log.Println("error storing message from handler", err)
	}
}