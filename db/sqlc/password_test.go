package db

import (
	"back-end/util"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPassword(t *testing.T) {
	password := util.RandomString(6)
	hashedpsw, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedpsw)

	err = util.CheckPassword(password, hashedpsw)
	require.NoError(t, err)
}
func TestWrongPassword(t *testing.T) {
	password := util.RandomString(6)

	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	wrongPassword := util.RandomString(6)
	err = util.CheckPassword(wrongPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
