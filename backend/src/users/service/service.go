package service

import (
	"context"
	"github.com/aflores04/chat/backend/src/jwt"
	"github.com/aflores04/chat/backend/src/users/db"
	request2 "github.com/aflores04/chat/backend/src/users/request"
	"github.com/aflores04/chat/backend/src/users/response"
)

type UserServiceModule struct{}

type UserService interface {
	RegisterUser(ctx context.Context, req *request2.RegisterUserRequest) (*response.RegisterUserResponse, error)
	LoginAttempt(ctx context.Context, req *request2.LoginRequest) *response.LoginResponse
}

type userService struct {
	repo      db.UserRepository
	jwtClient jwt.JwtClient
}

func (*UserServiceModule) ProvideUserService(repo db.UserRepository, jwtClient jwt.JwtClient) UserService {
	return &userService{
		repo:      repo,
		jwtClient: jwtClient,
	}
}
