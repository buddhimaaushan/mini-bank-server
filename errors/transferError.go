package app_error

type transferError struct {
	*baseError
}

// Wrap returns a wrapped error.
func (e *transferError) Wrap(err error) *transferError {
	e.err = err
	return &transferError{
		&baseError{
			kind: e.kind,
			err:  e.err,
		},
	}
}

// As returns whether the error matches the target error.
func (e *transferError) As(target interface{}) bool {
	if _, ok := target.(*transferError); ok {
		return true
	}
	return false
}

// Error returns the error message.
func (e *transferError) Error() string {
	switch e.kind {
	case FromAccountNotActiveError:
		return e.Format("your account is not active")
	case ToAccountNotActiveError:
		return e.Format("the account you are trying to transfer is not active")
	case InsufficientAccountBalanceError:
		return e.Format("your account balance is not sufficient for this transaction")
	case UniqueAccountError:
		return e.Format("cannot transfer to the same account")
	case InvalidAmountError:
		return e.Format("transfer amount should be greater than 0")
	default:
		return e.err.Error()
	}
}

// TrasferError contains all transfer error types.
var TransferError = struct {
	FromAccountNotActiveError       *transferError
	ToAccountNotActiveError         *transferError
	InsufficientAccountBalanceError *transferError
	UniqueAccountError              *transferError
	InvalidAmountError              *transferError
}{
	FromAccountNotActiveError: &transferError{
		&baseError{
			kind: FromAccountNotActiveError,
		},
	},
	ToAccountNotActiveError: &transferError{
		&baseError{
			kind: ToAccountNotActiveError,
		},
	},
	InsufficientAccountBalanceError: &transferError{
		&baseError{
			kind: InsufficientAccountBalanceError,
		},
	},
	UniqueAccountError: &transferError{
		&baseError{
			kind: UniqueAccountError,
		},
	},
	InvalidAmountError: &transferError{
		&baseError{
			kind: InvalidAmountError,
		},
	},
}

// NewTransferError returns a new transferError.
func NewTransferError(kind errKind, code int, err error) *transferError {
	return &transferError{
		&baseError{
			kind: kind,
			err:  err,
		},
	}
}
