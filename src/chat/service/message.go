package service

import (
	"context"
	"encoding/json"
	"github.com/aflores04/chat/src/chat/domain"
)

const messagesQueue = "messages"

func (s *chatService) StoreMessage(ctx context.Context, message *domain.Message) (*domain.Message, error) {
	storedMessage, err := s.repo.StoreMessage(ctx, message)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	queue := s.rabbit.CreateQueue(messagesQueue)
	s.rabbit.Publish(ctx, queue, string(b))

	return storedMessage, nil
}
