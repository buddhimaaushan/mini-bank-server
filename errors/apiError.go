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
	case InvalidUsernameOrPasswordError:
		return e.Format("invalid username or password")
	case InvalidRequestError:
		return e.Format("invalid request")
	case FetchingDataError:
		return e.Format("error while fetching data")
	default:
		return e.err.Error()
	}
}

// TrasferError contains all transfer error types.
var ApiError = struct {
	InvalidUsernameOrPasswordError *apiError
	InvalidRequestError            *apiError
	FetchingDataError              *apiError
}{
	InvalidUsernameOrPasswordError: &apiError{
		&baseError{
			kind: InvalidUsernameOrPasswordError,
		},
	},
	InvalidRequestError: &apiError{
		&baseError{
			kind: InvalidRequestError,
		},
	},
	FetchingDataError: &apiError{
		&baseError{
			kind: FetchingDataError,
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
