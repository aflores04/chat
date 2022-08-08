package request

import "github.com/aflores04/chat/src/users/domain"

type RegisterUserRequest struct {
	User *domain.User `json:"user"`
}
