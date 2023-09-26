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
	createAccountSql := `INSERT INTO Account (id, hashed_password, email, email_verified, salt) VALUES ($1, $2, $3, $4, $5)`
	_, err := a.client.Exec(createAccountSql, account.Id, account.HashedPassword, account.Email, account.EmailVerified, account.Salt)
	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	var newAccount domain.Account
	getAccountSql := `SELECT * FROM Account WHERE id = $1`
	err = a.client.Get(&newAccount, getAccountSql, account.Id)
	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &newAccount, nil
}

func (a AuthRepositoryDB) CreateUser(user domain.User) (*domain.User, *errs.AppError) {
	createUserSql := `INSERT INTO User (id, account_id, username, image_url) VALUES ($1, $2, $3, $4)`
	_, err := a.client.Exec(createUserSql, user.Id, user.AccountId, user.Username, user.ImageUrl)
	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	var newUser domain.User
	getUserSql := `SELECT * FROM User WHERE id = $1`
	err = a.client.Get(&newUser, getUserSql, user.Id)
	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &newUser, nil
}

func NewAuthRepositoryDB(conn *sqlx.DB) AuthRepositoryDB {
	return AuthRepositoryDB{client: conn}
}
