package models

type Bill struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	DueDateDay int64   `json:"due_date_day"`
	Amount     float64 `json:"amount"`
	IsAutoPay  bool    `json:"is_autopay"`
}
