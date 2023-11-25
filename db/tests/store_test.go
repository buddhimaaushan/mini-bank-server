package db_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/buddhimaaushan/mini_bank/db"
	"github.com/buddhimaaushan/mini_bank/db/sqlc"
	app_error "github.com/buddhimaaushan/mini_bank/errors"
	"github.com/stretchr/testify/require"
)

func createRandomAccountForTransferTx(t *testing.T, ctx context.Context, balance int64, status sqlc.Status) sqlc.Account {
	account, err := testQueries.CreateAccount(ctx, sqlc.CreateAccountParams{Type: "savings", Balance: balance, AccStatus: status})
	if err != nil {
		t.Fatal(err)
	}
	return account
}

func getUpdatedAccountByTransferTx(t *testing.T, ctx context.Context, accountID int64) sqlc.Account {
	updatedAccount, err := testQueries.GetAccount(ctx, accountID)
	require.NoError(t, err)

	return updatedAccount
}

func deleteUserCreatedByTransferTx(t *testing.T, ctx context.Context, accountID int64) {
	err := testQueries.DeleteAccount(ctx, accountID)
	require.NoError(t, err)
}

func TestTransferTx(t *testing.T) {
	// Create a new context
	ctx := context.Background()

	// Create a new store using the testDB
	store := db.NewStore(testDB)

	// Create two new accounts
	account1 := createRandomAccountForTransferTx(t, ctx, 10000, "active")
	account2 := createRandomAccountForTransferTx(t, ctx, 10000, "active")

	// Print the balances of the accounts before the transfer
	fmt.Println(">> Before:", account1.Balance, account2.Balance)

	// Number of transfers to perform
	n := 10

	// Amount to transfer
	amount := int64(100)

	// Create channels to collect errors and results from the goroutines
	errs := make(chan error)
	results := make(chan db.TransferTxResult)

	// Perform transfers concurrently
	for i := 0; i < n; i++ {
		txName := fmt.Sprintf("tx: %02d", i+1)

		go func() {
			ctx := context.WithValue(context.Background(), db.TxKey, txName)

			// Perform the transfer using the store object
			result, err := store.TransferTx(ctx, db.TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}

	// Collect the results and validate them
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		// Get the error from the channel
		err := <-errs
		require.NoError(t, err)

		// Get the result from the channel
		result := <-results
		require.NotEmpty(t, result)

		// Get the transfer from the result
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)

		// Get the transfer from the store
		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// Get the entries from the store
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)
		// Get the entry from the store
		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)
		//  Get the entry from the store
		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// Get the accounts from the store
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		// Compare the balances of the accounts
		fmt.Printf(">> %s, From: %d, To: %d\n", result.TxName, fromAccount.Balance, toAccount.Balance)
		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true

		// Delete the transfer from the store
		err = store.DeleteTransfer(context.Background(), result.Transfer.ID)
		require.NoError(t, err)

		// Delete the 'from' entry
		err = store.DeleteEntry(context.Background(), result.FromEntry.ID)
		require.NoError(t, err)

		// Delete the 'to' entry
		err = store.DeleteEntry(context.Background(), result.ToEntry.ID)
		require.NoError(t, err)

	}

	// Retrieve the updated account information for accounts
	updatedAccount1 := getUpdatedAccountByTransferTx(t, ctx, account1.ID)
	updatedAccount2 := getUpdatedAccountByTransferTx(t, ctx, account2.ID)

	// Print the balances of the updated accounts
	fmt.Println(">> After:", updatedAccount1.Balance, updatedAccount2.Balance)

	// Compare the updated balances
	require.Equal(t, updatedAccount1.Balance, account1.Balance-amount*int64(n))
	require.Equal(t, updatedAccount2.Balance, account2.Balance+amount*int64(n))

	// Delete the account1 from the store
	deleteUserCreatedByTransferTx(t, ctx, account1.ID)

	// Delete the account2 from the store
	deleteUserCreatedByTransferTx(t, ctx, account2.ID)
}

func TestTransferTxBetwenTwoAccountsDeadlock(t *testing.T) {
	// Create a new context
	ctx := context.Background()

	// Create a new store using the testDB
	store := db.NewStore(testDB)

	// Create two new accounts
	account1 := createRandomAccountForTransferTx(t, ctx, 10000, "active")
	account2 := createRandomAccountForTransferTx(t, ctx, 10000, "active")

	// Print the balances of the accounts before the transfer
	fmt.Println(">> Before:", account1.Balance, account2.Balance)

	// Number of transfers to perform
	n := 10

	// Amount to transfer
	amount := int64(100)

	// Create channels to collect errors and results from the goroutines
	errs := make(chan error)

	// Perform transfers concurrently
	for i := 0; i < n; i++ {
		txName := fmt.Sprintf("tx: %02d", i+1)
		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}

		go func() {
			ctx := context.WithValue(context.Background(), db.TxKey, txName)

			// Perform the transfer using the store object
			result, err := store.TransferTx(ctx, db.TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})
			errs <- err
			// Compare the balances of the accounts
			fmt.Printf(">> %s, From: %d, To: %d\n", result.TxName, result.FromEntry.Amount, result.ToEntry.Amount)

			// Delete the transfer from the store
			err = store.DeleteTransfer(context.Background(), result.Transfer.ID)
			require.NoError(t, err)

			// Delete the 'from' entry
			err = store.DeleteEntry(context.Background(), result.FromEntry.ID)
			require.NoError(t, err)

			// Delete the 'to' entry
			err = store.DeleteEntry(context.Background(), result.ToEntry.ID)
			require.NoError(t, err)
		}()
	}

	for i := 0; i < n; i++ {
		// Get the error from the channel
		err := <-errs
		require.NoError(t, err)

	}

	// Retrieve the updated account information for accounts
	updatedAccount1 := getUpdatedAccountByTransferTx(t, ctx, account1.ID)
	updatedAccount2 := getUpdatedAccountByTransferTx(t, ctx, account2.ID)

	// Print the balances of the updated accounts
	fmt.Println(">> After:", updatedAccount1.Balance, updatedAccount2.Balance)

	// Compare the updated balances
	require.Equal(t, updatedAccount1.Balance, account1.Balance)
	require.Equal(t, updatedAccount2.Balance, account2.Balance)

	// Delete the account1 from the store
	deleteUserCreatedByTransferTx(t, ctx, account1.ID)

	// Delete the account2 from the store
	deleteUserCreatedByTransferTx(t, ctx, account2.ID)
}
func TestTransferTxAmountLTEQZero(t *testing.T) {
	// Create a new context
	ctx := context.Background()

	// Create a new store using the testDB
	store := db.NewStore(testDB)

	// Create two new accounts
	account1 := createRandomAccountForTransferTx(t, ctx, 10000, "active")
	account2 := createRandomAccountForTransferTx(t, ctx, 10000, "active")

	// Print the balances of the accounts before the transfer
	fmt.Println(">> Before:", account1.Balance, account2.Balance)

	// Number of transfers to perform
	n := 10

	// Amount to transfer
	amount := int64(0)
	if n%2 == 1 {
		amount = int64(-100)
	}

	// Create channels to collect errors and results from the goroutines
	errs := make(chan error)

	// Perform transfers concurrently
	for i := 0; i < n; i++ {
		txName := fmt.Sprintf("tx: %02d", i+1)
		fromAccountID := account1.ID
		toAccountID := account2.ID

		go func() {
			ctx := context.WithValue(context.Background(), db.TxKey, txName)

			// Perform the transfer using the store object
			result, err := store.TransferTx(ctx, db.TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})
			errs <- err
			// Compare the balances of the accounts
			fmt.Printf(">> %s, From: %d, To: %d\n", result.TxName, result.FromEntry.Amount, result.ToEntry.Amount)

			// Delete the transfer from the store
			err = store.DeleteTransfer(context.Background(), result.Transfer.ID)
			require.NoError(t, err)

			// Delete the 'from' entry
			err = store.DeleteEntry(context.Background(), result.FromEntry.ID)
			require.NoError(t, err)

			// Delete the 'to' entry
			err = store.DeleteEntry(context.Background(), result.ToEntry.ID)
			require.NoError(t, err)
		}()
	}

	for i := 0; i < n; i++ {
		// Get the error from the channel
		err := <-errs
		require.Error(t, err)

		// Check if the error is an invalid amount error
		require.Equal(t, err, app_error.TransferError.InvalidAmount)

	}

	// Retrieve the updated account information for accounts
	updatedAccount1 := getUpdatedAccountByTransferTx(t, ctx, account1.ID)
	updatedAccount2 := getUpdatedAccountByTransferTx(t, ctx, account2.ID)

	// Print the balances of the updated accounts
	fmt.Println(">> After:", updatedAccount1.Balance, updatedAccount2.Balance)

	// Compare the updated balances
	require.Equal(t, updatedAccount1.Balance, account1.Balance)
	require.Equal(t, updatedAccount2.Balance, account2.Balance)

	// Delete the account1 from the store
	deleteUserCreatedByTransferTx(t, ctx, account1.ID)

	// Delete the account2 from the store
	deleteUserCreatedByTransferTx(t, ctx, account2.ID)
}
