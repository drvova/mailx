package model

type PlanUsage struct {
	UserID       string
	Email        string
	Tier         string
	AliasCount   int64
	MaxAliases   int64
	RecipientCount int64
	MaxRecipients  int64
	CredentialCount int64
	MaxCredentials  int64
	SessionCount   int64
	MaxSessions    int64
}
