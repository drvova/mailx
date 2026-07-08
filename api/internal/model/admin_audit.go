package model

import "time"

type AdminAudit struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time `json:"created_at"`
	AdminEmail string    `json:"admin_email" gorm:"index"`
	Action     string    `json:"action" gorm:"index"`
	Target     string    `json:"target"`
	Details    string    `json:"details"`
}
