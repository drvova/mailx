package model

type AliasTrend struct {
	Date     string `json:"date"`
	Forwards int64  `json:"forwards"`
	Blocks   int64  `json:"blocks"`
}
