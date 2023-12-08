package app_error

type hashError struct {
	*baseError
}

// Wrap returns a wrapped error.
func (e *hashError) Wrap(err error) *hashError {
	e.err = err
	return &hashError{
		&baseError{
			kind: e.kind,
			err:  e.err,
		},
	}
}

// As returns whether the error matches the target error.
func (e *hashError) As(target interface{}) bool {
	if _, ok := target.(*hashError); ok {
		return true
	}
	return false
}

// Error returns the error message.
func (e *hashError) Error() string {
	switch e.kind {
	case HashPasswordError:
		return e.Format("unable to hash password")

	default:
		return e.err.Error()
	}
}

// TrasferError contains all transfer error types.
var HashError = struct {
	HashPasswordError *hashError
}{
	HashPasswordError: &hashError{
		&baseError{
			kind: HashPasswordError,
		},
	},
}

// NewhashError returns a new hashError.
func NewHashError(kind errKind, code int, err error) *hashError {
	return &hashError{
		&baseError{
			kind: kind,
			err:  err,
		},
	}
}
