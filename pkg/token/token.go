package token

import (
	"time"

	"github.com/google/uuid"
)

type Store interface {
	GenerateJWT(userID uuid.UUID, duration time.Duration) (string, error)
	VerifyJWTToken(token string) (*Payload, error)
}
