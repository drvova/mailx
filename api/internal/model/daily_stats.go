package model

type DailyStats struct {
	Date     string `json:"date"`
	Total    int64  `json:"total"`
	Forwards int64  `json:"forwards"`
	Blocks   int64  `json:"blocks"`
	Replies  int64  `json:"replies"`
	Sends    int64  `json:"sends"`
	Signups  int64  `json:"signups"`
}
