package model

import "time"

type OnboardingUser struct {
	UserID         string    `json:"user_id"`
	Email          string    `json:"email"`
	CreatedAt      time.Time `json:"created_at"`
	AliasCount     int64     `json:"alias_count"`
	RecipientCount int64     `json:"recipient_count"`
	Status         string    `json:"status"` // empty, no_aliases, no_recipients, complete
}
