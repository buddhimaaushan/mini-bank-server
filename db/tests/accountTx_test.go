package db_test

import (
	"context"
	"testing"

	"github.com/buddhimaaushan/mini_bank/db"
	"github.com/buddhimaaushan/mini_bank/db/sqlc"
	"github.com/stretchr/testify/require"
)

func TestAccountTx(t *testing.T) {
	ctx := context.Background()

	store := db.NewStore(testDB)

	// Create two new users
	_, user1 := createRandomUser(t)
	_, user2 := createRandomUser(t)

	arg := db.AccountTxParams{
		Type:      "deposit",
		Balance:   1000,
		AccStatus: sqlc.StatusActive,
		UserIDs:   []int64{user1.ID, user2.ID},
	}

	// Create account and account holders
	result, err := store.AccountTx(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// Check account
	account, err := store.GetAccount(ctx, result.Account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	// Assert account details
	require.Equal(t, arg.Type, account.Type)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.AccStatus, account.AccStatus)

	// Check account holders
	require.Len(t, result.AccountHolders, len(arg.UserIDs))
	for _, account_holder := range result.AccountHolders {
		// Assert account holder details
		require.NotEmpty(t, account_holder)
		require.Equal(t, account.ID, account_holder.AccID)
		require.Contains(t, arg.UserIDs, account_holder.UserID)
	}

	// Delete account holders
	deleteAccountHolders(t, ctx, result.AccountHolders)

	// Delete Account
	deleteAccount(t, ctx, result.Account.ID)

	// Delete users
	deleteUser(t, user1)
	deleteUser(t, user2)

}
