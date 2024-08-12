package models

import "time"

type Event struct {
	ID        string
	Name      string
	Location  string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
