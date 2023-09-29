package jwt

import (
	"auth-service/app/auth/domain"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type DefaultJwtService struct{}

func (s DefaultJwtService) GenerateAuthToken(userId string, accountId string) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userId,
		"accountId": accountId,
	})

	token, err := claims.SignedString(secret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s DefaultJwtService) ValidateToken(token string) (bool, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s DefaultJwtService) DecodeToken(token string) (*domain.TokenPayload, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	return &domain.TokenPayload{
		UserId:    claims["userId"].(string),
		AccountId: claims["accountId"].(string),
	}, nil
}
