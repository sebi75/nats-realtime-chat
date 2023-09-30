package domain

import (
	"api/app/account/dto"
	"api/utils"
	"database/sql"
	"time"
)

type Account struct {
	Id             string         `json:"id" db:"id"`
	CreatedAt      time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt      sql.NullTime   `json:"updated_at" db:"updated_at"`
	LoginCount     int            `json:"login_count" db:"login_count"`
	LastLogin      sql.NullTime   `json:"last_login" db:"last_login"`
	LastIp         sql.NullString `json:"last_ip" db:"last_ip"`
	HashedPassword string         `json:"hashed_password" db:"hashed_password"`
	Salt           string         `json:"salt" db:"salt"`
	Email          string         `json:"email" db:"email"`
	EmailVerified  bool           `json:"email_verified" db:"email_verified"`
}

func (a Account) ToResponseDTO() *dto.AccountDTO {
	return &dto.AccountDTO{
		Id:            a.Id,
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     utils.NullTimeToPtr(a.UpdatedAt),
		Email:         a.Email,
		LastLogin:     utils.NullTimeToPtr(a.LastLogin),
		EmailVerified: a.EmailVerified,
	}
}
