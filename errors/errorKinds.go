package app_error

type errKind int

const (
	_ errKind = iota
	ToAccountNotActiveError
	FromAccountNotActiveError
	InsufficientAccountBalanceError
	UniqueAccountError
	InvalidAmountError
	InvalidUsernameOrPasswordError
	InvalidRequestError
	CreateTokenError
	HashPasswordError
	CreateUserError
	UserNotFoundError
	FetchingDataError
	AccountNotFoundError
	AccountHoldersNotFoundError
	CreateTokenMakerError
	CreateSessionError
)
