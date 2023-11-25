package db

import (
	"back-end/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func EntryTest(t *testing.T) {
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NotEmpty(t, entry)
	require.NoError(t, err)
	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)
}
