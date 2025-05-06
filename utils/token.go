package utils

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenToken(claims jwt.MapClaims) (string, error) {

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString([]byte(os.Getenv("APP_KEY")))

	return token, err
}

func DecodeToken(token string) (*jwt.Token, jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	jwt, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		key := os.Getenv("APP_KEY")

		if key == "" {
			return []byte(key), errors.New("key is empty")
		}
		return []byte(key), nil
	})

	return jwt, claims, err
}
