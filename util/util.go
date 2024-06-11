package util

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
)

func HandleError(msg string, err error) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}

func OpenFile(configFileName string) (*os.File, error) {
	file, err := os.Open(configFileName)
	if err!=nil{
		return nil, err
	}
	return file, nil
}

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}

func DecodePassword(password string) (string, error) {
	pass, err := base64.StdEncoding.DecodeString(password)
	if err!=nil{
		return "", errors.New("error decoding db password")
	}
	return string(pass), nil
}
