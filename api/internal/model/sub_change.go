package model

import "time"

type SubscriptionChange struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time `json:"created_at"`
	UserID     string    `json:"user_id" gorm:"index"`
	AdminEmail string    `json:"admin_email"`
	OldTier    string    `json:"old_tier"`
	NewTier    string    `json:"new_tier"`
	OldPlanID  *string   `json:"old_plan_id"`
	NewPlanID  *string   `json:"new_plan_id"`
	Reason     string    `json:"reason"`
}
