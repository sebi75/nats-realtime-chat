package domain

type Account struct {
	Id             string `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	LoginCount     int    `json:"login_count"`
	LastLogin      string `json:"last_login"`
	LastIp         string `json:"last_ip"`
	HashedPassword string `json:"hashed_password"`
	Salt           string `json:"salt"`
	Email          string `json:"email"`
	EmailVerified  bool   `json:"email_verified"`
}
