package db_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	sqlc "github.com/buddhimaaushan/mini_bank/db/sqlc"
	"github.com/buddhimaaushan/mini_bank/utils"
	"github.com/stretchr/testify/require"
)

// Create a random user
func createRandomUser(t *testing.T) sqlc.User {

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
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	// Check user
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Nic, user.Nic)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Phone, user.Phone)

	// Check user auto-generated field values
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.Equal(t, sqlc.Status(utils.INACTIVE), user.AccStatus)
	require.Equal(t, sqlc.Rank(utils.BRONZE), user.CustomerRank)
	require.False(t, user.IsEmailVerified)
	require.False(t, user.IsPhoneVerified)
	require.False(t, user.IsAnEmployee)
	require.False(t, user.IsACustomer)
	require.False(t, user.CreatedAt.Time.IsZero())
	require.True(t, user.PasswordChangedAt.Time.IsZero())
	require.True(t, user.EmailChangedAt.Time.IsZero())

	// Return user
	return user
}

// Compare two users
func compareUserEquality(t *testing.T, user1, user2 sqlc.User) {
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Nic, user2.Nic)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.PasswordChangedAt, user2.PasswordChangedAt)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.IsEmailVerified, user2.IsEmailVerified)
	require.Equal(t, user1.Phone, user2.Phone)
	require.Equal(t, user1.IsPhoneVerified, user2.IsPhoneVerified)
	require.Equal(t, user1.AccStatus, user2.AccStatus)
	require.Equal(t, user1.CustomerRank, user2.CustomerRank)
	require.Equal(t, user1.IsAnEmployee, user2.IsAnEmployee)
	require.Equal(t, user1.IsACustomer, user2.IsACustomer)
	require.Equal(t, user1.Role, user2.Role)
	require.Equal(t, user1.Department, user2.Department)

	require.WithinDuration(t, user1.EmailChangedAt.Time, user2.EmailChangedAt.Time, time.Second)
	require.WithinDuration(t, user1.PhoneChangedAt.Time, user2.PhoneChangedAt.Time, time.Second)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}

func deleteUser(t *testing.T, user sqlc.User) {
	// Delete user
	delUser, err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, delUser)

	// Check user
	compareUserEquality(t, user, delUser)

}

func TestCreateAndDeleteUser(t *testing.T) {
	randUser := createRandomUser(t)
	deleteUser(t, randUser)
}

func TestGetUserByID(t *testing.T) {
	// Create a random user
	randUser := createRandomUser(t)

	// Get the "randUser" created in db from db
	dbUser, err := testQueries.GetUserByID(context.Background(), randUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, dbUser)

	// Check user
	compareUserEquality(t, randUser, dbUser)

	// Delete user
	deleteUser(t, randUser)
}
