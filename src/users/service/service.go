package service

import (
	"context"
	"github.com/aflores04/chat/src/users/db"
	"github.com/aflores04/chat/src/users/request"
	"github.com/aflores04/chat/src/users/response"
)

type UserServiceModule struct{}

type UserService interface {
	RegisterUser(ctx context.Context, req *request.RegisterUserRequest) (*response.RegisterUserResponse, error)
	LoginAttempt(ctx context.Context, req *request.LoginRequest) bool
}

type userService struct {
	repo db.UserRepository
}

func (*UserServiceModule) ProvideUserService(repo db.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}
