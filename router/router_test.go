package router

import (
	"io"
	"testing"

	"github.com/sohail-9098/ms-user-auth/user"
	"github.com/stretchr/testify/require"
)

// Empty username
func TestRouter_ValidateUserFields_01(t *testing.T) {
	user := user.Credentials{
		Username: "",
		Password: "test",
	}
	err := validateUserFields(user)
	require.Error(t, err, "username should not be empty")

}
// Empty password both
func TestRouter_ValidateUserFields_02(t *testing.T) {
	user := user.Credentials{
		Username: "test",
		Password: "",
	}
	err := validateUserFields(user)
	require.Error(t, err, "password should not be empty")
}

// Empty username
func TestRouter_ValidateUserFields_03(t *testing.T) {
	user := user.Credentials{
		Username: "",
		Password: "",
	}
	err := validateUserFields(user)
	require.Error(t, err, "password and username should not be empty")
}

// Valid input
func TestRouter_ValidateUserFields_04(t *testing.T) {
	user := user.Credentials{
		Username: "test",
		Password: "test",
	}
	err := validateUserFields(user)
	require.NoError(t, err, "password and username is not be empty")
}
