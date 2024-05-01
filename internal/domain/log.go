package domain

import "time"

type Log struct {
	UserId    int       `json:"user_id" db:"user_id"`
	Message   string    `json:"message" db:"message"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type GetAdminLogsOut struct {
	Username  string    `json:"username" db:"username"`
	Message   string    `json:"message" db:"message"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
