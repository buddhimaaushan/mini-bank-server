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
	case ErrCreateToken:
		return e.Format("unable to create auth token")
	case ErrCreateTokenMaker:
		return e.Format("unable to create token maker")
	case ErrInvalidToken:
		return e.Format("invalid token")
	default:
		return e.err.Error()
	}
}

// TrasferError contains all transfer error types.
var TokenError = struct {
	ErrCreateToken      *tokenError
	ErrCreateTokenMaker *tokenError
	ErrInvalidToken     *tokenError
}{
	ErrCreateToken: &tokenError{
		&baseError{
			kind: ErrCreateToken,
		},
	},
	ErrCreateTokenMaker: &tokenError{
		&baseError{
			kind: ErrCreateTokenMaker,
		},
	},
	ErrInvalidToken: &tokenError{
		&baseError{
			kind: ErrInvalidToken,
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
