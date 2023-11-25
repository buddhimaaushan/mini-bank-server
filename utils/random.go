package utils

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/buddhimaaushan/mini_bank/constants"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomType() string {
	types := [...]string{"Current account", "Savings account", "cheque", "Certificate of deposit", "Recurring Deposit account", "Fixed deposit", "Overseas", "High-yield savings account", "Individual retirement account", "Business checking account", "Money market fund"}
	return types[rand.Intn(len(types))]
}

func RandomBalance() string {
	return strconv.Itoa(RandomInt(1000, 1000000000))
}

func RandomCurrency() constants.Currency {
	currencyList := constants.CurrencyList
	return currencyList[rand.Intn(len(currencyList))]
}
