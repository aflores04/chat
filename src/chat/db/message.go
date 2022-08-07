package db

import (
	"context"
	"github.com/aflores04/chat/src/chat/domain"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	messagesCollection = "messages"
)

func (r chatRepository) StoreMessage(ctx context.Context, message *domain.Message) (*domain.Message, error) {
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
