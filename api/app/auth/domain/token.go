package domain

type TokenPayload struct {
	UserId    string `json:"userId"`
	AccountId string `json:"accountId"`
}
