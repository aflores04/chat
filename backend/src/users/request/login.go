package request

type LoginRequest struct {
	Username *string `json:"username" bson:"username,omitempty"`
	Password *string `json:"password" bson:"password,omitempty"`
}
