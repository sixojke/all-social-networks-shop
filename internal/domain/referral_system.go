package domain

import "time"

type ReferralSystem struct {
	ReferralCode  string    `json:"referral_code" db:"referral_code"`
	TotalVisitors int       `json:"total_visitors" db:"total_visitors"`
	Description   string    `json:"description" db:"description"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}
