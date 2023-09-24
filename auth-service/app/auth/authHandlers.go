package auth

import (
	"auth-service/app/auth/dto"
	"auth-service/errs"
	"encoding/json"
	"net/http"
)

type AuthHandlers struct {
	service AuthService
}

func (ah AuthHandlers) Signup(w http.ResponseWriter, request *http.Request) {
	var signupRequest dto.SignupRequest
	err := json.NewDecoder(request.Body).Decode(&signupRequest)
	if err != nil {
		errs.NewBadRequestError("Invalid request body")
	}
	// newUser = ah.service.Signup(&signupRequest)
}

func NewAuthHandlers(authService AuthService) *AuthHandlers {
	return &AuthHandlers{
		service: authService,
	}
}
