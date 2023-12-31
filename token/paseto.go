package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/buddhimaaushan/mini_bank/db/sqlc"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetoMaker creates a new PasetoMaker.
func NewPasetoMaker(symmetricKey string) (Maker, error) {

	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly 32 characters")
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil

}

// CreateToken creates a new token for a specific username and duration
func (maker *PasetoMaker) CreateToken(tokenID uuid.UUID, userID int64, username string, role string, accStatus sqlc.Status, duration time.Duration) (string, *Payload, error) {

	Payload, err := NewPayload(tokenID, userID, username, role, accStatus, duration)
	if err != nil {
		return "", Payload, err
	}

	// return maker.paseto.Sign(maker.symmetricKey, payload, nil)
	token, err := maker.paseto.Encrypt(maker.symmetricKey, Payload, nil)
	return token, Payload, err
}

var ErrInvalidToken = fmt.Errorf("token is invalid")

// VerifyToken checks if the token is valid or not
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {

	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil

}
