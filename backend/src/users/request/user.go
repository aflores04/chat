package request

import (
	"github.com/aflores04/chat/backend/src/users/domain"
)

type RegisterUserRequest struct {
	User *domain.User `json:"user"`
}
