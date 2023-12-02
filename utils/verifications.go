package utils

import (
	"regexp"
	"slices"

	"github.com/buddhimaaushan/mini_bank/db/sqlc"
)

// Minimum balance for an account to be considered valid.
var MinAccountBalance int64 = 100

// verifyTxBalance verifies the balance of an account.
func IsAccountBalanceSufficient(account sqlc.Account) bool {
	return account.Balance >= MinAccountBalance
}

// IsAccountActive verifies the status of an account.
func IsAccountActive(account sqlc.Account) bool {
	return account.AccStatus == "active"
}

// EmailRegex is a regex for email
var EmailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@([a-z0-9.\-]+\.[a-z]{2,4})$`)

// validEmailTypes is a list of valid email types
var validEmailTypes = []string{"gmail.com", "yahoo.com", "hotmail.com", "outlook.com"}

// verifyTxBalance verifies the email type is valid or not.
func IsAValidEmailType(email string) bool {
	//Create a regex
	r := EmailRegex

	//get the match
	matchSlice := r.FindStringSubmatch(email)
	if len(matchSlice) == 0 {
		return false
	}

	//get the email type
	match := matchSlice[1]

	//check if the email type is valid
	return slices.Contains(validEmailTypes, match)

}
