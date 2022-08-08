package service

import (
	"context"
	"github.com/aflores04/chat/src/chat/db"
	"github.com/aflores04/chat/src/chat/domain"
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
