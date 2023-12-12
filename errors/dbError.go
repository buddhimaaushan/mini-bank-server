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
	case ErrCreateUser:
		return e.Format("unable to create user")
	case ErrUserNotFound:
		return e.Format("user not found")
	case ErrAccountNotFound:
		return e.Format("account not found")
	case ErrAccountHoldersNotFound:
		return e.Format("account holders not found")
	case ErrCreateSession:
		return e.Format("unable to create session")
	default:
		return e.err.Error()
	}
}

// TrasferError contains all transfer error types.
var DbError = struct {
	ErrCreateUser             *dbError
	ErrUserNotFound           *dbError
	ErrAccountNotFound        *dbError
	ErrAccountHoldersNotFound *dbError
	ErrCreateSession          *dbError
}{
	ErrCreateUser: &dbError{
		&baseError{
			kind: ErrCreateUser,
		},
	},
	ErrUserNotFound: &dbError{
		&baseError{
			kind: ErrUserNotFound,
		},
	},
	ErrAccountNotFound: &dbError{
		&baseError{
			kind: ErrAccountNotFound,
		},
	},
	ErrAccountHoldersNotFound: &dbError{
		&baseError{
			kind: ErrAccountHoldersNotFound,
		},
	},
	ErrCreateSession: &dbError{
		&baseError{
			kind: ErrCreateSession,
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
