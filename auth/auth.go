package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sohail-9098/ms-user-auth/db"
)

var (
	secretKey = []byte(os.Getenv("SECRET_KEY"))
)

func CreateToken(username string, tokenExpiryMinutes time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * tokenExpiryMinutes).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AuthenticateUser(username, password string) bool {
	return db.FetchUser(username).Password == password
}
