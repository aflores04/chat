package db

import (
	"context"
	"github.com/aflores04/chat/backend/src/chat_messages/domain"
	"github.com/aflores04/chat/backend/src/mongodb"
)

type ChatMessagesRepository interface {
	StoreMessage(ctx context.Context, message *domain.Message) (*domain.Message, error)
	GetMessages(ctx context.Context, amount int64, order int64, sortKey string) ([]*domain.Message, error)
}

type ChatMessagesRepositoryModule struct{}

type chatMessagesRepository struct {
	client *mongodb.MongoDB
}

func (*ChatMessagesRepositoryModule) ProvideUserRepository(client *mongodb.MongoDB) ChatMessagesRepository {
	return &chatMessagesRepository{
		client: client,
	}
}
