package app_error

type err int

const (
	_ err = iota
	toAccountNotActive
	fromAccountNotActive
	insufficientAccountBalance
	sameAccount
	invalidAmount
)
