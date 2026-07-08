package model

type DomainStats struct {
	Domain     string `json:"domain"`
	Enabled    bool   `json:"enabled"`
	Verified   bool   `json:"verified"`
	AliasCount int64  `json:"alias_count"`
}
