package auth

import (
	"auth-service/app/auth/domain"
	"auth-service/errs"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type AuthRepositoryDB struct {
	client *sqlx.DB
}

func (a AuthRepositoryDB) FindUserById(id string) (*domain.User, *errs.AppError) {
	var user domain.User
	getUserSql := `SELECT * FROM User WHERE id = $1`
	err := a.client.Get(&user, getUserSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("User not found")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &user, nil
}

func (a AuthRepositoryDB) FindUserByEmail(email string) (*domain.User, *errs.AppError) {
	var user domain.User
	getUserSql := `SELECT u.* FROM User u
	 			JOIN Account ac ON u.account_id = ac.id
				WHERE ac.email = $1`
	err := a.client.Get(&user, getUserSql, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("User not found")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &user, nil
}

func (a AuthRepositoryDB) FindAccountById(id string) (*domain.Account, *errs.AppError) {
	var account domain.Account
	getAccountSql := `SELECT * FROM Account WHERE id = $1`
	err := a.client.Get(&account, getAccountSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Account not found")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &account, nil
}

func (a AuthRepositoryDB) FindAccountByEmail(email string) (*domain.Account, *errs.AppError) {
	var account domain.Account
	getAccountSql := `SELECT * FROM Account WHERE email = $1`
	err := a.client.Get(&account, getAccountSql, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Account not found")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &account, nil
}

func (a AuthRepositoryDB) FindUserByUsername(username string) (*domain.User, *errs.AppError) {
	var user domain.User
	getUserSql := `SELECT * FROM User WHERE username = $1`
	err := a.client.Get(&user, getUserSql, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("User not found")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &user, nil
}

func (a AuthRepositoryDB) CreateAccount(account domain.Account) (*domain.Account, *errs.AppError) {
	// implement, good night
	return nil, nil
}

func NewAuthRepositoryDB(conn *sqlx.DB) AuthRepositoryDB {
	return AuthRepositoryDB{client: conn}
}
