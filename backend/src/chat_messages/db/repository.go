package db

import (
	"context"
	"github.com/aflores04/chat/backend/src/chat_messages/domain"
	"github.com/aflores04/chat/backend/src/mongodb"
)

type ChatRepository interface {
	StoreMessage(ctx context.Context, message *domain.Message) (*domain.Message, error)
}

type ChatRepositoryModule struct{}

type chatRepository struct {
	client *mongodb.MongoDB
}

func (*ChatRepositoryModule) ProvideUserRepository(client *mongodb.MongoDB) ChatRepository {
	return &chatRepository{
		client: client,
	}
}
