package app_error

import "fmt"

type transferError struct {
	kind err
	code int
	err  error
}

// Wrap returns a wrapped error.
func (e *transferError) Wrap(err error) *transferError {
	e.err = err
	return &transferError{
		kind: e.kind,
		code: e.code,
		err:  e.err,
	}
}

// Unwrap returns the underlying error.
func (e *transferError) Unwrap() error {
	return e.err
}

// Is returns whether the error matches the target error.
func (e *transferError) Is(target error) bool {
	if err, ok := target.(*transferError); ok {
		return e.kind == err.kind
	}
	return false
}

// As returns whether the error matches the target error.
func (e *transferError) As(target interface{}) bool {
	if err, ok := target.(*transferError); ok {
		*err = *e
		return true
	}
	return false
}

// AsError returns the error.
func (e *transferError) AsError() error {
	return e
}

// String returns the error message.
func (e *transferError) String() string {
	return e.Error()
}

// Kind returns the error kind.
func (e *transferError) Kind() err {
	return e.kind
}

// Code returns the error code.
func (e *transferError) Code() int {
	return e.code
}

// Error returns the error message.
func (e *transferError) Error() string {
	switch e.kind {
	case fromAccountNotActive:
		return fmt.Sprintf("your account is not active [%d]: %s", e.code, e.err.Error())
	case toAccountNotActive:
		return fmt.Sprintf("the account you are trying to transfer is not active [%d]: %s", e.code, e.err.Error())
	case insufficientAccountBalance:
		return fmt.Sprintf("your account balance is not sufficient for this transaction [%d]: %s", e.code, e.err.Error())
	case sameAccount:
		return fmt.Sprintf("cannot transfer to the same account [%d]: %s", e.code, e.err.Error())
	case invalidAmount:
		return fmt.Sprintf("transfer amount should be greater than 0 [%d]: %s", e.code, e.err.Error())
	default:
		return e.err.Error()
	}
}

// TrasferError contains all transfer error types.
var TransferError = struct {
	FromAccountNotActive       *transferError
	ToAccountNotActive         *transferError
	InsufficientAccountBalance *transferError
	SameAccount                *transferError
	InvalidAmount              *transferError
}{
	FromAccountNotActive: &transferError{
		kind: fromAccountNotActive,
		code: 1001,
	},
	ToAccountNotActive: &transferError{
		kind: toAccountNotActive,
		code: 1002,
	},
	InsufficientAccountBalance: &transferError{
		kind: insufficientAccountBalance,
		code: 1003,
	},
	SameAccount: &transferError{
		kind: sameAccount,
		code: 1004,
	},
	InvalidAmount: &transferError{
		kind: invalidAmount,
		code: 1005,
	},
}

// NewTransferError returns a new transferError.
func NewTransferError(kind err, code int, err error) *transferError {
	return &transferError{
		kind: kind,
		code: code,
		err:  err,
	}
}
