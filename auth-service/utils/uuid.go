package utils

import (
	"github.com/google/uuid"
)

// GenerateUUID generates a new UUID
func GenerateUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
