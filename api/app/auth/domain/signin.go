package domain

import "api/errs"

type SigninRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

func (sr *SigninRequest) Validate() *errs.AppError {
	if sr.Username == "" {
		return errs.NewValidationError("Username is required")
	}

	if sr.Password == "" {
		return errs.NewValidationError("Password is required")
	}

	return nil
}

type SigninResponse struct {
	Token string `json:"token,omitempty"`
}
