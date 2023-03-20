package domain

import "time"

type (
	Qoute struct {
		ID        string `db:"id" json:"id"`
		Body      string `db:"body" json:"body"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
