package token

import (
	"time"

	"github.com/buddhimaaushan/mini_bank/db/sqlc"
	"github.com/google/uuid"
)

type Maker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(tokenID uuid.UUID, userID int64, username string, role string, accStatus sqlc.Status, duration time.Duration) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
