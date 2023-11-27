package db_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/buddhimaaushan/mini_bank/db"
	"github.com/buddhimaaushan/mini_bank/db/sqlc"
	"github.com/buddhimaaushan/mini_bank/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

var store db.Store
var testDB *pgxpool.Pool
var config utils.Config

func TestMain(m *testing.M) {

	var err error

	// Load environment variables
	config, err = utils.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	log.Println("environment variables loaded")
	log.Println(config)

	// Connect to database
	testDB, err = db.ConnectToDb(config.DatabaseURL)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	log.Println("database connection established")

	store = db.NewStore(testDB)
	os.Exit(m.Run())
}

// Create a random user
func createRandomUser(t *testing.T) (sqlc.CreateUserParams, sqlc.User) {
	// Gemnerate a random user
	arg := sqlc.CreateUserParams{
		FirstName:      utils.RandomString(6),
		LastName:       utils.RandomString(6),
		Username:       utils.RandomString(6),
		Nic:            strconv.Itoa(utils.RandomInt(1000000000, 9999999999)),
		HashedPassword: "secret",
		Email:          fmt.Sprintf("%s@%s.com", utils.RandomString(6), utils.RandomString(6)),
		Phone:          strconv.Itoa(utils.RandomInt(1000000000, 9999999999)),
	}

	// Create user
	user, err := store.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	// Return user
	return arg, user
}

// Delete an user
func deleteUser(t *testing.T, user sqlc.User) {
	// Delete user
	delUser, err := store.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, delUser)

	// Check user
	compareUserEquality(t, user, delUser)

}

// Create random account
func createRandomAccount(t *testing.T, ctx context.Context, balance int64, status sqlc.Status) sqlc.Account {
	account, err := store.CreateAccount(ctx, sqlc.CreateAccountParams{Type: "savings", Balance: balance, AccStatus: status})
	if err != nil {
		t.Fatal(err)
	}
	return account
}

// Get updated account
func getUpdatedAccount(t *testing.T, ctx context.Context, accountID int64) sqlc.Account {
	updatedAccount, err := store.GetAccount(ctx, accountID)
	require.NoError(t, err)

	return updatedAccount
}

// Delete an account
func deleteAccount(t *testing.T, ctx context.Context, accountID int64) {
	err := store.DeleteAccount(ctx, accountID)
	require.NoError(t, err)
}

func deleteAccountHolders(t *testing.T, ctx context.Context, accountHolders []sqlc.AccountHolder) {
	for _, account_holder := range accountHolders {

		arg := sqlc.DeleteAccountHolderParams{
			AccID:  account_holder.AccID,
			UserID: account_holder.UserID,
		}

		// Delete account holder
		err := store.DeleteAccountHolder(ctx, arg)
		require.NoError(t, err)
	}
}
