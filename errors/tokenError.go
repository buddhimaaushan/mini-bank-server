package app_error

type tokenError struct {
	*baseError
}

// Wrap returns a wrapped error.
func (e *tokenError) Wrap(err error) *tokenError {
	e.err = err
	return &tokenError{
		&baseError{
			kind: e.kind,
			err:  e.err,
		},
	}
}

// As returns whether the error matches the target error.
func (e *tokenError) As(target interface{}) bool {
	if _, ok := target.(*tokenError); ok {
		return true
	}
	return false
}

// Error returns the error message.
func (e *tokenError) Error() string {
	switch e.kind {
	case CreateTokenError:
		return e.Format("unable to create auth token")
	case CreateTokenMakerError:
		return e.Format("unable to create token maker")
	default:
		return e.err.Error()
	}
}

// TrasferError contains all transfer error types.
var TokenError = struct {
	CreateTokenError      *tokenError
	CreateTokenMakerError *tokenError
}{
	CreateTokenError: &tokenError{
		&baseError{
			kind: CreateTokenError,
		},
	},
	CreateTokenMakerError: &tokenError{
		&baseError{
			kind: CreateTokenMakerError,
		},
	},
}

// NewtokenError returns a new tokenError.
func NewTokenError(kind errKind, code int, err error) *tokenError {
	return &tokenError{
		&baseError{
			kind: kind,
			err:  err,
		},
	}
}
