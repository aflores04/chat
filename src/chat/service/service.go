package service

import (
	"context"
	"github.com/aflores04/chat/src/chat/db"
	"github.com/aflores04/chat/src/chat/domain"
	"github.com/aflores04/chat/src/rabbitmq"
)

type ChatServiceModule struct{}

func (*ChatServiceModule) ProvideChatService(
	rabbit rabbitmq.RabbitQueue,
	chatRepository db.ChatRepository) ChatService {
	return &chatService{
		rabbit: rabbit,
		repo:   chatRepository,
	}
}

type ChatService interface {
	StoreMessage(ctx context.Context, message *domain.Message) (*domain.Message, error)
}

type chatService struct {
	rabbit rabbitmq.RabbitQueue
	repo   db.ChatRepository
}
