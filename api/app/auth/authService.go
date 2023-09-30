package auth

import (
	"api/app/auth/domain"
	"api/env"
	"api/errs"
	"api/utils/logger"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type AuthService struct {
	config *env.Config
}

func (as *AuthService) Login(signinRequest *domain.SigninRequest) (*domain.SigninResponse, *errs.AppError) {
	validateErr := signinRequest.Validate()
	if validateErr != nil {
		return nil, validateErr
	}
	authServiceUrl := as.config.AUTH.Url
	loginUrl := authServiceUrl + "/login"
	payload, err := json.Marshal(signinRequest)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}
	payloadReader := bytes.NewBuffer(payload)
	req, err := http.Post(loginUrl, "application/json", payloadReader)
	if err != nil {
		logger.Error(err.Error())
		if req.StatusCode == http.StatusBadRequest || req.StatusCode == http.StatusUnauthorized {
			// return the error message returned by auth-service
			return nil, errs.NewBadRequestError(err.Error())
		} else {
			return nil, errs.NewUnexpectedError("Unexpected error")
		}
	}
	defer req.Body.Close()

	var signinResponse domain.SigninResponse
	err = json.NewDecoder(req.Body).Decode(&signinResponse)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}
	return &signinResponse, nil
}

func (as *AuthService) Signup(body *io.ReadCloser) (*domain.SignupResponse, *errs.AppError) {
	authServiceUrl := as.config.AUTH.Url
	signupUrl := authServiceUrl + "/signup"
	req, err := http.Post(signupUrl, "application/json", *body)
	if req.StatusCode == http.StatusBadRequest {
		var res struct {
			Message string `json:"message"`
		}
		err := json.NewDecoder(req.Body).Decode(&res)
		if err != nil {
			return nil, errs.NewBadRequestError("Invalid auth service request body")
		}
		return nil, errs.NewBadRequestError(res.Message)
	}
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}
	defer req.Body.Close()

	var signupResponse domain.SignupResponse
	err = json.NewDecoder(req.Body).Decode(&signupResponse)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}

	return &signupResponse, nil
}

func (as AuthService) Verify(token string) (*domain.TokenPayload, *errs.AppError) {
	authServiceUrl := as.config.AUTH.Url
	verifyUrl := authServiceUrl + "/verify" + "?token=" + token
	req, err := http.Get(verifyUrl)
	if err != nil {
		if req.StatusCode == http.StatusBadRequest {
			// return the error message returned by the auth service
			return nil, errs.NewBadRequestError(err.Error())
		} else {
			return nil, errs.NewUnexpectedError("Unexpected error")
		}
	}
	defer req.Body.Close()

	var verifyResponse domain.TokenPayload
	err = json.NewDecoder(req.Body).Decode(&verifyResponse)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}

	return &verifyResponse, nil
}

func NewAuthService(config *env.Config) *AuthService {
	return &AuthService{
		config: config,
	}
}
