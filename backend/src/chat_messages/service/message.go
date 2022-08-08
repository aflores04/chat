package service

import (
	"context"
	"github.com/aflores04/chat/src/chat/domain"
)

func (s *chatService) StoreMessage(ctx context.Context, message *domain.Message) (*domain.Message, error) {
	storedMessage, err := s.repo.StoreMessage(ctx, message)
	if err != nil {
		return nil, err
	}

	return storedMessage, nil
}
