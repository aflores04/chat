package domain

import (
	"time"
)

type Message struct {
	Username  *string    `json:"username" bson:"username"`
	RoomId    *string    `bson:"room_id" bson:"roomId"`
	Timestamp *time.Time `json:"timestamp" bson:"timestamp"`
	Body      *string    `json:"body" bson:"body"`
}
