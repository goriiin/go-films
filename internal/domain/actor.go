package domain

import "time"

type Actor struct {
	ID          string
	Name        string
	Birthday    time.Time
	Description string
	Gender      string
}
