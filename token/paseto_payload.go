package token

import (
	"time"

	"github.com/google/uuid"
)

// Payload contains the payload data of the token
type PayloadPaseto struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	// Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayloadPaseto(username string, duration time.Duration) (*PayloadPaseto, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &PayloadPaseto{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *PayloadPaseto) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
