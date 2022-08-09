package service

import (
	"context"
	"github.com/aflores04/chat/backend/src/chat_messages/domain"
	"github.com/aflores04/chat/backend/src/chat_messages/request"
	"github.com/aflores04/chat/backend/src/chat_messages/response"
)

const (
	defaultSortKey = "timestamp"
)

func (s *chatMessagesService) StoreMessage(ctx context.Context, message *domain.Message) (*domain.Message, error) {
	storedMessage, err := s.repo.StoreMessage(ctx, message)
	if err != nil {
		return nil, err
	}

	return storedMessage, nil
}

func (s *chatMessagesService) GetMessages(ctx context.Context, req *request.ListMessagesRequest) (*response.ListMessagesResponse, error) {
	messages, err := s.repo.GetMessages(ctx, req.Amount, req.SortOrder, req.SortKey)
	if err != nil {
		return nil, err
	}

	return &response.ListMessagesResponse{Messages: messages}, nil
}
