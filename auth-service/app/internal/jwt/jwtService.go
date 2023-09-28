package jwt

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(userId string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

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

func (s DefaultJwtService) ValidateToken(token string) (*jwt.Token, error) {
	return nil, nil
}
