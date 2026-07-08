package model

type HourlyVolume struct {
	Hour  int   `json:"hour"`
	Count int64 `json:"count"`
}
