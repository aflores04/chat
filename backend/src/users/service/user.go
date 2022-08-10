package service

import (
	"context"
	"github.com/aflores04/chat/backend/src/users/errors"
	"github.com/aflores04/chat/backend/src/users/helpers"
	"github.com/aflores04/chat/backend/src/users/request"
	"github.com/aflores04/chat/backend/src/users/response"
)

const (
	usernameKey = "username"
)

func (s userService) RegisterUser(ctx context.Context, req *request.RegisterUserRequest) (*response.RegisterUserResponse, error) {
	req.User.Password = helpers.Hash(req.User.Password)

	isUsernameInDatabase, _ := s.repo.GetUserBy(ctx, usernameKey, req.User.Username)
	if isUsernameInDatabase != nil {
		return nil, &errors.UsernameAlreadyExistsError{}
	}

	user, err := s.repo.Register(ctx, req.User)
	if err != nil {
		return nil, err
	}

	return &response.RegisterUserResponse{
		User: user,
	}, nil
}

func (s userService) LoginAttempt(ctx context.Context, req *request.LoginRequest) *response.LoginResponse {
	user := s.repo.Login(ctx, req.Username, req.Password)
	if user == nil {
		return nil
	}

	jwtToken := s.jwtClient.CreateJWT(map[string]interface{}{usernameKey: user.Username})

	return &response.LoginResponse{Token: jwtToken}
}
