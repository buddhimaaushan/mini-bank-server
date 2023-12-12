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
	case ErrFromAccountNotActive:
		return e.Format("your account is not active")
	case ErrToAccountNotActive:
		return e.Format("the account you are trying to transfer is not active")
	case ErrInsufficientAccountBalance:
		return e.Format("your account balance is not sufficient for this transaction")
	case ErrUniqueAccount:
		return e.Format("cannot transfer to the same account")
	case ErrInvalidAmount:
		return e.Format("transfer amount should be greater than 0")
	default:
		return e.err.Error()
	}
}

// TrasferError contains all transfer error types.
var TransferError = struct {
	ErrFromAccountNotActive       *transferError
	ErrToAccountNotActive         *transferError
	ErrInsufficientAccountBalance *transferError
	ErrUniqueAccount              *transferError
	ErrInvalidAmount              *transferError
}{
	ErrFromAccountNotActive: &transferError{
		&baseError{
			kind: ErrFromAccountNotActive,
		},
	},
	ErrToAccountNotActive: &transferError{
		&baseError{
			kind: ErrToAccountNotActive,
		},
	},
	ErrInsufficientAccountBalance: &transferError{
		&baseError{
			kind: ErrInsufficientAccountBalance,
		},
	},
	ErrUniqueAccount: &transferError{
		&baseError{
			kind: ErrUniqueAccount,
		},
	},
	ErrInvalidAmount: &transferError{
		&baseError{
			kind: ErrInvalidAmount,
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
