package domain

import (
	userDomain "api/app/user/domain"
)

type Friend struct {
	Id          string `db:"id"`
	RequesterId string `db:"requester_id" json:"requester_id"`
	AddresseeId string `db:"addressee_id" json:"addressee_id"`
	CreatedAt   string `db:"created_at" json:"created_at"`
	Status      string `db:"status" json:"status"`
}

type FriendWithUser struct {
	Friend
	User *userDomain.User `json:"user" db:"user"`
}
