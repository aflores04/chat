package service

import (
	"context"
	"github.com/aflores04/chat/src/users/errors"
	"github.com/aflores04/chat/src/users/helpers"
	"github.com/aflores04/chat/src/users/request"
	"github.com/aflores04/chat/src/users/response"
)

const (
	usernameKey = "username"
	emailKey    = "email"
)

func (s userService) RegisterUser(ctx context.Context, req *request.RegisterUserRequest) (*response.RegisterUserResponse, error) {
	req.User.Password = helpers.Hash(req.User.Password)

	isUsernameInDatabase, _ := s.repo.GetUserBy(ctx, usernameKey, req.User.Username)
	if isUsernameInDatabase != nil {
		return nil, &errors.UsernameAlreadyExistsError{}
	}

	isEmailInDatabase, _ := s.repo.GetUserBy(ctx, emailKey, req.User.Email)
	if isEmailInDatabase != nil {
		return nil, &errors.EmailAlreadyExistsError{}
	}

	user, err := s.repo.Register(ctx, req.User)
	if err != nil {
		return nil, err
	}

	return &response.RegisterUserResponse{
		User: user,
	}, nil
}

func (s userService) LoginAttempt(ctx context.Context, req *request.LoginRequest) bool {
	if s.repo.Login(ctx, req.Username, req.Password) {
		return true
	}

	return false
}
