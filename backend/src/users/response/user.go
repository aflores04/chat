package response

import (
	"github.com/aflores04/chat/backend/src/users/domain"
)

type RegisterUserResponse struct {
	User *domain.User `json:"user"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
