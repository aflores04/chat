package db

import (
	"context"
	"github.com/aflores04/chat/src/mongodb"
	"github.com/aflores04/chat/src/users/domain"
)

type UserRepository interface {
	Login(ctx context.Context, username *string, password *string) *domain.User
	Register(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserBy(ctx context.Context, key string, value *string) (*domain.User, error)
}

type UserRepositoryModule struct{}

type userRepository struct {
	client *mongodb.MongoDB
}

func (*UserRepositoryModule) ProvideUserRepository(client *mongodb.MongoDB) UserRepository {
	return &userRepository{
		client: client,
	}
}
