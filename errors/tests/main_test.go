package app_error_test

import (
	"testing"

	app_error "github.com/buddhimaaushan/mini_bank/errors"
	"github.com/stretchr/testify/require"
)

func TestErrorAsErrorType(t *testing.T) {
	err1 := app_error.TransferError.FromAccountNotActiveError
	err2 := app_error.ApiError.InvalidUsernameOrPasswordError
	err3 := app_error.ApiError.InvalidUsernameOrPasswordError

	require.False(t, err1.As(err2))
	require.True(t, err2.As(err3))
}
