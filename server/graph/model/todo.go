package model

import (
	"fmt"
	"github.com/boourns/dblib"
)

type Todo struct {
	ID   int64    `json:"id"`
	Text string    `json:"text"`
	Done bool      `json:"done"`
	UserID int64  `json:"user"`
}

func (t *Todo) User(db dblib.DBLike) (*User, error) {
	users, err := SelectUser(db, "WHERE ID = ?", t.UserID)
	if err != nil {
		return nil, err
	}
	if len(users) != 1 {
		return nil, fmt.Errorf("user not found: %d", t.UserID)
	}
	return users[0], nil
}
