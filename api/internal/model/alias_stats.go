package model

type AliasForwardStats struct {
	AliasID   string `json:"alias_id"`
	AliasName string `json:"alias_name"`
	UserEmail string `json:"user_email"`
	Forwards  int64  `json:"forwards"`
	Blocks    int64  `json:"blocks"`
	Replies   int64  `json:"replies"`
	Sends     int64  `json:"sends"`
}
