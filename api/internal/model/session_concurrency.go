package model

type SessionConcurrency struct {
	UserID         string `json:"user_id"`
	Email          string `json:"email"`
	ActiveSessions int64  `json:"active_sessions"`
}
