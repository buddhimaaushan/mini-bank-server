package utils

import "github.com/buddhimaaushan/mini_bank/db/sqlc"

// Account statuses
const (
	ACTIVE   = "active"
	INACTIVE = "inactive"
	DELETED  = "deleted"
	HOLDED   = "holded"
)

// IsVerifiedAccStatus verifies the account status.
func IsVerifiedAccStatus(status string) bool {
	switch status {
	case ACTIVE, INACTIVE, DELETED, HOLDED:
		return true
	}
	return false
}

// Customer ranks
const (
	BRONZE   = "bronze"
	SILVER   = "silver"
	GOLD     = "gold"
	PLATINUM = "platinum"
	GARNET   = "garnet"
	RUBY     = "ruby"
	DIAMOND  = "diamond"
)

// IsVerifiedCustomerRank verifies the customer rank.
func IsVerifiedCustomerRank(status string) bool {
	switch status {
	case BRONZE, SILVER, GOLD, PLATINUM, GARNET, RUBY, DIAMOND:
		return true
	}
	return false
}

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
