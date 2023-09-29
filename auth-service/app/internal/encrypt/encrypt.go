package encrypt

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

type EncryptService struct{}

func (es EncryptService) GenerateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

// 2 Layer salted password ( bcrypt uses internally a salt when hashing passwords )
func (es EncryptService) HashPassword(password, salt string) (string, error) {
	saltedPassword := password + salt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (es EncryptService) ComparePasswords(password, salt string, encryptedPassword string) (bool, error) {
	saltedPassword := password + salt
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(saltedPassword))
	if err != nil {
		return false, err
	}

	return true, nil
}
