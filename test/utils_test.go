package util

import (
	"testing"

	"github.com/sohail-9098/ms-user-auth/util"
)

func TestDb_OpenFile(t *testing.T) {
	configFileName := "../config.yaml"
	util.OpenFile(configFileName)
}
