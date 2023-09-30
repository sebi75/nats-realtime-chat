package dto

import "time"

type AccountDTO struct {
	Id            string     `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	Email         string     `json:"email"`
	LastLogin     *time.Time `json:"last_login"`
	EmailVerified bool       `json:"email_verified"`
}
