package db

import (
	"context"
	"github.com/aflores04/chat/backend/src/chat_messages/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	messagesCollection = "messages"
)

func (r chatMessagesRepository) StoreMessage(ctx context.Context, message *domain.Message) (*domain.Message, error) {
	collection := r.client.MongoClient.Database(r.client.Database).Collection(messagesCollection)

	result, err := collection.InsertOne(ctx, message)
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(ctx, bson.M{"_id": result.InsertedID})
	if err != nil {
		return nil, err
	}

	_ = cursor.Decode(&message)

	return message, nil
}

func (r chatMessagesRepository) GetMessages(ctx context.Context, amount int64, order int64, sortKey string) ([]*domain.Message, error) {
	var messages []*domain.Message

	collection := r.client.MongoClient.Database(r.client.Database).Collection(messagesCollection)

	findOptions := &options.FindOptions{}
	findOptions.SetSort(bson.D{{sortKey, order}})
	findOptions.SetLimit(amount)

	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		message := &domain.Message{}

		_ = cursor.Decode(message)

		messages = append(messages, message)
	}

	return messages, nil
}
