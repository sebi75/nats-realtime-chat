package dto

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// todo: schema validation for each input request
