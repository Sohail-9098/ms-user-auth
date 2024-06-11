package auth

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sohail-9098/ms-user-auth/db"
)

var (
	secretKey = []byte(os.Getenv("SECRET_KEY"))
)

func CreateToken(username string, tokenExpiryMinutes time.Duration) (string, error) {
	if secretKey == nil {
		log.Printf("SECRET_KEY is not set")
		return "", errors.New("SECRET_KEY is not set")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * tokenExpiryMinutes).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Printf("error creating secret: %v", err)
		return "", err
	}
	return tokenString, nil
}

func AuthenticateUser(username, password string) (bool, error) {
	// fetch user and
	// check password
	user, err := db.FetchUser(username)
	if err != nil || user.Password != password {
		return false, err
	}
	return true, nil
}
