package service

import (
	"context"
	"github.com/aflores04/chat/backend/src/chat_messages/db"
	"github.com/aflores04/chat/backend/src/chat_messages/domain"
	"github.com/aflores04/chat/backend/src/chat_messages/request"
	"github.com/aflores04/chat/backend/src/chat_messages/response"
)

type ChatMessagesServiceModule struct{}

func (*ChatMessagesServiceModule) ProvideChatService(
	chatRepository db.ChatMessagesRepository) ChatMessagesService {
	return &chatMessagesService{
		repo: chatRepository,
	}
}

type ChatMessagesService interface {
	StoreMessage(ctx context.Context, message *domain.Message) (*domain.Message, error)
	GetMessages(ctx context.Context, req *request.ListMessagesRequest) (*response.ListMessagesResponse, error)
}

type chatMessagesService struct {
	repo db.ChatMessagesRepository
}
