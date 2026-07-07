package model

import "time"

// MaxInboxMessageSize is the upper bound for a stored raw MIME message (5 MB).
const MaxInboxMessageSize = 5 * 1024 * 1024

// MaxInboxMessagesPerAlias caps stored messages per inbox alias (oldest are evicted).
const MaxInboxMessagesPerAlias = 50

// InboxMessage stores a full incoming email for an inbox-type alias.
// Raw holds the unmodified MIME message; it is parsed on read.
type InboxMessage struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserID    string    `gorm:"index" json:"-"`
	AliasID   string    `gorm:"index" json:"alias_id"`
	From      string    `json:"from"`
	FromName  string    `json:"from_name"`
	Subject   string    `json:"subject"`
	Read      bool      `gorm:"default:false" json:"read"`
	Size      int       `json:"size"`
	Raw       []byte    `gorm:"type:mediumblob" json:"-"`
}

// RenderedMessage is a parsed, sanitized representation of an InboxMessage
// safe to display in a browser.
type RenderedMessage struct {
	ID          uint         `json:"id"`
	CreatedAt   time.Time    `json:"created_at"`
	From        string       `json:"from"`
	FromName    string       `json:"from_name"`
	Subject     string       `json:"subject"`
	HTML        string       `json:"html"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

// Attachment describes an attachment by name and size only; bodies are not served.
type Attachment struct {
	Name string `json:"name"`
	Size int    `json:"size"`
}
