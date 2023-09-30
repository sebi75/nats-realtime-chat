package auth

import (
	"api/app/auth/domain"
	"api/utils"
	"encoding/json"
	"net/http"
)

type AuthHandlers struct {
	service *AuthService
}

func (ah AuthHandlers) Signin(w http.ResponseWriter, request *http.Request) {
	var signinRequest domain.SigninRequest
	err := json.NewDecoder(request.Body).Decode(&signinRequest)
	if err != nil {
		utils.ResponseWriter(w, http.StatusBadRequest, err.Error())
		return
	}
	signinResponse, appErr := ah.service.Login(&signinRequest)
	if appErr != nil {
		utils.ResponseWriter(w, appErr.Code, appErr.Message)
		return
	}
	utils.ResponseWriter(w, http.StatusOK, signinResponse)
	return
}

func (ah AuthHandlers) Signup(w http.ResponseWriter, request *http.Request) {
	signupResponse, serviceErr := ah.service.Signup(&request.Body)
	if serviceErr != nil {
		utils.ResponseWriter(w, serviceErr.Code, serviceErr.Message)
		return
	}

	utils.ResponseWriter(w, http.StatusOK, signupResponse)
	return
}

func (ah AuthHandlers) Verify(w http.ResponseWriter, request *http.Request) {
	token := request.Header.Get("Authorization")[7:]
	if token == "" {
		utils.ResponseWriter(w, http.StatusBadRequest, "Token is required")
		return
	}
	verifyResponse, serviceErr := ah.service.Verify(token)
	if serviceErr != nil {
		utils.ResponseWriter(w, serviceErr.Code, serviceErr.Message)
		return
	}

	utils.ResponseWriter(w, http.StatusOK, verifyResponse)
	return
}

func NewAuthHandlers(service *AuthService) *AuthHandlers {
	return &AuthHandlers{
		service: service,
	}
}
