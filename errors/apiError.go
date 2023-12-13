package app_error

type apiError struct {
	*baseError
}

// Wrap returns a wrapped error.
func (e *apiError) Wrap(err error) *apiError {
	e.err = err
	return &apiError{
		&baseError{
			kind: e.kind,
			err:  e.err,
		},
	}
}

// As returns whether the error matches the target error.
func (e *apiError) As(target interface{}) bool {
	if _, ok := target.(*apiError); ok {
		return true
	}
	return false
}

// Error returns the error message.
func (e *apiError) Error() string {
	switch e.kind {
	case ErrErrInvalidUsernameOrPassword:
		return e.Format("invalid username or password")
	case ErrInvalidRequest:
		return e.Format("invalid request")
	case ErrDataFetching:
		return e.Format("error while fetching data")
	case ErrMissingAuthHeader:
		return e.Format("authorization header is not provided")
	case ErrInvalidAuthHeader:
		return e.Format("invalid authorization header format")
	case ErrInvalidAuthType:
		return e.Format("unsupported authorization type")
	case ErrUnauthorized:
		return e.Format("you are not authorized")
	default:
		return e.err.Error()
	}
}

// TrasferError contains all transfer error types.
var ApiError = struct {
	ErrErrInvalidUsernameOrPassword *apiError
	ErrInvalidRequest               *apiError
	ErrDataFetching                 *apiError
	ErrMissingAuthHeader            *apiError
	ErrInvalidAuthHeader            *apiError
	ErrInvalidAuthType              *apiError
	ErrUnauthorized                 *apiError
}{
	ErrErrInvalidUsernameOrPassword: &apiError{
		&baseError{
			kind: ErrErrInvalidUsernameOrPassword,
		},
	},
	ErrInvalidRequest: &apiError{
		&baseError{
			kind: ErrInvalidRequest,
		},
	},
	ErrDataFetching: &apiError{
		&baseError{
			kind: ErrDataFetching,
		},
	},
	ErrMissingAuthHeader: &apiError{
		&baseError{
			kind: ErrMissingAuthHeader,
		},
	},
	ErrInvalidAuthHeader: &apiError{
		&baseError{
			kind: ErrInvalidAuthHeader,
		},
	},
	ErrInvalidAuthType: &apiError{
		&baseError{
			kind: ErrInvalidAuthType,
		},
	},
	ErrUnauthorized: &apiError{
		&baseError{
			kind: ErrUnauthorized,
		},
	},
}

// NewapiError returns a new apiError.
func NewApiError(kind errKind, code int, err error) *apiError {
	return &apiError{
		&baseError{
			kind: kind,
			err:  err,
		},
	}
}
