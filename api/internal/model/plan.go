package model

type Plan struct {
	BaseModel
	Name              string `gorm:"uniqueIndex" json:"name"`
	DisplayName       string `json:"display_name"`
	PriceCents        int    `json:"price_cents"`
	Currency          string `gorm:"default:usd" json:"currency"`
	Interval          string `gorm:"default:monthly" json:"interval"` // monthly, yearly, one_time
	MaxRecipients     int    `json:"max_recipients"`
	MaxCredentials    int    `json:"max_credentials"`
	MaxDailyAliases   int    `json:"max_daily_aliases"`
	MaxDailySendReply int    `json:"max_daily_send_reply"`
	MaxSessions       int    `json:"max_sessions"`
	IsActive          bool   `gorm:"default:true" json:"is_active"`
	SortOrder         int    `gorm:"default:0" json:"sort_order"`
}
