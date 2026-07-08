package model

type UserQuota struct {
	UserID          string `json:"user_id"`
	Tier            string `json:"tier"`
	AliasCount      int64  `json:"alias_count"`
	RecipientCount  int64  `json:"recipient_count"`
	CredentialCount int64  `json:"credential_count"`
	SessionCount    int64  `json:"session_count"`
	MaxAliases      int64  `json:"max_aliases"`
	MaxRecipients   int64  `json:"max_recipients"`
	MaxCredentials  int64  `json:"max_credentials"`
	MaxSessions     int64  `json:"max_sessions"`
}
