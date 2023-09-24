package domain

type AuthRepository interface {
	Signup(user *User) error
	Signin(user *User) error
}
