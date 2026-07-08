package model

import "time"

// UserWithSub is a read-only projection for enriched CSV export.
type UserWithSub struct {
	ID           string     `json:"id"`
	Email        string     `json:"email"`
	IsActive     bool       `json:"is_active"`
	IsAdmin      bool       `json:"is_admin"`
	CreatedAt    time.Time  `json:"created_at"`
	Tier         string     `json:"tier"`
	SubType      string     `json:"sub_type"`
	SubActive    bool       `json:"sub_active"`
	ActiveUntil  *time.Time `json:"active_until"`
}
