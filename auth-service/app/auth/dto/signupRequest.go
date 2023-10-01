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
	if len(sr.Username) > 0 && len(sr.Username) < 3 {
		return errs.NewValidationError("Username must be at least 3 characters")
	}
	if len(sr.Username) > 15 {
		return errs.NewValidationError("Username must be at most 15 characters")
	}
	if len(sr.Email) > 0 && len(sr.Email) < 5 {
		return errs.NewValidationError("Email must be at least 5 characters")
	}
	if len(sr.Email) > 50 {
		return errs.NewValidationError("Email must be at most 50 characters")
	}
	if len(sr.Password) < 5 {
		return errs.NewValidationError("Password must be at least 5 characters")
	}
	if len(sr.Password) > 35 {
		return errs.NewValidationError("Password must be at most 35 characters")
	}
	return nil
}
