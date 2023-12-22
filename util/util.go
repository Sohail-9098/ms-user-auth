package util

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

func HandleError(msg string, err error) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}

func OpenFile(configFileName string) *os.File {
	file, err := os.Open(configFileName)
	HandleError("error open file: ", err)
	return file
}

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}
