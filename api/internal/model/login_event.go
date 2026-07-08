package model

import "time"

type LoginEvent struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UserID    string    `json:"user_id" gorm:"index"`
	Success   bool      `json:"success"`
	IP        string    `json:"ip"`
}
