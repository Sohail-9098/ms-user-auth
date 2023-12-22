package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/sohail-9098/ms-user-auth/test/e2e/utils"
	"github.com/sohail-9098/ms-user-auth/user"
	"github.com/stretchr/testify/require"
)

func TestE2E_Login(t *testing.T) {
	utils.StartApp()
	user := &user.User{Username: "Chek", Password: "qwerty"}
	userJson, err := json.Marshal(user)
	require.NoError(t, err, nil)
	req, err := http.NewRequest("GET", "http://localhost:4000/login", bytes.NewBuffer(userJson))
	require.NoError(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	require.NoError(t, err, nil)
	require.Equal(t, http.StatusOK, res.StatusCode, nil)
}
