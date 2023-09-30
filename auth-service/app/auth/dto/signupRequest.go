package dto

import "auth-service/errs"

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (sr *SignupRequest) Validate() *errs.AppError {
	if sr.Username == "" {
		return errs.NewValidationError("Username is required")
	}
	if sr.Password == "" {
		return errs.NewValidationError("Password is required")
	}
	if sr.Email == "" {
		return errs.NewValidationError("Email is required")
	}
	return nil
}
