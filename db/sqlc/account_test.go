package db

import (
	"back-end/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

// sudo docker exec -it postgres12 psql --username=root --dbname=simple_bank
// docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine
// migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
// sudo service postgresql stop
func createRandomAccount(t *testing.T) Account {
	account := CreateAccountParams{
		Owner:    util.RandomOwner(),
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
