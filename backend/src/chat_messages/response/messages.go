package response

import "github.com/aflores04/chat/backend/src/chat_messages/domain"

type ListMessagesResponse struct {
	Messages []*domain.Message `json:"messages"`
}
