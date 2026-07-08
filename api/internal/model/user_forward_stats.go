package model

type UserForwardStats struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Forwards int64  `json:"forwards"`
	Blocks   int64  `json:"blocks"`
	Replies  int64  `json:"replies"`
	Sends    int64  `json:"sends"`
}
