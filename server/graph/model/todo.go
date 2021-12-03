package model

type Todo struct {
	ID   int64    `json:"id"`
	Text string    `json:"text"`
	Done bool      `json:"done"`
	UserID int64  `json:"user"`
}

