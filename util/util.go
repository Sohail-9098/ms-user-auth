package util

import (
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
