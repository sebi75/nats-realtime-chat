package domain

import "auth-service/app/auth/dto"

type UserWithAccountDTO struct {
	User
	Account dto.AccountDTO `json:"account"`
}

type UserWithAccount struct {
	User
	Account Account `json:"account"`
}

type User struct {
	Id        string `json:"id" db:"id"`
	AccountId string `json:"account_id" db:"account_id"`
	Username  string `json:"username" db:"username"`
	ImageUrl  string `json:"image_url" db:"image_url"`
}
