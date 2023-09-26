package auth

import (
	"auth-service/app/auth/domain"
	"auth-service/app/auth/dto"
	"auth-service/app/internal/encrypt"
	"auth-service/errs"
	"auth-service/utils"
	"auth-service/utils/logger"
	"net/http"
)

type AuthService struct {
	repo AuthRepositoryDB
}

func (as AuthService) Signup(reqInput dto.SignupRequest) (*domain.User, *errs.AppError) {
	accountByUsername, err := as.repo.FindUserByUsername(reqInput.Username)
	if accountByUsername != nil {
		return nil, errs.NewBadRequestError("Username already exists")
	}
	if err != nil {
		if err.Code != http.StatusNotFound {
			logger.Error("Error on DB operation while checking username")
			return nil, err
		}
	}
	accountByEmail, err := as.repo.FindAccountByEmail(reqInput.Email)
	if accountByEmail != nil {
		return nil, errs.NewBadRequestError("Email already exists")
	}
	if err != nil {
		if err.Code != http.StatusNotFound {
			return nil, err
		}
	}

	encryptService := encrypt.EncryptService{}
	salt, encrErr := encryptService.GenerateSalt()
	if encrErr != nil {
		logger.Error("Error while generating salt")
		return nil, errs.NewUnexpectedError("Unexpected error")
	}
	hashedPassword, encrErr := encryptService.HashPassword(reqInput.Password, salt)
	if encrErr != nil {
		logger.Error("Error while hashing password")
		return nil, errs.NewUnexpectedError("Unexpected error")
	}
	accountUUID, uuidErr := utils.GenerateUUID()
	if uuidErr != nil {
		logger.Error("Error while generating UUID")
		return nil, errs.NewUnexpectedError("Unexpected error")
	}
	newAccount := domain.Account{
		Id:             accountUUID,
		HashedPassword: hashedPassword,
		Email:          reqInput.Email,
		EmailVerified:  false,
		Salt:           salt,
	}
	newAccountDB, err := as.repo.CreateAccount(newAccount)
	if err != nil {
		return nil, err
	}
	userUUID, uuidErr := utils.GenerateUUID()
	if uuidErr != nil {
		logger.Error("Error while generating UUID")
		return nil, errs.NewUnexpectedError("Unexpected error")
	}
	newUser := domain.User{
		Id:        userUUID,
		AccountId: newAccountDB.Id,
		Username:  reqInput.Username,
		ImageUrl:  "",
	}
	newUserDB, err := as.repo.CreateUser(newUser)
	if err != nil {
		return nil, err
	}
	return newUserDB, nil
}

func NewAuthService(repo AuthRepositoryDB) AuthService {
	return AuthService{
		repo: repo,
	}
}
