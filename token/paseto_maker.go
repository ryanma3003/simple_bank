package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	sysmetricKey []byte
}

// NewPasetoMaker creates a new PasetoMaker
func NewPasetoMaker(sysmetricKey string) (MakerPaseto, error) {
	if len(sysmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		sysmetricKey: []byte(sysmetricKey),
	}
	return maker, nil
}

func (maker *PasetoMaker) CreatePasetoToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayloadPaseto(username, duration)
	if err != nil {
		return "", err
	}

	return maker.paseto.Encrypt(maker.sysmetricKey, payload, nil)
}

func (maker *PasetoMaker) VerifyPasetoToken(token string) (*PayloadPaseto, error) {
	payload := &PayloadPaseto{}

	err := maker.paseto.Decrypt(token, maker.sysmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
