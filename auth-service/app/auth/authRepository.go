package auth

import (
	"auth-service/app/auth/domain"

	"github.com/jmoiron/sqlx"
)

type AuthRepositoryDB struct {
	client *sqlx.DB
}

func (a AuthRepositoryDB) Signup(user *domain.User) (*domain.User, error) {
	return nil, nil
}

func NewAuthRepositoryDB(conn *sqlx.DB) AuthRepositoryDB {
	return AuthRepositoryDB{client: conn}
}
