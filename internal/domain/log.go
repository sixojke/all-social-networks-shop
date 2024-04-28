package domain

import "time"

type Log struct {
	Message     string    `json:"message" db:"message"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
