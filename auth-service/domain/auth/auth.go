package domain

type AuthRepository interface {
	Signup(user *User) error
	Signin(user *User) error
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
