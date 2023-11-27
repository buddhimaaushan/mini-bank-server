package db

import (
	"context"
	"fmt"

	"github.com/buddhimaaushan/mini_bank/db/sqlc"
	app_error "github.com/buddhimaaushan/mini_bank/errors"
	"github.com/buddhimaaushan/mini_bank/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectToDb establishes a connection to the database and returns the connection object.
func ConnectToDb(dbUri string) (conn *pgxpool.Pool, err error) {
	// Create a new connection pool using the DATABASE_URL environment variable.
	conn, err = pgxpool.New(context.Background(), dbUri)
	return conn, err
}

type Store interface {
	sqlc.Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	AccountTx(ctx context.Context, arg AccountTxParams) (AccountTxResult, error)
}

// SQLSrore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*sqlc.Queries
	db *pgxpool.Pool
}

// NewStore creates a new Store instance.
// It takes a db *pgx.Conn parameter and returns a pointer to a Store.
func NewStore(db *pgxpool.Pool) Store {
	// Create a new Store instance and initialize its fields.
	return &SQLStore{
		db:      db,
		Queries: sqlc.New(db),
	}
}

// execTx executes a transaction using the provided function.
// The function receives a *sqlc.Queries object to execute the queries within the transaction.
// It returns an error if the transaction or any of the queries fail.
func (store *SQLStore) execTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	// Begin a new transaction with the database connection.
	tx, err := store.db.BeginTx(ctx, pgx.TxOptions{
		// IsoLevel: pgx.RepeatableRead,
	})
	if err != nil {
		return err
	}

	// Create a new *sqlc.Queries object to execute queries within the transaction.
	q := sqlc.New(tx)

	// Execute the provided function with the *sqlc.Queries object.
	err = fn(q)
	if err != nil {
		// Rollback the transaction if the provided function returns an error.
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	// Commit the transaction if the provided function completes successfully.
	return tx.Commit(ctx)
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    sqlc.Transfer `json:"transfer"`
	FromAccount sqlc.Account  `json:"from_account"`
	ToAccount   sqlc.Account  `json:"to_account"`
	FromEntry   sqlc.Entry    `json:"from_entry"`
	ToEntry     sqlc.Entry    `json:"to_entry"`
	TxName      any           `json:"tx_name"`
}

var TxKey = struct{}{}

// TransferTx performs a transfer transaction.
// It takes a context and TransferTxParams as input and returns a TransferTxResult and an error.
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult // Declare a variable of type TransferTxResult. This variable will hold the result of the transaction.

	// Execute the transaction using the execTx method of the store.
	err := store.execTx(ctx, func(q *sqlc.Queries) error {
		var err error

		// Get the transaction name from the context.
		result.TxName = ctx.Value(TxKey)

		// Verify to and from accounts
		if arg.FromAccountID == arg.ToAccountID {
			return app_error.TransferError.SameAccount
		}

		// Verify the ammount is greater than 0
		if arg.Amount <= 0 {
			return app_error.TransferError.InvalidAmount
		}

		// Create a transfer using the CreateTransfer method of the queries.
		result.Transfer, err = q.CreateTransfer(ctx, sqlc.CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		// Create an entry for the 'from' account with a negative amount.
		result.FromEntry, err = q.CreateEntry(ctx, sqlc.CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		// Create an entry for the 'to' account with a positive amount.
		result.ToEntry, err = q.CreateEntry(ctx, sqlc.CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.FromAccountID < arg.ToAccountID {

			// Update the balance of the 'from' and 'to' account
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
			if err != nil {
				return err
			}

		} else {

			// Update the balance of the 'to' and 'from' account
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return result, err
}

// addMoney updates the balance of two accounts and returns the updated accounts.
func addMoney(ctx context.Context, q *sqlc.Queries, accountID1 int64, amount1 int64, accountID2 int64, amount2 int64) (account1 sqlc.Account, account2 sqlc.Account, err error) {
	// Update the balance of the first account
	account1, err = q.UpdateAccountBalance(ctx, sqlc.UpdateAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	//  Verify the status of the first account
	if !utils.IsAccountActive(account1) {
		if amount1 <= 0 {
			err = app_error.TransferError.FromAccountNotActive
		}
		err = app_error.TransferError.ToAccountNotActive
	}
	if err != nil {
		return
	}

	// Verify the balance of the first account
	if !utils.IsAccountBalanceSufficient(account1) {
		err = app_error.TransferError.InsufficientAccountBalance
		return
	}

	// Update the balance of the second account
	account2, err = q.UpdateAccountBalance(ctx, sqlc.UpdateAccountBalanceParams{
		ID:     accountID2,
		Amount: amount2,
	})
	if err != nil {
		return
	}

	//  Verify the status of the second account
	if !utils.IsAccountActive(account2) {
		if amount2 <= 0 {
			err = app_error.TransferError.FromAccountNotActive
		}
		err = app_error.TransferError.ToAccountNotActive
	}
	if err != nil {
		return
	}

	// Verify the balance of the second account
	if !utils.IsAccountBalanceSufficient(account2) {
		err = app_error.TransferError.InsufficientAccountBalance
		return
	}

	return
}

type AccountTxParams struct {
	Type      string      ` json:"type"`
	Balance   int64       ` json:"balance"`
	AccStatus sqlc.Status ` json:"acc_status"`
	UserIDs   []int64     ` json:"user_id"`
}

type AccountTxResult struct {
	Account        sqlc.Account         `json:"account"`
	AccountHolders []sqlc.AccountHolder `json:"account_holders"`
}

// AccountTx creates a new account with account holders
func (store *SQLStore) AccountTx(ctx context.Context, arg AccountTxParams) (AccountTxResult, error) {
	var result AccountTxResult

	// Execute the transaction using the execTx method of the store.
	err := store.execTx(ctx, func(q *sqlc.Queries) error {
		var err error

		// Create an account
		result.Account, err = q.CreateAccount(ctx, sqlc.CreateAccountParams{
			Type:      arg.Type,
			Balance:   arg.Balance,
			AccStatus: arg.AccStatus,
		})
		if err != nil {
			return err
		}

		//Generate account holders from user IDs
		accountHolders := generateAccountHolders(result.Account.ID, arg.UserIDs)

		// Create account holders
		limit, err := q.CreateAccountHolders(ctx, accountHolders)
		if err != nil {
			return err
		}

		// Get the account holders
		result.AccountHolders, err = q.GetAccountHoldersByAccountID(ctx, sqlc.GetAccountHoldersByAccountIDParams{
			AccID:  result.Account.ID,
			Limit:  int32(limit),
			Offset: 0,
		})
		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}

// Generate account holders
func generateAccountHolders(accountID int64, userIDs []int64) []sqlc.CreateAccountHoldersParams {
	var accountHolders []sqlc.CreateAccountHoldersParams
	for _, userID := range userIDs {
		accountHolders = append(accountHolders, sqlc.CreateAccountHoldersParams{
			AccID:  accountID,
			UserID: userID,
		})
	}
	return accountHolders

}
