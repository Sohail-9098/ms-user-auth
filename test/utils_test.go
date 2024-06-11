package test

import (
	"testing"

	"github.com/sohail-9098/ms-user-auth/util"
	"github.com/stretchr/testify/require"
)

func TestDb_OpenCloseFile(t *testing.T) {
	// Arrage
	configFileName := "../config.yaml"
	// Act
	file, err := util.OpenFile(configFileName)
	err := file.Close()
	// Assert
	require.NoError(t, err, "error closing file")
}
