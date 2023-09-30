package dto

import "auth-service/errs"

type SigninRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

func (sr *SigninRequest) Validate() *errs.AppError {
	if len(sr.Username) > 0 && len(sr.Email) > 0 {
		return errs.NewValidationError("Provide either username or email, not both")
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
	// password required in any case
	if len(sr.Password) < 5 {
		return errs.NewValidationError("Password must be at least 5 characters")
	}
	if len(sr.Password) > 35 {
		return errs.NewValidationError("Password must be at most 35 characters")
	}
	return nil
}
