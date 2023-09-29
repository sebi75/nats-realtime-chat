package auth

import (
	"auth-service/app/auth/domain"
	"auth-service/app/auth/dto"
	"auth-service/app/internal/encrypt"
	"auth-service/app/internal/jwt"
	"auth-service/errs"
	"auth-service/utils"
	"auth-service/utils/logger"
	"net/http"
)

type AuthService struct {
	repo AuthRepositoryDB
}

func (as AuthService) Signup(reqInput *dto.SignupRequest) (*domain.UserWithAccountDTO, *errs.AppError) {
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
	logger.Info("New user created: " + newAccount.Email)
	return &domain.UserWithAccountDTO{
		User:    *newUserDB,
		Account: *newAccountDB.ToResponseDTO(),
	}, nil
}

func (as AuthService) Signin(reqInput *dto.SigninRequest) (string, *errs.AppError) {
	result, err := as.repo.FindAccountByUsername(reqInput.Username)
	if err != nil {
		return "", err
	}
	encryptService := encrypt.EncryptService{}
	if res, err := encryptService.ComparePasswords(reqInput.Password, result.Account.Salt, result.Account.HashedPassword); err != nil || !res {
		return "", errs.NewUnauthorizedError("Invalid credentials")
	}
	jwtService := jwt.DefaultJwtService{}
	token, tokenErr := jwtService.GenerateAuthToken(result.Id, result.AccountId)
	if tokenErr != nil {
		return "", errs.NewUnexpectedError("Unexpected error")
	}
	return token, nil
}

func NewAuthService(repo AuthRepositoryDB) AuthService {
	return AuthService{
		repo: repo,
	}
}
