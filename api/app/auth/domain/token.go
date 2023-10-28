package domain

import "api/app/user/domain"

type TokenPayload struct {
	UserId    string `json:"userId"`
	AccountId string `json:"accountId"`
}

type VerifyResponse struct {
	domain.UserWithAccountDTO
}
