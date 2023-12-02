package db

import (
	"back-end/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomAccount(t *testing.T) Account {
	user := CreateRandomUser(t)
	account := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	acc, err := testQueries.CreateAccount(context.Background(), account)

	require.NoError(t, err)
	require.NotEmpty(t, acc)
	require.Equal(t, acc.Owner, account.Owner)
	require.Equal(t, acc.Balance, account.Balance)
	require.Equal(t, acc.Currency, account.Currency)
	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)
	return acc
}
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}
