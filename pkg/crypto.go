package pkg

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GenerateToken() string {
	uuid := uuid.New()
	token, err := HashPassword(uuid.String())
	if err != nil {
		return uuid.String()
	}
	return token
}
