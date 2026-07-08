package model

type BounceStats struct {
	AliasID   string `json:"alias_id"`
	AliasName string `json:"alias_name"`
	UserEmail string `json:"user_email"`
	Bounces   int64  `json:"bounces"`
}
