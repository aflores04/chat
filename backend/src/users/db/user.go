package db

import (
	"context"
	"github.com/aflores04/chat/src/users/domain"
	"github.com/aflores04/chat/src/users/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

const (
	usersCollection = "users"

	usernameKey = "username"
	passwordKey = "password"
)

func (r userRepository) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	collection := r.client.MongoClient.Database(r.client.Database).Collection(usersCollection)

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(ctx, bson.M{"_id": result.InsertedID})
	if err != nil {
		return nil, err
	}

	_ = cursor.Decode(&user)

	return user, nil
}

// GetUserBy find user by available keys [username]
func (r *userRepository) GetUserBy(ctx context.Context, key string, value *string) (*domain.User, error) {
	user := &domain.User{}

	collection := r.client.MongoClient.Database(r.client.Database).Collection(usersCollection)

	result := collection.FindOne(ctx, bson.M{key: value})

	err := result.Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Login(ctx context.Context, username *string, password *string) *domain.User {
	user := &domain.User{}

	collection := r.client.MongoClient.Database(r.client.Database).Collection(usersCollection)

	hashedPassword := helpers.Hash(password)
	result := collection.FindOne(ctx, bson.M{usernameKey: username, passwordKey: hashedPassword})

	err := result.Decode(user)
	if err != nil {
		log.Println("error decoding user login attempt")
		return nil
	}

	if user != nil {
		return user
	}

	return nil
}
