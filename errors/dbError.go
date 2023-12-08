package app_error

type dbError struct {
	*baseError
}

// Wrap returns a wrapped error.
func (e *dbError) Wrap(err error) *dbError {
	e.err = err
	return &dbError{
		&baseError{
			kind: e.kind,
			err:  e.err,
		},
	}
}

// As returns whether the error matches the target error.
func (e *dbError) As(target interface{}) bool {
	if _, ok := target.(*dbError); ok {
		return true
	}
	return false
}

// Error returns the error message.
func (e *dbError) Error() string {
	switch e.kind {
	case CreateUserError:
		return e.Format("unable to create user")
	case UserNotFoundError:
		return e.Format("user not found")
	case AccountNotFoundError:
		return e.Format("account not found")
	case AccountHoldersNotFoundError:
		return e.Format("account holders not found")
	default:
		return e.err.Error()
	}
}

// TrasferError contains all transfer error types.
var DbError = struct {
	CreateUserError             *dbError
	UserNotFoundError           *dbError
	AccountNotFoundError        *dbError
	AccountHoldersNotFoundError *dbError
}{
	CreateUserError: &dbError{
		&baseError{
			kind: CreateUserError,
		},
	},
	UserNotFoundError: &dbError{
		&baseError{
			kind: UserNotFoundError,
		},
	},
	AccountNotFoundError: &dbError{
		&baseError{
			kind: AccountNotFoundError,
		},
	},
	AccountHoldersNotFoundError: &dbError{
		&baseError{
			kind: AccountHoldersNotFoundError,
		},
	},
}

// NewdbError returns a new dbError.
func NewDbError(kind errKind, code int, err error) *dbError {
	return &dbError{
		&baseError{
			kind: kind,
			err:  err,
		},
	}
}
