package domain

import "github.com/jmoiron/sqlx"

type AuthRepositoryDB struct {
	client *sqlx.DB
}

func (a AuthRepositoryDB) Signup(user *User) (*User, error) {
	//implementation
	return nil, nil
}

func NewAuthRepositoryDB(conn *sqlx.DB) AuthRepositoryDB {
	return AuthRepositoryDB{client: conn}
}
