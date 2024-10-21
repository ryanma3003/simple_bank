package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  int64     `json:"issued_at"`
	ExpiredAt int64     `json:"expired_at"`
}

// NewPayload creates a new token payload with a spesific username and duration
func NewPayload(username string, duration time.Duration) (jwt.MapClaims, error) {
	payload := &Payload{}

	var err error
	payload.ID, err = uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload.Username = username
	payload.ExpiredAt = time.Now().Add(duration).Unix()
	payload.IssuedAt = time.Now().Unix()

	payloads := jwt.MapClaims{}

	payloads["sub"] = payload.ID
	payloads["aud"] = payload.Username
	payloads["iat"] = payload.IssuedAt
	payloads["exp"] = payload.ExpiredAt

	return payloads, nil
}

// valid checks if the token payload is valid
func (payload *Payload) Valid() error {
	t := time.Unix(0, payload.ExpiredAt)
	if time.Now().After(t) {
		return ErrExpiredToken
	}

	return nil
}
