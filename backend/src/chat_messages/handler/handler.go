package handler

import (
	"github.com/aflores04/chat/backend/src/chat_messages/service"
	"net/http"
)

type ChatMessagesHandlerModule struct{}

func (*ChatMessagesHandlerModule) ProvideChatMessagesHandler(service service.ChatMessagesService) ChatMessagesHandler {
	return &chatMessagesHandler{
		service: service,
	}
}

type ChatMessagesHandler interface {
	GetMessages(w http.ResponseWriter, r *http.Request)
}

type chatMessagesHandler struct {
	service service.ChatMessagesService
}
