package app_error

import "fmt"

type baseError struct {
	kind errKind
	err  error
}

// Wrap returns a wrapped error.
func (e *baseError) Wrap(err error) *baseError {
	e.err = err
	return &baseError{
		kind: e.kind,
		err:  e.err,
	}
}

// Unwrap returns the underlying error.
func (e *baseError) Unwrap() error {
	return e.err
}

// Is returns whether the error matches the target error.
func (e *baseError) Is(target error) bool {
	if err, ok := target.(*baseError); ok {
		return e.kind == err.kind
	}
	return false
}

// As returns whether the error matches the target error.
func (e *baseError) As(target interface{}) bool {
	if _, ok := target.(*baseError); ok {
		return true
	}
	return false
}

// AsError returns the error.
func (e *baseError) AsError() error {
	return e
}

// String returns the error message.
func (e *baseError) String() string {
	return e.Error()
}

// Kind returns the error kind.
func (e *baseError) Kind() errKind {
	return e.kind
}

// Error returns the error message.
func (e *baseError) Error() string {

	return e.err.Error()

}

// FormatError returns the formated error message.
func (e *baseError) Format(msg string) string {
	if e.err != nil {
		return fmt.Sprintf("%s: %s", msg, e.err.Error())
	}
	return fmt.Sprintf("%s", msg)
}
