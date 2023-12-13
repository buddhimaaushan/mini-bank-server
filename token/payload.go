package token

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID         uuid.UUID `json:"id"`
	UserID     int64     `json:"user_id"`
	Username   string    `json:"username"`
	Role       string    `json:"role"`
	Department string    `json:"department"`
	IssuedAt   time.Time `json:"issued_at"`
	ExpiresAt  time.Time `json:"expired_at"`
}

// CreateToken creates a new token for a specific username and duration
func NewPayload(tokenID uuid.UUID, userID int64, username string, role string, department string, duration time.Duration) (*Payload, error) {

	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:         tokenID,
		UserID:     userID,
		Username:   username,
		Role:       role,
		Department: department,
		IssuedAt:   time.Now(),
		ExpiresAt:  time.Now().Add(duration),
	}

	return payload, nil
}

var ErrExpiredToken = fmt.Errorf("token has expired")

// VerifyToken checks if the token is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return ErrExpiredToken
	}

	return nil
}
