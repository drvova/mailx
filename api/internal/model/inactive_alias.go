package model

import "time"

type InactiveAlias struct {
	AliasID      string    `json:"alias_id"`
	AliasName    string    `json:"alias_name"`
	UserID       string    `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	DaysInactive int       `json:"days_inactive"`
}
