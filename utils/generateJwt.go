package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GenerateJwt(ID uuid.UUID) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 168)),
		Issuer:    "EngineerMandi",
		Subject:   ID.String(),
	}

	jwtKeyString := os.Getenv("JWT_KEY")
	if jwtKeyString == "" {
		return "", errors.New("JWT_KEY is not set")
	}

	JWT_KEY := []byte(jwtKeyString)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWT_KEY)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
