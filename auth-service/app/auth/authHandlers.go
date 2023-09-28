package auth

import (
	"auth-service/app/auth/dto"
	"auth-service/errs"
	"auth-service/utils"
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
		resErrr := errs.NewBadRequestError("Invalid request body")
		utils.ResponseWriter(w, resErrr.Code, resErrr.AsMessage())
		return
	}
	newUser, serviceErr := ah.service.Signup(&signupRequest)
	if serviceErr != nil {
		utils.ResponseWriter(w, serviceErr.Code, serviceErr.AsMessage())
		return
	}
	utils.ResponseWriter(w, http.StatusCreated, newUser)
	return
}

func (ah AuthHandlers) Signin(w http.ResponseWriter, request *http.Request) {
	var signinRequest dto.SigninRequest
	err := json.NewDecoder(request.Body).Decode(&signinRequest)
	if err != nil {
		resErr := errs.NewBadRequestError("Invalid request body")
		utils.ResponseWriter(w, resErr.Code, resErr.AsMessage())
		return
	}
	token, serviceErr := ah.service.Signin(&signinRequest)
	if serviceErr != nil {
		utils.ResponseWriter(w, serviceErr.Code, serviceErr.AsMessage())
		return
	}
	utils.ResponseWriter(w, http.StatusOK, token)
	return
}

func NewAuthHandlers(authService AuthService) *AuthHandlers {
	return &AuthHandlers{
		service: authService,
	}
}
