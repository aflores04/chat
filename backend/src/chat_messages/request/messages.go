package request

type ListMessagesRequest struct {
	Amount    int64  `json:"amount"`
	SortKey   string `json:"sort_key"`
	SortOrder int64  `json:"sort_order"`
}
