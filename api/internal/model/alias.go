package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrDuplicateAlias       = errors.New("alias already exists")
	ErrDuplicateAliasDomain = errors.New("wildcard aliases limit reached for this domain")
)

type AliasType int

const (
	AliasRelay AliasType = 0
	AliasInbox AliasType = 1
)

type Alias struct {
	BaseModel
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	Name             string         `gorm:"unique" json:"name"`
	UserID           string         `json:"-"`
	Type             AliasType      `gorm:"default:0" json:"type"`
	ExpiresAt        *time.Time     `json:"expires_at,omitempty"`
	Enabled          bool           `json:"enabled"`
	Description      string         `gorm:"default:''" json:"description"`
	Recipients       string         `gorm:"default:''" json:"recipients"`
	FromName         string         `gorm:"default:''" json:"from_name"`
	CatchAll         bool           `json:"catch_all"`
	Stats            AliasStats     `gorm:"-" json:"stats"`
	IsCustomDomain   bool           `gorm:"-" json:"is_custom_domain"`
	IsDomainVerified *bool          `gorm:"-" json:"is_domain_verified"`
	IsDomainEnabled  bool           `gorm:"-" json:"is_domain_enabled"`
}

type AliasStats struct {
	Forwards int `json:"forwards"`
	Blocks   int `json:"blocks"`
	Replies  int `json:"replies"`
	Sends    int `json:"sends"`
}

type AliasList struct {
	Aliases []Alias `json:"aliases"`
	Total   int     `json:"total"`
}
