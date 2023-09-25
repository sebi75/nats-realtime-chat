package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(userId string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type DefaultJwtService struct {
}

func (s DefaultJwtService) GenerateAuthToken(userId string) (string, error) {
	return "", nil
}

func (s DefaultJwtService) ValidateToken(token string) (*jwt.Token, error) {
	return nil, nil
}
