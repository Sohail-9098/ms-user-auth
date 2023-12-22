package util

import (
	"testing"
)

func TestDb_OpenFile(t *testing.T) {
	configFileName := "../config/config.yaml"
	OpenFile(configFileName)
}
