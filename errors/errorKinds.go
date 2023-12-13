package app_error

type errKind int

const (
	_ errKind = iota
	ErrToAccountNotActive
	ErrFromAccountNotActive
	ErrInsufficientAccountBalance
	ErrUniqueAccount
	ErrInvalidAmount
	ErrErrInvalidUsernameOrPassword
	ErrInvalidRequest
	ErrCreateToken
	ErrHashPassword
	ErrCreateUser
	ErrUserNotFound
	ErrDataFetching
	ErrAccountNotFound
	ErrAccountHoldersNotFound
	ErrCreateTokenMaker
	ErrCreateSession
	ErrMissingAuthHeader
	ErrInvalidAuthHeader
	ErrInvalidAuthType
	ErrInvalidToken
	ErrUnauthorized
)
