package utils

import "github.com/buddhimaaushan/mini_bank/db/sqlc"

// Minimum balance for an account to be considered valid.
var MinAccountBalance int64 = 100

// verifyTxBalance verifies the balance of an account.
func IsAccountBalanceSufficient(account sqlc.Account) bool {
	if account.Balance < MinAccountBalance {
		return false
	}
	return true
}

// IsAccountActive verifies the status of an account.
func IsAccountActive(account sqlc.Account) bool {
	if account.AccStatus != "active" {
		return false
	}
	return true
}
