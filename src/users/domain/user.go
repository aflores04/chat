package domain

type User struct {
	Username *string `json:"username" bson:"username,omitempty"`
	Email    *string `json:"email" bson:"email,omitempty"`
	Password *string `json:"password" bson:"password,omitempty"`
}
