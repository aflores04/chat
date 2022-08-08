package service

import (
	"context"
	"github.com/aflores04/chat/backend/src/chat_messages/db"
	"github.com/aflores04/chat/backend/src/chat_messages/domain"
)

type ChatServiceModule struct{}

func (*ChatServiceModule) ProvideChatService(
	chatRepository db.ChatRepository) ChatService {
	return &chatService{
		repo: chatRepository,
	}
}

type ChatService interface {
	StoreMessage(ctx context.Context, message *domain.Message) (*domain.Message, error)
}

type chatService struct {
	repo db.ChatRepository
}
