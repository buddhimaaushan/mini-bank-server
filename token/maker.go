package token

import (
	"time"

	"github.com/google/uuid"
)

type Maker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(tokenID uuid.UUID, userID int64, username string, role string, department string, duration time.Duration) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
