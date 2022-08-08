package handler

import (
	"context"
	"github.com/AlekSi/pointer"
	"github.com/aflores04/chat/src/chat/domain"
	"github.com/aflores04/chat/src/chat/service"
	"github.com/aflores04/chat/src/websocket"
	"log"
)

type MessageHandlerModule struct{}

func (*MessageHandlerModule) ProvideMessageHandler(service service.ChatService) MessageHandler {
	return &messageHandler{chatService: service}
}

type MessageHandler interface {
	Handle(message websocket.Message)
}

type messageHandler struct {
	chatService service.ChatService
}

func (h messageHandler) Handle(wsMessage websocket.Message) {
	ctx := context.Background()

	message := transformWebsocketMessage(wsMessage)

	_, err := h.chatService.StoreMessage(ctx, message)
	if err != nil {
		log.Println("error storing message from handler", err)
	}
}

func transformWebsocketMessage(wsMessage websocket.Message) *domain.Message {
	message := &domain.Message{}
	message.Timestamp = pointer.ToTime(wsMessage.Payload.Timestamp)
	message.RoomId = pointer.ToString(wsMessage.Payload.RoomId)
	message.Username = pointer.ToString(wsMessage.Payload.Username)
	message.Body = pointer.ToString(wsMessage.Payload.Body)

	return message
}
